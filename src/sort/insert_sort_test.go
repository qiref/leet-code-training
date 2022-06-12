package sort

import (
	"testing"
)

func Test_insetSort(t *testing.T) {
	i := insetSort([]int{2, 4, 6, 1, 2, 988, 23, 465, 3})
	t.Logf("%v", i)
}

func Test_insetSort1(t *testing.T) {
	i := insetSort1([]int{2, 4, 6, 1, 2, 988, 23, 465, 3})
	t.Logf("%v", i)
}
