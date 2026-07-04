package batch

import (
	"net/http"
)

type BatchItem struct {
	Request  *http.Request
	Response chan *BatchItemResult
}

type BatchItemResult struct {
	Index      int
	StatusCode int
	Body       []byte
	Headers    http.Header
}

type RequestBox struct {
	ID       string
	Key      string
	Requests []*BatchItem
}

func SplitResults(box *RequestBox, results []BatchItemResult) {
	for _, result := range results {
		if result.Index < len(box.Requests) {
			item := box.Requests[result.Index]
			item.Response <- &result
		}
	}
}
