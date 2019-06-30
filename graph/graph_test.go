package graph

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://ja.wikipedia.org/wiki/隣接行列

func TestGraph(t *testing.T) {
	matrix := [][]int{
		{0, 1, 0, 0, 1, 0},
		{1, 0, 1, 0, 1, 0},
		{0, 1, 0, 1, 0, 0},
		{0, 0, 1, 0, 1, 1},
		{1, 1, 0, 1, 0, 0},
		{0, 0, 0, 1, 0, 0},
	}

	list := matrixToList(matrix)
	for i, node := range list {
		fmt.Printf("%d: ", i+1)
		for k := range node.adjacent {
			fmt.Printf("%d ", k+1)
		}
		fmt.Println()
	}
	got := listToMatrix(list)

	assert.Equal(t, matrix, got)
}
