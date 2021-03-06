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

ちなみにtypeはinterfaceを宣言するだけでなく
型を宣言するために使用する。

*/
package main

import "fmt"

type MyInt int

type describe interface {
	description() string
}

type Square struct {
	edgeLength int
}

type Product struct {
	id    uint
	name  string
	price uint
	PR    PRStatement
}

type PRStatement func() string

func (pr PRStatement) description() string {
	return pr()
}

func (s *Square) description() string {
	return fmt.Sprintf("The lengths of all four sides are equal. [edgeLength: %d]", s.edgeLength)
}

// 組み込み型の別名の型をレシーバにして対象インターフェースのメソッドを定義
func (i MyInt) description() string {
	return fmt.Sprintf("MyInt is actually int. value is %d", i)
}

func printDescription(d describe) {
	fmt.Printf("Description: %s\n", d.description())
}

func main() {

	// 既存の組み込み型にインターフェイスを実装させる
	val1 := MyInt(100)
	printDescription(val1)

	// 構造体にインターフェイスを実装させる
	sq := &Square{edgeLength: 10}
	printDescription(sq)

	// 関数にインターフェイスを実装させる
	p1 := &Product{id: 1, name: "Google", price: 1000}
	p1.PR = func() string {
		return fmt.Sprintf("%d", p1.name)
	}
	printDescription(p1.PR)

}
