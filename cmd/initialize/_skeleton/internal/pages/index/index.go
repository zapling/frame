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
		<link rel="icon" type="image/x-icon" href="/assets/favicon.ico">
    </head>
    <body>
        <p>Hello World</p>
    </body>
	<script type="text/javascript">
		const _connectToDevServerUpdates = () => {
			const socket = new WebSocket('ws://localhost:4000/ws');

			socket.onopen = () => {
				if (window._reconnectToDevServerInterval) {
					setTimeout(() => {
						window.location.href = window.location.href;
					}, 500);
				}

				clearInterval(window._reconnectToDevServerInterval);

				socket.onclose = () => {
					window._reconnectToDevServerInterval= setInterval(() => {
						_connectToDevServerUpdates();
					}, 5000);
				}
			}

			socket.onmessage = () => {
				window.location.href = window.location.href;
			}

			socket.onerror = () => {
				socket.close();
			}
		}

		_connectToDevServerUpdates();
	</script>
</html>
`)

func GetPage(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write(page)
}
