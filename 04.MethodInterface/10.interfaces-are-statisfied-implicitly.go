/*
interfaceはメソッドの集まり。
インターフェイス型の宣言時に指定したメソッドを実装することで
インターフェイスを実装することができる。

下記例だと関数にインターフェイスを実装させている

*/
package main

import "fmt"

type person struct {
	name string
	age  int
}

type i interface {
	age() int
}

func (p person) m() int {
	return p.age
}

func main() {
	i := person{name: "Taro", age: 20}
	fmt.Println(i.m())
}
