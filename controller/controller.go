package controller

import (
	"encoding/json"
	"fmt"
	"strconv"

	"iteung/config"

	"github.com/aiteung/atdb"
	"github.com/aiteung/simpati"

	"github.com/gin-gonic/gin"
	"github.com/whatsauth/whatsauth"
)

func WsWhatsAuthQR(c *gin.Context) {
	whatsauth.ServeWs(c.Writer, c.Request)
}

func PostWhatsAuthMessage(c *gin.Context) {
	var ws whatsauth.WhatsauthStatus
	if c.Request.Host == config.Internalhost {
		var m whatsauth.WhatsauthMessage
		c.BindJSON(&m)
		msg := m.Message
		b, err := json.Marshal(msg)
		if err != nil {
			fmt.Printf("Error: %s", err)
			return
		}
		ws.Status = strconv.FormatBool(whatsauth.SendMessageTo(m.Id, string(b)))
	} else {
		ws.Status = c.Request.Host
	}
	c.JSON(200, ws)

}

func PostWhatsAuthRequest(c *gin.Context) {
	if c.Request.Host == config.Internalhost {
		var req whatsauth.WhatsauthRequest
		c.BindJSON(&req)
		mariaconn := atdb.MariaConnect(config.Ulbimariaconn)
		ntfbtn := simpati.RunModule(req, config.Usertables[:], mariaconn)
		c.JSON(200, ntfbtn)
	}
}
