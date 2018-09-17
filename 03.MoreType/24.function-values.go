/*
hypot := func(x, y float64) float64 { return math.Sqrt(x*x + y*y) }
上記のように関数も変数として扱える。
※無名関数(lambda)みたい。

浮動小数点xとyを引数として渡して
x*x+y*yの平方根を求める関数を返す関数をhypot変数として定義する

fmt.Println(hypot(5,12))だと√(5*5+12*12)が帰ってくる
fmt.Println(compute(hypot))だと、上記で定義したhypotを引数として
compute関数にfnで渡して、fn(3,4)をreturnとする。

*/
package main

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}
