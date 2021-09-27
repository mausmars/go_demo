package cacheline_demo

import (
	//"flag"
	"sync"
	"testing"
)

//https://learnku.com/go/t/45683

//go test -v -bench=. -benchmem -run=none

type SimpleStruct struct {
	n int
}

//M 是等于 1 百万的常数
func BenchmarkStructureFalseSharing(b *testing.B) {
	M := 1000000
	structA := SimpleStruct{}
	structB := SimpleStruct{}
	wg := sync.WaitGroup{}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		wg.Add(2)
		go func() {
			for j := 0; j < M; j++ {
				structA.n += j
			}
			wg.Done()
		}()
		go func() {
			for j := 0; j < M; j++ {
				structB.n += j
			}
			wg.Done()
		}()
		wg.Wait()
	}
}

const CacheLinePadSize = 64
const M = 1000000

type PaddedStruct struct {
	n int

	pad [CacheLinePadSize - 32%CacheLinePadSize]byte
}

func BenchmarkStructurePadding(b *testing.B) {
	structA := PaddedStruct{}
	structB := SimpleStruct{}
	wg := sync.WaitGroup{}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		wg.Add(2)
		go func() {
			for j := 0; j < M; j++ {
				structA.n += j
			}
			wg.Done()
		}()
		go func() {
			for j := 0; j < M; j++ {
				structB.n += j
			}
			wg.Done()
		}()
		wg.Wait()
	}
}
