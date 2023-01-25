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

// PreOrder Node-Left-Right
func (tree *BinaryTree) PreOrder() {
	if tree.root == nil {
		return
	}

	fmt.Printf("%d->", tree.root.data)

	leftSubTree := &BinaryTree{root: tree.root.left}
	leftSubTree.PreOrder()

	rightSubTree := &BinaryTree{root: tree.root.right}
	rightSubTree.PreOrder()
}

// InOrder Left-Node-Right
func (tree *BinaryTree) InOrder() {
	if tree.root == nil {
		return
	}

	leftSubTree := &BinaryTree{root: tree.root.left}
	leftSubTree.InOrder()

	fmt.Printf("%d->", tree.root.data)

	rightSubTree := &BinaryTree{root: tree.root.right}
	rightSubTree.InOrder()
}

// PostOrder Left-Right-Node
func (tree *BinaryTree) PostOrder() {
	if tree.root == nil {
		return
	}

	leftSubTree := &BinaryTree{root: tree.root.left}
	leftSubTree.PostOrder()

	rightSubTree := &BinaryTree{root: tree.root.right}
	rightSubTree.PostOrder()

	fmt.Printf("%d->", tree.root.data)

}

func (tree *BinaryTree) LevelOrder() {

	if tree.root == nil {
		return
	}
	var queue []*Node

	queue = append(queue, tree.root)

	for len(queue) != 0 {
		curr := queue[0]

		fmt.Printf("%d->", curr.data)

		queue = queue[1:]

		if curr.left != nil {
			queue = append(queue, curr.left)
		}

		if curr.right != nil {
			queue = append(queue, curr.right)

		}

	}

}

func (tree *BinaryTree) Insert(newNode *Node) *Node {
	switch {
	case tree.root == nil:
		tree.root = newNode

		return tree.root

	case newNode.data > tree.root.data:
		rightSubTree := &BinaryTree{root: tree.root.right}

		tree.root.right = rightSubTree.Insert(newNode)

	case newNode.data < tree.root.data:
		leftSubTree := &BinaryTree{root: tree.root.left}

		tree.root.left = leftSubTree.Insert(newNode)
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

func (tree *BinaryTree) Search(data byte) {
	switch {

	case tree.root == nil:
		fmt.Println(data, "not found")

	case tree.root.data == data:
		fmt.Println(data, "found")

	case data > tree.root.data:
		rightSubTree := &BinaryTree{root: tree.root.right}

		rightSubTree.Search(data)

	case data < tree.root.data:
		leftSubTree := &BinaryTree{root: tree.root.left}

		leftSubTree.Search(data)
	}
}

func binaryTree() {

	binaryTree := &BinaryTree{}

	for _, val := range []byte{24, 20, 27, 22, 21, 30, 38, 35, 40, 18, 32, 33} {
		binaryTree.Insert(&Node{data: val})
	}

	fmt.Print("\n\nPre Order Traversal: ")
	binaryTree.PreOrder()

	fmt.Print("\n\nIn Order Traversal: ")
	binaryTree.InOrder()

	fmt.Print("\n\nPost Order Traversal: ")
	binaryTree.PostOrder()

	fmt.Print("\n\nLevel Order Traversal: ")
	binaryTree.LevelOrder()

	binaryTree.Delete(24)

	fmt.Print("\n\nAfter deleting 24 (Pre Order): ")
	binaryTree.PreOrder()

	fmt.Println()
	binaryTree.Search(40)

}
