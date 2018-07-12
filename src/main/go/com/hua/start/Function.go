/**
 * 描述: 递归/函数
 * @filename Function.go
 * @author qianye.zheng
 * @version 1.0
 */
package main

import "fmt"

/**
 *
 * func function_name(param list) (return_types) {
 * 	// function body
 * }
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

	fmt.Println(max(1, 2))
	fmt.Println(swap("a", "b"))
	// fmt.Print("")
}

/**
 * 最大值
 *
 */
func max(num1 int, num2 int) int {
	if num1 > num2 {
		return num1
	} else {
		return num2
	}
}

/**
 * [swap1 description]
 * @param  {[type]} x string        [description]
 * @param  {[type]} y string)       (string,      string [description]
 * @return {[type]}   [description]
 */
func swap(x string, y string) (string, string) {
	return y, x

}
