/*
struct型だけじゃなく、typeに対してもメソッドを定義することが可能。
*/
package main

import (
	"fmt"
	"math"
)

type MyFloat float64

func (f MyFloat) Abs() float64 { // MyFloatが値レシーバ、-π(-3.14159265.....)を取得する。
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())

	f1 := MyFloat(-1.11111111)
	fmt.Println(f1.Abs())

	f2 := MyFloat(2.22222222)
	fmt.Println(f2.Abs())
}
