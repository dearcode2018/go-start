/**
 * 描述: 常量
 * @filename Constant.go
 * @author qianye.zheng
 * @version 1.0
 */
package main

import "fmt"
import "unsafe"

/**
 * 1.显式类型定义
 * const const_name type = value
 *
 * 2.隐式类型定义
 * const const_name = value
 *
 *
 * iota: 特殊常量，可以被编译器修改的常量，在常量中每出现一次会自增1
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

	// 1.显式类型定义
	const SIZE int = 10
	fmt.Println(SIZE)

	// 2.隐式类型定义
	const NAME = "zheng"
	fmt.Println(NAME)

	const LENGTH = unsafe.Sizeof(NAME)
	fmt.Println(LENGTH)

	/* 0一直自增 */
	const (
		a = iota
		b = iota
		c = iota
	)
	fmt.Println(a, b, c)
	// fmt.Print("")
}
