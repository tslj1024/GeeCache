package main

import (
	"fmt"
	"lru"
)

type String string

func (d String) Len() int {
	return len(d)
}

func main() {
	lru := lru.New(int64(0), nil)
	lru.Add("key1", String("1234"))
	val, _ := lru.Get("key1")
	fmt.Print(val)
}
