package utils

import (
	"bytes"
	"sync"
	"time"
)

// 指定长度随机中文字符(包含复杂字符)
func GenFixedLengthChineseChars(length int) string {

	var buf bytes.Buffer

	RandIntHandler(40869-19968, length, func(num, i int) {
		buf.WriteRune(rune(num + 19968))
	})
	return buf.String()
}

// 指定范围随机中文字符
func GenRandomLengthChineseChars(start, end int) string {
	length := RandInt(start, end)
	return GenFixedLengthChineseChars(length)
}

// 随机英文小写字母
func RandStr(len int) string {
	data := make([]byte, len)
	RandIntHandler(25, len, func(num, i int) {
		data[i] = byte(num + 97)
	})
	return string(data)
}

// 指定范围随机 int
var rng RNG

func RandInt(min, max int) int {
	return min + RUint(max-min)
}

func RUint(n int) int {
	return int(rng.Uint32n(uint32(n)))
}

func RandIntHandler(maxN, i int, hander func(num, i int)) {
	var y uint32 = rng.Uint32()

	for {
		if i == 0 {
			return
		}
		i--
		// 运算方法
		hander(int((uint64(y)*uint64(maxN))>>32), i)

		y ^= y << 13
		y ^= y >> 17
		y ^= y << 5
	}
}

// 指定范围随机 int64
func RandInt64(min, max int64) int64 {
	return int64(rng.Uint64n(uint64(max - min)))
}

// 反转字符串
func ReverseString(s string) string {
	runes := []rune(s)
	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}
	return string(runes)
}

func Uint32() uint32 {
	v := rngPool.Get()
	if v == nil {
		v = &RNG{}
	}
	r := v.(*RNG)
	x := r.Uint32()
	rngPool.Put(r)
	return x
}

var rngPool sync.Pool

type RNG struct {
	x uint32
	y uint64
}

func (r *RNG) Uint32() uint32 {
	for r.x == 0 {
		r.x = getRandomUint32()
	}

	// See https://en.wikipedia.org/wiki/Xorshift
	x := r.x
	x ^= x << 13
	x ^= x >> 17
	x ^= x << 5
	r.x = x
	return x
}

func (r *RNG) Uint32n(maxN uint32) uint32 {
	x := r.Uint32()
	// See http://lemire.me/blog/2016/06/27/a-fast-alternative-to-the-modulo-reduction/
	return uint32((uint64(x) * uint64(maxN)) >> 32)
}

func (r *RNG) Uint64() uint64 {
	for r.y == 0 {
		r.y = getRandomUint64()
	}

	y := r.y
	y ^= y << 13
	y ^= y >> 7
	y ^= y << 5
	r.y = y
	return y
}
func (r *RNG) Uint64n(maxN uint64) uint64 {
	x := r.Uint64()
	return x % (maxN + 1)
}

func getRandomUint32() uint32 {
	x := time.Now().UnixNano()
	return uint32((x >> 32) ^ x)
}

func getRandomUint64() uint64 {
	x := time.Now().UnixNano()
	return uint64(x)
}
