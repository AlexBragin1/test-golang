package middleware

import (
	"fmt"
	"net/http"
	"strconv"

	"test/domain"
	"test/domain/groups"
	http2 "test/http"

	"github.com/gorilla/mux"
)

type RequestTokenService interface {
	ReadFromRequest(r *http.Request, keys ...string) (map[string]string, error)
}

type AuthMiddleware struct {
	tokenService RequestTokenService
}

func NewAuthMiddleware(tokenService RequestTokenService) *AuthMiddleware {
	return &AuthMiddleware{tokenService: tokenService}
}

func (mw *AuthMiddleware) WithToken(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenMap, err := mw.tokenService.ReadFromRequest(r, "flow")
		if err != nil {
			http2.WriteResponse(w, http.StatusUnauthorized, err)
			return
		}

		if tokenMap["flow"] != string(domain.FLOW_AUTHORIZATION) {
			http2.WriteResponse(w, http.StatusUnauthorized, fmt.Errorf("token type error"))
			return
		}

		next(w, r)
	}
}

func (mw *AuthMiddleware) RouterWithGroup(resourceGroups groups.Groups) mux.MiddlewareFunc {
	return func(h http.Handler) http.Handler {
		return mw.withGroup(resourceGroups, h.ServeHTTP)
	}
}

type RouteWrap func(handlerFunc http.HandlerFunc) http.HandlerFunc

func (mw *AuthMiddleware) WrapWithGroup(resourceGroups groups.Groups) RouteWrap {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return mw.withGroup(resourceGroups, next)
	}
}

func (mw *AuthMiddleware) withGroup(resourceGroups groups.Groups, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenMap, err := mw.tokenService.ReadFromRequest(r, "flow", "groups")
		if err != nil {
			http2.WriteResponse(w, http.StatusUnauthorized, err)
			return
		}

		if domain.FLOW_AUTHORIZATION != domain.Flow(tokenMap["flow"]) {
			http2.WriteResponse(w, http.StatusUnauthorized, fmt.Errorf("token type error"))
			return
		}

		groupsInt, _ := strconv.Atoi(tokenMap["groups"])
		if !groups.Groups(groupsInt).HasAccessTo(resourceGroups) {
			http2.WriteResponse(w, http.StatusUnauthorized, fmt.Errorf("access denied: incorrect group"))
			return
		}

		next(w, r)
	}
}

func (mw *AuthMiddleware) WrapWithPathMatch() RouteWrap {
	return func(h http.HandlerFunc) http.HandlerFunc {
		return mw.pathMatchToken(h.ServeHTTP)
	}
}

func (mw *AuthMiddleware) pathMatchToken(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		if vars["user_id"] == "" {
			http2.WriteResponse(w, http.StatusUnauthorized, fmt.Errorf("used_id is empty"))
			return
		}

		tokenMap, err := mw.tokenService.ReadFromRequest(r, "flow", "user_id")
		if err != nil {
			http2.WriteResponse(w, http.StatusUnauthorized, err)
			return
		}

		if domain.FLOW_AUTHORIZATION != domain.Flow(tokenMap["flow"]) {
			http2.WriteResponse(w, http.StatusUnauthorized, fmt.Errorf("token type error"))
			return
		}

		if vars["user_id"] != tokenMap["user_id"] {
			http2.WriteResponse(w, http.StatusUnauthorized, fmt.Errorf("access denied: incorrect user"))
			return
		}

		next(w, r)
	}
}

func (mw *AuthMiddleware) WrapWithPathMatchOrAdmin() RouteWrap {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return mw.pathMatchOrAdminToken(next)
	}
}

func (mw *AuthMiddleware) pathMatchOrAdminToken(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		if vars["user_id"] == "" {
			http2.WriteResponse(w, http.StatusUnauthorized, fmt.Errorf("used_id is empty"))
			return
		}

		tokenMap, err := mw.tokenService.ReadFromRequest(r, "flow", "user_id", "groups")
		if err != nil {
			http2.WriteResponse(w, http.StatusUnauthorized, err)
			return
		}

		if domain.FLOW_AUTHORIZATION != domain.Flow(tokenMap["flow"]) {
			http2.WriteResponse(w, http.StatusUnauthorized, fmt.Errorf("token type error"))
			return
		}

		groupsInt, _ := strconv.Atoi(tokenMap["groups"])
		if groups.Groups(groupsInt).HasAccessTo(groups.Admin) {
			next(w, r)
			return
		}

		if vars["user_id"] != tokenMap["user_id"] {
			http2.WriteResponse(w, http.StatusUnauthorized, fmt.Errorf("access denied: incorrect user"))
			return
		}

		next(w, r)
	}
}
