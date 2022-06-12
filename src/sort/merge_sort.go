package sort

func mergSort(arr []int) {
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
