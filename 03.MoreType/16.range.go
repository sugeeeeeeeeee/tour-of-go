/*
rangeはsliceやマップを一つずつ反復処理する
sliceをrangeで反復処理する場合、indexとvalueが取得される。
*/
package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
