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
	root   *BinaryNode
	height int
}

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

func InOrderBinarySubTreeFirst(node *BinaryNode) *BinaryNode {
	if node.left != nil {
		return InOrderBinarySubTreeFirst(node.left)
	} else {
		return node
	}
}

func InOrderBinarySubTreeLast(node *BinaryNode) *BinaryNode {
	if node.right != nil {
		return InOrderBinarySubTreeLast(node.right)
	} else {
		return node
	}
}

func (bt *BinaryTree) InOrderInsertAfter(node, new *BinaryNode) {
	if node.right == nil {
		node.right = new
	} else {
		tmp := InOrderBinarySubTreeFirst(node.right)
		tmp.left = new
	}
}

// Delete a item in the binary tree, but still maintain in-order
// func (bt *BinaryTree) Delete(node *BinaryNode) {
// 	if node.left == nil && node.right == nil {

// 	}
// }

// func InOrderNodeSuccessor(node *BinaryNode) *BinaryNode {
// 	if node.right != nil {
// 		return node.right
// 	} else {
// 		for node != nil {

// 		}
// 	}
// }
