package codewars

// DiffLists Returns the first slice after removing all items
// also contained in the second slice.
func DiffLists(a, b []int) []int {
	for _, v := range b {
		size := len(a)
		for i := 0; i < size; i++ {
			if a[i] == v {
				// remove item by re-slicing
				a = append(a[:i], a[i+1:]...)
				// since item was removed, decreases the index by 1
				i -= 1
				size = len(a)
			}
		}
	}
	return a
}
