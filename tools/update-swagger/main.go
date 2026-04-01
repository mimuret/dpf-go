package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	data, err := os.ReadFile("api/swagger.json")
	if err != nil {
		log.Fatal(err)
	}

	var sw map[string]interface{}
	if err := json.Unmarshal(data, &sw); err != nil {
		log.Fatal(err)
	}

	components, ok := sw["components"].(map[string]interface{})
	if !ok {
		components = make(map[string]interface{})
		sw["components"] = components
	}

	schemas, ok := components["schemas"].(map[string]interface{})
	if !ok {
		schemas = make(map[string]interface{})
		components["schemas"] = schemas
	}

	// Define common schemas
	errorModels := []string{"NotFound", "TooManyRequests", "SystemError", "ServiceUnavailable", "GatewayTimeout"}
	for _, model := range errorModels {
		if _, ok := schemas[model]; !ok {
			schemas[model] = map[string]interface{}{
				"type":     "object",
				"required": []string{"request_id", "error_type", "error_message", "error_details"},
				"properties": map[string]interface{}{
					"request_id":    map[string]interface{}{"$ref": "#/components/schemas/RequestId"},
					"error_type":    map[string]interface{}{"$ref": "#/components/schemas/ErrorType"},
					"error_message": map[string]interface{}{"$ref": "#/components/schemas/ErrorMessage"},
					"error_details": map[string]interface{}{"$ref": "#/components/schemas/ErrorDetails"},
				},
			}
		}
	}

	// Define common responses
	responsesComp, ok := components["responses"].(map[string]interface{})
	if !ok {
		responsesComp = make(map[string]interface{})
		components["responses"] = responsesComp
	}

	commonResponses := map[string]struct {
		description string
		model       string
	}{
		"404": {"Specified resource not found.", "NotFound"},
		"429": {"Too many requests.", "TooManyRequests"},
		"500": {"System error occurred.", "SystemError"},
		"503": {"Service unavailable.", "ServiceUnavailable"},
		"504": {"Gateway timeout.", "GatewayTimeout"},
	}

	for _, info := range commonResponses {
		if _, ok := responsesComp[info.model]; !ok {
			responsesComp[info.model] = map[string]interface{}{
				"description": info.description,
				"content": map[string]interface{}{
					"application/json": map[string]interface{}{
						"schema": map[string]interface{}{
							"$ref": "#/components/schemas/" + info.model,
						},
					},
				},
			}
		}
	}

	// Paths
	paths, ok := sw["paths"].(map[string]interface{})
	if ok {
		for _, methods := range paths {
			methodsMap, ok := methods.(map[string]interface{})
			if !ok {
				continue
			}
			for method, opRaw := range methodsMap {
				if method == "parameters" {
					continue
				}
				op, ok := opRaw.(map[string]interface{})
				if !ok {
					continue
				}
				respRaw, ok := op["responses"]
				if !ok {
					respRaw = make(map[string]interface{})
					op["responses"] = respRaw
				}
				responses := respRaw.(map[string]interface{})

				for code, info := range commonResponses {
					if _, ok := responses[code]; !ok {
						responses[code] = map[string]interface{}{
							"$ref": "#/components/responses/" + info.model,
						}
					}
				}
			}
		}
	}

	out, err := json.MarshalIndent(sw, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	if err := os.WriteFile("api/swagger.json", out, 0644); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Updated api/swagger.json with common error responses (preserved other fields).")
}
