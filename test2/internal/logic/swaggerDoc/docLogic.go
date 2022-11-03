package swaggerDoc

import (
	"context"
	"fmt"
	"net/http"

	"github.com/zeromicro/go-zero/core/logx"
	"project-admin/test2/internal/svc"
)

type DocLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDocLogic(ctx context.Context, svcCtx *svc.ServiceContext) DocLogic {
	return DocLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DocLogic) Doc(w http.ResponseWriter) error {
	html := `
		<!DOCTYPE html>
		<html>
			<head>
			<link type="text/css" rel="stylesheet" href="https://cdn.jsdelivr.net/npm/swagger-ui-dist@3/swagger-ui.css">
			<link rel="shortcut icon" href="https://fastapi.tiangolo.com/img/favicon.png">
			<title>Swagger UI</title>
			</head>
			<body>
			<div id="swagger-ui">
			</div>
			<script src="https://cdn.jsdelivr.net/npm/swagger-ui-dist@3/swagger-ui-bundle.js"></script>
			<!-- 'SwaggerUIBundle' is now available on the page -->
			<script>
			const ui = SwaggerUIBundle({
				url: '/docData',
				// oauth2RedirectUrl: window.location.origin + '/docs/oauth2-redirect',
				dom_id: '#swagger-ui',
				presets: [
				SwaggerUIBundle.presets.apis,
				SwaggerUIBundle.SwaggerUIStandalonePreset
				],
				layout: "BaseLayout",
				deepLinking: true,
				showExtensions: true,
				showCommonExtensions: true
			})
			</script>
			</body>
		</html>
	`
	fmt.Fprintf(w, html)
	return nil
}
