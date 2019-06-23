package binary_tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTraversePreOrder(t *testing.T) {
	cases := []struct {
		name string
		n    *Node
		want []int
	}{
		{
			n: &Node{
				value: 0,
				left: &Node{
					value: 1,
					left:  &Node{value: 2},
					right: &Node{value: 3},
				},
				right: &Node{
					value: 4,
					left: &Node{
						value: 5,
						left:  &Node{value: 6},
						right: &Node{value: 7},
					},
					right: &Node{value: 8},
				},
			},
			want: []int{0, 1, 2, 3, 4, 5, 6, 7, 8},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var got []int
			TraversePreOrder(tc.n, func(v interface{}) {
				got = append(got, v.(int))
			})
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestTraverseInOrder(t *testing.T) {
	cases := []struct {
		name string
		n    *Node
		want []int
	}{
		{
			n: &Node{
				value: 0,
				left: &Node{
					value: 1,
					left:  &Node{value: 2},
					right: &Node{value: 3},
				},
				right: &Node{
					value: 4,
					left: &Node{
						value: 5,
						left:  &Node{value: 6},
						right: &Node{value: 7},
					},
					right: &Node{value: 8},
				},
			},
			want: []int{2, 1, 3, 0, 6, 5, 7, 4, 8},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var got []int
			TraverseInOrder(tc.n, func(v interface{}) {
				got = append(got, v.(int))
			})
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestTraversePostOrder(t *testing.T) {
	cases := []struct {
		name string
		n    *Node
		want []int
	}{
		{
			n: &Node{
				value: 0,
				left: &Node{
					value: 1,
					left:  &Node{value: 2},
					right: &Node{value: 3},
				},
				right: &Node{
					value: 4,
					left: &Node{
						value: 5,
						left:  &Node{value: 6},
						right: &Node{value: 7},
					},
					right: &Node{value: 8},
				},
			},
			want: []int{2, 3, 1, 6, 7, 5, 8, 4, 0},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var got []int
			TraversePostOrder(tc.n, func(v interface{}) {
				got = append(got, v.(int))
			})
			assert.Equal(t, tc.want, got)
		})
	}
}
