package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"{{ .ModuleName }}/internal/pages/index"
)

func getServer() *http.Server {
	r := chi.NewRouter()

	r.Handle("/assets/*", http.StripPrefix("/assets/", getAssetsHandler("assets")))

	r.Get("/", index.GetPage)

	return &http.Server{Addr: ":3000", Handler: r}
}
