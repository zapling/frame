package serve

import (
	_ "embed"
	"fmt"
	"net/http"
	"os"

	"github.com/zapling/frame/pkg/run"
	"golang.org/x/net/websocket"
)

//go:embed refresh_page_ws_client.js
var refreshPageWSClientJS []byte

type developmentServer struct {
	wsConns map[*websocket.Conn]bool
}

func (s *developmentServer) start() {
	mux := http.NewServeMux()
	mux.HandleFunc("/refresh_page_ws_client.js", s.getWebsocketClientHandler)
	mux.Handle("/ws", websocket.Handler(s.websocketHandler))

	server := &http.Server{
		Addr:    ":4000",
		Handler: mux,
	}

	if err := run.Webserver(server); err != nil {
		fmt.Printf("Development server encountered an error: %v", err)
		os.Exit(1)
	}
}

func (s *developmentServer) websocketHandler(ws *websocket.Conn) {
	s.wsConns[ws] = true
	select {} // Block forever to not close the connection
}

func (s *developmentServer) getWebsocketClientHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/javascript")
	_, _ = w.Write(refreshPageWSClientJS)
}

func (s *developmentServer) notifyClients() {
	for ws := range s.wsConns {
		_, err := ws.Write([]byte("1"))
		if err != nil {
			delete(s.wsConns, ws)
		}
	}
}
