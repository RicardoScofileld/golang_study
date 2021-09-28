// listing07.go
package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	// 分配两个逻辑处理器给调度器使用
	runtime.GOMAXPROCS(2)

	// wg用来等待程序完成
	// 计数加2，表示要等待两个goroutine
	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("Start Goroutines")

	// 声明一个匿名函数，并创建一个goroutines
	go func() {
		defer wg.Done()

		// 显示字母表3次
		for count := 0; count < 3; count++ {
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c", char)
			}
		}
	}()

	go func() {
		defer wg.Done()
		for count := 0; count < 3; count++ {
			for char := 'A'; char < 'A' + 26; char++ {
				fmt.Printf("%c", char)
			}
		}
	}()

	fmt.Println("Waiting to Finsh")
	wg.Wait()
	fmt.Println("\nTerminating Program")
}
