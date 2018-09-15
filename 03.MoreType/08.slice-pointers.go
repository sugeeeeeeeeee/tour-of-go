/*
スライスは配列を参照しているだけで実際にはデータを持たない。
スライスの要素を変えると参照再起の配列の値が変わる。
*/
package main

import "fmt"

func main() {
	names := [4]string{
		"Tom",
		"Ken",
		"John",
		"George",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "xxx"
	fmt.Println(a, b)
	fmt.Println(names)
}
