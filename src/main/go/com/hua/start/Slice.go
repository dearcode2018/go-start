/**
 * 描述: 切片
 * @filename Sliece.go
 * @author qianye.zheng
 * @version 1.0
 */
package main

import "fmt"

/**
 * 定义切片
 * 1.声明一个未指定大小的数组
 * var slice_name []type
 * slice_name = make([]type, length[, capacity])
 *
 * 2.使用make函数创建切片，容量可省略
 * var slice_name = make([]type, length[, capacity])
 *
 * 3.切片初始化
 * slice_name := []type{member...}
 */

/**
 *
 * 描述:
 * @author qianye.zheng
 * @param
 * @return
 */
func main() {

	//slice2 = make([]int, 4, 1)
	//slice2 = []int{1, 2}
	//fmt.Println(slice2)

	slice3 := []int{1, 2, 6, 0}
	fmt.Println(slice3)
	// fmt.Print("")
}
