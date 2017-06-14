package wechat

import (
	"fmt"
	"net/http"

	"github.com/FoolishOldMan/wechat/util"
)

var (
	// AppID appid
	AppID string
	// AppSecret appsecret
	AppSecret string
)

const (
	accessTokenURL = "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s"
	ticketURL      = "https://api.weixin.qq.com/cgi-bin/ticket/getticket?access_token=%s&type=jsapi"
)

// AccessToken struct
type AccessToken struct {
	util.CommonError
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
}

// JSTicket 请求js_ticket返回结果
type JSTicket struct {
	util.CommonError
	Ticket    string `json:"ticket"`
	ExpiresIn int64  `json:"expires_in"`
}

// ConfigWechat 微信基本配置
func ConfigWechat(appID, appSecret string) {
	AppID = appID
	AppSecret = appSecret
}

// GetAccessToken 获取access_token
func GetAccessToken() (*AccessToken, error) {
	urlStr := fmt.Sprintf(accessTokenURL, AppID, AppSecret)
	response, err := http.Get(urlStr)
	if err != nil {
		return nil, fmt.Errorf("GetAccessToken error : 请求accessTokenURL异常")
	}
	result := AccessToken{}
	err = util.Response2Struct(response, &result)
	if err != nil {
		return nil, fmt.Errorf("GetAccessToken error : 解析accessToken错误 , %v", err.Error())
	}
	if result.ErrCode != 0 {
		return nil, fmt.Errorf("GetAccessToken error : errcode = %v , errmsg = %v", result.ErrCode, result.ErrMsg)
	}
	return &result, nil
}

// GetJSTicket 获取js_ticket
func GetJSTicket(accessToken string) (*JSTicket, error) {
	urlStr := fmt.Sprintf(ticketURL, accessToken)
	response, err := http.Get(urlStr)
	if err != nil {
		return nil, fmt.Errorf("GetJSTicket error : 请求ticketURL异常")
	}
	result := JSTicket{}
	err = util.Response2Struct(response, &result)
	if err != nil {
		return nil, fmt.Errorf("GetJSTicket error : 解析js_ticket错误 , %v", err.Error())
	}
	if result.ErrCode != 0 {
		return nil, fmt.Errorf("GetJSTicket error : errcode = %v , errmsg = %v", result.ErrCode, result.ErrMsg)
	}
	return &result, nil
}
