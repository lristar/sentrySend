package wxwork

import (
	"testing"
)

func TestHttpPostDo(t *testing.T) {
	mgs := MsgContent{ErrorMsg: "错误信息", OriginIp: "127.0.0.1"}
	str := TemplateExchange(mgs)
	body := SendMsg{Msgtype: "markdown", Markdown: struct {
		Content string "json:\"content\""
	}{Content: str}}
	HttpPostDo("", body)
}
