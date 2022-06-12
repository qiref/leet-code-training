package recursion

func getMax(arr []int) int {
	return process(arr, 0, len(arr)-1)
}

func process(arr []int, l int, r int) int {
	if l == r {
		return arr[l]
	}
	mid := l + (r-l)>>1 // (l+r)/2 可能会溢出， l+(r-l)/2 -> l+(r-l)>>1
	maxL := process(arr, l, mid)
	maxR := process(arr, mid+1, r)
	if maxL > maxR {
		return maxL
	} else {
		return maxR
	}
}
