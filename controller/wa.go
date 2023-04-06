package controller

import (
	"context"
	"fmt"
	"time"

	"iteung/config"

	"github.com/aiteung/atmessage"
	"github.com/aiteung/musik"
	_ "github.com/mattn/go-sqlite3"
	"go.mau.fi/whatsmeow"
	waProto "go.mau.fi/whatsmeow/binary/proto"
	"go.mau.fi/whatsmeow/store/sqlstore"
	"go.mau.fi/whatsmeow/types"
	"go.mau.fi/whatsmeow/types/events"
	waLog "go.mau.fi/whatsmeow/util/log"
)

func HandlingMessage(Info *types.MessageInfo, Message *waProto.Message) {
	go config.Client.MarkRead([]string{Info.ID}, time.Now(), Info.Chat, Info.Sender)
	if !Info.IsFromMe {
		atmessage.SendMessage("hai", Info.Sender, config.Client)
	}
}

func WAEventHandler(evt interface{}) {
	switch v := evt.(type) {
	case *events.Message:
		go HandlingMessage(&v.Info, v.Message)
	case *events.Receipt:
		fmt.Println(v)
	}
}

func RunWA() {
	fmt.Println("Starting Whatsapp")
	dbLog := waLog.Stdout("Database", "ERROR", true)
	musik.CreateFolderifNotExist("./session/")
	container, err := sqlstore.New("sqlite3", "file:./session/gowa.db?_foreign_keys=on", dbLog)
	if err != nil {
		panic(err)
	}
	deviceStore, err := container.GetFirstDevice()
	if err != nil {
		panic(err)
	}
	clientLog := waLog.Stdout("Client", "ERROR", true)
	config.Client = whatsmeow.NewClient(deviceStore, clientLog)
	config.Client.AddEventHandler(WAEventHandler)
	PairWA(config.Client)

}

func PairWA(wc *whatsmeow.Client) {
	var err error
	if wc.Store.ID == nil {
		// No ID stored, new login
		qrChan, _ := wc.GetQRChannel(context.Background())
		err = wc.Connect()
		if err != nil {
			panic(err)
		}
		for evt := range qrChan {
			if evt.Event == "code" {
				fmt.Println("QR code:", evt.Code)
			} else {
				fmt.Println("Login event:", evt.Event)
			}
		}
	} else {
		// Already logged in, just connect
		err = wc.Connect()
		if err != nil {
			panic(err)
		}
		fmt.Println("Client Connected")
	}
}
