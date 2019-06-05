package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	res := make(map[string]int)
	for _, c := range strings.Fields(s) {
		_, ok := res[c]
		if ok{
			res[c] +=1
		} else {
			res[c]=1
		}
	}
	return res
}

func main() {
	wc.Test(WordCount)
}