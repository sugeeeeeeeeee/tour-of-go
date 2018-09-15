/*
Mapは多言語でいうと事の連想配列、ハッシュ、辞書などのこと。
マップの初期値は何も指定しなければnil。nilマップはキーを持っておらず、キーを追加することはできない。
*/
package main

import "fmt"

type Vertex struct { // <<< これでstrustを定義、float64 64bitの浮動小数点のフィールドを持つ構造体
	Lat, Long float64
}

func main() {
	m := make(map[string]Vertex) // これでmap型の配列を初期化した
	m["Bell Labs"] = Vertex{
		40.68433, -74.39987,
	}
	m["test"] = Vertex{
		1.1111, -2.222222,
	}
	fmt.Println(m["Bell Labs"])
	fmt.Println(m["Bell Labs"].Lat)
	fmt.Println(m["Bell Labs"].Long)

	fmt.Println(m["test"])
	fmt.Println(m["test"].Lat)
	fmt.Println(m["test"].Long)

	fmt.Println(m)
}
