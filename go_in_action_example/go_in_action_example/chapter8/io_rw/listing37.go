// 这个程序展示来自不同标准库的不同函数是如何使用io.Writer接口的
package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	// 创建要给buffer值，并将一个字符串写入Buffer中
	// 使用io.Writer的Writer方法
	var b bytes.Buffer
	b.Write([]byte("Hello "))
	// 使用Fprintf来将一个字符串拼接到Buffer里
	// 将bytes.Buffer的地址作为io.Writer类型值传入
	fmt.Fprintf(&b, "World!")
	// 将Buffer的内容输出到标准输出设别
	// 将os.File值的地址作为io.Writer类型值传入
	b.WriteTo(os.Stdout)
}