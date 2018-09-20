/*
# インターフェイス型とは
- メソッドのシグニチャの集まり
	- どのようなメソッドを実装するべきかを規定したもの
- 下記構成。

type <型名> interface {
    メソッド名(引数の型, ...) (返り値の型, ...)
    ・
    ・
}

- 下記例だと
	- Myfloatをtypeでinterfaceとして定義する。
	- mainで-√(2)をMyfloatでfへ代入する。
	-


*/
package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f  // a MyFloat implements Abser
	a = &v // a *Vertex implements Abser

	// In the following line, v is a Vertex (not *Vertex)
	// and does NOT implement Abser.
	//a = v

	fmt.Println(a.Abs())
}

type MyFloat float64

func (f float64) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

/*
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
*/
