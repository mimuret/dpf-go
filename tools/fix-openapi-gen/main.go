package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func main() {
	err := filepath.Walk("api", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() || filepath.Ext(path) != ".go" {
			return nil
		}

		content, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		newContent := fixArrayDefault(content)
		newContent = addOtelImports(path, newContent)

		if filepath.Base(path) == "client.go" {
			newContent = fixClient(newContent)
		} else if strings.HasPrefix(filepath.Base(path), "api_") || filepath.Base(path) == "api_execute_all_gen.go" {
			newContent = fixApiMethods(newContent)
		}

		if !bytes.Equal(content, newContent) {
			if err := os.WriteFile(path, newContent, info.Mode()); err != nil {
				return err
			}
			fmt.Printf("Fixed %s\n", path)
		}

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
}

func fixArrayDefault(content []byte) []byte {
	re := regexp.MustCompile(`(?m)var defaultValue (\[\].+) = \[\](\s*)$`)
	content = re.ReplaceAllFunc(content, func(match []byte) []byte {
		submatches := re.FindSubmatch(match)
		if len(submatches) < 3 {
			return match
		}
		typeStr := string(submatches[1])
		space := string(submatches[2])
		return []byte(fmt.Sprintf("var defaultValue %s = %s{}%s", typeStr, typeStr, space))
	})
	// Handle cases where it's not at the end of the line (e.g. followed by a comment or just no newline in some contexts)
	re2 := regexp.MustCompile(`(?m)var defaultValue (\[\].+) = \[\](\s+[^ \r\n[])`)
	content = re2.ReplaceAllFunc(content, func(match []byte) []byte {
		submatches := re2.FindSubmatch(match)
		if len(submatches) < 3 {
			return match
		}
		typeStr := string(submatches[1])
		rest := string(submatches[2])
		return []byte(fmt.Sprintf("var defaultValue %s = %s{}%s", typeStr, typeStr, rest))
	})
	return content
}

func addOtelImports(path string, content []byte) []byte {
	if strings.Contains(string(content), "go.opentelemetry.io/otel") {
		return content
	}

	importRe := regexp.MustCompile(`(?m)^import \(\n`)
	if !importRe.Match(content) {
		return content
	}

	otelImports := "\t\"go.opentelemetry.io/otel\"\n\t\"go.opentelemetry.io/otel/trace\"\n\t\"go.opentelemetry.io/otel/codes\"\n"
	if filepath.Base(path) == "client.go" {
		otelImports += "\t\"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp\"\n"
	}

	return importRe.ReplaceAll(content, []byte("import (\n"+otelImports))
}

func fixClient(content []byte) []byte {
	if strings.Contains(string(content), "\ttracer trace.Tracer") {
		return content
	}

	// Add tracer field to APIClient
	structRe := regexp.MustCompile(`type APIClient struct \{`)
	content = structRe.ReplaceAll(content, []byte("type APIClient struct {\n\ttracer trace.Tracer"))

	// Initialize tracer in NewAPIClient
	newClientRe := regexp.MustCompile(`(?m)func NewAPIClient\(cfg \*Configuration\) \*APIClient \{\n\tif cfg\.HTTPClient == nil \{\n\t\tcfg\.HTTPClient = http\.DefaultClient\n\t\}`)
	content = newClientRe.ReplaceAll(content, []byte(`func NewAPIClient(cfg *Configuration) *APIClient {
	if cfg.HTTPClient == nil {
		cfg.HTTPClient = http.DefaultClient
	}
	if cfg.HTTPClient.Transport == nil {
		cfg.HTTPClient.Transport = http.DefaultTransport
	}
	cfg.HTTPClient.Transport = otelhttp.NewTransport(cfg.HTTPClient.Transport)`))

	// In NewAPIClient, assign tracer
	assignTracerRe := regexp.MustCompile(`c := &APIClient\{\}`)
	content = assignTracerRe.ReplaceAll(content, []byte(`c := &APIClient{}
	c.tracer = otel.Tracer("github.com/mimuret/dpf-go/api")`))

	// Instrument callAPI
	callApiRe := regexp.MustCompile(`(?m)func \(c \*APIClient\) callAPI\(request \*http\.Request\) \(\*http\.Response, error\) \{`)
	content = callApiRe.ReplaceAll(content, []byte(`func (c *APIClient) callAPI(request *http.Request) (*http.Response, error) {
	span := trace.SpanFromContext(request.Context())`))

	callApiDoRe := regexp.MustCompile(`(?m)resp, err := c\.cfg\.HTTPClient\.Do\(request\)\n\tif err != nil \{\n\t\treturn resp, err\n\t\}`)
	content = callApiDoRe.ReplaceAll(content, []byte(`resp, err := c.cfg.HTTPClient.Do(request)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return resp, err
	}
	if resp.StatusCode >= 400 {
		span.SetStatus(codes.Error, resp.Status)
	}`))
	// Add DPFError helper to GenericOpenAPIError
	genericErrorRe := regexp.MustCompile(`(?m)// Body returns the raw bytes of the response\nfunc \(e GenericOpenAPIError\) Body\(\) \[\]byte \{\n\treturn e\.body\n\}\n\n// Model returns the unpacked model of the error\nfunc \(e GenericOpenAPIError\) Model\(\) interface\{\} \{\n\treturn e\.model\n\}`)
	dpfErrorHelper := `// Body returns the raw bytes of the response
func (e GenericOpenAPIError) Body() []byte {
	return e.body
}

// Model returns the unpacked model of the error
func (e GenericOpenAPIError) Model() interface{} {
	return e.model
}

// DPFError returns the structured DPF error if available.
func (e GenericOpenAPIError) DPFError() DPFError {
	if e.model == nil {
		return nil
	}
	if v, ok := e.model.(DPFError); ok {
		return v
	}
	// If it's a value type, try to get a pointer to it
	val := reflect.ValueOf(e.model)
	if val.Kind() != reflect.Ptr {
		ptr := reflect.New(val.Type())
		ptr.Elem().Set(val)
		if v, ok := ptr.Interface().(DPFError); ok {
			return v
		}
	}
	return nil
}`
	content = genericErrorRe.ReplaceAll(content, []byte(dpfErrorHelper))

	// Update GenericOpenAPIError.Error() to show structured error
	errorMethodRe := regexp.MustCompile(`(?m)func \(e GenericOpenAPIError\) Error\(\) string \{\n\treturn e\.error\n\}`)
	newErrorMethod := `func (e GenericOpenAPIError) Error() string {
	if der := e.DPFError(); der != nil {
		return der.Error()
	}
	return e.error
}`
	content = errorMethodRe.ReplaceAll(content, []byte(newErrorMethod))

	return content
}

func fixApiMethods(content []byte) []byte {
	if strings.Contains(string(content), "tracer.Start(r.ctx,") {
		return content
	}

	// Standard Execute methods
	// match: func (a *ZonesAPIService) DeleteZoneChangesExecute(r ApiDeleteZoneChangesRequest) (*AsyncResponse, *http.Response, error) {
	methodRe := regexp.MustCompile(`(?m)^func \(a \*([A-Z][a-zA-Z0-9_]+APIService)\) ([A-Z][a-zA-Z0-9_]+Execute)\(r ([A-Z][a-zA-Z0-9_]+Request)\) \(([^)]+)\) \{`)

	content = methodRe.ReplaceAllFunc(content, func(match []byte) []byte {
		submatches := methodRe.FindSubmatch(match)
		serviceName := string(submatches[1])
		methodName := string(submatches[2])
		requestType := string(submatches[3])
		returnTypes := string(submatches[4])

		replacement := fmt.Sprintf(`func (a *%s) %s(r %s) (%s) {
	ctx, span := a.client.tracer.Start(r.ctx, "%s.%s")
	defer span.End()
	r.ctx = ctx`, serviceName, methodName, requestType, returnTypes, serviceName, methodName)

		return []byte(replacement)
	})

	// ExecuteAll methods (already have tracer code if generated by updated gen-execute-all, but this tool runs AFTER)
	// match: func (r ApiGetZoneListRequest) ExecuteAll() (*GetZones, *http.Response, error) {
	// If gen-execute-all hasn't been run yet, this will apply it.
	executeAllRe := regexp.MustCompile(`(?m)^func \(r ([A-Z][a-zA-Z0-9_]+Request)\) ExecuteAll\(\) \(([^)]+)\) \{`)
	content = executeAllRe.ReplaceAllFunc(content, func(match []byte) []byte {
		submatches := executeAllRe.FindSubmatch(match)
		requestType := string(submatches[1])
		returnTypes := string(submatches[2])

		types := strings.Split(returnTypes, ",")
		if len(types) != 3 {
			return match
		}

		replacement := fmt.Sprintf(`func (r %s) ExecuteAll() (res %s, lastResp %s, err error) {
	ctx, span := r.ApiService.client.tracer.Start(r.ctx, "%s.ExecuteAll")
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
		}
		span.End()
	}()
	r.ctx = ctx`, requestType, strings.TrimSpace(types[0]), strings.TrimSpace(types[1]), requestType)

		return []byte(replacement)
	})

	return content
}
