package api

import (
	"fmt"
	"net/http"
	"time"

	"go.opentelemetry.io/otel/codes"
)

// SyncWaitError は SyncWait の実行中に発生したエラーを表す構造体です。
type SyncWaitError struct {
	RequestId string
	Job       *GetJobs
	Response  *http.Response
	Err       error
}

func (e *SyncWaitError) Error() string {
	if e.Job != nil && e.Job.GetStatus() == "FAILED" {
		return fmt.Sprintf("job %s failed: %s: %s", e.RequestId, e.Job.GetErrorType(), e.Job.GetErrorMessage())
	}
	return fmt.Sprintf("sync wait for %s failed: %v", e.RequestId, e.Err)
}

func (e *SyncWaitError) Unwrap() error {
	return e.Err
}

// SyncWait は非同期リクエストを受け取った後、そのジョブが完了するまで待機します。
// ステータスが SUCCESSFUL になれば結果を返し、 FAILED の場合はエラーを返します。
// エラーが発生した場合は、詳細情報を含む *SyncWaitError 型のエラーを返します。
func (a *JobsAPIService) SyncWait(asyncResp *AsyncResponse, resp *http.Response, err error) (job *GetJobs, jobResp *http.Response, retErr error) {
	if err != nil {
		return nil, resp, err
	}
	if asyncResp == nil {
		return nil, resp, fmt.Errorf("asyncResp is nil")
	}

	ctx, span := a.client.tracer.Start(resp.Request.Context(), "JobsAPIService.SyncWait")
	defer func() {
		if retErr != nil {
			span.RecordError(retErr)
			span.SetStatus(codes.Error, retErr.Error())
		}
		span.End()
	}()

	reqID := asyncResp.RequestId
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return nil, nil, &SyncWaitError{
				RequestId: reqID,
				Response:  nil,
				Err:       ctx.Err(),
			}
		case <-ticker.C:
			res, resp, err := a.GetJob(ctx, reqID).Execute()
			if err != nil {
				return nil, resp, &SyncWaitError{
					RequestId: reqID,
					Response:  resp,
					Err:       err,
				}
			}
			if res.GetStatus() != "RUNNING" {
				if res.GetStatus() == "FAILED" {
					return res, resp, &SyncWaitError{
						RequestId: reqID,
						Job:       res,
						Response:  resp,
						Err:       fmt.Errorf("job failed"),
					}
				}
				return res, resp, nil
			}
		}
	}
}
