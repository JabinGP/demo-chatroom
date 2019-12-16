package reqo

// GetUser GET "/user" request object
type GetUser struct {
	Username string `json:"username"`
	ID       uint   `json:"id"`
}

// PutUser PUT "/user" request object
type PutUser struct {
	Gender   int64  `json:"gender"`
	Age      int64  `json:"age"`
	Interest string `json:"interest"`
}

// PostUser POST "/user" request object
type PostUser struct {
	Username string `json:"username"`
	Passwd   string `json:"passwd"`
	Gender   int64  `json:"gender"`
	Age      int64  `json:"age"`
	Interest string `json:"interest"`
}

// PostLogin POST "/login" request object
type PostLogin struct {
	Username string `json:"username"`
	Passwd   string `json:"passwd"`
}
