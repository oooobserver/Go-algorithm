package structure

import "fmt"

/*
Typical binary tree:
			7
		3
			6
	1
			5
		2
			4
*/

type BinaryNode struct {
	item  int
	left  *BinaryNode
	right *BinaryNode
}

type BinaryTree struct {
	root *BinaryNode
}

func SubTreeHeight() {

}

// Create a test binary tree
func newTestBinaryTree() BinaryTree {
	bt := BinaryTree{}
	nodes := make([]BinaryNode, 0)
	for i := 0; i < 8; i++ {
		nodes = append(nodes, BinaryNode{i, nil, nil})
	}
	nodes[1].left = &nodes[2]
	nodes[1].right = &nodes[3]
	nodes[2].left = &nodes[4]
	nodes[2].right = &nodes[5]
	nodes[3].left = &nodes[6]
	nodes[3].right = &nodes[7]
	bt.root = &nodes[1]
	return bt
}

func newTestBST() BinaryTree {
	bt := BinaryTree{}

	nodes := make([]BinaryNode, 0)
	for i := 1; i <= 8; i++ {
		nodes = append(nodes, BinaryNode{i, nil, nil})
	}

	nodes[4].left = &nodes[2]
	nodes[4].right = &nodes[6]
	nodes[2].left = &nodes[1]
	nodes[2].right = &nodes[3]
	nodes[6].left = &nodes[5]
	nodes[6].right = &nodes[7]
	nodes[1].left = &nodes[0]
	bt.root = &nodes[4]

	return bt
}

// Print the structure of a binary tree
func (bt *BinaryTree) PrintBinaryTree() {
	tmp := bt.root
	printBinaryTreeHelper(tmp, 0)
}

func printBinaryTreeHelper(root *BinaryNode, level int) {
	if root == nil {
		return
	}

	printBinaryTreeHelper(root.right, level+1)
	for i := 0; i < level; i++ {
		fmt.Print("    ")
	}
	fmt.Println(root.item)
	printBinaryTreeHelper(root.left, level+1)
}

func (bt *BinaryTree) PreOrderDisplay(node ...*BinaryNode) {
	var cur *BinaryNode
	if len(node) == 0 {
		cur = bt.root
	} else {
		cur = node[0]
	}

	if cur == nil {
		return
	}

	fmt.Println(cur.item)
	bt.PreOrderDisplay(cur.left)
	bt.PreOrderDisplay(cur.right)
}

func (bt *BinaryTree) InorderTrav() []int {
	var res []int
	inorderTravHelper(bt.root, &res)
	return res
}

func inorderTravHelper(node *BinaryNode, res *[]int) {
	if node == nil {
		return
	}

	inorderTravHelper(node.left, res)
	*res = append(*res, node.item)
	inorderTravHelper(node.right, res)
}

func (bt *BinaryTree) InOrderDisplay(node ...*BinaryNode) {
	var cur *BinaryNode
	if len(node) == 0 {
		cur = bt.root
	} else {
		cur = node[0]
	}

	if cur == nil {
		return
	}

	bt.InOrderDisplay(cur.left)
	fmt.Println(cur.item)
	bt.InOrderDisplay(cur.right)
}

func (bt *BinaryTree) PostOrderDisplay(node ...*BinaryNode) {
	var cur *BinaryNode
	if len(node) == 0 {
		cur = bt.root
	} else {
		cur = node[0]
	}

	if cur == nil {
		return
	}

	bt.PostOrderDisplay(cur.left)
	bt.PostOrderDisplay(cur.right)
	fmt.Println(cur.item)
}

// Return the first node of a sub-tree using in-order
func BinarySubTreeFirst(node *BinaryNode) *BinaryNode {
	if node.left != nil {
		return BinarySubTreeFirst(node.left)
	} else {
		return node
	}
}

// Return the last node of a sub-tree using in-order
func BinarySubTreeLast(node *BinaryNode) *BinaryNode {
	if node.right != nil {
		return BinarySubTreeLast(node.right)
	} else {
		return node
	}
}

func (bt *BinaryTree) BinarySearch(value int) bool {
	return binarySearchHelper(bt.root, value)
}

func binarySearchHelper(node *BinaryNode, value int) bool {
	if node == nil {
		return false
	}
	if value < node.item {
		return binarySearchHelper(node.left, value)
	} else if value > node.item {
		return binarySearchHelper(node.right, value)
	}

	return true
}

// Insert the node after specify node, keep in-order
func (bt *BinaryTree) InsertAfter(node *BinaryNode, new_value int) {
	new_node := &BinaryNode{item: new_value}
	if node.right == nil {
		node.right = new_node
	} else {
		tmp := BinarySubTreeFirst(node.right)
		tmp.left = new_node
	}
}

// Delete a node and keep the inorder, this is for tree with BST
func (bt *BinaryTree) BinaryDelete(value int) {
	// In case, the root change
	bt.root = binaryDeleteHelper(bt.root, value)
}

func binaryDeleteHelper(node *BinaryNode, value int) *BinaryNode {
	if node == nil {
		return nil
	}

	if value < node.item {
		node.left = binaryDeleteHelper(node.left, value)
	} else if value > node.item {
		node.right = binaryDeleteHelper(node.right, value)
	} else {
		// Case 1: No child or only one child
		// a: if no child, just return and assign nil to its parent's left or right
		// b: if has right node, return and its right node to replace itself
		// c: left situation is same as right
		if node.left == nil {
			return node.right
		} else if node.right == nil {
			return node.left
		}

		// Case 2: Node with two children
		// Find the minimum value in the right subtree
		// Exchange the two node
		// Alternativly, can exchange the last node of its left tree: pred := BinarySubTreeLast(node.left)
		successor := BinarySubTreeFirst(node.right)
		node.item = successor.item
		// Keep delete in its right sub-tree
		node.right = binaryDeleteHelper(node.right, successor.item)
	}

	return node
}
