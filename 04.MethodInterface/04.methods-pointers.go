/*
ポインタレシーバを持つメソッドは、レシーバが指す変数を変更できる。
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

func (v *Vertex) Scale(f float64) { // <- ポインタレシーバを持つメソッド
	v.X = v.X * f // <- vの値を変更できる。この例だと*10している
	v.Y = v.Y * f // <- 同上
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
}
