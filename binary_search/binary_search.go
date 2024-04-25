package binary_search

func BinarySearch(arr []int, value int) bool {
	begin := 0
	end := len(arr) - 1

	for begin <= end {
		mid := (begin + end) / 2

		guess := arr[mid]

		if guess < value {
			begin = mid + 1
		} else if guess > value {
			end = mid - 1
		} else {
			return true
		}
	}

	return false
}
