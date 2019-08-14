package api

import (
	"github.com/julienschmidt/httprouter"
	"net/http"

	"github.com/wexel-nath/authrouter"
)

func healthz(_ *http.Request, _ httprouter.Params, _ authrouter.User) (interface{}, interface{}, int) {
	result := struct{
		Status string `json:"status"`
	}{
		Status: "ok",
	}
	return result, nil, http.StatusOK
}
