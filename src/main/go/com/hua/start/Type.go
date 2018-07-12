/**
 * 描述: 数据类型: 数组、指针、转换
 * @filename Type.go
 * @author qianye.zheng
 * @version 1.0
 */
package main

import "fmt"

/**
 * 数组
 * 1.声明指定大小的数组
 * var array_name [SIZE]type
 *
 * 2.声明并初始化
 * var array_name = [SIZE]type {member ...}
 *
 */

/**
 *
 * 描述:
 * @author qianye.zheng
 * @param
 * @return
 */
func main() {

	var arr1 = [5]int{1, 3, 4, 0, -1}
	fmt.Println(arr1)
	for i := 0; i < 5; i++ {
		fmt.Println(arr1[i])
	}

	var str = "hhe"
	// 指针 var pointer_name *type
	var p1 *string = &str
	fmt.Println(p1)

	// fmt.Print("")
}
