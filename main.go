package main

import (
	"boilerplate/config"
	"boilerplate/url"

	"github.com/gin-gonic/gin"
	"github.com/whatsauth/whatsauth"
)

func main() {
	go whatsauth.HubRun(&whatsauth.Hub)
	site := gin.New()
	site.SetTrustedProxies(nil)
	site.Use(config.Cors())
	url.Web(site)
	site.Run()
}
