package mocks

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"strconv"
	"strings"
	"time"

	//"github.com/edwingeng/wuid/mysql/wuid"
	"golang.org/x/crypto/bcrypt"

	"project-admin/common/genid/generator"
	"project-admin/common/genid/utils"
	"project-admin/common/uniqueid"
)

type Tag interface {
	Name() string
	Handler(column *Column) string
	Desc() string
}

var (
	numLen = [][]int{
		1: []int{0, 9},
		2: []int{10, 99},
		3: []int{100, 999},
		4: []int{1000, 9999},
		5: []int{10000, 99999},
		6: []int{100000, 999999},
		7: []int{1000000, 9999999},
		8: []int{10000000, 99999999},
		9: []int{100000000, 999999999},
	}
	chars = []byte(`0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ`)
	genid = new(generator.GeneratorData)
)

var tags = []Tag{
	new(Uint),
	new(Char),
	new(Phone),
	new(Email),
	new(Name),
	new(Address),
	new(BankID),
	new(City),
	new(IdCart),
	new(ChineseChar),
	new(English),
	new(OrderNo),
	new(Password),
	new(Time),
	new(TimeStamp),
	new(Date),
	new(DateTime),
}

var TagMap = map[string]Tag{
	"uuid":        new(Uuid),
	"uint":        new(Uint),
	"char":        new(Char),
	"phone":       new(Phone),
	"email":       new(Email),
	"name":        new(Name),
	"address":     new(Address),
	"bankID":      new(BankID),
	"city":        new(City),
	"idCart":      new(IdCart),
	"chineseChar": new(ChineseChar),
	"english":     new(English),
	"orderNo":     new(OrderNo),
	"password":    new(Password),
	"time":        new(Time),
	"timeStamp":   new(TimeStamp),
	"date":        new(Date),
	"dateTime":    new(DateTime),
}

type Uuid struct{}

func (*Uuid) Name() string {
	return "uuid"
}
func (*Uuid) Handler(column *Column) string {
	return strconv.FormatInt(uniqueid.GenId(), 10)
}
func (*Uuid) Desc() string {
	return "唯一id(雪花)"
}

type Uint struct{}

func (*Uint) Name() string {
	return "uint"
}
func (*Uint) Handler(column *Column) string {
	// 直接为零
	if column.Max == 0 && column.Len == 0 || column.Len > 9 {
		return "0"
	}

	if len(column.Content) > 0 {
		return column.Content[rand.Intn(len(column.Content))].(string)
	}

	var (
		max = column.Max
		min = column.Min
	)

	// 固定长度为
	// 1 => [0 ~ 9]       => 0 ~ 9  + 0
	// 2 => [10 ~ 99]     => 0 ~ 89 + 10
	// 3 => [100 ~ 999]   => 0 ~ 899 + 100
	// 4 => [1000 ~ 9999] => 0 ~ 8999 + 1000
	// ....

	if column.Len > 0 && column.Len < 10 {
		// 规避不合理数据存在
		if min == 0 || min > numLen[column.Len][1] || min < numLen[column.Len][0] {
			min = numLen[column.Len][0]
		}

		if max == 0 || max > numLen[column.Len][1] || max < numLen[column.Len][0] {
			max = numLen[column.Len][1]
		}
	}
	r := utils.RUint(max-min+1) + min
	if r > max {
		r = max
	}
	return strconv.Itoa(r)
}
func (*Uint) Desc() string {
	return "随机整数"
}

type Char struct{}

func (*Char) Name() string {
	return "char"
}
func (*Char) Handler(column *Column) string {
	var ret strings.Builder

	chLen := column.Len

	// 固定长度, 固定长度优先级大于可变长度

	if column.FixedLen != nil {
		chLen = column.PrepareFixedLen()
	}

	ret.Grow(chLen)

	utils.RandIntHandler(62, chLen, func(num, i int) {
		ret.WriteByte(chars[num])
	})

	return ret.String()
}
func (*Char) Desc() string {
	return "随机字符 0~9, a~z, A~Z"
}

type Phone struct{}

func (*Phone) Name() string {
	return "phone"
}
func (*Phone) Handler(column *Column) string {
	return genid.GeneratorPhone()
}
func (*Phone) Desc() string {
	return "随机手机号码"
}

type Email struct{}

func (*Email) Name() string {
	return "email"
}
func (*Email) Handler(column *Column) string {
	return genid.GeneratorEmail()
}
func (*Email) Desc() string {
	return "随机邮件"
}

type Name struct{}

func (*Name) Name() string {
	return "name"
}
func (*Name) Handler(column *Column) string {
	return genid.GeneratorName()
}
func (*Name) Desc() string {
	return "随机名字"
}

type Address struct{}

func (*Address) Name() string {
	return "address"
}
func (*Address) Handler(column *Column) string {
	return genid.GeneratorAddress()
}
func (*Address) Desc() string {
	return "随机地址"
}

type BankID struct{}

func (*BankID) Name() string {
	return "bankid"
}
func (*BankID) Handler(column *Column) string {
	return genid.GeneratorBankID()
}
func (*BankID) Desc() string {
	return "随机银行卡号"
}

type City struct{}

func (*City) Name() string {
	return "city"
}
func (*City) Handler(column *Column) string {
	return genid.GeneratorProvinceAdnCityRand()
}
func (*City) Desc() string {
	return "随机城市"
}

type IdCart struct{}

func (*IdCart) Name() string {
	return "idcart"
}
func (*IdCart) Handler(column *Column) string {
	ret, _, _, _, _ := genid.GeneratorIDCart()
	return ret
}
func (*IdCart) Desc() string {
	return "随机身份证号码"
}

type ChineseChar struct{}

func (*ChineseChar) Name() string {
	return "chinese_char"
}
func (*ChineseChar) Handler(column *Column) string {
	return utils.GenRandomLengthChineseChars(6, column.PrepareFixedLen())
	//return utils.GenFixedLengthChineseChars(column.PrepareFixedLen())
}
func (*ChineseChar) Desc() string {
	return "随机中文字符"
}

type English struct{}

func (*English) Name() string {
	return "english"
}
func (*English) Handler(column *Column) string {
	return utils.RandStr(column.PrepareFixedLen())
}
func (*English) Desc() string {
	return "随机a~z,A~Z"
}

//type Wuids struct {
//	g *wuid.WUID
//}
//func NewWuids(db *sql.DB, err error) *Wuids {
//	w := new(Wuids)
//	w.g = wuid.NewWUID("default", nil)
//	_ = w.g.LoadH28FromMysql(func() (*sql.DB, bool, error) {
//		return db, false, nil
//	}, "wuid")
//	return w
//}
//func (*Wuids) Name() string {
//	return "wuid"
//}
//func (w *Wuids) Handler(column *Column) string {
//	return fmt.Sprintf("%#016x", w.g.Next())
//}
//func (*Wuids) Desc() string {
//	return "随机wuid"
//}

type OrderNo struct{}

func (*OrderNo) Name() string {
	return "order_no"
}
func (*OrderNo) Handler(column *Column) string {
	timeNow := time.Now().Format("20060102150405")
	rand := utils.RUint(899999) + 100000
	return timeNow + strconv.Itoa(rand)
}
func (*OrderNo) Desc() string {
	return "随机订单编号"
}

type Password struct {
	pass string
}

func (*Password) Name() string {
	return "password"
}
func (p *Password) Handler(column *Column) string {
	if p.pass != "" {
		return p.pass
	}

	param, ok := column.Params.(string)
	if !ok {
		param = "hash"
	}
	pass, ok := column.Default.Value.(string)
	if !ok {
		pass = "0000"
	}

	switch param {
	case "md5":
		p.pass = p.Md5(pass)
	default:
		// 默认哈希
		p.pass = p.Hash(pass)
	}
	return p.pass
}
func (*Password) Desc() string {
	return "密码支持md5、hash加密方式；配置Default.Value为指定密码，默认为0000，目前仅支持密码为固定统一密码"
}
func (p *Password) Hash(pass string) string {
	bpass, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return ""
	}
	return string(bpass)
}

func (p *Password) Md5(pass string) string {
	h := md5.New()
	h.Write([]byte(p.pass))
	return hex.EncodeToString(h.Sum(nil))
}

// 时间值或持续时间
type Time struct{}

func (*Time) Name() string {
	return "time"
}
func (*Time) Handler(column *Column) string {
	if now, ok := column.Default.Value.(string); ok && now == "now" {
		return time.Now().Format("15:04:05")
	}
	return genid.GeneratorDate().Format("15:04:05")
}
func (*Time) Desc() string {
	return "随机时间，格式：15:04:05,配置Default.Value为now获取当前时间"
}

// 日期
type Date struct{}

func (*Date) Name() string {
	return "date"
}
func (*Date) Handler(column *Column) string {
	if now, ok := column.Default.Value.(string); ok && now == "now" {
		return time.Now().Format("2006-01-02")
	}
	return genid.GeneratorDate().Format("2006-01-02")
}
func (*Date) Desc() string {
	return "随机时间，格式：2006-01-02,配置Default.Value为now获取当前时间"
}

// 时间戳
type TimeStamp struct{}

func (*TimeStamp) Name() string {
	return "timestamp"
}
func (s *TimeStamp) Handler(column *Column) string {

	if now, ok := column.Default.Value.(string); ok && now == "now" {
		return strconv.Itoa(int(time.Now().Unix()))
	}

	return strconv.Itoa(int(genid.GeneratorDate().Unix()))
}
func (s *TimeStamp) Desc() string {
	return "随机时间戳,配置Default.Value为now获取当前时间"
}

type DateTime struct{}

func (*DateTime) Name() string {
	return "datetime"
}
func (t *DateTime) Handler(column *Column) string {
	if now, ok := column.Default.Value.(string); ok && now == "now" {
		return time.Now().Format("2006-01-02 15:04:05")
	}
	return genid.GeneratorDate().Format("2006-01-02 15:04:05")
}
func (t *DateTime) Desc() string {
	return "随机时间，格式：2006-01-02 15:04:05,配置Default.Value为now获取当前时间"
}
