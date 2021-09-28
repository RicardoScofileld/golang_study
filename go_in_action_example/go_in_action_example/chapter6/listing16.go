// 这个示例展示如何使用互斥锁来定义一段需要同步访问的代码临界区资源的同步访问
package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	counter int
	wg      sync.WaitGroup
	// mutex用来定义一段代码临界区
	mutex sync.Mutex
)

func main() {
	wg.Add(2)

	go incCounter(1)
	go incCounter(2)

	wg.Wait()
	fmt.Printf("Final Counter: %d \n", counter)
}

// incCounter使用互斥锁来同步并保证安全访问，增加包里的counter的值
func incCounter(id int) {
	defer wg.Done()

	for count := 0; count < 2; count++ {
		// 同一时刻只允许一个goroutine进入这个临界区
		mutex.Lock()
		{
			// 捕获counter的值
			value := counter
			// 当前goroutine从线程退出，并放回到队列
			runtime.Gosched()
			// 增加本地value变量的值
			value++
			// 将该值保存回counter
			counter = value
		}
		mutex.Unlock()
		// 释放锁，允许其他正在等待的goroutine进入临界区

	}
}
