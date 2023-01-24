package url

import (
	"boilerplate/controller"

	"github.com/gin-gonic/gin"
)

func Web(page *gin.Engine) {
	page.POST("/api/whatsauth/message", controller.PostWhatsAuthMessage)
	page.POST("/api/whatsauth/request", controller.PostWhatsAuthRequest)
	page.GET("/ws/whatsauth/qr", controller.WsWhatsAuthQR)

}
