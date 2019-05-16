package wxopenapi

import (
	"encoding/json"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/wxopencrypt"
)

var STOKEN string
var SENCODINGAESKEY string
var SAPPID string
var SAPPSECRET string
var URL_AUTH_NOTIFY string

type WxOpen struct {
	mout_info func(string, string, int64, int)
	mmutex    *sync.RWMutex
	minfos    map[string]OpenToken
	mcrypt    *wxopencrypt.WXBizMsgCrypt
}

func newWxOpen() *WxOpen {
	wo := &WxOpen{}
	wo.mcrypt = wxopencrypt.NewWXBizMsgCrypt()
	wo.minfos = make(map[string]OpenToken)
	wo.mmutex = new(sync.RWMutex)
	return wo
}

var GWxOpen = newWxOpen()

func (this *WxOpen) Init(stoken, sencodingaeskey, sappid, sappsecret string) {
	STOKEN = stoken
	SENCODINGAESKEY = sencodingaeskey
	SAPPID = sappid
	SAPPSECRET = sappsecret
	this.mcrypt.Init(STOKEN, SENCODINGAESKEY, SAPPID)
	return
}

func (this *WxOpen) SetOutInfoFunc(outf func(string, string, int64, int)) {
	this.mout_info = outf
}

func (this *WxOpen) GetInfo(typ string) (val OpenToken) {
	this.mmutex.Lock()
	defer this.mmutex.Unlock()

	val = this.minfos[typ]
	return
}

func (this *WxOpen) SetInfo(typ, info string, ca int64, ei int) {
	this.mmutex.Lock()
	defer this.mmutex.Unlock()

	var val OpenToken
	val.Typ = typ
	val.Info = info
	val.CreatedAt = ca
	if val.CreatedAt <= 0 {
		val.CreatedAt = time.Now().Unix()
	}
	val.ExpiredIn = ei
	if val.ExpiredIn > 0 {
		val.ExpiredAt = val.CreatedAt + int64(val.ExpiredIn)
	}
	this.minfos[typ] = val
	return
}

func (this *WxOpen) UpdateTokenLoop() {
	time.Sleep(3 * time.Second)
	for {
		tn := time.Now().Unix()
		fmt.Println("UpdateTokenLoop ", tn)

		token := this.GetInfo(PRE_AUTH_CODE)
		if len(token.Typ) <= 0 || tn+100 >= token.ExpiredAt {
			this.UpdatePreAuthCode()
		}

		token = this.GetInfo(COMPONENT_ACCESS_TOKEN)
		if len(token.Typ) <= 0 || tn+600 > token.ExpiredAt {
			this.UpdateAccessToken()
		}

		time.Sleep(10 * time.Second)
	}
}

func (this *WxOpen) UpdateAccessToken() {
	fmt.Println("UpdateAccessToken ", time.Now().Unix())
	ticket := this.GetInfo(COMPONENT_VERIFY_TICKET)
	if len(ticket.Typ) > 0 {
		fmt.Println("ticket is ", ticket)
		var req ReqAccessToken
		req.ComponentAppid = SAPPID
		req.ComponentAppsecret = SAPPSECRET
		req.ComponentVerifyTicket = ticket.Info
		reqstr, _ := json.Marshal(req)
		rsp, err := PostJsonByte(URL_COMPONENT_ACCESS_TOKEN, reqstr)
		if err != nil || strings.Contains(string(rsp), "errcode") {
			fmt.Println("请求 access 失败", string(rsp))
			return
		}
		var rspobj RspAccessToken
		json.Unmarshal(rsp, &rspobj)
		fmt.Println(rspobj)
		this.SetInfo(COMPONENT_ACCESS_TOKEN, rspobj.ComponentAccessToken, 0, rspobj.ExpiresIn)
		if this.mout_info != nil {
			this.mout_info(COMPONENT_ACCESS_TOKEN, rspobj.ComponentAccessToken, 0, rspobj.ExpiresIn)
		}
	} else {
		fmt.Println("ticket is empty")
	}
}

func (this *WxOpen) UpdatePreAuthCode() {
	fmt.Println("UpdatePreAuthCode ", time.Now().Unix())
	access := this.GetInfo(COMPONENT_ACCESS_TOKEN)
	if len(access.Typ) > 0 {
		fmt.Println("access is ", access)
		var req ReqPreAuthCode
		req.ComponentAppid = SAPPID
		reqstr, _ := json.Marshal(req)
		rsp, err := PostJsonByte(fmt.Sprintf(URL_PRE_AUTH_CODE, access.Info), reqstr)
		if err != nil {
			fmt.Println("请求 preauthcode 失败 err", err.Error())
			return
		}
		if strings.Contains(string(rsp), "errcode") {
			fmt.Println("请求 preauthcode 失败 rsp", string(rsp))
			if strings.Contains(string(rsp), "access_token expired") {
				this.UpdateAccessToken()
			}
			return
		}
		var rspobj RspPreAuthCode
		json.Unmarshal(rsp, &rspobj)
		fmt.Println(rspobj)
		this.SetInfo(PRE_AUTH_CODE, rspobj.PreAuthCode, 0, rspobj.ExpiresIn)
		if this.mout_info != nil {
			this.mout_info(PRE_AUTH_CODE, rspobj.PreAuthCode, 0, rspobj.ExpiresIn)
		}
	} else {
		fmt.Println("access is empty")
	}
}

func (this *WxOpen) GetPreAuthUrl() string {
	preauthcode := this.GetInfo(PRE_AUTH_CODE)
	if len(SAPPID) <= 0 || len(preauthcode.Typ) <= 0 || len(URL_AUTH_NOTIFY) <= 0 {
		return ""
	}
	url := fmt.Sprintf(FMT_URL_PRE_AUTH_CODE, SAPPID, preauthcode.Info, URL_AUTH_NOTIFY, 3)
	return url
}

func (this *WxOpen) Decrypt(sMsgSignature string, sTimeStamp string, sNonce string, sPostData string) (ret int, sMsg string) {
	return this.mcrypt.DecryptMsg(sMsgSignature, sTimeStamp, sNonce, sPostData)
}
