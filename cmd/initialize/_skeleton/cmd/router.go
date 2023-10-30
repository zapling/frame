package main

import (
	"net/http"

	"{{ .ModuleName }}/internal/components/app"
	"{{ .ModuleName }}/internal/pages/index"

	"github.com/go-chi/chi/v5"
	"github.com/zapling/frame/pkg/cfg"
)

func getServer(frameCfg cfg.Values) *http.Server {
	conf := app.Config{
		IsRunningInDevMode: frameCfg.IsDevMode,
		DevServerPort:      frameCfg.DevServerPort,
	}

	r := chi.NewRouter()

	r.Handle("/assets/*", http.StripPrefix("/assets/", getAssetsHandler("assets")))

	r.Get("/", app.Handler(conf, index.Component()))

	return &http.Server{Addr: ":3000", Handler: r}
}
