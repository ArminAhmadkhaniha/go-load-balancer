package proxy

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"



)

func NewProxy(targetUrl string)(http.Handler, error){
	url, err :=  url.Parse(targetUrl)
	if err != nil{
		return nil, fmt.Errorf("invalid url %v", err)
	}
	proxy := httputil.NewSingleHostReverseProxy(url)

	originalDirector := proxy.Director

	customDirector := func(req *http.Request){
		originalDirector(req)

		req.Host = url.Host
		req.URL.Scheme = url.Scheme
		req.URL.Host = url.Host
		

	}
	proxy.Director = customDirector

	return proxy, nil
	
}

