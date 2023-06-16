package binarytree

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
