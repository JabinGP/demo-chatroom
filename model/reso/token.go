package reso

// GetTokenInfo GET "/token/info" response object
type GetTokenInfo struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
}
