package app

import (
	"net/http"

	"github.com/a-h/templ"
)

func Handler(conf Config, body templ.Component) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := Component(conf, body).Render(r.Context(), w); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}
