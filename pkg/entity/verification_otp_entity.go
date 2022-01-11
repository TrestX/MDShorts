package entity

type MessageStructure struct {
	To     string `json:"to"`
	Body   string `json:"body"`
	Source string `json:"source"`
}
type MessageBody struct {
	Messages []MessageStructure `json:"messages"`
}
type ClickSendResponseCode struct {
	ResponseCode string `json:"response_code"`
}
