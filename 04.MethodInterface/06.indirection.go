/*
ScaleFuncは引数にポインタを渡してあげる必要がある。値渡しだとスケールアップできない。
## 値渡しとポインタ渡しの違い
- 値渡し
	- 渡された変数の値が別のアドレスにコピーされる。
	- なので渡された値を別関数で変更しても、元の値は影響なし。
	- こんな感じ。
```
func main() {
	a := 100 // pointer 0x10414020
}
func test(a int) {
	fmt.Println(a) // pointer 0x10414040
}
```

- ポインタ渡し
	- 渡された変数のアドレス(変数が格納されている場所)が渡される。
	- 渡されたアドレスを通じて元の値を変更できる。
	- こんな感じ。
```
func main() {
	a := 100 // pointer 0x10414020
}
func test(a *int) {
	fmt.Println(a) // pointer 0x10414020
}
```
*/
package main

import "fmt"

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(2)
	ScaleFunc(&v, 10)

	p := &Vertex{4, 3}
	p.Scale(3)
	ScaleFunc(p, 8)

	fmt.Println(v, p)
}
