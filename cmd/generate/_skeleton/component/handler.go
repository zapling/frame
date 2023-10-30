package component

import "net/http"

func Handler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := Component().Render(r.Context(), w); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}
