/*
rangeで取得するindexやvalueは"_"で捨てることができる。
valueを省略すると、indexのみが取得できる。
indexのみの省略はできない。
*/
package main

import "fmt"

func main() {
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i)
	}
	for _, v := range pow {
		fmt.Printf("%d\n", v)
	}
	for i := range pow {
		fmt.Printf("%d\n", i)
	}
}
