/*
リテラルの作り方

arrayで記載したが
[4]int{0,1,2,3}
[]int{0,1,2,3}
なので上が配列を作成。下が配列を参照するリテラルを作成。
*/
package main

import "fmt"

func main() {
	primes := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
}
