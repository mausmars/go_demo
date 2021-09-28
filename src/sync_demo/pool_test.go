package main

import (
	"go_demo/src/sync_demo/module"
	"testing"
)
//https://pkg.go.dev/cmd/go#hdr-Testing_flags
//go test -bench=. -benchmem -cpu 4

func BenchmarkWithoutPool(b *testing.B) {
	var s *module.Hello
	for i := 0; i < b.N; i++ {
		s = &module.Hello{A: 1}
		b.StartTimer()
		module.Say(s)
		b.StopTimer()
	}
}

func BenchmarkWithPool(b *testing.B) {
	var s *module.Hello
	for i := 0; i < b.N; i++ {
		s = module.Pool.Get().(*module.Hello)
		s.A = 1
		b.StartTimer()
		module.Say(s)
		b.StopTimer()
		module.Pool.Put(s)
	}
}
