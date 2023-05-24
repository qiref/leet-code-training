package sort

// 冒泡排序算法 从左往右，两两相互比较大小，左边的大就交换位置，循环往复，把较大的放后面。
// 冒泡算法的核心就是相邻两个元素进行比较，验证方法就是第一轮会找到最大元素，放到最后，所以内层循环第一轮一定会比较所有元素
func bubbleSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	for i := 0; i < len(arr); i++ {
		for j := 1; j < len(arr)-i; j++ {
			if arr[j-1] > arr[j] {
				temp := arr[j-1]
				arr[j-1] = arr[j]
				arr[j] = temp
			}
		}
	}
	return arr
}

// 最外层循环 i 从后往前，此时 j 只需要按照 i 的值循环，就能找到相邻两个元素进行比较
func bubbleSort1(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	for i := len(arr) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

// 最外层 i 从 0-n，要想内层循环第一轮比较所有元素，则j < n-i，然后保证第一个元素不越界，视情况+1 或 -1
func bubbleSort2(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}
