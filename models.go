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

type ReqAuthAccessToken struct {
	ComponentAppid    string `json:"component_appid"`
	AuthorizationCode string `json:"authorization_code"`
}
type FuncscopeCategory struct {
	Id int `json:"id"`
}
type FuncInfo struct {
	FuncCat FuncscopeCategory `json:"funcscope_category"`
}
type AuthAccessTokenInfo struct {
	AuthorizerAppid        string     `json:"authorizer_appid"`
	AuthorizerAccessToken  string     `json:"authorizer_access_token"`
	ExpiresIn              int        `json:"expires_in"`
	AuthorizerRefreshToken string     `json:"authorizer_refresh_token"`
	FuncInfos              []FuncInfo `json:"func_info"`
}
type RspAuthAccessToken struct {
	AuthorizationInfo AuthAccessTokenInfo `json:"authorization_info"`
}

type AuthAccessToken struct {
	ComponentAppid         string
	AuthorizerAppid        string
	AuthorizerAccessToken  string
	ExpiresIn              int
	AuthorizerRefreshToken string
	FuncInfos              []FuncInfo
}

type ReqUpdateAuthAccessToken struct {
	ComponentAppid         string `json:"component_appid"`
	AuthorizerAppid        string `json:"authorizer_appid"`
	AuthorizerRefreshToken string `json:"authorizer_refresh_token"`
}
type RspUpdateAuthAccessToken struct {
	AuthorizerAccessToken  string `json:"authorizer_access_token"`
	ExpiresIn              int    `json:"expires_in"`
	AuthorizerRefreshToken string `json:"authorizer_refresh_token"`
}

type UpdateAuthAccessToken struct {
	ComponentAppid         string
	AuthorizerAppid        string
	AuthorizerAccessToken  string
	ExpiresIn              int
	AuthorizerRefreshToken string
}
