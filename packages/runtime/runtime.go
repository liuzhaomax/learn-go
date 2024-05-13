package runtime

import (
	"fmt"
	"reflect"
	"runtime"
)

func Run() {
	var a = new([]int)
	fmt.Println(reflect.TypeOf(a))
	b := append(*a, 0)
	fmt.Println(b[0])
	fmt.Println("================")
	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.NumGoroutine())
	fmt.Println(runtime.Version())
	fmt.Println(runtime.Stack([]byte("5"), true)) // 获取当前 Goroutine 的调用栈信息
	fmt.Println(runtime.Caller(1))
	fmt.Println(runtime.GOMAXPROCS(2)) // 用于设置或获取当前程序的最大 CPU 核心数
	// 锁定和解锁 Goroutine 到 OS 线程
	// 使用场景
	// 1. 跨平台调用 C 代码：当你在 Go 程序中调用 C 语言函数或使用 CGo 进行 C 语言与 Go 语言的交互时，
	// 通常需要将当前 Goroutine 锁定到操作系统线程上。这是因为在某些情况下，C 函数可能会在特定的 OS 线程上执行，
	// 为了确保与 C 函数的交互正常，需要将 Goroutine 锁定到相应的 OS 线程上。
	// 2. 控制并发执行顺序：在一些特定的情况下，你可能希望在 Goroutine 中手动控制代码的执行顺序，
	// 确保某些操作在特定的 OS 线程上执行。通过使用 LockOSThread() 和 UnlockOSThread() 函数，
	// 可以手动锁定 Goroutine 到特定的 OS 线程上，从而控制代码的执行顺序。
	// 3. 防止 Goroutine 被抢占：在某些情况下，你可能希望防止某个 Goroutine 被 Go 运行时系统抢占，
	// 以确保其能够持续执行一段时间。通过使用 LockOSThread() 函数，可以将 Goroutine 锁定到 OS 线程上，防止其被抢占，
	// 直到调用 UnlockOSThread() 函数将其解锁为止。
	runtime.LockOSThread()
	runtime.UnlockOSThread()
	// 让出线程，使用场景
	// 1. 协作式多任务处理：在一些需要协作式多任务处理的情况下，例如在长时间运行的循环中，可以在适当的时候调用 Gosched() 函数，
	// 让出当前 Goroutine 的执行权限，以避免独占 CPU 资源，从而使其他 Goroutine 有机会执行。
	// 2. 避免 Goroutine 长时间占用 CPU：如果有一些需要长时间运行的任务，为了避免其中的一个 Goroutine 长时间占用 CPU，
	// 可以在适当的时候调用 Gosched() 函数，让出执行权限。
	// 3. 解决死锁问题：在某些情况下，可能会出现 Goroutine 之间的死锁问题，调用 Gosched() 函数有助于打破死锁循环，
	// 让其他 Goroutine 有机会执行，从而使程序恢复正常运行。
	runtime.Gosched() // 让出线程
	runtime.Goexit()  // 终止goroutine
}
