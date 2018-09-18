/*
クラスという概念はないが、struct(構造体)にメソッドを定義できる
メソッドとは、任意の型に関数を紐付けるための機能

https://blog.y-yuki.net/entry/2017/05/05/000000

--
メソッドの定義は以下のように記述します。レシーバーに関する記述をする部分が、関数と異なる箇所になります

func (<レシーバー>) <関数名>([引数]) [戻り値の型] {
    [関数の本体]
}
例えばPersonという構造体を定義して、そのPersonのプロフィールを返すメソッドを定義する場合は以下のような記述をします

type Person struct {
    name string
    age  uint
}

func (p *Person) profile(header string) string {
    return fmt.Sprintf("%s Profile[name: %s, age: %d]", header, p.name, p.age)
}

func (p Person) profile2(header string) string {
    return fmt.Sprintf("%s Profile[name: %s, age: %d]", header, p.name, p.age)
}
--
上記例だと、
func main() {
	p := &Person{name:"Taro",age:"10"}
	fmt.Println(p.profile("header string"))
}
とすることでstructにメソッドを定義した形で利用できる。



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

func (v Vertex) Add() float64 {
	return v.X + v.Y
}

func main() {
	fmt.Println("値レシーバー")
	v := &Vertex{3, 4}
	fmt.Println(v.Abs())

	fmt.Println("ポインタレシーバー")
	vv := &Vertex{3, 4}
	fmt.Println(vv.Add())
}
