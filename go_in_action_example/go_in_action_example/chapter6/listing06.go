// listing04.go
package main
import (
	"fmt"
	"runtime"
	"sync"
)

// wg用来等待程序完成
var wg sync.WaitGroup

// main函数是程序的入口
func main() {
	// 分配一个逻辑处理器给调度器用
	runtime.GOMAXPROCS(1)

	// 计数加2，表示要等待两个goroutine
	wg.Add(2)

	// 创建两个goroutine
	fmt.Println("Create Goroutines")
	go printPrime("A")
	go printPrime("B")

	// 等待goroutine结束
	fmt.Println("Waiting To Finsh")
	wg.Wait()

	fmt.Println("Terminating Program")
}

// printPrime 显示5000以内的素数
func printPrime(prefix string) {
	defer wg.Done()

next:
	for outer := 2; outer < 5000; outer++ {
		for inner :=2; inner < outer; inner ++ {
			if outer%inner == 0 {
				continue next
			}
		}
		fmt.Printf("%s:%d\n", prefix, outer)
	}
	fmt.Println("Completed", prefix)
}
