package simpati

import (
	"fmt"

	"boilerplate/helper/wacipher"
	"boilerplate/model"

	"github.com/whatsauth/whatsauth"
)

func RunModule(req model.WhatsauthRequest) model.NotifButton {
	var content string
	delay := req.Delay
	if GetUsernamefromPhonenumber(req.Phonenumber) != "" {
		infologin := GetLoginInfofromPhoneNumber(req.Phonenumber)
		infologin.Uuid = req.Uuid
		fmt.Println(infologin)
		status := whatsauth.SendStructTo(req.Uuid, infologin)
		if status {
			content = fmt.Sprintf("Hai kak , login aplikasi sukses, silahkan kakak melihat kembali ke aplikasi. Waktu scan %v detik.", delay)
		} else {
			content = fmt.Sprintf("Maaf kak login gagal. Kemungkinan qr code tidak valid atau qr code nya sudah expire kak. Silahkan scan ulang kembali ya kak. Atau kakak terlalu lama mengirim kodenya, kakak butuh waktu %v detik untuk mengirim kode authentikasi ini. Semoga selanjutnya bisa lebih cekatan ya kak. Semangat kak.", delay)
		}
	} else {
		content = fmt.Sprintf("Hai kak , Nomor whatsapp ini *tidak terdaftar* di sistem kami, silahkan silahkan gunakan nomor yang terdftar ya kak. Waktu scan %v detik.", delay)
	}
	header := "WhatsAuth Single Sign On"
	footer := fmt.Sprintf("Login Aplikasi : %v", wacipher.GetAppHost(req.Uuid))
	btm := GenerateButtonMessage(header, content, footer)
	var notifbtn model.NotifButton
	notifbtn.User = req.Phonenumber
	notifbtn.Server = "s.whatsapp.net"
	notifbtn.Message = btm
	return notifbtn
}

func GenerateButtonMessage(header string, content string, footer string) (btnmsg model.ButtonsMessage) {
	btnmsg.Message.HeaderText = header
	btnmsg.Message.ContentText = content
	btnmsg.Message.FooterText = footer
	btnmsg.Buttons = []model.WaButton{{
		ButtonId:    "whatsauth|challange1",
		DisplayText: "Sama Sama",
	},
		{
			ButtonId:    "whatsauth|challange3",
			DisplayText: "Sawangsulna",
		},
		{
			ButtonId:    "whatsauth|challange2",
			DisplayText: "Mangga",
		},
	}
	return btnmsg
}
