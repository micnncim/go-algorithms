package binary_search

import (
	"testing"
)

func TestSearch(t *testing.T) {
	sorted := make([]int, 10000)
	for i := 0; i < 10000; i++ {
		sorted[i] = 2 * i
	}

	for i := 0; i < 10000; i++ {
		index := Search(sorted, 2*i)
		if index != i {
			t.Fatal(index)
		}
	}

	if Search(sorted, 3) != NotFound {
		t.Fatal("not return NotFound")
	}
}
