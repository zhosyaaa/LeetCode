package main

//func main() {
//
//}

// ищет минимум и собирает в начало
func Selection(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		min_index := i
		for j := i; j < len(nums); j++ {
			if nums[j] < nums[min_index] {
				min_index = j
			}
		}
		nums[i], nums[min_index] = nums[min_index], nums[i]
	}
	return nums
}

// собирает максимум в конец
func BubbleSort(array []int) []int {
	for i := 0; i < len(array)-1; i++ {
		for j := 0; j < len(array)-i-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
	return array
}

// ///////////////////////////////////////////////////////////
func insertionsort(items []int) {
	for i := 1; i < len(items); i++ {
		j := i
		for j > 0 {
			if items[j-1] > items[j] {
				items[j-1], items[j] = items[j], items[j-1]
			}
			j = j - 1
		}
	}
}

// //////////////////////////////////////////////////////////
//func mergeSort(items []int) []int {
//	if len(items) < 2 {
//		return items
//	}
//	first := mergeSort(items[:len(items)/2])
//	second := mergeSort(items[len(items)/2:])
//	return merge(first, second)
//}

//func merge(a []int, b []int) []int {
//	final := []int{}
//	i := 0
//	j := 0
//	for i < len(a) && j < len(b) {
//		if a[i] < b[j] {
//			final = append(final, a[i])
//			i++
//		} else {
//			final = append(final, b[j])
//			j++
//		}
//	}
//	for ; i < len(a); i++ {
//		final = append(final, a[i])
//	}
//	for ; j < len(b); j++ {
//		final = append(final, b[j])
//	}
//	return final
//}

// ////////////////////////////////////////////////
func quickSort(arr []int, low, high int) []int {
	if low < high {
		var p int
		arr, p = partition(arr, low, high)
		arr = quickSort(arr, low, p-1)
		arr = quickSort(arr, p+1, high)
	}
	return arr
}
func partition(arr []int, low, high int) ([]int, int) {
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return arr, i
}
func quickSortStart(arr []int) []int {
	return quickSort(arr, 0, len(arr)-1)
}
