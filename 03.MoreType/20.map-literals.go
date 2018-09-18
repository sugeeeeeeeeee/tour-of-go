/*
mapリテラルはstructリテラルに似ているがkeyが必要
stringがキーで、structが値
*/
package main

import "fmt"

type V struct {
	Lat, Long float64
}

var m = map[string]V{
	"test1": V{
		1.111, -1.111,
	},
	"test2": V{
		2.222, -2.222,
	},
}

func main() {
	fmt.Println(m)
	fmt.Println(m["test1"])
}
