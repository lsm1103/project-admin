package middleware

import "net/http"

func AllowedFn (header http.Header) {
	//"*"表示接受任意域名的请求，这个值也可以根据自己需要，设置成不同域名
	header.Set("Access-Control-Allow-Origin", "*")
	//header.Set("Access-Control-Allow-Origin", "*.pj-ai.com")
	//w.Header().Add("Access-Control-Allow-Credentials", "true")
	//w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	header.Add("Access-Control-Allow-Headers", "DNT,X-CustomHeader,Cookies,Keep-Alive,Range,User-Agent,X-Requested-With,If-Modified-Since,Cache-Control,Content-Type,Access-Control-Allow,Token,Accept,Authorization,x-auth-token")
}

func NotAllowedFn (w http.ResponseWriter) {
	//写通过cors拒绝请求的逻辑
	return
}