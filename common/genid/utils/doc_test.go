package utils

import (
	"os"
	"runtime/trace"
	"sync"
	"testing"
	"time"
)

func Test_docGoRandInt32(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Logf("docGoRandInt32() = %v", docGoRandInt32(26))
	}
}

func Benchmark_docGoRandInt32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		docGoRandInt32(26)
	}
}

func Test_go_docRandInt32(t *testing.T) {
	setTrace("docGoRandInt32_trace.out", func(){
		goHander(10, func()  {
			docGoRandInt32(26)
			//t.Logf("docGoRandInt32() = %v", docGoRandInt32(26))
		})
	})
}

func Benchmark_go_docRandInt32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		goHander(10, func()  {
			docGoRandInt32(26)
		})
	}
}

func TestRNG_Int32(t *testing.T) {
	var rg RNG
	setTrace("docRngInt_trace.out", func(){
		goHander(10, func()  {
			for i := 0; i < 10; i++ {
				rg.Uint32n(26)
			}
		})
	})
}
func BenchmarkRNG_Int32(b *testing.B) {
	var rg RNG
	for i := 0; i < b.N; i++ {
		goHander(10, func()  {
			for i := 0; i < 10; i++ {
				rg.Uint32n(26)
			}
		})
	}
}

func TestRNG_Int32Hander(t *testing.T) {
	setTrace("docRngInt_trace.out", func(){
		goHander(10, func()  {
			RandIntHandler(26,10, func(a,i int) {})
		})
	})
}
func BenchmarkRNG_Int32Hander(b *testing.B) {
	for i := 0; i < b.N; i++ {
		goHander(10, func()  {
			RandIntHandler(26,10, func(a,i int) {})
		})
	}
}

func Test_timeUinx(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Logf("time.Now().UnixNano() = %v", time.Now().UnixNano())
	}
}

func goHander(workerNum int, handler func() )  {
	var sg sync.WaitGroup
	for i := 0; i < workerNum; i++ {
		sg.Add(1)
		go func() {
			defer sg.Done()
			handler()
		}()
	}
	sg.Wait()
}

func setTrace(name string, handler func()) {
	f, _ := os.Create(name)
	defer f.Close()
	// 启动trace 对程序的执行进行分析捕捉
	err := trace.Start(f)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()

	handler()
}
