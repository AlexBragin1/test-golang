package middleware

import (
	"fmt"
	"io"
	"net/http"

	http2 "test/http"
	"test/logger"
)

type DumpMiddleware struct {
}

func NewDumpMiddleware() *DumpMiddleware {
	return &DumpMiddleware{}
}

func (mw *DumpMiddleware) Dump(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		requestBytes, err := io.ReadAll(r.Body)
		if err != nil {
			http2.WriteResponse(w, http.StatusInternalServerError, fmt.Errorf("can't read request body: %w", err))
			return
		}
		logger.Info(fmt.Sprintf("request body: %s", requestBytes))

		next(w, r)
	}
}
