# GobyExample


# go并发编程

## Mutex



### Mutex中的data race竞争

```shell
go run -race Mutex.go 
#可以检测出数据竞争
```

+ 1.引入



```go
//count++ 不是一个原子操作，它至少包含几个步骤，比如读取变量 count 的当前值，对这个值加 1，把结果再保存到 count 中。因为不是原子操作，就可能有并发的问题。
//比如，10 个 goroutine 同时读取到 count 的值为 9527，接着各自按照自己的逻辑加 1，值变成了 9528，然后把这个结果再写回到 count 变量。但是，实际上，此时我们增加的总数应该是 10 才对，这里却只增加了 1，好多计数都被“吞”掉了。这是并发访问共享数据的常见错误。
func counterDateRace() {
	var count = 0
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 10000; j++ {
				count++
			}
		}()
	}
	wg.Wait()
	fmt.Println(count)
}
```


+ 2.解决


```go
//使用锁mutex，排他锁，可以解决data race问题
//Mutex 的零值是还没有 goroutine 等待的未加锁的状态，
//所以不需要额外的初始化，直接声明变量（如 var mu sync.Mutex）即可。
func counterNoDateRace() {
	var mu sync.Mutex
	var count = 0
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 10000; j++ {
				mu.Lock()
				count++
				mu.Unlock()
			}
		}()
	}
	wg.Wait()
	fmt.Println(count)
}
```


+ 3.优化


```go
type Counter struct {
	mu    sync.Mutex
	Count int64
}

func counterNoDateRaceStruct() {
	var count Counter
	var wg sync.WaitGroup

	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 10000; j++ {
				count.mu.Lock()
				count.Count++
				count.mu.Unlock()
			}
		}()

	}
	wg.Wait()
	fmt.Println(count.Count)
}
```


+ 4.封装方法


```go
type NewCounter struct {
	CounterType int
	Name        string
	mu          sync.Mutex
	count       int64
}

func (c *NewCounter) Incr() {
	c.mu.Lock()
	c.count++
	c.mu.Unlock()
}

func (c *NewCounter) Count() int64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

func counterNoDateRaceFunc() {
	var count NewCounter
	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 10000; j++ {
				count.Incr()
			}
		}()
	}
	wg.Wait()
	fmt.Println(count.count)
}
```
