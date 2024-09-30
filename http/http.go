package http

import (
	"encoding/json"
	"errors"
	"net/http"

	errors2 "gitlab.com/usdtkg/payout/errors"
	"gitlab.com/usdtkg/payout/logger"
	"go.uber.org/zap"
)

func WriteResponse(w http.ResponseWriter, code int, data any) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)

	if err, ok := data.(error); ok {
		data = err.Error()
	}

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		logger.Error("error when writing response", zap.Error(err))
	}
}

func WriteError(w http.ResponseWriter, err error) {
	var code int = 500

	var appError *errors2.AppError

	if errors.As(err, &appError) {
		code = appError.Code
	}

	WriteResponse(w, code, err.Error())
}
