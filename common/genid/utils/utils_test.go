package utils

import (
	"testing"
)

func TestRandInt64(t *testing.T) {
	type args struct {
		min int64
		max int64
	}
	tests := []struct {
		name string
		args args
	}{
		{"0", args{int64(0), int64(1640995200)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RandInt64(tt.args.min, tt.args.max); true {
				t.Logf("RandInt64() = %v ", got)
			}
		})
	}
}

func Test_getRandomUint64(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Logf("getRandomUint64() = %v", getRandomUint64())
	}
}

func Test_getRandomUint32(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Logf("getRandomUint32() = %v", getRandomUint32())
	}
}

func TestRandStr(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Logf("RandStr() = %v", RandStr(10))
	}
}
func BenchmarkRandStr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandStr(10)
	}
}

func TestGenFixedLengthChineseChars(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Logf("GenFixedLengthChineseChars() = %v", GenFixedLengthChineseChars(10))
	}
}
