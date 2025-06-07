package ws

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func ConstructWS(c *gin.Context, userId int64) {

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}

	client := &Client{
		ID:   userId,
		Conn: conn,
		Send: make(chan []byte, 256),
	}

	WsHub.RegisterClient(client)

	go client.WriteLoop()
	go client.ReadLoop(WsHub)
}
