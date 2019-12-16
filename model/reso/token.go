package reso

// GetTokenInfo GET "/token/info" response object
type GetTokenInfo struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
}
