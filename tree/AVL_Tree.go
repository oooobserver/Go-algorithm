package tree

type AVLBinaryNode struct {
	value  int
	left   *AVLBinaryNode
	right  *AVLBinaryNode
	height int
}

type AVLTree struct {
	root *AVLBinaryNode
}

func (node *AVLBinaryNode) skew() int {
	return node.left.height - node.right.height
}

func (node *AVLBinaryNode) subTreeUpdate() {
	node.height = 1 + max(node.left.height, node.right.height)
}

func BinarySubTreeRotateRight(node *AVLBinaryNode) {
	if node.left == nil {
		return
	}
	b, e := node.left, node.right
	a, c := b.left, b.right
	node.value, b.value = b.value, node.value
	node.left, node.right = a, b
	b.left, b.right = c, e
	node.subTreeUpdate()
	b.subTreeUpdate()
}

func BinarySubTreeRotateLeft(node *AVLBinaryNode) {
	if node.right == nil {
		return
	}
	a, d := node.left, node.right
	c, e := d.left, d.right
	node.value, d.value = d.value, node.value
	node.left, node.right = d, e
	d.left, d.right = a, c
	node.subTreeUpdate()
	d.subTreeUpdate()
}

func (node *AVLBinaryNode) rebalance() {
	if node.skew() == 2 {
		// Case 3
		if node.right.skew() < 0 {
			BinarySubTreeRotateRight(node.right)
		}
		BinarySubTreeRotateLeft(node)
	} else if node.skew() == -2 {
		if node.left.skew() > 0 {
			BinarySubTreeRotateLeft(node.left)
		}
		BinarySubTreeRotateRight(node)
	}
}

func (node *AVLBinaryNode) Maintain() {
	node.rebalance()
	node.subTreeUpdate()
}
