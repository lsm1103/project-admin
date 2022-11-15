package mocks

import (
	"fmt"
	"math/rand"
	"reflect"
	"strconv"
	"strings"
	"time"
	"unsafe"
)

func RespMock(resp interface{}) {
	dt := reflect.TypeOf(resp).Elem()
	dv := reflect.ValueOf(resp).Elem()
	mockCore(dt, dv)
	//fmt.Printf("*【resp: %+v】\n", resp)
}

func mockCore(dt reflect.Type, dv reflect.Value) {
	if dv.Kind() == reflect.Invalid {
		dv = reflect.New(dt).Elem()
	}
	// 解析字段
	for i := 0; i < dt.NumField(); i++ {
		// 取tag
		field := dt.Field(i)
		tag := field.Tag
		name := field.Name
		fieldType := field.Type.String()
		fieldValue := dv.FieldByName(name)
		tag_label := tag.Get("tag")

		switch tag_label {
		case "struct":
			//结构体类型处理
			struct_ := field.Type
			mockCore(struct_, fieldValue)
		case "list":
			//列表类型处理
			var val_ reflect.Value
			//设置随机数种子，由于种子数值，每次启动都不一样
			rand.Seed(time.Now().UnixNano())
			count := rand.Intn(10)
			if count == 0 {
				count = 1
			}
			for i_ := 0; i_ < count; i_++ {

				listItemStructT := fieldValue.Type().Elem()
				//确定列表的子项类型：指针｜结构体
				listItemKind := listItemStructT.Kind()
				// 列表的子项为指针时，listItemStructT需进行处理为结构体类型
				if listItemKind == reflect.Ptr {
					listItemStructT = listItemStructT.Elem()
				}
				listItemStructV := reflect.New(listItemStructT).Elem()
				mockCore(listItemStructT, listItemStructV)

				listItemStructV_ := reflect.NewAt(listItemStructT, unsafe.Pointer(listItemStructV.UnsafeAddr()))
				// 列表的子项为结构体时，listItemStructV_需进行处理为指针类型
				if listItemKind == reflect.Struct {
					listItemStructV_ = listItemStructV_.Elem()
				}
				if !val_.IsValid() {
					val_ = reflect.Append(fieldValue, listItemStructV_)
				} else {
					val_ = reflect.Append(val_, listItemStructV_)
				}
			}

			fieldValue.Set(val_)
			//fmt.Printf("*【list: %+v】\n", fieldValue)
		case "map":
			//字典类型处理
		default:
			len_label := tag.Get("len")
			fixed_len_label := tag.Get("fixed_len")
			min_label := tag.Get("min")
			max_label := tag.Get("max")
			content_label := tag.Get("content")

			var err error
			var len_, min_, max_ int
			if len_label != "" {
				len_, err = strconv.Atoi(len_label)
				if err != nil {
					fmt.Printf("err:%+v", err)
				}
			}
			var fixed_len_ interface{}
			if fixed_len_label != "" {
				fixed_len_ = fixed_len_label
			}
			if min_label != "" {
				min_, err = strconv.Atoi(min_label)
				if err != nil {
					fmt.Printf("err:%+v", err)
				}
			}
			if max_label != "" {
				max_, err = strconv.Atoi(max_label)
				if err != nil {
					fmt.Printf("err:%+v", err)
				}
			}
			content_ := []interface{}{}
			if content_label != "" {
				for _, item := range strings.Split(content_label, "|") {
					content_ = append(content_, item)
				}
			}

			//生成模拟的值
			value := TagMap[tag_label].Handler(&Column{
				Len:      len_,
				FixedLen: fixed_len_,
				Max:      max_,
				Min:      min_,
				Content:  content_,
			})

			//转化成反射格式
			var valueRF reflect.Value
			switch fieldType {
			case "int64":
				int64_, err := strconv.ParseInt(value, 10, 64)
				if err != nil {
					fmt.Printf("set int64 value err: %s", err.Error())
					valueRF = reflect.ValueOf(int64(0))
				}
				valueRF = reflect.ValueOf(int64_)
			case "int":
				int_, err := strconv.Atoi(value)
				if err != nil {
					fmt.Printf("set int value err: %s", err.Error())
					valueRF = reflect.ValueOf(int(0))
				}
				valueRF = reflect.ValueOf(int_)
			case "bool":
				bool_, err := strconv.ParseBool(value)
				if err != nil {
					fmt.Printf("set int value err: %s", err.Error())
					valueRF = reflect.ValueOf(false)
				}
				valueRF = reflect.ValueOf(bool_)
			default:
				valueRF = reflect.ValueOf(value)
			}
			//赋值
			fieldValue.Set(valueRF)
			//fmt.Printf("【%s: %+v】\n", name, valueRF.Interface())
		}

	}
}
