package minsum

//求{1,2,3,4}的左边最小和，比如1的最小和就是0（左边没有比1小的数），2的最小和就是1，3的最小和就是1+2=3，4的最小和就是1+2+3=6，那这个数组的最小和就是0+1+3+6=10；
//求array数组左边最小和？
func calMinSum(arr []int) (int, []int) {
	if len(arr) < 2 {
		return 0, arr
	}
	mid := len(arr) / 2
	sumL, arrL := calMinSum(arr[0:mid])
	sumR, arrR := calMinSum(arr[mid:])
	return merge(arrL, arrR, sumL+sumR)
}

// merge 合并左边数组和右边数组，利用归并排序的思路，在左边数组和右边数组合并时，计算最小和
// 计算思路，当 arrL[L] < arrR[R] 时，此时可以得到：arr[L] 元素小于整个 arrR[]，因为 arrL arrR 都分别有序
// 那么 arr[L] 一定是需要计算 len(arrR)-R 次最小和的，所以最小和计算方式就是 sum += arrL[pL] * (len(arrR) - pR)
func merge(arrL []int, arrR []int, sum int) (int, []int) {
	var result []int
	var pL, pR int
	for pL < len(arrL) && pR < len(arrR) {
		if arrL[pL] < arrR[pR] {
			result = append(result, arrL[pL])
			sum += arrL[pL] * (len(arrR) - pR)
			pL++
		} else {
			result = append(result, arrR[pR])
			pR++
		}
	}
	for pL < len(arrL) {
		result = append(result, arrL[pL])
		pL++
	}
	for pR < len(arrR) {
		result = append(result, arrR[pR])
		pR++
	}
	return sum, result
}
