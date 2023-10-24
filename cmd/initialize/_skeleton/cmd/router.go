package main

import (
	"net/http"

	"{{ .ModuleName }}/internal/pages/index"

	"github.com/go-chi/chi/v5"
	"github.com/zapling/frame/pkg/cfg"
)

func getServer(frameCfg cfg.Values) *http.Server {
	r := chi.NewRouter()

	r.Handle("/assets/*", http.StripPrefix("/assets/", getAssetsHandler("assets")))

	r.Get("/", index.GetPage(frameCfg.IsDevMode, frameCfg.DevServerPort))

	return &http.Server{Addr: ":3000", Handler: r}
}
