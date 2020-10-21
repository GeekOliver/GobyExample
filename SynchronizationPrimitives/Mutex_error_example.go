package main

import (
	"fmt"
	"sync"
	"time"
)

/*
介绍四大易错mutex场景
*/

//
//1.Lock/Unlock 不是成对出现
//

func example1() {
	var mu sync.Mutex
	defer mu.Unlock()
	fmt.Println("hhh")
}

// hhh
// fatal error: sync: unlock of unlocked mutex

// goroutine 1 [running]:
// runtime.throw(0x4c56cc, 0x1e)
//         /usr/local/go/src/runtime/panic.go:774 +0x72 fp=0xc00006ae50 sp=0xc00006ae20 pc=0x4297c2
// sync.throw(0x4c56cc, 0x1e)
//         /usr/local/go/src/runtime/panic.go:760 +0x35 fp=0xc00006ae70 sp=0xc00006ae50 pc=0x429745
// sync.(*Mutex).unlockSlow(0xc000014108, 0xc0ffffffff)
//         /usr/local/go/src/sync/mutex.go:196 +0xd6 fp=0xc00006ae98 sp=0xc00006ae70 pc=0x4657d6
// sync.(*Mutex).Unlock(0xc000014108)
//         /usr/local/go/src/sync/mutex.go:190 +0x48 fp=0xc00006aeb8 sp=0xc00006ae98 pc=0x4656e8
// main.example1()
//         /root/code/goCode/src/GobyExample/SynchronizationPrimitives/Mutex_error_example.go:17 +0xcb fp=0xc00006af50 sp=0xc00006aeb8 pc=0x48cf6b
// main.main()
//         /root/code/goCode/src/GobyExample/SynchronizationPrimitives/Mutex_error_example.go:21 +0x20 fp=0xc00006af60 sp=0xc00006af50 pc=0x48cfc0
// runtime.main()
//         /usr/local/go/src/runtime/proc.go:203 +0x21e fp=0xc00006afe0 sp=0xc00006af60 pc=0x42b15e
// runtime.goexit()
//         /usr/local/go/src/runtime/asm_amd64.s:1357 +0x1 fp=0xc00006afe8 sp=0xc00006afe0 pc=0x453591
// exit status 2
// You have new mail in /var/mail/root

//
//2.Copy 已使用的 Mutex
//

type Counter1 struct {
	sync.Mutex
	Count int
}

func example2() {
	var c Counter1
	c.Lock()
	defer c.Unlock()
	c.Count++
	foo(c) // 复制锁
}

// 这里Counter的参数是通过复制的方式传入的
func foo(c Counter1) {
	c.Lock()
	defer c.Unlock()
	fmt.Println("in foo")
}

//调用 foo 函数的时候，调用者会复制 Mutex 变量 c 作为 foo 函数的参数，不幸的是，复制之前已经使用了这个锁，这就导致，复制的 Counter 是一个带状态 Counter。

// fatal error: all goroutines are asleep - deadlock!

// goroutine 1 [semacquire]:
// sync.runtime_SemacquireMutex(0xc000014134, 0xc00006ae00, 0x1)
//         /usr/local/go/src/runtime/sema.go:71 +0x47
// sync.(*Mutex).lockSlow(0xc000014130)
//         /usr/local/go/src/sync/mutex.go:138 +0xfc
// sync.(*Mutex).Lock(...)
//         /usr/local/go/src/sync/mutex.go:81
// main.foo(0x1, 0x1)
//         /root/code/goCode/src/GobyExample/SynchronizationPrimitives/Mutex_error_example.go:69 +0x145
// main.example2()
//         /root/code/goCode/src/GobyExample/SynchronizationPrimitives/Mutex_error_example.go:64 +0x9c
// main.main()
//         /root/code/goCode/src/GobyExample/SynchronizationPrimitives/Mutex_error_example.go:76 +0x20
// exit status 2

//
//3.重入
//

//mutex是不允许重入的锁
//当一个线程获取锁时，如果没有其它线程拥有这个锁，那么，这个线程就成功获取到这个锁。之后，如果其它线程再请求这个锁，就会处于阻塞等待的状态。但是，如果拥有这把锁的线程再请求这把锁的话，不会阻塞，而是成功返回，所以叫可重入锁（有时候也叫做递归锁）。只要你拥有这把锁，你可以可着劲儿地调用，比如通过递归实现一些算法，调用者不会阻塞或者死锁。

func foo1(l sync.Locker) {
	fmt.Println("in foo")
	l.Lock()
	bar1(l)
	l.Unlock()
}

func bar1(l sync.Locker) {
	l.Lock()
	fmt.Println("in bar")
	l.Unlock()
}

func example3() {
	l := &sync.Mutex{}
	foo1(l)
}

// in foo
// fatal error: all goroutines are asleep - deadlock!

// goroutine 1 [semacquire]:
// sync.runtime_SemacquireMutex(0xc00001410c, 0x557c00, 0x1)
//         /usr/local/go/src/runtime/sema.go:71 +0x47
// sync.(*Mutex).lockSlow(0xc000014108)
//         /usr/local/go/src/sync/mutex.go:138 +0xfc
// sync.(*Mutex).Lock(0xc000014108)
//         /usr/local/go/src/sync/mutex.go:81 +0x47
// main.bar1(0x4db780, 0xc000014108)
//         /root/code/goCode/src/GobyExample/SynchronizationPrimitives/Mutex_error_example.go:111 +0x35
// main.foo1(0x4db780, 0xc000014108)
//         /root/code/goCode/src/GobyExample/SynchronizationPrimitives/Mutex_error_example.go:105 +0xa5
// main.example3(...)
//         /root/code/goCode/src/GobyExample/SynchronizationPrimitives/Mutex_error_example.go:119
// main.main()
//         /root/code/goCode/src/GobyExample/SynchronizationPrimitives/Mutex_error_example.go:125 +0x3d
// exit status 2

//
//4.死锁
//
func example4() {
	// 派出所证明
	var psCertificate sync.Mutex // 物业证明
	var propertyCertificate sync.Mutex
	var wg sync.WaitGroup
	wg.Add(2) // 需要派出所和物业都处理
	// 派出所处理goroutine
	go func() {
		defer wg.Done() // 派出所处理完成
		psCertificate.Lock()
		defer psCertificate.Unlock() // 检查材料
		time.Sleep(5 * time.Second)  // 请求物业的证明
		propertyCertificate.Lock()
		propertyCertificate.Unlock()
	}()
	// 物业处理goroutine
	go func() {
		defer wg.Done() // 物业处理完成
		propertyCertificate.Lock()
		defer propertyCertificate.Unlock() // 检查材料
		time.Sleep(5 * time.Second)        // 请求派出所的证明
		psCertificate.Lock()
		psCertificate.Unlock()
	}()
	wg.Wait()
	fmt.Println("成功完成")
}

func main() {
	example4()
}
