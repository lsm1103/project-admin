package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
	"project-admin/projectBuilds/project2/internal/config"
	"project-admin/projectBuilds/project2/internal/handler"
	"project-admin/projectBuilds/project2/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/project.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	// server := rest.MustNewServer(c.RestConf)
	server := rest.MustNewServer(c.RestConf, rest.WithCustomCors(func(header http.Header) {
		//"*"表示接受任意域名的请求，这个值也可以根据自己需要，设置成不同域名
		header.Set("Access-Control-Allow-Origin", "*")
		//header.Set("Access-Control-Allow-Origin", "*.pj-ai.com")
		//w.Header().Add("Access-Control-Allow-Credentials", "true")
		//w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		header.Add("Access-Control-Allow-Headers", "DNT,X-CustomHeader,Cookies,Keep-Alive,Range,User-Agent,X-Requested-With,If-Modified-Since,Cache-Control,Content-Type,Access-Control-Allow,Token,Accept,Authorization,x-auth-token")
	}, nil))
	defer server.Stop()

	server.Use(func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			logx.Info("global middleware")
			next(w, r)
			header := w.Header()
			logx.Infof("header:%+v\n", header)
		}
	})
	//// 自定义错误
	//httpx.SetErrorHandler(func(err error) (int, interface{}) {
	//	switch e := err.(type) {
	//	case *xerr.CodeError:
	//		return http.StatusOK, e.Msg
	//	default:
	//		return http.StatusInternalServerError, nil
	//	}
	//})

	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
