/*
04のScaleとAbsを、メソッドから関数に書き直したもの

func (v Vertex) Abs() float64 {
}
↓
func Asv(v Vertex) f float64 {
}

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

// func Scale(v *Vertex, f float64) {
func Scale(v Vertex, f float64) {
	// 上の*Vertexから*を削除し、ポインタから変数に変更して引数に渡す。
	// prog.go:23:8: cannot use &v (type *Vertex) as type Vertex in argument to Scale

	// となりコンパイルエラーとなる。
	// mainで呼び出してるScale関数でポインタを引数で渡しているため。
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	// Scale(&v, 10)
	// ここをポインタから変数に変えることでコンパイルは通るが意図した動き(10でスケール)にはならない。
	Scale(v, 10)
	fmt.Println(Abs(v))
}
