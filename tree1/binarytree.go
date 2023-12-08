package tree1

import (
	"golang.org/x/exp/constraints"
)

type Node[T constraints.Ordered] struct {
	data  T
	left  *Node[T]
	right *Node[T]
}

type BinaryTree[T constraints.Ordered] struct {
	root *Node[T]
}

func (tree *BinaryTree[T]) InvertTree() *Node[T] {
	if tree.root == nil {
		return nil
	}

	rightSubTree := &BinaryTree[T]{root: tree.root.right}

	leftSubTree := &BinaryTree[T]{root: tree.root.left}

	tree.root.left = rightSubTree.InvertTree()

	tree.root.right = leftSubTree.InvertTree()

	return tree.root
}

// PreOrder Node-Left-Right
func (tree *BinaryTree[T]) PreOrder(result []T) []T {
	if tree.root == nil {
		return result
	}

	result = append(result, tree.root.data)

	leftSubTree := &BinaryTree[T]{root: tree.root.left}
	result = leftSubTree.PreOrder(result)

	rightSubTree := &BinaryTree[T]{root: tree.root.right}
	result = rightSubTree.PreOrder(result)

	return result
}

// InOrder Left-Node-Right
func (tree *BinaryTree[T]) InOrder(result []T) []T {
	if tree.root == nil {
		return result
	}

	leftSubTree := &BinaryTree[T]{root: tree.root.left}
	result = leftSubTree.InOrder(result)

	result = append(result, tree.root.data)

	rightSubTree := &BinaryTree[T]{root: tree.root.right}
	result = rightSubTree.InOrder(result)

	return result
}

// PostOrder Left-Right-Node
func (tree *BinaryTree[T]) PostOrder(result []T) []T {
	if tree.root == nil {
		return result
	}

	leftSubTree := &BinaryTree[T]{root: tree.root.left}
	result = leftSubTree.PostOrder(result)

	rightSubTree := &BinaryTree[T]{root: tree.root.right}
	result = rightSubTree.PostOrder(result)

	result = append(result, tree.root.data)

	return result
}

func (tree *BinaryTree[T]) LevelOrder() []T {
	var result []T

	if tree.root == nil {
		return result
	}

	queue := []*Node[T]{tree.root}

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

func (tree *BinaryTree[T]) Insert(data T) *Node[T] {
	newNode := &Node[T]{data: data}

	switch {
	case tree.root == nil:
		tree.root = newNode

		return tree.root

	case newNode.data > tree.root.data || newNode.data == tree.root.data:
		rightSubTree := &BinaryTree[T]{root: tree.root.right}

		tree.root.right = rightSubTree.Insert(data)

	case newNode.data < tree.root.data:
		leftSubTree := &BinaryTree[T]{root: tree.root.left}

		tree.root.left = leftSubTree.Insert(data)
	}

	return tree.root

}

func (tree *BinaryTree[T]) Delete(data T) *Node[T] {
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

			rightSubTree := &BinaryTree[T]{root: tree.root.right}

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
		rightSubTree := &BinaryTree[T]{root: tree.root.right}

		tree.root.right = rightSubTree.Delete(data)

	case data < tree.root.data:
		leftSubTree := &BinaryTree[T]{root: tree.root.left}

		tree.root.left = leftSubTree.Delete(data)
	}

	return tree.root
}

func (tree *BinaryTree[T]) IsFull() bool {
	switch {
	case tree.root == nil:
		return true

	case tree.root.left == nil && tree.root.right == nil:
		return true

	case tree.root.left != nil && tree.root.right != nil:
		rightSubTree := &BinaryTree[T]{root: tree.root.right}

		leftSubTree := &BinaryTree[T]{root: tree.root.left}

		return leftSubTree.IsFull() && rightSubTree.IsFull()

	default:
		return false
	}
}

func (tree *BinaryTree[T]) IsPerfect(depth, level int) bool {
	if tree.root == nil {
		return true
	}

	if tree.root.left == nil && tree.root.right == nil {
		return depth == level
	}

	if tree.root.left == nil || tree.root.right == nil {
		return false
	}

	leftSubTree := &BinaryTree[T]{root: tree.root.left}

	rightSubTree := &BinaryTree[T]{root: tree.root.right}

	return leftSubTree.IsPerfect(depth, level+1) && rightSubTree.IsPerfect(depth, level+1)
}

func (tree *BinaryTree[T]) IsBalanced() bool {
	return tree.root == nil
}

func (tree *BinaryTree[T]) Depth() int {
	if tree.root == nil {
		return -1
	}

	lDepth := (&BinaryTree[T]{root: tree.root.left}).Depth()

	rDepth := (&BinaryTree[T]{root: tree.root.right}).Depth()

	if lDepth > rDepth {
		return (lDepth + 1)
	}

	return rDepth + 1
}

func (tree *BinaryTree[T]) Search(data T) bool {
	switch {
	case tree.root == nil:
		return false

	case tree.root.data == data:
		return true

	case data > tree.root.data:
		rightSubTree := &BinaryTree[T]{root: tree.root.right}

		return rightSubTree.Search(data)

	case data < tree.root.data:
		leftSubTree := &BinaryTree[T]{root: tree.root.left}

		return leftSubTree.Search(data)

	default:
		return false
	}
}
