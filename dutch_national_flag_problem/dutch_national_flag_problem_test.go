package dutchnationalflagproblem

import (
	"testing"
)

func Test_dutchNationalFlagProblem1(t *testing.T) {
	i := dutchNationalFlagProblem1([]int{6, 32, 45, 7, 2, 7, 2}, 6)
	t.Logf("%v", i)
}

func Test_dutchNationalFlagProblem2(t *testing.T) {
	i := dutchNationalFlagProblem2([]int{6, 32, 6, 7, 2, 7, 2}, 6)
	t.Logf("%v", i)
}

func Test_dutchNationalFlagProblem3(t *testing.T) {
	i := dutchNationalFlagProblem3([]int{6, 32, 6, 6, 6, 6, 64, 6, 4, 7, 2, 7, 2}, 6)
	t.Logf("%v", i)
}

func Test_dutchNationalFlagProblem4(t *testing.T) {
	i, i2, i3 := dutchNationalFlagProblem4([]int{1, 3, 4, 2, 5}, 5)
	t.Logf("%v", i)
	t.Log(i2)
	t.Log(i3)
}
