package wxopenapi

type NotifyAuth struct {
	AppId                        string
	CreateTime                   string
	InfoType                     string
	ComponentVerifyTicket        string
	AuthorizerAppid              string
	AuthorizationCode            string
	AuthorizationCodeExpiredTime string
	PreAuthCode                  string
}

type OpenToken struct {
	Typ       string
	CreatedAt int64
	ExpiredAt int64
	ExpiredIn int
	Info      string
}

// <AppId>第三方平台appid</AppId>
// <CreateTime>1413192760</CreateTime>
// <InfoType>updateauthorized</InfoType>
// <AuthorizerAppid>公众号appid</AuthorizerAppid>
// <AuthorizationCode>授权码（code）</AuthorizationCode>
// <AuthorizationCodeExpiredTime>过期时间</AuthorizationCodeExpiredTime>
// <PreAuthCode>预授权码</PreAuthCode>

type ReqAccessToken struct {
	ComponentAppid        string `json:"component_appid"`
	ComponentAppsecret    string `json:"component_appsecret"`
	ComponentVerifyTicket string `json:"component_verify_ticket"`
}

type RspAccessToken struct {
	ComponentAccessToken string `json:"component_access_token"`
	ExpiresIn            int    `json:"expires_in"`
}

type ReqPreAuthCode struct {
	ComponentAppid string `json:"component_appid"`
}

type RspPreAuthCode struct {
	PreAuthCode string `json:"pre_auth_code"`
	ExpiresIn   int    `json:"expires_in"`
}
