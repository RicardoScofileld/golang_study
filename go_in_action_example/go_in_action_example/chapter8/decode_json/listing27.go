// 这个示例程序展示如何解码JSON字符串
package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// Contact结构代表我们的JSON字符串
type Contact struct {
	Name    string `json:"name"`
	Title   string `json:"title"`
	Contact struct {
		Home string `json:"home"`
		Cell string `json:"cell"`
	} `json:"contact"`
}

// JSON包含用于反序列化的演示字符串
var JSON = `{
	"name": "Gopher",
	"title": "programmer",
	"contact": {
		"home": "415.33.33",
		"cell": "153.43.43"
}
}`

func main() {
	// 将JSON字符串反序列化到变量
	var c Contact
	err := json.Unmarshal([]byte(JSON), &c)
	if err != nil {
		log.Println("ERROR: ", err)
		return
	}
	fmt.Println(c)
}
