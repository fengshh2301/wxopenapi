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

// <AppId>第三方平台appid</AppId>
// <CreateTime>1413192760</CreateTime>
// <InfoType>updateauthorized</InfoType>
// <AuthorizerAppid>公众号appid</AuthorizerAppid>
// <AuthorizationCode>授权码（code）</AuthorizationCode>
// <AuthorizationCodeExpiredTime>过期时间</AuthorizationCodeExpiredTime>
// <PreAuthCode>预授权码</PreAuthCode>
