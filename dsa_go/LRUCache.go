package dsa

import "fmt"

type dllNode struct {
	key  int
	info int
	prev *dllNode
	next *dllNode
}

type Cache struct {
	tail     *dllNode
	head     *dllNode
	capacity int
	cache    map[int]*dllNode
}

func createCache(capacity int) *Cache {
	obj := &Cache{
		tail:     nil,
		head:     nil,
		capacity: capacity,
		cache:    make(map[int]*dllNode),
	}
	return obj
}
func (obj *Cache) insertNode(node *dllNode) {
	node.prev = nil
	node.next = obj.head
	if obj.head != nil {
		obj.head.prev = node
	}
	obj.head = node
	if obj.tail == nil {
		obj.tail = node
	}
}

func (obj *Cache) removeNode(node *dllNode) {
	if node.prev != nil {
		node.prev.next = node.next
	} else {
		obj.head = node.next
	}
	if node.next != nil {

		node.next.prev = node.prev
	} else {
		obj.tail = node.prev
	}
}

func (obj *Cache) put(key, data int) {
	if node, ok := obj.cache[key]; ok {
		node.info = data
		obj.removeNode(node)
		obj.insertNode(node)
	} else {
		if len(obj.cache) == obj.capacity {
			node = obj.tail
			obj.removeNode(node)
			delete(obj.cache, node.key)
		}
		node := &dllNode{key: key, info: data}
		obj.insertNode(node)
		obj.cache[key] = node
	}

}

func (obj *Cache) get(key int) int {
	if node, found := obj.cache[key]; found {
		obj.removeNode(node)
		obj.insertNode(node)
		return node.info
	} else {
		return -1
	}
}

func (obj *Cache) displayCache() {
	if obj.head == nil {
		fmt.Print("No data in Cache!!")
	} else {
		temp := obj.head
		for temp != nil {
			fmt.Println(temp)
			temp = temp.next
		}
	}
}

func LRUCache() {
	lru := createCache(2)
	lru.put(1, 1)
	lru.put(2, 4)
	lru.displayCache()
	fmt.Println(lru.get(2))
}
