package tree1_test

import (
	"reflect"
	"testing"

	t1 "github.com/rohanshrestha09/dsa-go/tree1"
)

func TestBinaryTree(t *testing.T) {
	bt := new(t1.BinaryTree)

	data := []byte{24, 20, 27, 22, 21, 30, 38, 35, 40, 18, 32, 33}

	levelOrder := []byte{24, 20, 27, 18, 22, 30, 21, 38, 35, 40, 32, 33}

	preOrder := []byte{24, 20, 18, 22, 21, 27, 30, 38, 35, 32, 33, 40}

	inOrder := []byte{18, 20, 21, 22, 24, 27, 30, 32, 33, 35, 38, 40}

	postOrder := []byte{18, 21, 22, 20, 33, 32, 35, 40, 38, 30, 27, 24}

	for _, val := range data {
		bt.Insert(val)
	}

	t.Run("test level order", func(t *testing.T) {
		result := bt.LevelOrder()

		if !reflect.DeepEqual(result, levelOrder) {
			t.Errorf("have %v, want %v", result, levelOrder)
		}
	})

	t.Run("test pre order", func(t *testing.T) {
		result := bt.PreOrder([]byte{})

		if !reflect.DeepEqual(result, preOrder) {
			t.Errorf("have %v, want %v", result, preOrder)
		}
	})

	t.Run("test in order", func(t *testing.T) {
		result := bt.InOrder([]byte{})

		if !reflect.DeepEqual(result, inOrder) {
			t.Errorf("have %v, want %v", result, inOrder)
		}
	})

	t.Run("test post order", func(t *testing.T) {
		result := bt.PostOrder([]byte{})

		if !reflect.DeepEqual(result, postOrder) {
			t.Errorf("have %v, want %v", result, postOrder)
		}
	})

	t.Run("search element in tree", func(t *testing.T) {
		var searchElement byte = 40

		found := bt.Search(searchElement)

		if !found {
			t.Errorf("%d must be present in %v", searchElement, data)
		}
	})

	t.Run("check depth of the tree", func(t *testing.T) {
		want := 5

		depth := bt.Depth()

		if depth != want {
			t.Errorf("Depth should be %d, but got %d", want, depth)
		}
	})

	bt.Insert(25)

	bt.Insert(28)

	t.Run("check if full binary tree", func(t *testing.T) {
		isFullBinaryTree := bt.IsFullBinaryTree()

		if !isFullBinaryTree {
			t.Errorf("Must be full binary tree")
		}
	})
}
