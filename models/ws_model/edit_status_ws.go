package ws_model

type EditStatus struct {
	Type    string `json:"type"`
	To      int64  `json:"to"`
	Content string `json:"content"` // 内容
}
