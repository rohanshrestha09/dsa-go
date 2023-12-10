package tree1_test

import (
	"math/rand"
	"slices"
	"testing"

	t1 "github.com/rohanshrestha09/dsa-go/tree1"
)

func TestBinaryTree(t *testing.T) {
	bt := new(t1.BinaryTree[int])

	data := []int{24, 20, 27, 22, 21, 30, 18, 36, 25, 28, 23, 26, 17, 24, 19}

	levelOrder := []int{24, 20, 27, 18, 22, 25, 30, 17, 19, 21, 23, 24, 26, 28, 36}

	preOrder := []int{24, 20, 18, 17, 19, 22, 21, 23, 27, 25, 24, 26, 30, 28, 36}

	inOrder := []int{17, 18, 19, 20, 21, 22, 23, 24, 24, 25, 26, 27, 28, 30, 36}

	postOrder := []int{17, 19, 18, 21, 23, 22, 20, 24, 26, 25, 28, 36, 30, 27, 24}

	for _, val := range data {
		bt.Insert(val)
	}

	t.Run("test level order", func(t *testing.T) {
		result := bt.LevelOrder()

		if !slices.Equal(result, levelOrder) {
			t.Errorf("have %v, want %v", result, levelOrder)
		}
	})

	t.Run("test pre order", func(t *testing.T) {
		result := bt.PreOrder([]int{})

		if !slices.Equal(result, preOrder) {
			t.Errorf("have %v, want %v", result, preOrder)
		}
	})

	t.Run("test in order", func(t *testing.T) {
		result := bt.InOrder([]int{})

		if !slices.Equal(result, inOrder) {
			t.Errorf("have %v, want %v", result, inOrder)
		}
	})

	t.Run("test post order", func(t *testing.T) {
		result := bt.PostOrder([]int{})

		if !slices.Equal(result, postOrder) {
			t.Errorf("have %v, want %v", result, postOrder)
		}
	})

	t.Run("search element in tree", func(t *testing.T) {
		searchElement := data[rand.Intn(len(data))]

		found := bt.Search(searchElement)

		if !found {
			t.Errorf("%d must be present in %v", searchElement, data)
		}
	})

	t.Run("check depth of the tree", func(t *testing.T) {
		want := 3

		depth := bt.Depth()

		if depth != want {
			t.Errorf("Depth should be %d, but got %d", want, depth)
		}
	})

	t.Run("check if full binary tree", func(t *testing.T) {
		isFullBinaryTree := bt.IsFull()

		if !isFullBinaryTree {
			t.Errorf("Must be full binary tree")
		}
	})

	t.Run("check if perfect binary tree", func(t *testing.T) {
		isPerfectBinaryTree := bt.IsPerfect(bt.Depth(), 0)

		if !isPerfectBinaryTree {
			t.Errorf("Must be perfect binary tree")
		}
	})

	t.Run("check if balanced binary tree", func(t *testing.T) {
		isBalancedBinaryTree := bt.IsBalanced()

		if !isBalancedBinaryTree {
			t.Errorf("Must be balanced binary tree")
		}
	})
}
