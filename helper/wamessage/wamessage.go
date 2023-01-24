package wamessage

import (
	"iteung/config"
	"iteung/helper/wahttp"
	"iteung/model"
)

func SendButtonMessage(btnmsg model.NotifButton) (response interface{}) {
	return wahttp.PostStructtoAPI(btnmsg, config.ApiWaButton)
}
