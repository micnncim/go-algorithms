package binary_search

const (
	NotFound = -1
)

// a is sorted slice.
func Search(a []int, want int) int {
	left, right := 0, len(a)

	for left < right {
		mid := (left + right) / 2
		if a[mid] == want {
			return mid
		}
		if a[mid] < want {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return NotFound
}
