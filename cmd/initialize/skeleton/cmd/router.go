package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"{{ .ModuleName }}/internal/pages/index"
)

func startRouter() error {
	r := chi.NewRouter()

	r.Handle("/assets/*", http.StripPrefix("/assets/", getAssetsHandler("assets")))

	r.Get("/", index.GetPage)

	return http.ListenAndServe(":3000", r)
}
