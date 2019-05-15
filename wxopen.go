package wxopenapi

import (
	"fmt"
	"sync"
	"time"

	"github.com/wxopencrypt"
)

type WxOpen struct {
	mmutex *sync.RWMutex
	minfos map[string]string
	mcrypt *wxopencrypt.WXBizMsgCrypt
}

func newWxOpen() *WxOpen {
	wo := &WxOpen{}
	wo.mcrypt = wxopencrypt.NewWXBizMsgCrypt()
	wo.minfos = make(map[string]string)
	wo.mmutex = new(sync.RWMutex)
	return wo
}

var GWxOpen = newWxOpen()

func (this *WxOpen) Init() {
	this.mcrypt.Init(STOKEN, SENCODINGAESKEY, SAPPID)
	return
}

func (this *WxOpen) GetInfo(key string) (val string) {
	this.mmutex.Lock()
	defer this.mmutex.Unlock()

	val = this.minfos[key]
	return
}

func (this *WxOpen) SetInfo(key, val string) {
	this.mmutex.Lock()
	defer this.mmutex.Unlock()

	this.minfos[key] = val
	return
}

func (this *WxOpen) update_token_loop() {
	for {
		this.update_token()
		time.Sleep(10 * time.Second)
		token := this.GetInfo(COMPONENT_ACCESS_TOKEN)
		if len(token) > 0 {
			time.Sleep(COMPONENT_ACCESS_TOKEN_UPDATE_SECOND * time.Second)
		} else {

		}
	}
}

func (this *WxOpen) update_token() {
	ticket := this.GetInfo(COMPONENT_VERIFY_TICKET)
	if len(ticket) > 0 {
		//DOPOST
	} else {
		fmt.Println("ticket is empty")
	}
}

func (this *WxOpen) Decrypt(sMsgSignature string, sTimeStamp string, sNonce string, sPostData string) (ret int, sMsg string) {
	return this.mcrypt.DecryptMsg(sMsgSignature, sTimeStamp, sNonce, sPostData)
}
