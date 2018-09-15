/*
Array = 配列
配列の長さも型となるので長さを変えることはできない。

ex)
var [10]int intの10個の配列
[]int{0,1,2,3,4,5,6,7,8,9] これもintの10個の配列を作成
*/
package main

import "fmt"

func main() {
	// int の2個の配列
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 13}
	fmt.Println(primes)

	ten := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(ten)
}
