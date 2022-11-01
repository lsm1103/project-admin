package mocks

import (
	"fmt"
	"project-admin/common/genid/utils"
	"strconv"
	"strings"
)

type Column struct {
	// 字段标签，确定字段生成方式
	Tag string `json:"tag"`
	// 字段固定长度
	Len int `json:"len"`
	// 长度范围,可变长度
	FixedLen interface{} `json:"fixed_len"`

	// 字段最大值
	Max int `json:"max"`
	// 字段最小值
	Min int `json:"min"`
	// 其它参数
	Params interface{} `json:"params"`
	// 默认值
	Default struct {
		Value     interface{} `json:"value"`
		Frequency int         `json:"frequency"`
	} `json:"default"`
	// 指定随机内容
	Content []interface{} `json:"content"`
}

//
//  PrepareFixedLen
//  @Description: 解析可变长度/优先级上固定长度大于可变长度
//  @receiver column
//  @return int
//
func (column *Column) PrepareFixedLen() int {

	if column.Len != 0 {
		return column.Len
	}

	if column.FixedLen == nil {
		return 0
	}

	switch column.FixedLen.(type) {
	case int:
		return column.FixedLen.(int)
	case string:
		rangelen := strings.Split(column.FixedLen.(string), "|")
		lens := make([]int, 0, len(rangelen))
		for _, v := range rangelen {
			tmp, _ := strconv.Atoi(v)
			lens = append(lens, tmp)
		}

		if len(lens) > 0 && lens[1] > lens[0] {
			return utils.RUint(lens[1]-lens[0]+1) + lens[0]
		}
	}

	return 0
}

//
//  PrepareDefult
//  @Description: 解析interface{}
//  @receiver column
//  @return string
//
func (column *Column) Marshaler(v interface{}) string {
	switch v.(type) {
	case float64, int, float32:
		return fmt.Sprintf("%v", v)
	case string:
		return v.(string)
	}
	return ""
}
