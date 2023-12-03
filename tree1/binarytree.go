package tree1

import "fmt"

type Node struct {
	data  byte
	left  *Node
	right *Node
}

type BinaryTree struct {
	root *Node
}

func (tree *BinaryTree) InvertTree() *Node {
	if tree.root == nil {
		return nil
	}

	rightSubTree := &BinaryTree{root: tree.root.right}

	leftSubTree := &BinaryTree{root: tree.root.left}

	tree.root.left = rightSubTree.InvertTree()

	tree.root.right = leftSubTree.InvertTree()

	return tree.root
}

// PreOrder Node-Left-Right
func (tree *BinaryTree) PreOrder(result []byte) []byte {
	if tree.root == nil {
		return result
	}

	result = append(result, tree.root.data)

	leftSubTree := &BinaryTree{root: tree.root.left}
	result = leftSubTree.PreOrder(result)

	rightSubTree := &BinaryTree{root: tree.root.right}
	result = rightSubTree.PreOrder(result)

	return result
}

// InOrder Left-Node-Right
func (tree *BinaryTree) InOrder(result []byte) []byte {
	if tree.root == nil {
		return result
	}

	leftSubTree := &BinaryTree{root: tree.root.left}
	result = leftSubTree.InOrder(result)

	result = append(result, tree.root.data)

	rightSubTree := &BinaryTree{root: tree.root.right}
	result = rightSubTree.InOrder(result)

	return result
}

// PostOrder Left-Right-Node
func (tree *BinaryTree) PostOrder(result []byte) []byte {
	if tree.root == nil {
		return result
	}

	leftSubTree := &BinaryTree{root: tree.root.left}
	result = leftSubTree.PostOrder(result)

	rightSubTree := &BinaryTree{root: tree.root.right}
	result = rightSubTree.PostOrder(result)

	result = append(result, tree.root.data)

	return result
}

func (tree *BinaryTree) LevelOrder() []byte {
	if tree.root == nil {
		return []byte{}
	}

	queue := []*Node{tree.root}

	var result []byte

	for len(queue) != 0 {
		curr := queue[0]

		result = append(result, curr.data)

		queue = queue[1:]

		if curr.left != nil {
			queue = append(queue, curr.left)
		}

		if curr.right != nil {
			queue = append(queue, curr.right)

		}
	}

	return result
}

func (tree *BinaryTree) Insert(data byte) *Node {
	newNode := &Node{data: data}

	switch {
	case tree.root == nil:
		tree.root = newNode

		return tree.root

	case newNode.data > tree.root.data || newNode.data == tree.root.data:
		rightSubTree := &BinaryTree{root: tree.root.right}

		tree.root.right = rightSubTree.Insert(data)

	case newNode.data < tree.root.data:
		leftSubTree := &BinaryTree{root: tree.root.left}

		tree.root.left = leftSubTree.Insert(data)
	}

	return tree.root

}

func (tree *BinaryTree) Delete(data byte) *Node {
	switch {
	case tree.root.data == data:

		switch {
		case tree.root.right == nil && tree.root.left == nil:
			tree.root = nil

		case tree.root.right == nil:
			tree.root = tree.root.left

		case tree.root.left == nil:
			tree.root = tree.root.right

		default:
			temp := tree.root.right

			for temp.left != nil {
				temp = temp.left
			}

			tree.root.data = temp.data

			rightSubTree := &BinaryTree{root: tree.root.right}

			tree.root.right = rightSubTree.Delete(temp.data)

			// temp := tree.root.left

			// for temp.right != nil {
			// 	temp = temp.right
			// }

			// leftSubtree := &BinaryTree{root: tree.root.left}

			// tree.root.left = leftSubtree.Delete(temp.data)
		}

		return tree.root

	case data > tree.root.data:
		rightSubTree := &BinaryTree{root: tree.root.right}

		tree.root.right = rightSubTree.Delete(data)

	case data < tree.root.data:
		leftSubTree := &BinaryTree{root: tree.root.left}

		tree.root.left = leftSubTree.Delete(data)
	}

	return tree.root
}

func (tree *BinaryTree) IsFullBinaryTree() bool {
	switch {
	case tree.root == nil:
		return true

	case tree.root.left == nil && tree.root.right == nil:
		return true

	case tree.root.left != nil && tree.root.right != nil:
		rightSubTree := &BinaryTree{root: tree.root.right}

		leftSubTree := &BinaryTree{root: tree.root.right}

		return leftSubTree.IsFullBinaryTree() && rightSubTree.IsFullBinaryTree()

	default:
		return false
	}
}

func (tree *BinaryTree) IsPerfectBinaryTree(level int) bool {
	if tree.root == nil {
		return true
	}

	if tree.root.left == nil && tree.root.right == nil {
		return tree.Depth() == level+1
	}

	if tree.root.left == nil || tree.root.right == nil {
		return false
	}

	leftSubTree := &BinaryTree{root: tree.root.left}

	rightSubTree := &BinaryTree{root: tree.root.right}

	return leftSubTree.IsPerfectBinaryTree(level+1) && rightSubTree.IsPerfectBinaryTree(level+1)
}

func (tree *BinaryTree) IsBalancedBinaryTree() bool {
	return tree.root == nil
}

func (tree *BinaryTree) Depth() int {
	if tree.root == nil {
		return -1
	}

	/* compute the depth of each subtree */
	lDepth := (&BinaryTree{root: tree.root.left}).Depth()

	rDepth := (&BinaryTree{root: tree.root.right}).Depth()

	/* use the larger one */
	if lDepth > rDepth {
		return (lDepth + 1)
	}

	return rDepth + 1
}

func (tree *BinaryTree) Search(data byte) bool {
	switch {
	case tree.root == nil:
		return false

	case tree.root.data == data:
		return true

	case data > tree.root.data:
		rightSubTree := &BinaryTree{root: tree.root.right}

		return rightSubTree.Search(data)

	case data < tree.root.data:
		leftSubTree := &BinaryTree{root: tree.root.left}

		return leftSubTree.Search(data)

	default:
		return false
	}
}

func Test() {
	tree := new(BinaryTree)

	for _, val := range []byte{24, 20, 27, 22, 21, 30, 38, 35, 40, 18, 32, 33} {
		tree.Insert(val)
	}

	tree.Delete(24)

	fmt.Println(tree.Depth())

	fmt.Println(tree.IsPerfectBinaryTree(0))

	tree.InvertTree()

	tree.LevelOrder()
}
