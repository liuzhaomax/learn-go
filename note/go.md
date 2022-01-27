# Golang
1. array和slice
2. Channel
3. 锁
4. 内存栈和堆
5. GC
6. 优雅的main
7. runtime包
8. reflect包
9. context包

## 1. array和slice
|     | array | slice                   |
|:---:|:------|:------------------------|
| 类型  | 值类型   | 引用类型                    |
| 内存  | 存值    | 存切片所切数组的头部地址、len、cap    |
| 传参  | 值传递   | 引用传递                    |
| 扩容  | 不可    | 小于1024扩1/2，大于等于1024扩1/4 |
作为参数传入函数时，array传入的是拷贝，是值传递，slice传入的是指针，是引用传递。<br/>
因此对参数（array）的操作会不改变原array，对参数（slice）的操作会改变原slice <br/>
slice扩容超过其定义的cap，会造成内存逃逸。
```go
arr1 := [3]int{1, 2, 3}
arr2 := [3]int{1, 2, 3}
fmt.Println(arr1 == arr2) // true
```
```go
slice1 := []int{1, 2, 3}
slice2 := []int{1, 2, 3}
fmt.Println(slice1 == slice2) // false
```

## 2. Channel
| 无缓存                       | 有缓存                          |
|:--------------------------|:-----------------------------|
| `ch := make(chan string)` | `ch := make(chan string, 1)` |
| 同步                        | 异步                           |
+ 给没人接收数据的chan发送数据，永久阻塞
+ 从没人发送数据的chan接收数据，永久阻塞
+ 给已经关闭的chan发送数据，panic
+ 从已经关闭的chan接收数据，无缓存，返回nil
+ 从已经关闭的chan接收数据，有缓存，返回零值

## 3. 锁
+ sync.Mutex 互斥锁（悲观锁） `locker := &sync.Mutex{}`
+ sync.RWMutex 读写锁 `locker := &sync.RWMutex{}`
  + 读锁（共享锁） `locker.RLock()` `locker.RUnlock()`
  + 写锁（排他锁） `locker.Lock()` `locker.Unlock()`
+ 自旋锁：自旋锁是指当一个线程在获取锁的时候，如果锁已经被其他线程获取，那么该线程将循环等待，然后不断地判断是否能够被成功获取，知直到获取到锁才会退出循环。
  获取锁的线程一直处于活跃状态，但是并没有执行任何有效的任务，使用这种锁会造成busy-waiting。
  它是为实现保护共享资源而提出的一种锁机制。其实，自旋锁与互斥锁比较类似，它们都是为了解决某项资源的互斥使用。无论是互斥锁，还是自旋锁，在任何时刻，最多只能由一个保持者，也就说，在任何时刻最多只能有一个执行单元获得锁。但是两者在调度机制上略有不同。**对于互斥锁，如果资源已经被占用，资源申请者只能进入睡眠状态。但是自旋锁不会引起调用者睡眠，如果自旋锁已经被别的执行单元保持，调用者就一直循环在那里看是否该自旋锁的保持者已经释放了锁**，“自旋”一词就是因此而得名。
```go
//不支持冲入的自旋锁的实现
type spinLock uint32
func (sl *spinLock) Lock() {
    for !atomic.CompareAndSwapUint32((*uint32)(sl), 0, 1) { //自旋
        runtime.Gosched()
    }
}
func (sl *spinLock) Unlock() {
    atomic.StoreUint32((*uint32)(sl), 0)
}
func NewSpinLock() sync.Locker {
    var lock spinLock
    return &lock
}
```
>更多自旋锁介绍 https://studygolang.com/articles/16480

## 4. 内存栈和堆
内存逃逸（栈->堆）：<br/>
golang程序变量会携带有一组校验数据，用来证明它的整个生命周期是否在运行时完全可知。如果变量通过了这些校验，它就可以在栈上分配。否则就说它 逃逸 了，必须在堆上分配。<br/>
+ **在方法内把局部变量指针返回** 局部变量原本应该在栈中分配，在栈中回收。但是由于返回时被外部引用，因此其生命周期大于栈，则溢出。
+ **发送指针或带有指针的值到 channel 中** 在编译时，是没有办法知道哪个 goroutine 会在 channel 上接收数据。所以编译器没法知道变量什么时候才会被释放。
+ **在一个切片上存储指针或带指针的值** 一个典型的例子就是 []*string 。这会导致切片的内容逃逸。尽管其后面的数组可能是在栈上分配的，但其引用的值一定是在堆上。
+ **slice 的背后数组被重新分配了，因为 append 时可能会超出其容量( cap )** slice 初始化的地方在编译时是可以知道的，它最开始会在栈上分配。如果切片背后的存储要基于运行时的数据进行扩充，就会在堆上分配。
+ **在 interface 类型上调用方法** 在 interface 类型上调用方法都是动态调度的 —— 方法的真正实现只能在运行时知道。想像一个 io.Reader 类型的变量 r , 调用 r.Read(b) 会使得 r 的值和切片b 的背后存储都逃逸掉，所以会在堆上分配。

## 5. GC