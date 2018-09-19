/*
06.indirectionとは逆に、下記関数は値が指定されているので
ポインタだとコンパイルエラーになる。
*/
package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs()) // メソッドが変数レシーバである場合、変数でも
	fmt.Println(AbsFunc(v))

	p := &Vertex{4, 3}
	fmt.Println(p.Abs()) // ポインタでもどちらのレシーバでも取ることが可能。
	fmt.Println(AbsFunc(*p))
}
