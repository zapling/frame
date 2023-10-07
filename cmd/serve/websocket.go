package serve

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/zapling/frame/pkg/run"
	"golang.org/x/net/websocket"
)

type wsServer struct {
	conns map[*websocket.Conn]bool
}

func (s *wsServer) start() {
	mux := http.NewServeMux()
	mux.Handle("/ws", websocket.Handler(func(ws *websocket.Conn) {
		s.conns[ws] = true

		for {
			time.Sleep(1 * time.Hour)
		}
	}))
	server := &http.Server{
		Addr:    ":4000",
		Handler: mux,
	}

	if err := run.Webserver(server); err != nil {
		fmt.Printf("Error Websocket server: %v", err)
		os.Exit(1)
	}
}

func (s *wsServer) notifyClients() {
	for ws := range s.conns {
		_, err := ws.Write([]byte("1"))
		if err != nil {
			delete(s.conns, ws)
		}
	}
}
