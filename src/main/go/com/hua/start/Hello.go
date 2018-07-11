/**
 * 描述:
 * @filename Hello.go
 * @author qianye.zheng
 * @version 1.0
 */
package main

import "fmt"

/**
 *
 * 描述:
 * @author qianye.zheng
 * @param
 * @return
 */

/* 别名 */
//import f "fmt"

/* 全局变量 */
var age int

func main() {
	//a := 10

	const PI = 3.14
	const c_a, c_b = 1, 2
	/* 没有括号 */
	for i := 1; i < 2; i++ {
		fmt.Println("i = ", i)
	}
	fmt.Printf("hello world go \n")
	fmt.Println("max = ", max(1, 2))

	var a, b = swap("a", "b")
	fmt.Println("swap = ", a, b)
	var point *string
	point = &a
	fmt.Println("pointer = ", point)
	//var blances [10]float32
	var book1 Book
	book1.auth = "zheng"
	book1.book_id = 1234
	book1.subject = "science"
	book1.title = "about science"

}

/**
 * 获取2个值中的最大值
 */
func max(num1, num2 int) int {
	var result int = 0
	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}

	return result
}

/**
 *
 */
func swap(x, y string) (string, string) {
	return y, x
}

/**
 *
 */
type Book struct {
	title   string
	auth    string
	subject string
	book_id int
}



fun testSlice() {
	// 切片声明
	var slice1 []int;
	/* 使用make来声明切片 make([]type, initLength, capacity) */
	var slice2 []int = make([]int, 10, 1)
 	// 初始化
 	slice3 := []int {1, 2, 3}
 	/* 使用数组初始化 var slice := arr[startIndex:endIndex(不含)] */
 	var arr [5]int = {2, 5, 1, 0, 6}
 	/* 从 1-3 */
 	slice4 := arr[1:4]
 	 /* 1 到 最后一个元素*/
 	slice5 := arr[1:]
 	/* 第一个元素 - 3 */
 	slice6 := arr[:4]
 	fmt.Printf("length = %d, capacity = %d", len(slice6), cap(slice6))
 	/* 空切片: 切片在初始化 */
 	var slice7 []int

 	/* 截取切片 slice[:] */
 	fmt.Print(slice6[0:2])

 	/* 向切片追加元素 */
 	slice7 = append(slice7, 34)
 	/* 追加多个元素 */
 	slice7 = append(slice7, 2, 12, 30)
 	/* 拷贝切片 */
 	var slice8 []int
 	/* 将slice7 拷贝到 slice8 */
 	copy(slice8, slice7)
}

















