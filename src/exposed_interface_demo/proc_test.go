package exposed_interface_demo

import (
	"runtime"
	"sync"
	"testing"
)

func TestProc(t *testing.T) {
	pid := sync.ProcPin()
	t.Logf("PID: %v", pid)
	sync.ProcUnpin()
}

func BenchmarkProc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sync.ProcPin()
		sync.ProcUnpin()
	}
}

func BenchmarkProcParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			sync.ProcPin()
			sync.ProcUnpin()
		}
	})
}

func BenchmarkProcSync(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sync.ProcPin()
		sync.ProcUnpin()
	}
}

func BenchmarkProcSyncParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			sync.ProcPin()
			sync.ProcUnpin()
		}
	})
}

func TestLockThread(t *testing.T) {
	runtime.LockOSThread()
	runtime.UnlockOSThread()
}

func BenchmarkLockThread(b *testing.B) {
	for i := 0; i < b.N; i++ {
		runtime.LockOSThread()
		runtime.UnlockOSThread()
	}
}

func BenchmarkLockThreadParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			runtime.LockOSThread()
			runtime.UnlockOSThread()
		}
	})
}

//func BenchmarkSemParallel(b *testing.B) {
//	var s uint32 = 8
//	b.RunParallel(func(pb *testing.PB) {
//		for pb.Next() {
//			sync.Semacquire(&s)
//			sync.Semrelease(&s)
//		}
//	})
//}
//
//func BenchmarkDoSpin(b *testing.B) {
//	b.RunParallel(func(pb *testing.PB) {
//		for pb.Next() {
//			sync.DoSpin()
//		}
//	})
//}