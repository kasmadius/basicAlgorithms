package basicAlgorithms

func binSearch(arr []int, key int) int {
	return binSearchRec(arr, key, 0, len(arr)-1)
}

func binSearchRec(arr []int, key int, low int, high int) int {
	if low > high {
		return -1
	}

	mid := low + ((high - low) / 2)
	if arr[mid] == key {
		return mid
	}

	if key < arr[mid] {
		return binSearchRec(arr, key, low, mid-1)
	}

	return binSearchRec(arr, key, mid+1, high)
}
