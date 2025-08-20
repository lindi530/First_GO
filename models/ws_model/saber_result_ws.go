package ws_model

type SaberResult struct {
	Type string `json:"type"`
	To   int64  `json:"to"`
	Win  bool   `json:"win"`
}
