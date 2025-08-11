package ws_model

type OnlineStatus struct {
	Type        string `json:"type"`
	From        int64  `json:"from"`
	OnlineState bool   `json:"online_state"`
}
