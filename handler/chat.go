package handler

import (
	"bytes"
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/feliux/chatai/view/chat"
	"github.com/gorilla/websocket"
)

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

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func HandleChatIndex(w http.ResponseWriter, r *http.Request) error {
	return render(w, r, chat.Index())
}

func HandleChat(w http.ResponseWriter, r *http.Request) error {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		slog.Error("error upgrading to websocket", "err", err)
	}
	defer ws.Close()
	for {
		Chat(w, r, ws)
	}
	return nil
}

func Chat(w http.ResponseWriter, r *http.Request, ws *websocket.Conn) {
	_, p, err := ws.ReadMessage()
	if err != nil {
		slog.Error("error reading message from websocket", "err", err)
	}
	var msg HTMXMessage
	json.Unmarshal(p, &msg)
	var buf bytes.Buffer
	// err = render(w, r, chat.SentAndRecv(msg.ChatMessage, "This is a message from the bot!"))
	// if err != nil {
	// 	slog.Error("error rendering SentAndRecv template", "err", err)
	// }
	chat.SentAndRecv(msg.ChatMessage, "This is a message from the bot!").Render(r.Context(), &buf)
	err = ws.WriteMessage(websocket.TextMessage, buf.Bytes())
	if err != nil {
		slog.Error("error writing message to websocket", "err", err)
	}
}
