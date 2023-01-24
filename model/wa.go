package model

type NotifButton struct {
	User    string         `json:"user"`
	Server  string         `json:"server"`
	Message ButtonsMessage `json:"button_messages"`
}

type WaButton struct {
	ButtonId    string `json:"button_id,omitempty"`
	DisplayText string `json:"display_text,omitempty"`
}

type WaButtonsMessage struct {
	HeaderText  string `json:"header_text,omitempty"`
	ContentText string `json:"content_text,omitempty"`
	FooterText  string `json:"footer_text,omitempty"`
}

type ButtonsMessage struct {
	Message WaButtonsMessage `json:"message,omitempty"`
	Buttons []WaButton       `json:"buttons,omitempty"`
}

type ListMessage struct {
	Title       string
	Description string
	FooterText  string
	ButtonText  string
	Sections    []WaListSection
}

type WaListSection struct {
	Title string
	Rows  []WaListRow
}

type WaListRow struct {
	Title       string
	Description string
	RowId       string
}

type PhoneList struct {
	Phones []string `json:"phones,omitempty"`
}

type Response struct {
	Response string `json:"response"`
}
