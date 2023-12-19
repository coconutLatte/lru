package lru

import "fmt"

type Element struct {
	Key   string
	Value interface{}
}

func (e Element) String() string {
	return fmt.Sprintf("Key: %s, Value: %v", e.Key, e.Value)
}

type Node struct {
	Ele  Element
	prev *Node
	next *Node
}

type Cache struct {
	root     *Node // head
	keys     map[string]struct{}
	capacity int
	size     int
}

func NewCache(capacity int) *Cache {
	lru := &Cache{
		root:     &Node{},
		capacity: capacity,
		keys:     make(map[string]struct{}),
	}

	return lru
}

func (lru *Cache) Get(key string) (interface{}, bool) {
	root := lru.root
	if root == nil {
		return "", false
	}

	_, exist := lru.keys[key]
	if !exist {
		return "", false
	}

	newHeadNode := &Node{
		Ele: Element{
			Key: key,
		},
		prev: root,
		next: root.next,
	}
	if root.next != nil {
		root.next.prev = newHeadNode
	}
	root.next = newHeadNode

	curr := newHeadNode.next
	for curr != nil {
		if curr.Ele.Key == key {
			break
		}
		curr = curr.next
	}
	newHeadNode.Ele.Value = curr.Ele.Value
	if curr.prev != nil {
		curr.prev.next = curr.next
	}
	if curr.next != nil {
		curr.next.prev = curr.prev
	}

	return curr.Ele.Value, true
}

func (lru *Cache) Put(key string, value interface{}) {
	root := lru.root
	if root == nil {
		return
	}

	// make an empty node to head
	newHeadNode := &Node{
		Ele: Element{
			Key:   key,
			Value: value,
		},
		prev: root,
		next: root.next,
	}
	if root.next != nil {
		root.next.prev = newHeadNode
	}
	root.next = newHeadNode

	_, exist := lru.keys[key]
	if exist {
		// delete the old one, search after head
		curr := newHeadNode.next
		for curr != nil {
			if curr.Ele.Key == key {
				break
			}
			curr = curr.next
		}
		if curr.prev != nil {
			curr.prev.next = curr.next
		}
		if curr.next != nil {
			curr.next.prev = curr.prev
		}

	} else {
		lru.keys[key] = struct{}{}
		lru.size++

		if lru.size > lru.capacity {
			curr := root
			for curr.next != nil {
				curr = curr.next
			}
			if curr.prev != nil {
				curr.prev.next = nil
			}

			delete(lru.keys, curr.Ele.Key)
			lru.size--
		}
	}
}

func printChain(head *Node) {
	curr := head.next

	for curr != nil {
		fmt.Println(curr.Ele)
		curr = curr.next
	}
}

func (lru *Cache) Print() {
	printChain(lru.root)
	println()
}
