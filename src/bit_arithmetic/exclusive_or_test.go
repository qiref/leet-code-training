package bitarithmetic

import (
	"testing"
)

func Test_printOddTimesNum1(t *testing.T) {
	i := printOddTimesNum1([]int{1, 1, 2, 2, 3})
	t.Log(i)
}

func Test_printOddTimesNum2(t *testing.T) {
	i, j := printOddTimesNum2([]int{1, 1, 2, 2, 3, 4})
	t.Log(i, j)
}
