package main



func merge(arr []int, left, mid, right int) []int {
	n1 := mid - left + 1
	n2 := right - mid

	leftArray := make([]int, n1)
	rightArray := make([]int, n2)

	
	for i := 0; i < n1; i++ {
		leftArray[i] = arr[left+i]
	}
	for j := 0; j < n2; j++ {
		rightArray[j] = arr[mid+1+j]
	}

	
	i, j, k := 0, 0, left
	for i < n1 && j < n2 {
		if leftArray[i] <= rightArray[j] {
			arr[k] = leftArray[i]
			i++
		} else {
			arr[k] = rightArray[j]
			j++
		}
		k++
	}

	
	for i < n1 {
		arr[k] = leftArray[i]
		i++
		k++
	}
	for j < n2 {
		arr[k] = rightArray[j]
		j++
		k++
	}

	return arr
}

func merge_sort(arr []int, left, right int) {
	if left < right {
		mid := (left + right) / 2
		merge_sort(arr, left, mid)
		merge_sort(arr, mid+1, right)
		merge(arr, left, mid, right)
	}
}