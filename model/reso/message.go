package reso

// GetMessage GET "/message" response object
type GetMessage struct {
	ID         int64  `json:"id"`
	SenderID   int64  `json:"senderId"`
	SenderName string `json:"senderName"`
	Content    string `json:"content"`
	SendTime   int64  `json:"sendTime"`
	Private    bool   `json:"private"`
}

// PostMessage POST "/message" response object
type PostMessage struct {
	ID int64 `json:"id"`
}
