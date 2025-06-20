package dsa

import "fmt"

type Node struct {
	left  *Node
	data  int
	right *Node
}

// create a Binary Tree
func (n *Node) CreateBinaryTree(val int) {
	if n.data < val {
		if n.right == nil {
			n.right = &Node{data: val}
		} else {
			n.right.CreateBinaryTree(val)
		}
	} else if n.data > val {
		if n.left == nil {
			n.left = &Node{data: val}
		} else {
			n.left.CreateBinaryTree(val)
		}
	}
}

// traverse Binary Tree
func (n *Node) TraverseBinaryTree() {
	if n != nil {
		n.left.TraverseBinaryTree()
		fmt.Println(n.data)
		n.right.TraverseBinaryTree()
	}
}

// search Binary Tree
func (n *Node) SearchBinaryTree(val int) bool {
	if n == nil {
		return false
	}
	if n.data < val {
		return n.right.SearchBinaryTree(val)
	} else if n.data > val {
		return n.left.SearchBinaryTree(val)
	}
	return true
}

func ImplementTree() {
	root := Node{data: 20}
	root.CreateBinaryTree(25)
	root.CreateBinaryTree(12)
	root.CreateBinaryTree(5)
	root.CreateBinaryTree(45)
	root.CreateBinaryTree(10)
	root.CreateBinaryTree(13)
	root.TraverseBinaryTree()
	fmt.Println(root.SearchBinaryTree(10))
	fmt.Println(root.SearchBinaryTree(20))
}
