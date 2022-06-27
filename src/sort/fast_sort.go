package sort

func fastSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	return sort1(arr, 0, len(arr)-1)
}

func sort1(arr []int, l int, r int) []int {
	if l < r {
		// 划分区间，从l到ltIdx区间是小于区间，mtIdx到r是大于区间，从ltIdx到mtIdx是等于区间
		ltIdx, mtIdx := partition1(arr, l, r)
		// 递归
		sort1(arr, l, ltIdx)
		sort1(arr, mtIdx, r)
	}
	return arr
}

func partition1(arr []int, l int, r int) (int, int) {
	// 这里的小于和大于区间的下标，初始范围应该是从 -1 到 len(arr)
	// 当发现有元素比num小是，小于区间第一个下标应该是0，此时应该替换 下标0(ltIdx+1) 和下标i 的元素
	ltIdx := l - 1
	mtIdx := r + 1
	num := arr[r]
	for i := l; i <= r; i++ {
		if i == mtIdx {
			// 如果i到大于num区域的下标，说明所有数据都已经对比完成
			break
		}
		if num == arr[i] {
			continue
		} else if num > arr[i] {
			// arr[i] 应该在小于num区域
			arr[i], arr[ltIdx+1] = arr[ltIdx+1], arr[i]
			ltIdx++
		} else if num < arr[i] {
			// arr[i] 应该在大于num区域
			arr[i], arr[mtIdx-1] = arr[mtIdx-1], arr[i]
			mtIdx--
			i--
		}
	}
	return ltIdx, mtIdx
}
