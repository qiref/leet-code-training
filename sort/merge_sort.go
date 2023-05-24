package sort

// 算法流程：先保证左边部分有序，然后保证右边部分有序，最后合并左右两部分
func mergeSort(arr []int) {
	process(arr, 0, len(arr)-1)
}

func process(arr []int, l int, r int) {
	if l == r {
		return
	}
	mid := l + (r-l)>>1
	process(arr, l, mid)
	process(arr, mid+1, r)
	merge(arr, l, mid, r)
}

// 合并过程中，分别用两个指针遍历左右集合，然后开辟一个新数组，分别比较左右集合的元素大小，直到其中一个集合遍历结束
// 一方集合遍历结束之后，需要继续将另一集合数据加入到新数组中
// 最后将新数组中的元素放到最终的数组中去，
func merge(arr []int, l int, mid int, r int) {
	newArr := make([]int, r-l+1)
	p1 := l
	p2 := mid + 1
	i := 0
	for p1 <= mid && p2 <= r {
		if arr[p2] < arr[p1] {
			newArr[i] = arr[p2]
			p2++
		} else {
			newArr[i] = arr[p1]
			p1++
		}
		i++
	}
	for p1 <= mid {
		newArr[i] = arr[p1]
		i++
		p1++
	}

	for p2 <= r {
		newArr[i] = arr[p2]
		i++
		p2++
	}
	for i := 0; i < len(newArr); i++ {
		arr[l+i] = newArr[i]
	}
}

func mergeSort1(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	mid := len(arr) / 2
	return merge1(mergeSort1(arr[0:mid]), mergeSort1(arr[mid:]))
}

func merge1(arrL []int, arrR []int) []int {
	var result []int
	for len(arrL) > 0 && len(arrR) > 0 {
		if arrL[0] < arrR[0] {
			result = append(result, arrL[0])
			arrL = arrL[1:]
		} else {
			result = append(result, arrR[0])
			arrR = arrR[1:]
		}
	}
	for len(arrL) > 0 {
		result = append(result, arrL[0])
		arrL = arrL[1:]
	}
	for len(arrR) > 0 {
		result = append(result, arrR[0])
		arrR = arrR[1:]
	}
	return result
}
