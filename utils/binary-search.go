package utils

// BinarySearch führt eine binäre Suche auf einer sortierten Liste durch
func BinarySearch(data []int, target int) int {
	low, high := 0, len(data)-1

	for low <= high {
		mid := (low + high) / 2

		if data[mid] < target {
			low = mid + 1
		} else if data[mid] > target {
			high = mid - 1
		} else {
			return mid
		}
	}

	return -1
}
