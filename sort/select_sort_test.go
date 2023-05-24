package sort

import (
	"fmt"
	"testing"
)

func Test_selectSor(t *testing.T) {
	arr := []int{234, 234, 2, 2, 5, 3456, 1324, 234, 234, 1, 6, 7, 3, 6}
	resultArr := selectSor(arr)
	t.Log(fmt.Sprintf("result %v", resultArr))
}
