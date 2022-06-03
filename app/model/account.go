package model

// AccountDecryptReq 微信小程序消息解密
type AccountDecryptReq struct {
	Code          string `json:"code,omitempty"`
	EncryptedData string `json:"encrypted_data,omitempty"`
	IvData        string `json:"iv_data,omitempty"`
}

// AccountPlainDataResp 小程序解密数据
type AccountPlainDataResp struct {
	OpenID          string `json:"openId"`
	UnionID         string `json:"unionId"`
	NickName        string `json:"nickName"`
	Gender          int    `json:"gender"`
	City            string `json:"city"`
	Province        string `json:"province"`
	Country         string `json:"country"`
	AvatarURL       string `json:"avatarUrl"`
	Language        string `json:"language"`
	PhoneNumber     string `json:"phoneNumber"`
	PurePhoneNumber string `json:"purePhoneNumber"`
	CountryCode     string `json:"countryCode"`
	Watermark       struct {
		Timestamp int64  `json:"timestamp"`
		AppID     string `json:"appid"`
	} `json:"watermark"`
}
