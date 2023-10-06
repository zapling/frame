package index

import (
	"net/http"
)

var page = []byte(`
<!DOCTYPE HTML>
<html lang="en">
    <head>
        <title>Hello World</title>
        <meta charset="UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1">
    </head>
    <body>
        <p>Hello World</p>
    </body>
</html>
`)

func GetPage(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write(page)
}
