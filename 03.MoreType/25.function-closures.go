/*
関数はclosureである
関数が変数として代入されたスコープの関数外の変数を使える

http://dqn.sakusakutto.jp/2009/01/javascript_5.html


メリット
関数終了後もローカル変数を参照できることのメリット１
⏩グローバル変数の節減

関数終了後もローカル変数を参照できることのメリット２
⏩クロージャを使うと１のメリットを享受しつつ計算量を増やさないことができる

*/
package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}
