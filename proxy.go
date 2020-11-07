package gopay

import (
	"net/http"
	"net/url"
)

var proxyFunc func(*http.Request) (*url.URL, error)

//UseProxy 设置代理地址
func UseProxy(rawURL string) {
	URL, err := url.Parse(rawURL)
	if err != nil {
		panic(err)
	}
	proxyFunc = http.ProxyURL(URL)
}

//GetProxy 获取全局代理
func GetProxy() func(*http.Request) (*url.URL, error) {
	return proxyFunc
}
