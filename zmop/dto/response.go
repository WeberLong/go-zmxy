package dto

type Response struct {
	Encrypted bool   `json:"encrypted"`
	Sign      string `json:"biz_response_sign"`
	Data      string `json:"biz_response"`
}

type ResponseScore struct {
	Success bool   `json:"success"`
	BizNo   string `json:"biz_no"`
	Score   string `json:"zm_score"`
}
