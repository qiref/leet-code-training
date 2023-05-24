package sort

import (
	"fmt"
	"testing"
)

func Test_bubbleSort(t *testing.T) {
	arr := []int{23, 23, 2, 1, 5, 7, 9, 3, 34534, 34}
	resultArr := bubbleSort(arr)
	t.Log(fmt.Sprintf("result %v", resultArr))
}

func Test_bubbleSort1(t *testing.T) {
	arr := []int{23, 23, 2, 1, 5, 7, 9, 3, 34534, 34}
	resultArr := bubbleSort1(arr)
	t.Log(fmt.Sprintf("result %v", resultArr))
}

func Test_bubbleSort2(t *testing.T) {
	arr := []int{23, 23, 2, 1, 5, 7, 9, 3, 34534, 34}
	resultArr := bubbleSort2(arr)
	t.Log(fmt.Sprintf("result %v", resultArr))
}
