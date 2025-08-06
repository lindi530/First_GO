package ws_model

var (
	MessageTypeMatchSuccess = "match_success"
	MessageTypeEditStatus   = "edit_status"
	MessageTypeChat         = "chat"
	MessageTypeError        = "error"
)

type MessageWs struct {
	Type string      `json:"type"`
	To   int64       `json:"to"`
	Data interface{} `json:"data"`
}
