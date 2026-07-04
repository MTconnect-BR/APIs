package batch

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"sync"
)

type Processor interface {
	ProcessBox(box *RequestBox) ([]BatchItemResult, error)
}

type DefaultProcessor struct {
	handler http.Handler
}

func NewDefaultProcessor(handler http.Handler) *DefaultProcessor {
	return &DefaultProcessor{handler: handler}
}

func (p *DefaultProcessor) ProcessBox(box *RequestBox) ([]BatchItemResult, error) {
	results := make([]BatchItemResult, len(box.Requests))
	var wg sync.WaitGroup

	for i, item := range box.Requests {
		wg.Add(1)
		go func(idx int, reqItem *BatchItem) {
			defer wg.Done()

			recorder := httptest.NewRecorder()

			newReq := reqItem.Request.Clone(reqItem.Request.Context())
			if reqItem.Request.Body != nil {
				body, _ := io.ReadAll(reqItem.Request.Body)
				newReq.Body = io.NopCloser(bytes.NewReader(body))
				newReq.ContentLength = int64(len(body))
			}

			p.handler.ServeHTTP(recorder, newReq)

			results[idx] = BatchItemResult{
				Index:      idx,
				StatusCode: recorder.Code,
				Body:       recorder.Body.Bytes(),
				Headers:    recorder.Header(),
			}
		}(i, item)
	}

	wg.Wait()

	return results, nil
}
