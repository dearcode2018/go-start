/**
 * 描述: 结构体
 * @filename Struct.go
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
func main() {

	// fmt.Print("")
	var book1 Book
	book1.title = "科学杂志"
	book1.author = "aiyinsitang"
	book1.subject = "science"
	book1.bookId = 1234
	fmt.Println(book1)
}

/**
 *结构体
 */
type Book struct {
	title   string
	author  string
	subject string
	bookId  int
}
