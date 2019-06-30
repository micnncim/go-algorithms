package graph

type Node struct {
	ID        int
	Neighbors map[int]*Node
}

func MatrixToList(matrix [][]int) []*Node {
	n := len(matrix)
	list := make([]*Node, 0, n)
	for i := 0; i < n; i++ {
		list = append(list, &Node{
			ID:        i,
			Neighbors: make(map[int]*Node),
		})
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 1 {
				list[i].Neighbors[j] = list[j]
			}
		}
	}
	return list
}

func ListToMatrix(list []*Node) [][]int {
	n := len(list)
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}
	for i, node := range list {
		for j := range node.Neighbors {
			matrix[i][j] = 1
		}
	}
	return matrix
}
