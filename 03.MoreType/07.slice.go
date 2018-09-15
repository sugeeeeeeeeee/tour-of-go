/*
配列は固定長。sliceは可変長
a[low:high]：low〜highを参照するスライスを作成
*/
package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11}
	fmt.Println(primes)
	var a []int = primes[1:4]
	fmt.Println(a)
}
