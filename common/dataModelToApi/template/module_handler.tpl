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
	post /create(create{{.tableName}}Req)
	
	@doc (
		summary:更新
		handlerType:update
	)
	@handler Update
	put /update(update{{.tableName}}Req)
	
	@doc (
		summary:删除
		handlerType:delete
	)
	@handler Delete
	delete /delete(DeleteReq)
	
	@doc (
		summary:查询一个
		handlerType:get
	)
	@handler Get
	get /get(GetReq) returns({{.tableName}})
	
	@doc (
		summary:查询列表
		handlerType:gets
	)
	@handler Gets
	post /gets(GetsReq) returns({{.tableName}}List)
}
