package dto

import (
	"fmt"
	"net/http"
	"strconv"
	"test/domain"
	"test/errors"

	"github.com/gorilla/mux"
)

type ListVariantsReq struct {
	UserID domain.UUID `json:"user_id"`
}

type ListVariantsRes struct {
	Variants []domain.Variant `json:"variants"`
}

func (d ListVariantsReq) Validate() error {
	if len(d.UserID) > 0 {

		return errors.NewValidationError("error validate")
	}

	return nil
}

type GetTaskReq struct {
	UserID    domain.UUID `json:"user_id"`
	VariantID domain.UUID `json:"variant_id"`
	TaskID    int         `json:"task_id"`
}

func NewGetTaskRequest(r *http.Request) (*GetTaskReq, error) {
	request := &GetTaskReq{}

	vars := mux.Vars(r)
	request.UserID = domain.UUID(vars["user_id"])
	request.VariantID = domain.UUID(vars["variant_id"])

	taskIDStr := vars["task_id"]

	var err error

	request.TaskID, err = strconv.Atoi(taskIDStr)
	if err != nil {
		return nil, fmt.Errorf("invalid task ID")

	}

	return request, nil
}

func (d GetTaskReq) Validate() error {
	if len(d.VariantID) > 0 && len(d.UserID) > 0 {

		return errors.NewValidationError("Error validate")
	}

	return nil
}

type GetTaskRes struct {
	TaskID      int    `json:"task_id"`
	Description string `json:"description"`
	Options     string `json:"options"`
}

type AnswerTaskReq struct {
	UserID    domain.UUID `json:"user_id"`
	VariantID domain.UUID `json:"variant_id"`
	TaskID    int         `json:"task_id"`
	AnswerID  domain.UUID `json:"answer_id"`
	Answer    string      `json:"answer"`
}

func NewAnswerTaskRequest(r *http.Request) (*AnswerTaskReq, error) {
	request := &AnswerTaskReq{}

	vars := mux.Vars(r)
	request.UserID = domain.UUID(vars["user_id"])
	request.VariantID = domain.UUID(vars["variant_id"])

	taskIDStr := vars["task_id"]

	var err error

	request.TaskID, err = strconv.Atoi(taskIDStr)
	if err != nil {
		return nil, fmt.Errorf("invalid task ID")

	}

	return request, nil
}

type AnswerTaskRes struct {
	AnswerID domain.UUID `json:"answer_id"`
}

func (d AnswerTaskReq) Validate() error {
	if len(d.VariantID) > 0 && len(d.UserID) > 0 {

		return errors.NewValidationError("")
	}

	return nil
}

type ResultVariantReq struct {
	UserID    domain.UUID `json:"user_id"`
	VariantID domain.UUID `json:"variant_id"`
	AnswerID  domain.UUID `json:"answer_id"`
}

func NewResultsVariantRequest(r *http.Request) (*ResultVariantReq, error) {
	request := &ResultVariantReq{}

	vars := mux.Vars(r)
	request.UserID = domain.UUID(vars["user_id"])
	request.VariantID = domain.UUID(vars["variant_id"])

	return request, nil
}

type ResultVariantRes struct {
	Percent string `json:"percent"`
}

func (d ResultVariantReq) Validate() error {
	if len(d.VariantID) > 0 && len(d.UserID) > 0 {

		return errors.NewValidationError("")
	}

	return nil
}
