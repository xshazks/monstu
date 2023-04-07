package controller

import (
	"fmt"
	"time"

	"iteung/config"

	"github.com/aiteung/atmessage"
	_ "github.com/mattn/go-sqlite3"
	waProto "go.mau.fi/whatsmeow/binary/proto"
	"go.mau.fi/whatsmeow/types"
	"go.mau.fi/whatsmeow/types/events"
)

func HandlingMessage(Info *types.MessageInfo, Message *waProto.Message) {
	go config.Client.MarkRead([]string{Info.ID}, time.Now(), Info.Chat, Info.Sender)
	go config.Client.SendChatPresence(Info.Chat, "composing", "")
	if !Info.IsFromMe {
		duration := time.Duration(10) * time.Second
		time.Sleep(duration)
		atmessage.SendMessage(Message.GetConversation(), Info.Chat, config.Client)
	}
}

func HandlingReceipt(Info *events.Receipt) {
	fmt.Println("Receipt", Info)
	fmt.Println(Info.MessageIDs)
	fmt.Println(Info.Type)
}

func WAEventHandler(evt interface{}) {
	switch v := evt.(type) {
	case *events.Message:
		go HandlingMessage(&v.Info, v.Message)
	case *events.Receipt:
		go HandlingReceipt(v)
	}
}
