package wxwork

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

var WXWORKROBOTURL = "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=d4d0ccfc-4d4d-4b7b-997c-a6045073b331"

func HttpPostDo(url string, body interface{}) (map[string]interface{}, error) {
	if url == "" {
		url = WXWORKROBOTURL
	}
	data, err := json.Marshal(&body)
	if err != nil {
		return nil, err
	}
	resp, err := http.Post(url, "application/json", strings.NewReader(string(data)))
	if err != nil {
		return nil, err
	}
	result := make(map[string]interface{}, 0)
	if err := jsonMar(resp.Body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func jsonMar(r io.Reader, obj interface{}) error {
	decoder := json.NewDecoder(r)
	if err := decoder.Decode(obj); err != nil {
		return err
	}
	return nil
}
