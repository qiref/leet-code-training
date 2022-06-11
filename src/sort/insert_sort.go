package sort

// 插入排序
// 排序过程，保证0-0位置上的有序，然后保证0-1位置上的有序，然后保证0-2位置上的有序
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
