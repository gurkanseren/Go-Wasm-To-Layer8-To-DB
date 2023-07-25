package utils

import (
	"context"
	"log"
	"time"

	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

func UpgradeConnToWebSocket(choice string) string {

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	c, _, err := websocket.Dial(ctx, "ws://localhost:8080/ws", nil)
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close(websocket.StatusInternalError, "error")

	err = wsjson.Write(ctx, c, choice)
	if err != nil {
		log.Fatal(err)
	}

	var v string
	err = wsjson.Read(ctx, c, &v)
	if err != nil {
		log.Fatal(err)
	}

	// log.Printf("Received from layer8-server: %v", v)

	c.Close(websocket.StatusNormalClosure, "")

	return v
}
