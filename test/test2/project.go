package main

import (
	"flag"
	"fmt"
	"net/http"

	"project-admin/test2/internal/config"
	"project-admin/test2/internal/handler"
	"project-admin/test2/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/project.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf, rest.WithCustomCors(func(header http.Header) {
		//"*"表示接受任意域名的请求，这个值也可以根据自己需要，设置成不同域名
		header.Set("Access-Control-Allow-Origin", "*")
		header.Set("Access-Control-Allow-Credentials", "true")
		header.Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		header.Add("Access-Control-Allow-Headers", "DNT,X-CustomHeader,Cookies,Keep-Alive,Range,User-Agent,X-Requested-With,If-Modified-Since,Cache-Control,Content-Type,Access-Control-Allow,Token,Accept,Authorization,x-auth-token")
	}, nil))
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
