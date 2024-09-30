package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"test/domain"
	"test/dto"
	http2 "test/http"

	"github.com/gorilla/mux"
)

type TestService interface {
	ListVariants(context.Context, dto.ListVariantsReq) (*dto.ListVariantsRes, error)
	AnswerTask(context.Context, dto.AnswerTaskReq) (*dto.AnswerTaskRes, error)
	GetTask(context.Context, dto.GetTaskReq) (*dto.GetTaskRes, error)
	Result(context.Context, dto.ResultVariantReq) (*dto.ResultVariantRes, error)
}

type TestController struct {
	service TestService
}

func NewTestController(s TestService) *TestController {
	return &TestController{s}
}

func (c *TestController) ListVariants(w http.ResponseWriter, r *http.Request) {
	var request dto.ListVariantsReq
	vars := mux.Vars(r)
	request.UserID = domain.UUID(vars["user_id"])

	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		http2.WriteResponse(w, http.StatusBadRequest, fmt.Errorf("could not parse JSON: %w", err))
		return
	}

	response, err := c.service.ListVariants(r.Context(), request)
	if err != nil {
		http2.WriteError(w, err)
	} else {
		http2.WriteResponse(w, http.StatusOK, response)
	}
}

func (c *TestController) GetTask(w http.ResponseWriter, r *http.Request) {
	request, err := dto.NewGetTaskRequest(r)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		http2.WriteResponse(w, http.StatusBadRequest, fmt.Errorf("could not parse JSON: %w", err))
		return
	}

	response, err := c.service.GetTask(r.Context(), *request)
	if err != nil {
		http2.WriteError(w, err)
	} else {
		http2.WriteResponse(w, http.StatusOK, response)
	}
}

func (c *TestController) AnswerTask(w http.ResponseWriter, r *http.Request) {
	request, err := dto.NewAnswerTaskRequest(r)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		http2.WriteResponse(w, http.StatusBadRequest, fmt.Errorf("could not parse JSON: %w", err))
		return
	}

	response, err := c.service.AnswerTask(r.Context(), *request)
	if err != nil {
		http2.WriteError(w, err)
	} else {
		http2.WriteResponse(w, http.StatusOK, response)
	}
}

func (c *TestController) Result(w http.ResponseWriter, r *http.Request) {
	request, err := dto.NewResultRequest(r)
	if err != nil {
		http.Error(w, "Invalid test user ID", http.StatusBadRequest)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		http2.WriteResponse(w, http.StatusBadRequest, fmt.Errorf("could not parse JSON: %w", err))
		return
	}

	response, err := c.service.Result(r.Context(), *request)
	if err != nil {
		http2.WriteError(w, err)
	} else {
		http2.WriteResponse(w, http.StatusOK, response)
	}
}
