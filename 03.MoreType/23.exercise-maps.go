/*
まず、下記を確認すると、
https://go-tour-jp.appspot.com/moretypes/23

文章が与えられて、その出現単語数を確認する。

strings.Filedsは文字列sを1つ以上の連続する空白文字を区切りとして
分割する関数である。

"Hello world !!"→["Hello","world","!!"]

こんな感じ。これをもとに
key:出現単語,value:出現回数となるmapをreturnすれば良い
*/
package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	count := strings.Fields(s)
	ret := map[string]int{}
	for _, v := range count {
		_, ok := ret[v]
		if ok == true {
			ret[v] += 1
		} else {
			ret[v] = 1
		}
	}
	return ret
}

func main() {
	wc.Test(WordCount)
}
