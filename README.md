# go-learn

---

**Golang优势：**

- 可直接编译成机器码，不依赖其他库，部署就是扔一个文件上去就完了
- 静态类型语言，但有动态语言的感觉，开发效率较高
- 语言层面支持并发
- 内置runtime，支持垃圾回收
- 丰富的标准库
- 内置了很多工具链，最好用的应该是gofmt工具，自动格式化代码
- 跨平台编译
- 内嵌C支持

----

```go
// 1）go以包作为管理单位
// 2）每个文件必须先声明包
// 3）一个工程必须有一个main包
package main

import "fmt"     // I/O包

// 唯一一个入口函数
func main() {     // 左括号必须和函数名在同一行
   fmt.Println("hello Golang!!!")     // Println()会自动换行
}
```

## 一、基本类型、流程控制

### 1.常量和变量

### 2.基本数据类型

### 3.fmt包的格式化输入输出

### 4.类型转换和别名

### 5.运算符

### 6.流程控制

----

## 二、函数、工程管理



## 三、复合类型



## 四、面向对象编程



## 五、异常、文本文件处理



## 六、并发编程-------天生支持



## 七、网络编程、socket编程



## 八、HTTP编程