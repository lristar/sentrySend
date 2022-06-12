package wxwork

func Send(msg string, errUrl string) {
	mgs := MsgContent{ErrorMsg: msg, ErrorUrl: errUrl, OriginIp: "127.0.0.1"}
	str := TemplateExchange(mgs)
	body := SendMsg{Msgtype: "markdown", Markdown: struct {
		Content string "json:\"content\""
	}{Content: str}}
	HttpPostDo("", body)
}
