package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"unsafe"
)

var x struct {
	a bool
	d *bool
	b int16
	c []int
}

//https://wizardforcel.gitbooks.io/gopl-zh/ch13/ch13-01.html

/**
unsafe.Alignof 函數返迴對應參數的類型需要對齊的倍數. 和 Sizeof 類似, Alignof 也是返迴一個常量表達式, 對應一個常量. 通常情況下布爾和數字類型需要對齊到它們本身的大小(最多8個字節), 其它的類型對齊到機器字大小.
*/

func main() {
	test1()
	fmt.Println("========================")
	test2()
	fmt.Println("========================")
	test3()
}

const mutexLocked = 1 << iota

func test3() {
	var mutex sync.Mutex
	p := (*int32)(unsafe.Pointer(&mutex))
	//直接通過指針操作內存
	atomic.CompareAndSwapInt32(p, 0, mutexLocked)
	fmt.Println(p)

	//uintptr类型类似c的指针
	p2 := (*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mutex)) + 4))
	atomic.CompareAndSwapUint32(p2, 0, mutexLocked)
	fmt.Println(mutex)
}

func test2() {
	//不要試圖引入一個uintptr類型的臨時變量，因爲它可能會破壞代碼的安全性
	// 和 pb := &x.b 等價
	pb := (*int16)(unsafe.Pointer(uintptr(unsafe.Pointer(&x)) + unsafe.Offsetof(x.b)))
	*pb = 42
	fmt.Println(x.b) // "42"
}

func test1() {
	//x sizeof 爲32 *內存對齊*
	fmt.Println("Sizeof ", unsafe.Sizeof(x))
	fmt.Println("Alignof ", unsafe.Alignof(x))
	//fmt.Println("Offsetof ",unsafe.Offsetof(x))
	fmt.Println("========================")
	fmt.Println("Sizeof a ", unsafe.Sizeof(x.a))
	fmt.Println("Alignof a ", unsafe.Alignof(x.a))
	fmt.Println("Offsetof a ", unsafe.Offsetof(x.a))
	fmt.Println("========================")
	fmt.Println("Sizeof b ", unsafe.Sizeof(x.b))
	fmt.Println("Alignof b ", unsafe.Alignof(x.b))
	fmt.Println("Offsetof b ", unsafe.Offsetof(x.b))
	fmt.Println("========================")
	fmt.Println("Sizeof c ", unsafe.Sizeof(x.c))
	fmt.Println("Alignof c ", unsafe.Alignof(x.c))
	fmt.Println("Offsetof c ", unsafe.Offsetof(x.c))
	fmt.Println("========================")
	fmt.Println("Sizeof d ", unsafe.Sizeof(x.d))
	fmt.Println("Alignof d ", unsafe.Alignof(x.d))
	fmt.Println("Offsetof d ", unsafe.Offsetof(x.d))
	fmt.Println("========================")
}
