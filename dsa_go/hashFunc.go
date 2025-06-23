package dsa

import "fmt"

type hashNode struct {
	info any
	next *hashNode
}

type hashMap struct {
	bucket []*hashNode
	size   int
}

func hashFunc(key string) int {
	cons := 31
	mod := 11
	hash_value := 0
	for _, char := range key {
		hash_value = (hash_value*cons + int(char)) % mod
	}
	return hash_value
}
func (h *hashMap) init(size int) {
	h.size = size
	h.bucket = make([]*hashNode, size)
}
func (h *hashMap) insert(key string, value any) {
	hashKey := hashFunc(key)
	node := &hashNode{info: value}
	bucket := h.bucket
	if hashKey >= h.size {
		fmt.Println("Inappropriate Size")
		return
	}
	if bucket[hashKey] != nil {
		head := bucket[hashKey]
		for head.next != nil {
			head = head.next
		}
		head.next = node
	} else {
		bucket[hashKey] = node
	}
}

func (h *hashMap) get(key string, value any) {
	bucket := h.bucket
	hashKey := hashFunc(key)
	hashlist := bucket[hashKey]
	if hashlist == nil {
		fmt.Println("Key is invalid!")
		return
	} else {
		head := hashlist
		if head.info == value {
			fmt.Println(value)
			return
		} else {
			for head.next != nil {
				if head.info == value {
					fmt.Println(value)
					return
				}
				head = head.next
			}
		}
	}
}

func Call() {
	h := &hashMap{}
	h.init(11)
	h.insert("name", "Alice")
	h.insert("age", 25)
	h.get("name", "Alice") // Alice
	h.get("age", 25)       // 25
	h.get("city", "ok")    // Key not found
}
