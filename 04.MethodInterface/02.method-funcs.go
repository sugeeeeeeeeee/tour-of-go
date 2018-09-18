/*
ポインタのメソッドレシーバを使うのは
メソッド呼び出し時に値をコピーによるオーバヘッドを避けるため
structのフィールドの内容を更新するためがある
*/
package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(Abs(v))
}
