package sqlUtils

import (
	"fmt"
	"strings"
)

type (
	ModelTool interface {
		BuildQuery(in *GetsReq, ListRows string, tableName string) string
	}

	defaultModelTool struct {}
)

func NewModelTool() ModelTool {
	return &defaultModelTool{}
}

func (t defaultModelTool) BuildQuery(in *GetsReq, ListRows string, tableName string) string {
	filter_ := ""
	tmp_ := ""
	if in.Query != nil {
		for index,value := range in.Query {
			Key := strings.TrimSpace(value.Key)
			Val := strings.TrimSpace(value.Val)
			Handle := strings.TrimSpace(value.Handle)
			NextHandle := strings.TrimSpace(value.NextHandle)
			switch Handle {
			case "between":
				t_ := strings.Split(Val, ",")
				tmp_ = fmt.Sprintf(" `%s` between '%s' and '%s' ", Key, t_[0], t_[1])
			case "in":
				t_ := ""
				for _, v := range strings.Split(Val, ",") {
					if len(t_) == 0 {
						t_ = fmt.Sprintf("'%s'", strings.ToLower(v))
					} else {
						t_ = fmt.Sprintf("%s, '%s'", t_, strings.ToLower(v))
					}
				}
				tmp_ = fmt.Sprintf(" `%s` in (%s) ", Key, t_)
			default:
				tmp_ = fmt.Sprintf(" `%s` %s '%s' ", Key, Handle, Val)
			}
			if index == 0 {
				filter_ = tmp_
			} else {
				filter_ = filter_ + NextHandle + tmp_
			}
		}
		if filter_ != "" {
			filter_ = fmt.Sprintf("where %s", filter_)
		}
	}
	//Fields_ := ""
	//for _, v_ := range strings.Split(in.Fields, ",") {
	//	if len(Fields_) == 0 {
	//		Fields_ = fmt.Sprintf("`%s`", strings.ToLower(v_))
	//	} else {
	//		Fields_ = fmt.Sprintf("%s, `%s`", Fields_, strings.ToLower(v_))
	//	}
	//}
	return fmt.Sprintf(
		"select %s from %s %s order by %s %v limit %v offset %v",
		ListRows,
		tableName,
		filter_,
		in.OrderBy,
		in.Sort,
		in.PageSize+1,
		(in.Current - 1) * in.PageSize,
	)
}
