package sort

// 插入排序
// 排序过程，保证0-0位置上的有序，然后保证0-1位置上的有序，然后保证0-2位置上的有序，以此类推，最后0-n-1位置都有序
// 注意，内层循环是从后向前的，
// 类似与扑克牌的插入，每取一张牌，往现有的有序手牌中从右往左插入
func insetSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	for i := 1; i < len(arr); i++ {
		for j := i; j > 0; j-- {
			if arr[j-1] > arr[j] {
				arr[j-1] = arr[j-1] ^ arr[j]
				arr[j] = arr[j-1] ^ arr[j]
				arr[j-1] = arr[j-1] ^ arr[j]
			}
		}
	}
	return arr
}

// 最外层倒序循环时，需要保证内层循环是从前往后，第一轮保证 0-0 位置上的有序，第二轮保证 0-1 位置上的有序
func insetSort1(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	for i := len(arr); i > 0; i-- {
		for j := len(arr) - i; j > 0 && arr[j] < arr[j-1]; j-- {
			arr[j], arr[j-1] = arr[j-1], arr[j]
		}
	}
	return arr
}
