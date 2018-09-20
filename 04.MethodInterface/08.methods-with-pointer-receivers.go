/*
# ポインタレシーバを使う理由
1. メソッドがレシーバが指す先の変数を変更するため
	- グローバル変数を少なくしましょが目的
	- 下記例だとmain関数のvはSacleメソッドには渡せない。
2. メソッドの呼び出し毎に変数のコピーを避けるため
	- ポインタレシーバだと変数値のアドレスを渡すので、変数のコピーが走らない
	- 値レシーバだと変数値を渡すので、アドレスが複製される。つまり変数がコピーされる。
	- つまり、オーバーヘッドをすくなくしましょが目的
*/
package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := &Vertex{3, 4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs())
	v.Scale(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.Abs())
}
