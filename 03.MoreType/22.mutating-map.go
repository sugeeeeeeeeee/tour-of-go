/*
map mへの値の挿入や更新
 m[key] = value
要素の取得
 elem = m[key]
要素の削除
 dele(m, key)
キーに対する要素の存在確認
 elem, ok = m[key]
 if ok == true {
	 fmt.Println(ok)
 }
*/
package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}
