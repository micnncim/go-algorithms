package binary_search_tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsert(t *testing.T) {
	tree := &Tree{root: NewNode(3)}
	tree.Insert(4)
	tree.Insert(1)
	tree.Insert(2)
	tree.Insert(5)

	assert.Equal(t, 4, tree.root.right.value)
	assert.Equal(t, 1, tree.root.left.value)
	assert.Equal(t, 2, tree.root.left.right.value)
	assert.Equal(t, 5, tree.root.right.right.value)
}

func TestSearch(t *testing.T) {
	tree := &Tree{root: NewNode(3)}
	tree.Insert(4)
	tree.Insert(1)
	tree.Insert(2)
	tree.Insert(5)

	found := tree.Search(1)
	t.Logf("%#v", found)
}
