package wechat

var (
	// AppID appid
	AppID string
	// AppSecret appsecret
	AppSecret string
)

// ConfigWechat user configWechat
func ConfigWechat(appID, appSecret string) {
	AppID = appID
	AppSecret = appSecret
}
