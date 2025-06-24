package dsa

import "fmt"

type hashNode struct {
	info string
	next *hashNode
}

type headNode struct {
	head *hashNode
}

type hashMap struct {
	bucket []*headNode
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
	h.bucket = make([]*headNode, size)
}
func (h *hashMap) insert(key, value string) {
	hashKey := hashFunc(key)
	node := &hashNode{info: value}
	bucket := h.bucket
	if hashKey >= h.size {
		fmt.Println("Inappropriate Size")
		return
	}
	if bucket[hashKey] != nil {
		pointer := bucket[hashKey]
		temp := pointer.head
		for temp.next != nil {
			temp = temp.next
		}
		temp.next = node
	} else {
		pointer := &headNode{}
		pointer.head = &hashNode{info: key}
		pointer.head.next = node
		bucket[hashKey] = pointer
	}
}

func (h *hashMap) get(key string, value string) {
	bucket := h.bucket
	hashKey := hashFunc(key)
	pointer := bucket[hashKey]
	if pointer == nil {
		fmt.Println("Key is invalid!")
		return
	} else {
		temp := pointer.head
		if temp.next.info == value {
			fmt.Print(temp.next.info)
		} else {
			for temp.next != nil {
				if temp.info == value {
					fmt.Println(temp.info)
					return
				}
				temp = temp.next
			}
		}
	}
}

func Call() {
	h := &hashMap{}
	h.init(11)
	h.insert("name", "Alice")
	h.insert("name", "Raji")
	h.insert("name", "Harsh")
	h.insert("age", "25")
	fmt.Println("\n")
	h.get("name", "Alice")
	fmt.Println("\n")
	h.get("name", "Raji")
	fmt.Println("\n")
	h.get("name", "Harsh")
	fmt.Println("\n")
	h.get("age", "25")
	fmt.Println("\n")   // 25
	h.get("city", "ok") // Key not found
}
