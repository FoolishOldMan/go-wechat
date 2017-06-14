# go-wechat

go语言版微信SDK-开发阶段

## 快速开始

拉取代码

`go get github.com/FoolishOldMan/go-wechat`

基本配置

`wechat.ConfigWechat("appID", "appSecret")`

## API

- [全局参数](#全局参数)
    - AccessToken
		- 获取AccessToken
    - JSTicket
		- 获取JSTicket
- [微信网页开发](#微信网页开发)
    - Oauth2 授权
		- 获取授权地址
		- 通过code换取access_token
		- 拉取用户信息
		- 刷新access_token
		- 检验access_token

## 全局参数

### AccessToken

**1.获取AccessToken**

AccessToken的有效期目前为2个小时，需定时刷新，重复获取将导致上次获取的access_token失效。

`GetAccessToken() 获取AccessToken`

### JSTicket

**1.获取JSTicket**

由于获取JSTicket的api调用次数非常有限，频繁刷新JSTicket会导致api调用受限，影响自身业务，开发者必须在自己的服务全局缓存JSTicket。

`GetJSTicket(accessToken string) 获取JSTicket`

## 微信网页开发

### Oauth2 授权

**1.获取授权地址**

`GetRedirectURL(redirectURI, scope, state string) 获取跳转的url地址`

快速方式

`GetRedirectBaseURL(redirectURI, state string) 获取跳转的url地址，无法获取用户信息`

`GetRedirectUserInfoURL(redirectURI, state string) 获取跳转的url地址，可用户信息`

**2.通过code换取access_token**

`GetUserAccessToken(code string) 通过网页授权的code 换取access_token`

**3.拉取用户信息**

`GetUserInfo(userAccessToken UserAccessToken) 通过userAccessToken换取用户信息，如果scope为 snsapi_userinfo 则可以通过此方法获取到用户基本信`

**4.刷新access_token**

`RefreshUserAccessToken(refreshToken string) 刷新access_token`

**5.检验access_token**

`CheckUserAccessToken(accessToken, openID string) 检验access_token是否有效`
