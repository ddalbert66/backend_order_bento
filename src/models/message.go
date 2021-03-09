package models

type Message struct {
	Id      int64  `json:"id,omitempty form:"id"`           //消息ID
	Userid  int64  `json:"userid,omitempty form:"userid"`   //誰發的
	Cmd     int    `json:"cmd,omitempty form:"cmd"`         //群聊還是私聊
	Dstid   int64  `json:"dstid,omitempty form:"dstid"`     //對端ID/群ID
	Media   int    `json:"media,omitempty form:"media"`     //消息樣式
	Content string `json:"content,omitempty form:"content"` //消息內容
	Pic     string `json:"pic,omitempty form:"pic"`         //預覽圖片
	Url     string `json:"url,omitempty form:"url"`         //服務的URL
	Memo    string `json:"memo,omitempty form:"memo"`       //簡單描述
	Amount  int    `json:"amount,omitempty form:"amount"`   //和數字相關的
}
