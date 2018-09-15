/*
struct(構造体)はフィールド=変数の集まり
*/
package main

import "fmt"

type Clock struct {
	h int
	m int
}

func main() {
	fmt.Println(Clock{1, 2})
}
