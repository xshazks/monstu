package main

import (
	"log"

	"github.com/aiteung/musik"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gitlab.com/informatics-research-center/auth-service/config"

	"github.com/whatsauth/whatsauth"

	"github.com/gofiber/fiber/v2"
	"gitlab.com/informatics-research-center/auth-service/url"
)

func main() {
	go whatsauth.RunHub()
	site := fiber.New()
	site.Use(cors.New(config.Cors))
	url.Web(site)
	log.Fatal(site.Listen(musik.Dangdut()))
}
