package utils

import (
	"context"
	"log"
	"time"

	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

func UpgradeConnToWebSocket() {
	go func() {
		ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
		defer cancel()

		c, _, err := websocket.Dial(ctx, "ws://localhost:8080/ws", nil)
		if err != nil {
			log.Fatal(err)
		}
		defer c.Close(websocket.StatusInternalError, "the sky is falling")

		err = wsjson.Write(ctx, c, "hi")
		if err != nil {
			log.Fatal(err)
		}

		c.Close(websocket.StatusNormalClosure, "")
	}()
}
