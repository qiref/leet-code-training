package bitmap

import (
	"fmt"
	"testing"
)

func TestNewBitmap(t *testing.T) {
	bitMap := NewBitMap(1)
	bitMap.set(1)
	bitMap.set(2)
	bitMap.set(3)
	bitMap.set(4)
	bitMap.set(0)
	fmt.Println(bitMap.String())

	bitMap.set(32)
	bitMap.set(33)
	fmt.Println(bitMap.String())

	bitMap.set(89)
	fmt.Println(bitMap.String())
}
