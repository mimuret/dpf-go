package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSyncWaitSuccessful(t *testing.T) {
	count := 0
	reqID := "782d746ac3cb46499b31708fa80e8660"
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/jobs/"+reqID {
			w.Header().Set("Content-Type", "application/json")
			count++
			if count < 2 {
				_ = json.NewEncoder(w).Encode(GetJobs{
					RequestId: reqID,
					Status:    "RUNNING",
				})
			} else {
				_ = json.NewEncoder(w).Encode(GetJobs{
					RequestId: reqID,
					Status:    "SUCCESSFUL",
				})
			}
			return
		}
		w.WriteHeader(http.StatusNotFound)
	}))
	defer ts.Close()

	cfg := NewConfiguration()
	cfg.Servers = ServerConfigurations{
		{
			URL: ts.URL,
		},
	}
	client := NewAPIClient(cfg)

	asyncResp := &AsyncResponse{
		RequestId: reqID,
		JobsUrl:   ts.URL + "/jobs/" + reqID,
	}

	req, _ := http.NewRequestWithContext(context.Background(), "PATCH", "/xxx", nil)
	resp := &http.Response{
		StatusCode: http.StatusAccepted,
		Request:    req,
	}

	job, _, err := client.JobsAPI.SyncWait(asyncResp, resp, nil)
	if err != nil {
		t.Fatalf("SyncWait failed: %v", err)
	}
	if job.GetStatus() != "SUCCESSFUL" {
		t.Errorf("expected SUCCESSFUL, got %s", job.GetStatus())
	}
	if count != 2 {
		t.Errorf("expected 2 calls, got %d", count)
	}
}

func TestSyncWaitFailed(t *testing.T) {
	reqID := "782d746ac3cb46499b31708fa80e8660"
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/jobs/"+reqID {
			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(GetJobs{
				RequestId:    reqID,
				Status:       "FAILED",
				ErrorType:    PtrString("TestError"),
				ErrorMessage: PtrString("test error message"),
			})
			return
		}
		w.WriteHeader(http.StatusNotFound)
	}))
	defer ts.Close()

	cfg := NewConfiguration()
	cfg.Servers = ServerConfigurations{
		{
			URL: ts.URL,
		},
	}
	client := NewAPIClient(cfg)

	asyncResp := &AsyncResponse{
		RequestId: reqID,
		JobsUrl:   ts.URL + "/jobs/" + reqID,
	}

	req, _ := http.NewRequestWithContext(context.Background(), "PATCH", "/xxx", nil)
	resp := &http.Response{
		StatusCode: http.StatusAccepted,
		Request:    req,
	}

	_, _, err := client.JobsAPI.SyncWait(asyncResp, resp, nil)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	syncErr, ok := err.(*SyncWaitError)
	if !ok {
		t.Fatalf("expected *SyncWaitError, got %T", err)
	}
	if syncErr.RequestId != reqID {
		t.Errorf("expected RequestId %s, got %s", reqID, syncErr.RequestId)
	}
	if syncErr.Job == nil || syncErr.Job.GetStatus() != "FAILED" {
		t.Errorf("expected FAILED job in error, got %v", syncErr.Job)
	}
	if syncErr.Response == nil {
		t.Error("expected Response in error, but got nil")
	}

	expected := "job 782d746ac3cb46499b31708fa80e8660 failed: TestError: test error message"
	if syncErr.Error() != expected {
		t.Errorf("expected %q, got %q", expected, syncErr.Error())
	}
}

func TestSyncWaitError(t *testing.T) {
	cfg := NewConfiguration()
	client := NewAPIClient(cfg)

	expectedErr := fmt.Errorf("pre-existing error")
	_, _, err := client.JobsAPI.SyncWait(nil, nil, expectedErr)
	if err != expectedErr {
		t.Errorf("expected %v, got %v", expectedErr, err)
	}

	_, _, err = client.JobsAPI.SyncWait(nil, &http.Response{}, nil)
	if err == nil || err.Error() != "asyncResp is nil" {
		t.Errorf("expected asyncResp is nil error, got %v", err)
	}
}

func TestSyncWaitCancelled(t *testing.T) {
	reqID := "782d746ac3cb46499b31708fa80e8660"
	ctx, cancel := context.WithCancel(context.Background())
	cancel() // Cancel immediately

	cfg := NewConfiguration()
	client := NewAPIClient(cfg)

	asyncResp := &AsyncResponse{
		RequestId: reqID,
		JobsUrl:   "http://localhost/jobs/" + reqID,
	}

	req, _ := http.NewRequestWithContext(ctx, "PATCH", "/xxx", nil)
	resp := &http.Response{
		StatusCode: http.StatusAccepted,
		Request:    req,
	}

	_, _, err := client.JobsAPI.SyncWait(asyncResp, resp, nil)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	syncErr, ok := err.(*SyncWaitError)
	if !ok {
		t.Fatalf("expected *SyncWaitError, got %T", err)
	}
	if syncErr.RequestId != reqID {
		t.Errorf("expected RequestId %s, got %s", reqID, syncErr.RequestId)
	}
	if syncErr.Err != context.Canceled {
		t.Errorf("expected context.Canceled, got %v", syncErr.Err)
	}
}
