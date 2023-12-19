package lru

import "testing"

func TestLru(t *testing.T) {
	lru := NewCache(3)
	lru.Put("key1", 1)
	lru.Put("key2", "value2")
	lru.Put("key3", "value3")
	lru.Print()
	lru.Put("key4", "value4")
	lru.Print()
	lru.Get("key3")
	lru.Print()
}
