package main

import (
	"flag"
	"fmt"
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
			Handler: DocData(),
		},
		{
			Method: http.MethodGet,
			Path:   "/buildAPP",
			Handler: BuildAPP(),
		},
	})

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
