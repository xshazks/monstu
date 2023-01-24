package model

type LoginInfo struct {
	Userid   string `json:"user_id,omitempty"`
	Username string `json:"user_name,omitempty"`
	Password string `json:"user_pass,omitempty"`
	Login    string `json:"login,omitempty"`
	Uuid     string `json:"uuid,omitempty"`
}

type Simpati struct {
	Userid   string `json:"user_id,omitempty"`
	Username string `json:"user_name,omitempty"`
	Password string `json:"user_pass,omitempty"`
	Login    string `json:"login,omitempty"`
	Uuid     string `json:"uuid,omitempty"`
}

type WhatsauthRequest struct {
	Uuid        string `json:"uuid,omitempty"`
	Phonenumber string `json:"phonenumber,omitempty"`
	Delay       uint32 `json:"delay,omitempty"`
}

type WhatsauthMessage struct {
	Id      string  `json:"id,omitempty"`
	Message Simpati `json:"message,omitempty"`
}

type WhatsauthStatus struct {
	Status string `json:"status,omitempty"`
}
