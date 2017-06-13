package oauth

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/FoolishOldMan/wechat"
	"github.com/FoolishOldMan/wechat/util"
)

var (
	redirectOauthURL      = "https://open.weixin.qq.com/connect/oauth2/authorize?appid=%s&redirect_uri=%s&response_type=code&scope=%s&state=%s#wechat_redirect"
	accessTokenURL        = "https://api.weixin.qq.com/sns/oauth2/access_token?appid=%s&secret=%s&code=%s&grant_type=authorization_code"
	refreshAccessTokenURL = "https://api.weixin.qq.com/sns/oauth2/refresh_token?appid=%s&grant_type=refresh_token&refresh_token=%s"
	checkAccessTokenURL   = "https://api.weixin.qq.com/sns/auth?access_token=%s&openid=%s"
	userInfoURL           = "https://api.weixin.qq.com/sns/userinfo?access_token=%s&openid=%s&lang=zh_CN"
)

const (
	// OAuthBaseScope 静默授权
	OAuthBaseScope = "snsapi_base"
	// OAuthUserInfoScope 基本信息授权
	OAuthUserInfoScope = "snsapi_userinfo"
)

// UserAccessToken 用户授权access_token的返回结果
type UserAccessToken struct {
	util.CommonError
	AccessToken  string `json:"access_token"`
	ExpiresIn    int64  `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	OpenID       string `json:"openid"`
	Scope        string `json:"scope"`
}

// UserInfo 用户授权获取到用户信息
type UserInfo struct {
	util.CommonError
	OpenID     string   `json:"openid"`
	Nickname   string   `json:"nickname"`
	Sex        int32    `json:"sex"`
	Province   string   `json:"province"`
	City       string   `json:"city"`
	Country    string   `json:"country"`
	HeadImgURL string   `json:"headimgurl"`
	Privilege  []string `json:"privilege"`
	Unionid    string   `json:"unionid"`
}

// GetRedirectURL 获取跳转的url地址
func GetRedirectURL(redirectURI, scope, state string) string {
	urlStr := url.QueryEscape(redirectURI)
	return fmt.Sprintf(redirectOauthURL, wechat.AppID, urlStr, scope, state)
}

// GetRedirectBaseURL 获取跳转的url地址，无用户信息
func GetRedirectBaseURL(redirectURI, state string) string {
	urlStr := url.QueryEscape(redirectURI)
	return fmt.Sprintf(redirectOauthURL, wechat.AppID, urlStr, OAuthBaseScope, state)
}

// GetRedirectUserInfoURL 获取跳转的url地址，用户信息
func GetRedirectUserInfoURL(redirectURI, state string) string {
	urlStr := url.QueryEscape(redirectURI)
	return fmt.Sprintf(redirectOauthURL, wechat.AppID, urlStr, OAuthUserInfoScope, state)
}

// GetUserAccessToken 通过网页授权的code 换取access_token(区别于context中的access_token)
func GetUserAccessToken(code string) (*UserAccessToken, error) {
	urlStr := fmt.Sprintf(accessTokenURL, wechat.AppID, wechat.AppSecret, code)
	response, err := http.Get(urlStr)
	if err != nil {
		return nil, fmt.Errorf("GetUserAccessToken error : 请求accessTokenURL异常")
	}
	result := UserAccessToken{}
	err = util.Response2Struct(response, &result)
	if err != nil {
		return nil, fmt.Errorf("GetUserAccessToken error : 解析userAccessToken错误 , %v", err.Error())
	}
	if result.ErrCode != 0 {
		return nil, fmt.Errorf("GetUserAccessToken error : errcode = %v , errmsg = %v", result.ErrCode, result.ErrMsg)
	}
	return &result, nil
}

// RefreshUserAccessToken 刷新access_token
func RefreshUserAccessToken(refreshToken string) (*UserAccessToken, error) {
	urlStr := fmt.Sprintf(refreshAccessTokenURL, wechat.AppID, refreshToken)
	response, err := http.Get(urlStr)
	if err != nil {
		return nil, fmt.Errorf("RefreshUserAccessToken error : 请求refreshAccessTokenURL异常")
	}
	result := UserAccessToken{}
	err = util.Response2Struct(response, &result)
	if err != nil {
		return nil, fmt.Errorf("RefreshUserAccessToken error : 解析userAccessToken错误 , %v", err.Error())
	}
	if result.ErrCode != 0 {
		return nil, fmt.Errorf("RefreshUserAccessToken error : errcode = %v , errmsg = %v", result.ErrCode, result.ErrMsg)
	}
	return &result, nil
}

// CheckUserAccessToken 检验access_token是否有效
func CheckUserAccessToken(accessToken, openID string) (bool, error) {
	urlStr := fmt.Sprintf(checkAccessTokenURL, accessToken, openID)
	response, err := http.Get(urlStr)
	if err != nil {
		return false, fmt.Errorf("CheckUserAccessToken error : 请求checkAccessTokenURL异常")
	}
	result := util.CommonError{}
	err = util.Response2Struct(response, &result)
	if err != nil {
		return false, fmt.Errorf("CheckUserAccessToken error : 解析commonError错误 , %v", err.Error())
	}
	if result.ErrCode != 0 {
		return false, fmt.Errorf("CheckUserAccessToken error : errcode = %v , errmsg = %v", result.ErrCode, result.ErrMsg)
	}
	return true, nil
}

// GetUserInfo 如果scope为 snsapi_userinfo 则可以通过此方法获取到用户基本信息
func GetUserInfo(userAccessToken UserAccessToken) (*UserInfo, error) {
	if userAccessToken.Scope != OAuthUserInfoScope {
		return nil, fmt.Errorf("GetUserInfo error : scope不为snsapi_userinfo , 不能获取用户信息 , 请重新以snsapi_userinfo授权登录")
	}
	urlStr := fmt.Sprintf(userInfoURL, userAccessToken.AccessToken, userAccessToken.OpenID)
	response, err := http.Get(urlStr)
	if err != nil {
		return nil, fmt.Errorf("GetUserInfo error : 请求userInfoURL异常")
	}
	result := UserInfo{}
	err = util.Response2Struct(response, &result)
	if err != nil {
		return nil, fmt.Errorf("GetUserInfo error : 解析userInfo错误 , %v", err.Error())
	}
	if result.ErrCode != 0 {
		return nil, fmt.Errorf("GetUserInfo error : errcode = %v , errmsg = %v", result.ErrCode, result.ErrMsg)
	}
	return &result, nil
}
