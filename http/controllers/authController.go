package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"test/dto"
	http2 "test/http"
)

type AuthService interface {
	Register(context.Context, dto.RegisterReq) (*dto.RegisterRes, error)
	Login(context.Context, dto.LoginReq) (*dto.LoginRes, error)
	LoginOut(context.Context, dto.LoginOutReq) (*dto.LoginOutRes, error)
}

type AuthController struct {
	service AuthService
}

func NewAuthController(s AuthService) *AuthController {
	return &AuthController{s}
}

func (c *AuthController) Register(w http.ResponseWriter, r *http.Request) {
	var request dto.RegisterReq
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http2.WriteResponse(w, http.StatusBadRequest, fmt.Errorf("could not parse JSON: %w", err))
		return
	}
	response, err := c.service.Register(r.Context(), request)
	if err != nil {
		http2.WriteError(w, err)
	} else {
		http2.WriteResponse(w, http.StatusOK, response)
	}
}

func (c *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	var request dto.LoginReq
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http2.WriteResponse(w, http.StatusBadRequest, fmt.Errorf("could not parse JSON: %w", err))
		return
	}

	response, err := c.service.Login(r.Context(), request)
	if err != nil {
		http2.WriteError(w, err)
	} else {
		http2.WriteResponse(w, http.StatusOK, response)
	}
}

func (c *AuthController) LoginOut(w http.ResponseWriter, r *http.Request) {

	var request dto.LoginOutReq
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http2.WriteResponse(w, http.StatusBadRequest, fmt.Errorf("could not parse JSON: %w", err))
		return
	}

	response, err := c.service.LoginOut(r.Context(), request)
	if err != nil {
		http2.WriteError(w, err)
	} else {
		http2.WriteResponse(w, http.StatusOK, response)
	}
}
