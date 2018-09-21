/*i
インターフェイスの値は具体的なタプルのように考えることができる。
(value, type)と考えられるが、値がこうはならない。

基底値と同じ名前のメソッドが実行される。

つまり、下記例だと
mainでstring pointerが基底値の場合
func (t *T) M() {}
が実行され
mainで浮動小数点が基底値の場合
func (f F) M() {}
が実行される。

*/
package main

import (
	"fmt"
	"math"
)

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func main() {
	var i I

	i = &T{"Hello"}
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
	fmt.Println(i)
}
