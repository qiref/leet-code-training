package sort

import "testing"

func Test_mergSort(t *testing.T) {
	arr := []int{2, 4, 5, 45, 7, 56, 2, 24, 354, 345, 21, 1, 6, -2}
	mergSort(arr)
	t.Logf("%v", arr)
}
