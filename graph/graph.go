package graph

type Node struct {
	adjacent map[int]*Node
}

func matrixToList(matrix [][]int) []*Node {
	n := len(matrix)
	list := make([]*Node, 0, n)
	for i := 0; i < n; i++ {
		list = append(list, &Node{
			adjacent: make(map[int]*Node),
		})
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 1 {
				list[i].adjacent[j] = list[j]
			}
		}
	}
	return list
}

func listToMatrix(list []*Node) [][]int {
	n := len(list)
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}
	for i, node := range list {
		for j := range node.adjacent {
			matrix[i][j] = 1
		}
	}
	return matrix
}
