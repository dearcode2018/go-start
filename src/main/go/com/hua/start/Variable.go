/**
 * 描述: 变量:全局/局部
 * @filename Variable.go
 * @author qianye.zheng
 * @version 1.0
 */
package main

import "fmt"

/*
 * 1.指定变量类型
 * var var_name type = value
 * 或
 * var var_name type
 * var_name = value
 *
 * 2.自动判断类型
 * var var_name = value
 *
 * 3.省略 var，声明并初始化
 *  var_name := value
 *
 *
 */

/**
 *
 * 描述:
 * @author qianye.zheng
 * @param
 * @return
 */

// 全局变量
var global_var = 10

func main() {

	// 1.指定变量类型
	var var1 int = 1
	fmt.Printf("var1 = %d", var1)
	fmt.Println()
	// 或者
	var var2 int
	var2 = 2
	fmt.Printf("var2 = %d", var2)
	fmt.Println()

	// 2.自动判断类型
	var var3 = 3
	fmt.Printf("var3 = %d", var3)
	fmt.Println()

	// 3.省略 var，声明并初始化
	var4 := 4
	fmt.Printf("var4 = %d", var4)
	fmt.Println()

	// 内存地址
	fmt.Printf("var4 addr = %d", &var4)

	fmt.Println()
	fmt.Printf("global_var = %d", global_var)

	// fmt.Print("")
}
