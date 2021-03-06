// 这个示例程序演示如何使用通道来监视程序运行的时间，以在程序运行时间过长时如何终止程序
package main

import (
	"github.com/goinaction/code/chapter7/patterns/runner"
	"log"
	"os"
	"time"
)

// timeout规定了必须在多少秒内处理完成
const timeout = 3 * time.Second

func main() {
	log.Println("Starting work.")
	// 为本次执行分配超时时间
	r := runner.New(timeout)
	// 要加入执行的任务
	r.Add(createTask(), createTask(), createTask())
	// 执行任务并处理结果
	if err := r.Start(); err != nil {
		switch err {
		case runner.ErrTimeout:
			log.Println("Terminating due to timeout.")
			os.Exit(1)
		case runner.ErrInterrupt:
			log.Println("Terminating due to interrupt.")
			os.Exit(2)
		}
	}
	log.Println("Process ended.")
}

// createTask返回一个根据id休眠指定秒数的任务
func createTask() func(int)  {
	return func(id int) {
		log.Println("Process - Task #%d.", id)
		time.Sleep(time.Duration(id) * time.Second)
	}
}

