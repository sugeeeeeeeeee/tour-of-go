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
