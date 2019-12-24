package reqo

// GetMessage GET "/message" request object
type GetMessage struct {
	BeginID   int64 `json:"beginId"`
	BeginTime int64 `json:"beginTime"`
	EndTime   int64 `json:"endTime"`
}

// PostMessage POST "/message" request object
type PostMessage struct {
	ReceiverID int64  `json:"receiverId"`
	Content    string `json:"content"`
}
