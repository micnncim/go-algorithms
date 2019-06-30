package binary_search_tree

// ref) https://flaviocopes.com/golang-data-structure-binary-search-tree/

type Node struct {
	value               int
	left, right, parent *Node
}

type Tree struct {
	root *Node
	size int
}

func NewNode(v int) *Node {
	return &Node{value: v}
}

func (t *Tree) Insert(v int) {
	node := NewNode(v)
	head := t.root
	for {
		if node.value < head.value {
			if head.left == nil {
				node.parent = head
				head.left = node
				return
			}
			head = head.left
		} else {
			if head.right == nil {
				node.parent = head
				head.right = node
				return
			}
			head = head.right
		}
	}
}

func (t *Tree) Search(v int) *Node {
	head := t.root
	for {
		if head == nil {
			return nil
		}
		if v == head.value {
			return head
		}
		if head.left == nil && head.right == nil {
			return nil
		}
		if head.left != nil && v < head.left.value {
			head = head.left
			continue
		}
		if head.right != nil && v < head.right.value {
			head = head.right
		}
	}
}

func (t *Tree) Delete(v int) bool {
	return false
}
