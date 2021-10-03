package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	r := httprouter.New()

	r.HandlerFunc(http.MethodGet, "/status", app.statusHandler)

	r.HandlerFunc(http.MethodPost, "/v1/signin", app.signupHandler)

	return r
}
