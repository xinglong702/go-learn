// 1）go以包作为管理单位
// 2）每个文件必须先声明包
// 3）一个工程必须有一个main包
package main

import "fmt"     // I/O包

// go函数能返回多个值
func fun(a, b, c int)  {
	return 1,
}


// 唯一一个入口函数
func main() {     // 左括号必须和函数名在同一行
	fmt.Println("hello Golang!!!")     // Println()会自动换行

	// 只是声明没有初始化的变量，值默认是0
	// 有类型推导，int可以省略，但必须初始化（显而易见）
	var a int = 10
	fmt.Println("a =",a)

	var c = 30
	// 等价于c := 30
	fmt.Println("c =",c)

	// printf是格式化输出的意思,println的ln是换行的意思
	fmt.Printf("%T", c)
}
