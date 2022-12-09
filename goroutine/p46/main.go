package main

import "fmt"

// select有一个default分支和只有一个分支 的块case称为 try-send 或 try-receive 通道操作
// 标准的 Go 编译器对 try-send 和 try-receive select blocks 进行了特殊优化，其执行效率远高于 multi-case select blocks
func main() {
	type Book struct{ id int }
	bookshelf := make(chan Book, 3)

	for i := 0; i < cap(bookshelf)*2; i++ {
		select {
		case bookshelf <- Book{id: i}:
			fmt.Println("succeeded to put book", i)
		default:
			fmt.Println("failed to put book")
		}
	}

	for i := 0; i < cap(bookshelf)*2; i++ {
		select {
		case book := <-bookshelf:
			fmt.Println("succeeded to get book", book.id)
		default:
			fmt.Println("failed to get book")
		}
	}
}
