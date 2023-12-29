package main

import (
	"fmt"
	"module13_task3/internal/lru"
)

func main() {

	cache := lru.NewLRUCache(3)

	cache.Add("key1", "value1")
	cache.Add("key2", "value2")
	cache.Add("key3", "value3")

	value1, ok := cache.Get("key1")
	fmt.Printf("Get value1 res:%v , %v\n", ok, value1)

	value2, ok := cache.Get("key2")
	fmt.Printf("Get value2 res:%v , %v\n", ok, value2)

	value3, ok := cache.Get("key3")
	fmt.Printf("Get value3 res:%v , %v\n", ok, value3)

	cache.Add("key4", "value4")

	value1, ok = cache.Get("key1")
	fmt.Printf("Get value1 res:%v , %v\n", ok, value1)

	value4, ok := cache.Get("key4")
	fmt.Printf("Get value4 res:%v , %v\n", ok, value4)

	cache.Remove("key2")
	value2, ok = cache.Get("key2")
	fmt.Printf("Get value2 res:%v , %v\n", ok, value2)

}
