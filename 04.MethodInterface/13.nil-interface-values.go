/*
nilインターフェイスは値も具体的な型も保持しない。
具体的なメソッドを占める方がインターフェイスのタプルに存在しないので
ランタイムエラーになる。
*/
package main

import "fmt"

type I interface {
	M()
}

func main() {
	var i I
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
