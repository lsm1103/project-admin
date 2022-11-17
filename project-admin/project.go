package main

import (
	"flag"
	"fmt"
	"github.com/bitly/go-simplejson"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"

	"project-admin/project-admin/internal/config"
	"project-admin/project-admin/internal/handler"
	"project-admin/project-admin/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"project-admin/common/middleware"
)

var configFile = flag.String("f", "etc/project.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf, rest.WithCustomCors(middleware.AllowedFn, nil))

	defer server.Stop()

	handler.RegisterHandlers(server, ctx)

	//swagger-doc
	server.AddRoutes([]rest.Route{
		{
			Method:  http.MethodGet,
			Path:    "/doc",
			Handler: Doc("", "dev"),
		},
		{
			Method: http.MethodGet,
			Path:   "/docData",
			Handler: func(w http.ResponseWriter, r *http.Request) {
				r.ParseForm()
				projectName := r.Form.Get("name")
				path := "/Users/xm/Desktop/go_package/project-admin/project-admin/swagger.json"
				if projectName != ""{
					path = fmt.Sprintf("../projectBuilds/%s/swagger.json", projectName)
				}
				buf, err := Read2Byte(path)
				if err != nil {
					httpx.Error(w, err)
					return
				}
				json, err := simplejson.NewJson(buf)
				if err != nil {
					httpx.Error(w, err)
					return
				}
				httpx.OkJson(w, json)
			},
		},
	})

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
