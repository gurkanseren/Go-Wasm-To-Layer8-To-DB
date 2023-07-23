package utils

import (
	"context"
	"log"
	"time"

	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

func UpgradeConnToWebSocket() string {

	data := "Connection Test Message"

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	c, _, err := websocket.Dial(ctx, "ws://localhost:8080/ws", nil)
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close(websocket.StatusInternalError, "error")

	err = wsjson.Write(ctx, c, data)
	if err != nil {
		log.Fatal(err)
	}

	c.Close(websocket.StatusNormalClosure, "")

	return data
}
