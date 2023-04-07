package main

import (
	"log"

	"iteung/config"
	"iteung/controller"

	"github.com/aiteung/atmessage"
	"github.com/aiteung/musik"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/whatsauth/whatsauth"

	"iteung/url"

	"github.com/gofiber/fiber/v2"
)

func main() {
	go whatsauth.RunHub()
	config.Client = atmessage.RunWA(controller.WAEventHandler)
	site := fiber.New(config.Iteung)
	site.Use(cors.New(config.Cors))
	url.Web(site)
	log.Fatal(site.Listen(musik.Dangdut()))
}
