package wamessage

import (
	"boilerplate/config"
	"boilerplate/helper/wahttp"
	"boilerplate/model"
)

func SendButtonMessage(btnmsg model.NotifButton) (response interface{}) {
	return wahttp.PostStructtoAPI(btnmsg, config.ApiWaButton)
}
