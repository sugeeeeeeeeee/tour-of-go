/*
Structへの初期値の割り当て方は変数定義と同じように{}で割り当てるが
&をつけるとポインタを戻す
*/
package main

import "fmt"

type Clock struct {
	h, m int
}

var (
	c1 = Clock{1, 2}

	// ↓は構造体Clockのhに1を入れて、mは指定しない。
	c2 = Clock{h: 1}
	c3 = Clock{}
	p  = &Clock{1, 2}
)

func main() {
	fmt.Println(c1, c2, c3, p)
}
