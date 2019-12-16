package model

// ResModel response model
type ResModel struct {
	Code string      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// WithData set model success and data
func (res *ResModel) WithData(data interface{}) *ResModel {
	res.Code = "200"
	res.Msg = "success"
	res.Data = data
	return res
}

// WithError set error message
func (res *ResModel) WithError(errCode string) *ResModel {
	res.Code = errCode
	res.Msg = errCode
	return res
}
