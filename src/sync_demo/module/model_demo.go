package module

import "sync"

type Hello struct {
	A int
}

func Say(self *Hello) { self.A++ }

var Pool = sync.Pool{
	New: func() interface{} { return new(Hello) },
}
