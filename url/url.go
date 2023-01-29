package url

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"gitlab.com/informatics-research-center/auth-service/controller"
)

func Web(page *fiber.App) {
	page.Post("/api/whatsauth/request", controller.PostWhatsAuthRequest)  //API from user whatsapp message from iteung gowa
	page.Get("/ws/whatsauth/qr", websocket.New(controller.WsWhatsAuthQR)) //websocket whatsauth

}
