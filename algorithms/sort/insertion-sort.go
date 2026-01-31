package sort

func InsertionSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1

		for j > 0 && arr[j] > key {
			arr[j+1] = key
			j--
		}

		arr[i] = key
	}

	return arr
}
