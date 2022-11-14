syntax = "{{.Syntax}}"

{{ .Info }}
type (
	DeleteReq {
		Id int64 `json:"id"` // 主键
	}

	GetReq {
		Id int64 `form:"id"` // 主键
	}

    // 查询列表；
    GetsQueryItem {
        Key        string `json:"key"`                                   // key
        Val        string `json:"val"`                                   // val int/str('xxx')/list('xx,xx,xx')
        Handle     string `json:"handle"`                                // 操作方法 =/>/</like
        NextHandle string `json:"nextHandle,options=or|and,default=and"` // 与下一个条件的操作方法 or/and
    }
    GetsReq {
        Query    []*GetsQueryItem `json:"query"`                              //查询
        OrderBy  string           `json:"orderBy,default=id"`                 //分组
        Sort     string           `json:"sort,options=desc|asc,default=desc"` //排序 desc/asc
        Current  int64            `json:"current,default=1"`                  //当前页
        PageSize int64            `json:"pageSize,default=10"`                //页面大小
    }
)

{{ .Types }}
{{ .ModuleHandler }}