package _http

import (
	"testing"
)

func TestHttpGet(t *testing.T) {
	url := "http://www.baidu.com"
	_, _, _, err := HttpGet(url)
	if err != nil {
		t.Error(err)
	} else {
		t.Log("HttpGet success")
	}
}

func TestHttpPost(t *testing.T) {
	url := "http://www.baidu.com"
	_, _, _, err := HttpPost(url, "", "")
	if err != nil {
		t.Error(err)
	} else {
		t.Log("HttpPost success")
	}
}
