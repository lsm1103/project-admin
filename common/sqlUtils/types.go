package sqlUtils

import "errors"

type (
	GetsQueryItem struct {
		Key        string `json:"key"`                                   // key
		Val        string `json:"val"`                                   // val int/str('xxx')/list('xx,xx,xx')
		Handle     string `json:"handle"`                                // 操作方法 =/>/</like
		NextHandle string `json:"nextHandle,options=or|and,default=and"` // 与下一个条件的操作方法 or/and
	}

	GetsReq struct {
		Query    []*GetsQueryItem `json:"query"`                              //查询
		OrderBy  string           `form:"orderBy,default=id"`                 //分组
		Sort     string           `form:"sort,options=desc|asc,default=desc"` //排序 desc/asc
		Current  int64            `form:"current,default=1"`                  //当前页
		PageSize int64            `form:"pageSize,default=10"`                //页面大小
	}
)

var (
	ErrNotFound = errors.New("找不到数据")
	ErrNotState = errors.New("没有state字段")

	//状态
	Del     int64 = -2 //删除
	Disable int64 = -1 //禁用
	Audited int64 = 1  //待审核
	Enable  int64 = 2  //启用
)