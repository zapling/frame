package main

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func getAssetsHandler(directory string) http.HandlerFunc {
	workDir, _ := os.Getwd()
	assetsDir := http.Dir(filepath.Join(workDir, directory))
	fs := http.FileServer(assetsDir)

	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path

		if path == "" || strings.HasSuffix(r.URL.Path, "/") {
			http.NotFound(w, r)
			return
		}

		fs.ServeHTTP(w, r)
	}
}
