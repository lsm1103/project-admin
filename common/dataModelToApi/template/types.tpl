type (
	create{{.tableName}}Req {
	    {{.createFields}}
	}

	update{{.tableName}}Req {
	    {{.updateFields}}
	}
	{{.tableName}} {
	    {{.allFields}}
	}

	// 查询列表结果
	{{.tableName}}List {
		List     []*{{.tableName}} `json:"list"tag:"list"`                       // 数据列表
		Current  int64        `json:"current"tag:"uint"min:"1"max:"1000"`     // 当前页
		PageSize int64        `json:"pageSize"tag:"uint"min:"10"max:"20"`    // 页面大小
		IsNext   bool         `json:"isNext"tag:"bool"` // 是否有下一页
		// total           int64        `json:"total"`  // 总数
	}
)
