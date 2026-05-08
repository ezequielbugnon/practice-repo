package binarysearch

func BinarySearch(arr []int, target int) bool {
	if len(arr) < 1 {
		return false
	}

	left := 0
	right := len(arr) - 1

	for left <= right {
		mid := left + (right-left)/2

		if arr[mid] == target {
			return true
		}

		if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return false
}
