/*
structのフィールドのポインタもドットでアクセス可能
(*p).vでも、p.sでもアクセス可能
*/
package main

import "fmt"

type Clock struct {
	h int
	m int
}

func main() {
	v := Clock{1, 2}
	p := &v
	p.h = 1e9
	fmt.Println(v)
	(*p).m = 100
	fmt.Println(v)
}
