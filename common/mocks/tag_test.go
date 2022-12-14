package mocks

import (
	"encoding/json"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"reflect"
	"strconv"
	"strings"
	"testing"
	"unsafe"
)

func TestNum_Handler(t *testing.T) {
	tests := []struct {
		name string
		args *Column
	}{
		//{"1", &Column{Max: 10, Min: 0}},
		//{"2", &Column{Len: 7}},
		//{"3", &Column{Len: 9}},
		//{"4", &Column{Max: 100, Min: 99}},
		//{"5", &Column{Len: 2, Max: 98}},
		{"6", &Column{Len: 20}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nu := &Uint{}
			got := nu.Handler(tt.args)
			t.Logf("Handler() = %v ", got)
		})
	}
}

func BenchmarkNum_Handler(b *testing.B) {
	nu := &Uint{}
	args := &Column{Max: 10, Min: 0}
	for i := 0; i < b.N; i++ {
		nu.Handler(args)
	}
}

func TestChar_Handler(t *testing.T) {
	tests := []struct {
		name string
		args *Column
	}{
		{"1", &Column{Len: 5}},
		{"1", &Column{Len: 5, FixedLen: 8}},
		{"1", &Column{FixedLen: 8}},
		{"1", &Column{FixedLen: "2|3"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ch := &Char{}
			got := ch.Handler(tt.args)
			t.Logf("Handler() = %v ", got)
		})
	}
}

func BenchmarkChar_Handler(b *testing.B) {
	tag := &Char{}
	args := &Column{FixedLen: "2|3"}
	for i := 0; i < b.N; i++ {
		r := tag.Handler(args)
		b.Logf("Handler() = %v", r)
	}
}

func TestDate_Handler(t *testing.T) {
	type args struct {
		column *Column
	}
	tests := []struct {
		name string
		args args
	}{
		{"1", args{&Column{}}},
		{"2", args{&Column{}}},
		{"3", args{&Column{}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			da := &Date{}
			if got := da.Handler(tt.args.column); true {
				t.Logf("Handler() = %v", got)
			}
		})
	}
}
func BenchmarkDate_Handler(b *testing.B) {
	da := &Date{}
	args := &Column{}
	for i := 0; i < b.N; i++ {
		r := da.Handler(args)
		b.Logf("Handler() = %v", r)
	}
}

func TestPassword_Handler(t *testing.T) {
	tests := []struct {
		name string
		args *Column
	}{
		{"1", &Column{
			Params: "md5",
			Default: struct {
				Value     interface{} `json:"value"`
				Frequency int         `json:"frequency"`
			}{"9999", 1},
		}},
		{"2", &Column{
			Params: nil,
			Default: struct {
				Value     interface{} `json:"value"`
				Frequency int         `json:"frequency"`
			}{},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Password{}
			if got := p.Handler(tt.args); true {
				t.Logf("Handler() = %v", got)
			}
		})
	}
}

func TestTimeStamp_Handler(t *testing.T) {
	tests := []struct {
		name string
		args *Column
	}{
		{"1", &Column{
			Default: struct {
				Value     interface{} `json:"value"`
				Frequency int         `json:"frequency"`
			}{"now", 1},
		}},
		{"2", &Column{}},
		{"3", &Column{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &TimeStamp{}
			if got := s.Handler(tt.args); true {
				t.Logf("Handler() = %v", got)
			}
		})
	}
}

type SingleMsgList struct {
	List     []*SingleMsg `json:"list"tag:"list"`                           // ????????????
	Current  int64        `json:"current"tag:"uint"len:"2"min:"1"max:"10"`  // ?????????
	PageSize int64        `json:"pageSize"tag:"uint"len:"2"content:"10|20"` // ????????????
	IsNext   bool         `json:"isNext"tag:"char"content:"true|false"`     // ??????????????????
}

type SingleMsg struct {
	Id               int64  `json:"id"tag:"uuid"`                                   // ?????????
	SenderType       string `json:"sender_type"tag:"uint"min:"1"max:"3"`            // ??????????????????1????????????2????????????3??????
	SenderId         int64  `json:"sender_id"tag:"uuid"`                            // ?????????id
	SenderDeviceId   string `json:"sender_device_id"tag:"uuid"`                     // ????????????id
	ReceiverId       int64  `json:"receiver_id"tag:"uuid"`                          // ?????????id
	ReceiverDeviceId string `json:"receiver_device_id"tag:"char"fixed_len:"18|100"` // ????????????id???????????????id"???"?????????*??????????????????
	ParentId         int64  `json:"parent_id"tag:"uuid"`                            // ??????id???????????????
	SendTime         string `json:"send_time"tag:"dateTime"`                        // ??????????????????
	MsgType          string `json:"msg_type"tag:"char"content:"text|img|link"`      // ????????????
	MsgContent       string `json:"msg_content"tag:"chineseChar"fixed_len:"18|150"` // ????????????
	Status           int64  `json:"status"tag:"uint"content:"-1|0|1"`               // ???????????????-1?????????0????????????1??????
	CreateTime       string `json:"createTime"tag:"dateTime"`                       // ????????????
	UpdateTime       string `json:"updateTime"tag:"dateTime"`                       // ????????????
	//Min int `json:"min"tag:"uint"`
	//Params interface{} `json:"params"tag:"char"`
	//Struct_       	 sss `json:"struct_"tag:"struct"`         // struct_
	//List_            []sss `json:"list_"tag:"list"`         // list_
	//Map_       		 map[string]string `json:"map_"tag:"map"`         // map_
	//Func_       	 func() `json:"func_"`         // func_,????????????
}

type sss struct {
	D     string `json:"msg_content"tag:"chineseChar"fixed_len:"18|150"`
	BTime string `json:"updateTime"tag:"dateTime"` // ????????????
}

func Test_RespMock(t *testing.T) {
	resp := &SingleMsgList{}
	RespMock(resp)
	marshal, err := json.Marshal(resp)
	if err != nil {
		t.Logf("err:%s", err.Error())
	}
	t.Logf("resp: %+v, %s\n", resp, marshal)
	return
}

//?????????????????????mock??????????????????????????????2022.10.19
//????????????????????????????????????????????????????????????tag????????????????????????????????????????????????????????????????????????????????????????????????
func interfaceT0(data interface{}) {
	ty := reflect.TypeOf(data).Elem()
	v := reflect.ValueOf(data).Elem()
	mockT0(ty, v)
	fmt.Printf("*???data: %+v???\n", data)
}
func mockT0(dt reflect.Type, dv reflect.Value) {
	//if dt.Kind() == reflect.Ptr {
	//	if dv.IsNil() {
	//		panic("nil ptr")
	//	}
	//	// ?????????????????????????????????????????????struct
	//	originType := dv.Elem().Type()
	//	if originType.Kind() != reflect.Struct {
	//		return
	//	}
	//	// ?????????
	//	dv = dv.Elem()
	//	dt = dt.Elem()
	//}

	// ????????????
	for i := 0; i < dt.NumField(); i++ {
		// ???tag
		field := dt.Field(i)
		tag := field.Tag
		name := field.Name
		fieldType := field.Type.String()
		fieldValue := dv.FieldByName(name)
		//if !fieldValue.IsValid() {
		//	continue
		//}
		//if fieldValue.CanInterface() {
		//	fmt.Printf("fieldValue.Interface():%+v \n", fieldValue.Interface())
		//}
		tag_label := tag.Get("tag")
		switch tag_label {
		case "struct": //?????????????????????
			struct_ := field.Type
			mockT0(struct_, fieldValue)
		case "list": //??????????????????
			// ??????????????????????????????????????????
			listItemStructT := fieldValue.Type().Elem()
			listItemKind := listItemStructT.Kind()
			if listItemKind == reflect.Ptr {
				listItemStructT = listItemStructT.Elem()
			}
			listItemStructV := reflect.New(listItemStructT).Elem()
			mockT0(listItemStructT, listItemStructV)

			listItemStructV_ := reflect.NewAt(listItemStructT, unsafe.Pointer(listItemStructV.UnsafeAddr()))
			if listItemKind == reflect.Struct {
				listItemStructV_ = listItemStructV_.Elem()
			}
			val_ := reflect.Append(fieldValue, listItemStructV_)
			fieldValue.Set(val_)
			fmt.Printf("*???list: %+v???\n", fieldValue)
			//// ???????????????????????????????????????????????????
			//newArr := make([]reflect.Value, 0)
			//newArr = append(newArr, listItemStructV)
			//// ???????????????????????????????????????
			//resArr := reflect.Append(fieldValue, newArr...)
			//// ????????????????????????
			//fieldValue.Set(resArr)

			//fmt.Printf("newArr:%+v,resArr:%+v", newArr, resArr)

			//dv_.Elem().Set(listItemStructV)
			//pkg_ := list_.Elem().Elem().String()
			//fmt.Printf("pkg_:%s, %+v", pkg_, dt.FieldByIndex([]int{i,0}))
			//pkgPaths := strings.Split(pkgPath, "/")
			//pkg_s := strings.Split(pkg_, ".")
			//if pkg_s[0] == pkgPaths[len(pkgPaths)-1]{
			//	objs := getObjByFile(pkgPath)
			//	fmt.Printf("objs:%+v", objs)
			//}
			//mockT1(list_, dv_)
		case "map":
			//??????????????????
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

			//??????????????????
			value := TagMap[tag_label].Handler(&Column{
				Len:      len_,
				FixedLen: fixed_len_,
				Max:      min_,
				Min:      max_,
				Content:  content_,
			})

			//?????????????????????
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
			default:
				valueRF = reflect.ValueOf(value)
			}
			//??????
			fieldValue.Set(valueRF)
			fmt.Printf("???%s: %+v???\n", name, valueRF.Interface())
		}

	}
}

//??????go????????????????????????????????????????????????????????? interface{}???struct???var?????????
func getObjByFile(source string) *ast.File {
	fs := token.NewFileSet()
	file, err := parser.ParseFile(fs, source, nil, 0)
	if err != nil {
		fmt.Print(err)
		return nil
	}
	return file
}
