package wechat

var (
	// AppID appid
	AppID string
	// AppSecret appsecret
	AppSecret string
)

// ConfigWechat 微信基本配置
func ConfigWechat(appID, appSecret string) {
	AppID = appID
	AppSecret = appSecret
}
