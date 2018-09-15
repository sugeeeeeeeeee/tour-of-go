/*
まず、ここを確認する。
https://github.com/golang/tour/blob/master/pic/pic.go

mainで設定されているpicは下記

func Show(f func(int, int) [][]uint8) {
	const (
		dx = 256 <<< dx/dyはpic.Show関数を読み出すと固定値256で設定される。
		dy = 256
	)
	data := f(dx, dy)
	m := image.NewNRGBA(image.Rect(0, 0, dx, dy))
	for y := 0; y < dy; y++ {
		for x := 0; x < dx; x++ {
			v := data[y][x]
			i := y*m.Stride + x*4
			m.Pix[i] = v
			m.Pix[i+1] = v
			m.Pix[i+2] = 255
			m.Pix[i+3] = 255
		}
	}
	ShowImage(m)
}

func Showで定義されている必要な引数は[][]uint8の2重配列。
なので、Pic関数では[][]uint8を出力する関数を作成する。
中身は問わないので自分で好きなようにいじることが可能。
*/
package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	ret := make([][]uint8, dx)

	for x := range ret {
		ret[x] = make([]uint8, dy)
		for y := range ret[x] {
			ret[x][y] = uint8(x * y)
		}
	}
	return ret
}

func main() {
	pic.Show(Pic)
}
