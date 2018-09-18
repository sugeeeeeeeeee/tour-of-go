/*
リテラルの要素から推測できる場合は、以下のように宣言を省略可能

省略前
var m = map[string]Vertex{
	"testA" : Vertex{
		1.000, 2.000,
	},
	"testB": Vertex{
		3.000, 4.000,
	}

省略後
var m = map[string]Vertex{
	"testA": {1.000, 2.000},
	"testB": {3.000, 4.000},
}
*/
package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": {1.111, -1.111},
	"Google":    {2.222, -2.222},
}

func main() {
	fmt.Println(m)
}
