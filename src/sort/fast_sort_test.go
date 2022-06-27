package sort

import (
	"testing"
)

func Test_fastSort(t *testing.T) {
	i := fastSort([]int{1, 3, 2, 4, 5})
	t.Logf("%v", i)
}

func Test_fastSort1(t *testing.T) {
	i := fastSort([]int{1, 3, 2, 4, 5, 5, 67, 0, 345, 223, -2, 234, 2, 2})
	t.Logf("%v", i)
}

func Test_fastSort2(t *testing.T) {
	i := fastSort([]int{1, -3, 2, -4, 5, 5, 67, 0, 11, 0, 0, -345, 223, -2, 234, 2, 2})
	t.Logf("%v", i)
}
