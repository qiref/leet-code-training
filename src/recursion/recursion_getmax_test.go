package recursion

import "testing"

func Test_getMax(t *testing.T) {
	i := getMax([]int{1, 2, 34, 5, 4, 7, 567, 2, 2})
	t.Log(i)
}
