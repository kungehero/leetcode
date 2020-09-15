package tree

func quickSortReverse(arr []int, l, r int) {
	if l >= r {
		return
	}
	i, j, key := l, r, arr[l]
	for j > i {
		for i < j && arr[j] < key {
			j--
		}
		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
		for i < j && arr[i] > key {
			i++
		}
		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
			j--
		}
		arr[i] = key
		quickSortReverse(arr, l, i-1)
		quickSortReverse(arr, i+1, r)
	}
}
