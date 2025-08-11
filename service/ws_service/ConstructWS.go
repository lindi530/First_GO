package ws_service

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

func ConstructWS(c *gin.Context, userid int64) {

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}

	client := &Client{
		ID:   userid,
		Conn: conn,
		Send: make(chan []byte, 256),
	}

	start(userid, client)
}

func start(userid int64, client *Client) {
	WsHub.RegisterClient(client)
	WsHub.SendOnlineData(userid, true)
	client.WriteOfflineMsg(WsHub, userid)

	go client.WriteLoop()
	go client.ReadLoop(WsHub)
}
