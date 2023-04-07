package controller

import (
	"fmt"
	"time"

	"iteung/config"

	"github.com/aiteung/atmessage"
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
		//sendButtonMessage(Info.Chat)
		//atmessage.SendMessage(Message.GetConversation(), Info.Chat, config.Client)
		var btnmsg = atmessage.WaButton{
			ButtonId:    "idbutton",
			DisplayText: "ok",
		}
		var btnmsg2 = atmessage.WaButton{
			ButtonId:    "idbutton2",
			DisplayText: "no",
		}
		var btn = atmessage.ButtonsMessage{
			Message: atmessage.WaButtonsMessage{
				HeaderText:  "judul",
				ContentText: "isi",
				FooterText:  "kaki",
			},
			Buttons: []atmessage.WaButton{btnmsg, btnmsg2},
		}
		atmessage.SendButtonMessage(btn, Info.Chat, config.Client)
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
