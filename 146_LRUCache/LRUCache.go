package main

import "fmt"

type LRUCache struct {
	size       int
	cap        int
	cache      map[int]*DListNode
	head, tail *DListNode
}

type DListNode struct {
	key, val   int
	prev, next *DListNode
}

func Constructor(capacity int) LRUCache {
	cache := LRUCache{
		size:  0,
		cap:   capacity,
		cache: make(map[int]*DListNode),
		head:  new(DListNode),
		tail:  new(DListNode),
	}
	cache.head.next = cache.tail
	cache.tail.prev = cache.head
	return cache
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.cache[key]; ok {
		this.removeNode(node)
		this.pushFront(node)
		return node.val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.cache[key]; ok {
		node.val = value
		this.removeNode(node)
		this.pushFront(node)
	} else {
		node = &DListNode{
			key: key,
			val: value,
		}
		this.cache[key] = node
		this.pushFront(node)
		this.size++
		if this.size > this.cap {
			this.removeTail()
			this.size--
		}
	}
}

func (this *LRUCache) pushFront(node *DListNode) {
	node.next = this.head.next
	this.head.next.prev = node

	this.head.next = node
	node.prev = this.head
}

func (this *LRUCache) removeNode(node *DListNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (this *LRUCache) removeTail() {
	prev := this.tail.prev
	delete(this.cache, prev.key)
	this.removeNode(prev)
}

func main() {
	lRUCache := Constructor(2)
	lRUCache.Put(1, 1)
	lRUCache.Put(2, 2)
	fmt.Println(lRUCache.Get(1))
	lRUCache.Put(3, 3)
	fmt.Println(lRUCache.Get(2))
	lRUCache.Put(4, 4)
	fmt.Println(lRUCache.Get(1))
	fmt.Println(lRUCache.Get(3))
	fmt.Println(lRUCache.Get(4))
}
