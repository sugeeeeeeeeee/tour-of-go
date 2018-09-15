/*
pointer = メモリアドレス = メモリ空間の中にある番地
&でpointerを引き出し、*で実際の値を読み出し、値に代入する。
*/
package main

import "fmt"

func main() {
	i, j := 1, 10
	var p *int

	// 変数iのポインタを読み出す = &i
	p = &i

	// *pで実際の値を読み出す = 変数iが表示される。i
	// なのでpは変数iのpointerが表示される。
	fmt.Println(*p) // output:21
	fmt.Println(p)  // output:0xc00006a010
	*p = 21
	fmt.Println(i)
	fmt.Println(p)

	p = &j
	*p = *p / 5
	fmt.Println(j)
	fmt.Println(p)
}
