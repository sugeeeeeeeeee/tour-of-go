/*
構造体のフィールドはドットでアクセス可能
書式：構造体.変数
*/
package main

import "fmt"

type Clock struct {
	h int
	m int
}

func main() {
	v := Clock{1, 2}
	fmt.Println(v)
	v.h = 10
	fmt.Println(v.h)
}
