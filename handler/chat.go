package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/feliux/chatai/view/chat"
	"github.com/gorilla/websocket"
)

func HandleChatIndex(w http.ResponseWriter, r *http.Request) error {
	return render(w, r, chat.Index())
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func HandleChat(w http.ResponseWriter, r *http.Request) error {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("error handle chat", err)
		return err
	}
	defer ws.Close()
	for {
		ChatLoop(ws, r.Context())
	}
	return nil
}

type HTMXMessage struct {
	ChatMessage string `json:"chat_message"`
	Headers     struct {
		HXRequest     string `json:"HX-Request"`
		HXTrigger     string `json:"HX-Trigger"`
		HXTriggerName string `json:"HX-Trigger-Name"`
		HXTarget      string `json:"HX-Target"`
		HXCurrentURL  string `json:"HX-Current-URL"`
	} `json:"HEADERS"`
}

func ChatLoop(ws *websocket.Conn, ctx context.Context) {
	_, p, err := ws.ReadMessage()
	if err != nil {
		fmt.Println(err)
	}
	var msg HTMXMessage
	json.Unmarshal(p, &msg)
	var buf bytes.Buffer
	chat.SentAndRecv(msg.ChatMessage, "You're an idiot").Render(ctx, &buf)
	err = ws.WriteMessage(websocket.TextMessage, buf.Bytes())
	if err != nil {
		log.Println("error in ws ", err)
	}
}
