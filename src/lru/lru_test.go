package lru

import (
	"fmt"
	"testing"
)

func TestNewLRU(t *testing.T) {
	lru, err := NewLruCache(10)
	if err != nil {
		fmt.Println("err")
	}
	for i := 0; i < 15; i++ {
		lru.Add(i, i)
		for _, v := range lru.Keys() {
			fmt.Print(v.(int), "\t")
		}
		lru.Get(3)
		fmt.Println()
	}
}
