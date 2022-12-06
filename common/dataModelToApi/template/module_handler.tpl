@server(
	prefix: admin/{{.tableName}}/v1
	group: {{.tableName}}
	module: {{.tableName}}
	serviceType: {{.serviceType}}
)
service project {
	@doc (
		summary:创建
		handlerType:create
	)
	@handler Create
	post / (create{{.tableName}}Req) returns({{.tableName}})

	@doc (
		summary:删除
		handlerType:delete
	)
	@handler Delete
	delete / (DeleteReq)
	
	@doc (
		summary:更新
		handlerType:update
	)
	@handler Update
	put / (update{{.tableName}}Req) returns({{.tableName}})
	
	@doc (
		summary:查询一个
		handlerType:get
	)
	@handler Get
	get /:id (GetReq) returns({{.tableName}})
	
	@doc (
		summary:查询列表
		handlerType:gets
	)
	@handler Gets
	post /gets (GetsReq) returns({{.tableName}}List)
}
