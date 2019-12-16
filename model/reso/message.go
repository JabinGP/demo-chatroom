package reso

// GetMessage GET "/message" response object
type GetMessage struct {
	ID         uint   `json:"id"`
	SenderID   uint   `json:"senderId"`
	SenderName string `json:"senderName"`
	Content    string `json:"content"`
	SendTime   int64  `json:"sendTime"`
	Private    bool   `json:"private"`
}

// PostMessage POST "/message" response object
type PostMessage struct {
	ID uint `json:"id"`
}
