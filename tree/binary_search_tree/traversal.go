package binary_search_tree

func TraversePreOrder(n *Node, f func(int)) {
	if n == nil {
		return
	}
	f(n.value)
	TraversePreOrder(n.left, f)
	TraversePreOrder(n.right, f)
}

func TraverseInOrder(n *Node, f func(int)) {
	if n == nil {
		return
	}
	TraverseInOrder(n.left, f)
	f(n.value)
	TraverseInOrder(n.right, f)
}

func TraversePostOrder(n *Node, f func(int)) {
	if n == nil {
		return
	}
	TraversePostOrder(n.left, f)
	TraversePostOrder(n.right, f)
	f(n.value)
}
