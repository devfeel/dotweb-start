package _http

import (
	"bytes"
	"crypto/tls"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"time"
)

const (
	MaxIdleConns        int = 100
	MaxIdleConnsPerHost int = 50
	IdleConnTimeout     int = 90
	InsecureSkipVerify      = true
)

var currClient *http.Client

func getHttpClient() *http.Client {
	return currClient
}

func init() {
	transport := &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		MaxIdleConns:        MaxIdleConns,
		MaxIdleConnsPerHost: MaxIdleConnsPerHost,
		IdleConnTimeout:     time.Duration(IdleConnTimeout) * time.Second,
		TLSClientConfig:     &tls.Config{InsecureSkipVerify: InsecureSkipVerify},
	}
	currClient = &http.Client{Transport: transport}
}

func HttpGet(url string) (body string, contentType string, intervalTime int64, errReturn error) {
	startTime := time.Now()
	intervalTime = 0
	contentType = ""
	body = ""
	errReturn = nil

	resp, err := getHttpClient().Get(url)
	if err != nil {
		intervalTime = int64(time.Now().Sub(startTime) / time.Millisecond)
		errReturn = err
		return
	}
	defer resp.Body.Close()

	bytebody, err := ioutil.ReadAll(resp.Body)
	intervalTime = int64(time.Now().Sub(startTime) / time.Millisecond)
	if err != nil {
		intervalTime = int64(time.Now().Sub(startTime) / time.Millisecond)
		errReturn = err
		return
	}
	body = string(bytebody)
	contentType = resp.Header.Get("Content-Type")
	intervalTime = int64(time.Now().Sub(startTime) / time.Millisecond)
	return
}

func HttpPost(url string, postbody string, bodyType string) (body string, contentType string, intervalTime int64, errReturn error) {
	startTime := time.Now()
	intervalTime = 0
	contentType = ""
	body = ""
	errReturn = nil
	postbytes := bytes.NewBuffer([]byte(postbody))
	//resp, err := currClient.Post(url, "application/x-www-form-urlencoded", postbytes)
	//resp, err := currClient.Post(url, "application/json", postbytes)
	if bodyType == "" {
		bodyType = "application/x-www-form-urlencoded"
	}
	resp, err := getHttpClient().Post(url, bodyType, postbytes)
	if err != nil {
		intervalTime = int64(time.Now().Sub(startTime) / time.Millisecond)
		errReturn = err
		return
	}
	defer resp.Body.Close()

	bytebody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		intervalTime = int64(time.Now().Sub(startTime) / time.Millisecond)
		errReturn = err
		return
	}
	body = string(bytebody)
	contentType = resp.Header.Get("Content-Type")
	intervalTime = int64(time.Now().Sub(startTime) / time.Millisecond)
	return

}

//从指定query集合获取指定key的值
func GetQuery(querys url.Values, key string) string {
	if len(querys[key]) > 0 {
		return querys[key][0]
	}
	return ""
}
