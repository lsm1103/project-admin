package mocks

type Tags interface {
	Uuid(column *Column) int64
	Uint(column *Column) int
	Uint64(column *Column) int64
	Char(column *Column) string
	Phone(column *Column) int64
	//Email(column *Column) string
	//Name(column *Column) string
	//Address(column *Column) string
	//BankID(column *Column) int64
	//City(column *Column) string
	//IdCart(column *Column) string
	//ChineseChar(column *Column) string
	//English(column *Column) string
	//OrderNo(column *Column) string
	//Password(column *Column) string
	//Time(column *Column) string
	//TimeStamp(column *Column) string
	//Date(column *Column) string
	//DateTime(column *Column) string
}

//type Tag_ struct {}
//
////唯一id(雪花)
//func (*Tag_) Uuid(column *Column) int64 {
//	return uniqueid.GenId()
//}
//
////随机整数,最大9位数
//func (*Tag_) Uint(column *Column) int {
//	// 直接为零
//	if column.Max == 0 && column.Len == 0 || column.Len > 9 {
//		return 0
//	}
//
//	var (
//		max = column.Max
//		min = column.Min
//	)
//
//	// 固定长度为
//	// 1 => [0 ~ 9]       => 0 ~ 9  + 0
//	// 2 => [10 ~ 99]     => 0 ~ 89 + 10
//	// 3 => [100 ~ 999]   => 0 ~ 899 + 100
//	// 4 => [1000 ~ 9999] => 0 ~ 8999 + 1000
//	// ....
//
//	if column.Len > 0 && column.Len < 10 {
//		// 规避不合理数据存在
//		if min == 0 || min > numLen[column.Len][1] || min < numLen[column.Len][0] {
//			min = numLen[column.Len][0]
//		}
//
//		if max == 0 || max > numLen[column.Len][1] || max < numLen[column.Len][0] {
//			max = numLen[column.Len][1]
//		}
//	}
//
//	return utils.RUint(max-min+1) + min
//}
//
////随机字符 0~9, a~z, A~Z
//func (*Tag_) Char(column *Column) string {
//	var ret strings.Builder
//
//	chLen := column.Len
//
//	// 固定长度, 固定长度优先级大于可变长度
//
//	if column.FixedLen != nil {
//		chLen = column.PrepareFixedLen()
//	}
//
//	ret.Grow(chLen)
//
//	utils.RandIntHandler(62, chLen, func(num, i int) {
//		ret.WriteByte(chars[num])
//	})
//
//	return ret.String()
//}
//
////随机手机号码
//func (*Tag_) Phone(column *Column) int64 {
//	parseInt, err := strconv.ParseInt(genid.GeneratorPhone(), 10, 64)
//	if err != nil {
//		return 13100001111
//	}
//	return parseInt
//}
