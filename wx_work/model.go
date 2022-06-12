package wxwork

import (
	"bytes"
	_ "embed"
	"html/template"
	"strings"
)

//go:embed template.tpl
var tpl string

// wechat
type SendMsg struct {
	Msgtype  string `json:"msgtype"`
	Markdown struct {
		Content string `json:"content"`
	} `json:"markdown"`
}

type MsgContent struct {
	ErrorMsg string `json:"error_msg"`
	OriginIp string `json:"origin_ip"`
	ErrorUrl string `json:"error_url"`
}

// template exchange
func TemplateExchange(msg MsgContent) string {
	buf := new(bytes.Buffer)
	tmpl, err := template.New("http").Parse(strings.TrimSpace(tpl))
	if err != nil {
		panic(err)
	}
	if err := tmpl.Execute(buf, msg); err != nil {
		panic(err)
	}
	return buf.String()
}
