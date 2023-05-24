package minsum

import "testing"

func Test_calMinSum(t *testing.T) {
	i, arr := calMinSum([]int{1, 3, 4, 2, 5})
	t.Log(i)
	t.Logf("%v", arr)
}

func Test_calMinSum1(t *testing.T) {
	i, arr := calMinSum([]int{3, 10, 6, 17, 9})
	t.Log(i)
	t.Logf("%v", arr)
}
