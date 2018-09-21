/*
下記例だとTがnilの場合、メソッドはレシーバである(t *T)をnilで呼び出す。
他の言語だとこれはnullポインター例外を引き起こすが、
Goだと適切にハンドリングすることが一般的である。
*/
package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func main() {
	var i I

	var t *T
	i = t
	describe(i)
	i.M()

	i = &T{"hello"}
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
