package index

import (
	"bytes"
	"fmt"
	"net/http"
)

var page = []byte(`
<!DOCTYPE HTML>
<html lang="en">
    <head>
        <title>Hello World</title>
        <meta charset="UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1">
		<link rel="icon" type="image/x-icon" href="/assets/favicon.ico">
		{#DEV_PAGE_REFRESH_JS}
    </head>
    <body>
        <p>Hello World</p>
    </body>
</html>
`)

func GetPage(isDevMode bool, devServerPort int) http.HandlerFunc {
	pageBytes := page

	var replaceBytes []byte
	if isDevMode {
		replaceBytes = []byte(fmt.Sprintf(
			`<script src="http://localhost:%d/refresh_page_ws_client.js"></script>`,
			devServerPort,
		))
	}

	pageBytes = bytes.ReplaceAll(pageBytes, []byte("{#DEV_PAGE_REFRESH_JS}"), replaceBytes)

	return func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write(pageBytes)
	}
}
