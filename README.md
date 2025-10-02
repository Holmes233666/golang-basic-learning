# golang-basic-learning

## 一、Golang简介

go语言的优势：

（1）极其简单的部署方式：① 可以直接编译为机器码，操作系统可以直接通过`./`执行代码； ② 不依赖其他库；③ 直接运行即可部署

（2）静态类型语言：编译的时候检查出来大多数的问题：`go build`

（3）语言层面的并发：天生的基因支持；充分利用多核

（4）强大的标准库：runtime系统调度机制（垃圾回收，调度机制）；高效的GC垃圾回收（GC1.8后增加了三色标记和混合写屏障）；丰富的标准库

（5）简单易学：25个关键字，C语言简洁的基因，内嵌C语法支持；面向对象特征（封装，继承，多态）；跨平台语言

![image-20250705161117304](https://cdn.jsdelivr.net/gh/Holmes233666/blogImage/images/image-20250705161117304.png)

Go语言的安装：

go path查询：`go env GOPATH`

![image-20250926183650024](https://cdn.jsdelivr.net/gh/Holmes233666/blogImage/images/image-20250926183650024.png)

bin：编译结果

Pkg：源码包

Src：自定义项目包地址

## 二、Golang基础语法

### 从main函数初见golang语法

编译 + 运行： `go run hello.go`

![image-20250705170411643](https://cdn.jsdelivr.net/gh/Holmes233666/blogImage/images/image-20250705170411643.png)

本main.go 中隐含的点：

```go
package main // 程序的包名，含有main函数的包一定属于main包

// 导包方式
import (
	"fmt"
	"time"
)

//import "fmt"
//import "time"

func main() { // 函数的函数名后面的 { 一定和函数名是在同一行的
	// go语言中的表达式加 “;” 和 不加 “;” 都可以，建议不加
	fmt.Println("Hello World")

	time.Sleep(1 * time.Second)
}
```



### 四种声明的方式



```go
package main

import "fmt"

// 声明全局变量，方法1-3是可以的

var g_A int = 100
var g_B = 200

// := 只能用在函数体内声明 + 定义一个变量
//g_C := 300， 会报错
/*
四种变量的声明方式
*/
func main() {
	// -------------------单变量的声明----------------------
	// 方法1：声明一个变量，默认的值是 0
	var a int
	fmt.Printf("type of a: %T\n", a)

	// 方法2：声明一个变量，给出初始化值
	var b int = 100
	fmt.Printf("type of b: %T\n", b)

	// 方法3：在初始化的时候省去对变量类型的说明，根据值自动匹配当前的数据类型
	var c = 1000
	fmt.Printf("type of c: %T\n", c)

	var test_str string = "hello world"
	fmt.Printf("type of %s: %T\n", test_str, test_str)

	var test_str2 = "hello world"
	fmt.Printf("type of %s: %T\n", test_str2, test_str2)

	// 方法4：【最常用的写法】省去var关键字，直接自动匹配
	e := 100 // 既声明， 又初始化
	fmt.Printf("type of %s: %T\n", e, e)

	g := 3.14
	fmt.Printf("type of %s: %T\n", g, g)

	s := "string s"
	fmt.Printf("type of %s: %T\n", s, s)

	fmt.Println("g_A = ", g_A, "g_B = ", g_B)

	// -------------------多变量的声明----------------------
	var aa, bb int = 1, 2
	fmt.Println("aa = ", aa, "bb = ", bb)
	var num, str = 100, "hello world"
	fmt.Println("num = ", num, "str = ", str)

	// 多行的多变量声明
	var (
		am = 1
		bm = "boolean type"
	)
	fmt.Println("am = ", am, "bm = ", bm)
}

```



### 常量与iota

```go
package main

import "fmt"

// const定义枚举类型
const (
	BEIJING  = 1
	NANJING  = 0
	SHENGHAI = 2
)

// 可以在const()添加一个关键字 iota，每行的iota都会累加1，第一行iota默认值是0

const (
	BANANA = 10 * iota
	WATERMALEN
	APPLE
)

const (
	aa, bb = iota + 1, iota + 2
	cc, dd
	ee, ff
	gg, hh = iota + 5, iota + 7
	ll, mm
)

func main() {
	// 常量的定义（只读属性）
	const len int = 3
	fmt.Println(len)
	// len = 10， 常量是不允许修改的

	fmt.Println("BANANA = ", BANANA, "WATERMALEN = ", WATERMALEN, "APPLE = ", APPLE)

	// iota只能配合const一起使用
	// var a int = iota

}

```



### 函数

#### 多返回值

单个返回值的写法：

![image-20250705210259895](/Users/sunyunqi/Library/Application Support/typora-user-images/image-20250705210259895.png)

多个返回值的写法：注意起名的多参数每个参数只是该函数的形参，离开函数体不再能访问。

```go
// 多返回值，匿名
func tool2(a string, b int) (int, int) {
	fmt.Println("--------tool2--------")
	fmt.Println("a:", a, "b:", b)
	c := b + 100
	d := b + 1000
	return c, d
}

// 多返回值，起名
func tool3(a string, b int) (r1 int, r2 int) {
	fmt.Println("--------tool3--------")
	fmt.Println("a:", a, "b:", b)
	fmt.Println("r1, r2是属于tool3的形参，默认初始化值是0")
	fmt.Println("r1:", r1, "r2:", r2)
	r1 = b + 1000
	r2 = b + 100
	return
}

// 多返回值，相同类型合并
func tool4(a string, b int) (r1, r2 int) {
	fmt.Println("--------tool4--------")
	fmt.Println("a:", a, "b:", b)
	r1 = b + 1000
	r2 = b + 100
	return
}
```

#### init函数与导包

Import导入包的时候的执行流程如下：

![image-20250706143118131](https://cdn.jsdelivr.net/gh/Holmes233666/blogImage/images/image-20250706143118131.png)

注意，向外提供的函数，函数名一定要大写。

导包了一定要使用，否则会报错：

![image-20250706200343719](https://cdn.jsdelivr.net/gh/Holmes233666/blogImage/images/image-20250706200343719.png)

如果只希望执行`init()`，但是不希望执行接口函数，应该如何操作？：`_ `别名：

![image-20250706201110510](https://cdn.jsdelivr.net/gh/Holmes233666/blogImage/images/image-20250706201110510.png)

### 指针

需要注意指针的“*”和C/C++的定义位置正好相反

<img src="https://cdn.jsdelivr.net/gh/Holmes233666/blogImage/images/image-20250706205656545.png" alt="image-20250706205656545" style="zoom:33%;" /><img src="https://cdn.jsdelivr.net/gh/Holmes233666/blogImage/images/image-20250706210730272.png" alt="image-20250706210730272" style="zoom:10%;" />



golang是支持二级指针的。

### defer关键字

defer：在函数体结束后执行defer后的语句，相当于析构函数/finnally；

- 如果有多个defer语句，按照栈的形式执行

  ![image-20250706211736777](https://cdn.jsdelivr.net/gh/Holmes233666/blogImage/images/image-20250706211736777.png)

- defer和return语句的先后：先执行`return`，再执行`defer`

  ```go
  import "fmt"
  
  func test() int {
  	fmt.Println("return test")
  	return 0
  }
  
  func returnAndDefer() int {
  	defer fmt.Println("main end1")
  	return test()
  }
  
  func main() {
  	returnAndDefer()
  }
  
  // output：
  return test
  main end1
  ```



### 切片：slice

#### 数组

在go中[10]int和[4]int是两种数据类型，在方法调用时，不同长度的数组不能使用相同的接口调用，并且函数内部是值传递，即在函数体内对数组内容进行修改，不影响原数组的内容。更建议使用动态数组

```go
package main

import "fmt"

func printArray(nums [10]int) { // 值传递：这一点和C/C++不一样了！
	nums[0] = 100
	for _, num := range nums {
		fmt.Println(num)
	}
}

func main() {
	// 固定长度的数组
	var nums [10]int

	nums2 := [10]int{1, 2, 3, 4}

	var nums3 = [4]int{1, 2, 3, 4}

	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}

	for idx, num := range nums2 {
		fmt.Println(idx, num)
	}

	// 查看数组的数据类型
	fmt.Printf("nums数组的数据类型为：%T\n", nums[0])
	fmt.Printf("nums2数组的数据类型为：%T\n", nums2[0])
	fmt.Printf("nums3数组的数据类型为：%T\n", nums3[0])

	// 尝试用函数printArray打印nums3
	printArray(nums)
	// printArray(nums3) // 报错： cannot use nums3 (variable of type [4]int) as [10]int value in argument to printArray
	fmt.Printf("out of printArry")
	for idx, num := range nums {
		fmt.Println(idx, num)
	}
}
```


#### 动态数组：slice

##### slice的声明方式

slice一共有4种声明方式，

```go
package main

import (
	"fmt"
)

func main() {
	// slice的声明方式：
	// 方式1：声明 + 定义
	slice1 := []int{1, 2, 3, 4, 5}
	fmt.Println("当前slice1的长度为", len(slice1))

	// 方式2：只声明
	slice2 := []int{}
	slice2 = make([]int, 3)    // 开辟空间
	slice2 = append(slice2, 1) // 追加一个元素
	fmt.Println("当前slice2的长度为", len(slice2))

	// 方式3：声明 + 初始化 [方式1 + 方式2 的合并方式]
	var slice3 = make([]int, 4)
	fmt.Println(slice3)

	// 方式4：直接使用推断
	slice4 := make([]int, 5)
	fmt.Println(slice4)

	var slice5 []int
	if slice5 == nil {
		fmt.Println("slice5 is nil")
	} else {
		fmt.Println("slice5 is not nil")
	}

}

// 结果输出为：
当前slice1的长度为 5
当前slice2的长度为 4
[0 0 0 0]
[0 0 0 0 0]
slice5 is nil
```



##### 元素的追加和切片

（1）元素的追加

<img src="https://cdn.jsdelivr.net/gh/Holmes233666/blogImage/images/image-20250706225036468.png" alt="image-20250706225036468" style="zoom:30%;" />

- `len`和`cap`的不同

len表示目前有效的元素有多少个，cap表示可以装的容量。初始化一个slice时可以用`make([]int, 3, 5)`的形式分别设置len和cap，如果使用`make([]int, 3)`，那么只有一个len和cap相同。

- 元素的追加

元素的追加使用`append(nums, 1)`的形式，向有效元素后追加一个元素；如果超过cap，那么会先扩容（目前观察到的是扩容为原来容量的2倍）

![image-20250706235642923](https://cdn.jsdelivr.net/gh/Holmes233666/blogImage/images/image-20250706235642923.png)

（2）元素的切片

<img src="https://cdn.jsdelivr.net/gh/Holmes233666/blogImage/images/image-20250707002146194.png" alt="image-20250707002146194" style="zoom:53%;" />



### 哈希表map



```go
package main

import "fmt"

func main() {
	// 声明myMap为key为int，value为string
	var myMap map[int]string // key是int, value是string
	if myMap == nil {
		fmt.Println("map is nil")
	}
	// 使用map前必须得初始化空间大小
	myMap = make(map[int]string, 10)
	myMap[1] = "hello"
	myMap[2] = "world"

	fmt.Println("myMap len is ", len(myMap), ", the content is", myMap)

	// 第二种声明方式
	myMap2 := make(map[string]string, 10)
	myMap2["hello"] = "world"
	fmt.Println("myMap2 len is ", len(myMap2), ", the content is", myMap2)

	// 第3种初始化方式
	myMap3 := map[string]string{
		"one": "cpp",		// 注意要有括号
	}
	fmt.Println("myMap3 len is ", len(myMap3), ", the content is", myMap3)

}
```






