package model

type PaymentConfig struct {
	Default string     `json:"default"`
	Weixin  WxPayInfo  `json:"weixin"`
	Alipay  AliPayInfo `json:"alipay"`
}

type WxPayInfo struct {
	NotifyUrl    string `json:"notify_url"`
	QinAppId     string `json:"qin_app_id"`
	QinAppSecret string `json:"qin_app_secret"`
	QinMchId     string `json:"qin_mch_id"`
	QinMd5Key    string `json:"qin_md5_key"`
}

type AliPayInfo struct {
	AppId     string `json:"app_id"`
	AppSecret string `json:"app_secret"`
	MchId     string `json:"mch_id"`
	Md5key    string `json:"md_5_key"`
}
