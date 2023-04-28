package bitmap

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type BitMap struct {
	data []int64
}

func NewBitMap(size int) *BitMap {
	return &BitMap{
		data: make([]int64, size),
	}
}

func ParseFromJsonStr(str string) *BitMap {
	bitMap := NewBitMap(0)
	err := json.Unmarshal([]byte(str), &bitMap.data)
	if err != nil {
		_ = fmt.Errorf("parse from json str return error %+v", err)
	}
	return bitMap
}

func (b *BitMap) set(index uint) {
	dataIdx := index >> 5 // index/32
	if dataIdx >= uint(len(b.data)) {
		b.data = append(b.data, 1<<(index&31)) // index&31 = index%32
	} else {
		b.data[dataIdx] |= 1 << (index & 31)
	}
}

func (b *BitMap) get(index uint) bool {
	dataIdx := index >> 5
	if len(b.data) <= 0 || uint(len(b.data)) < dataIdx {
		return false
	}
	return b.data[dataIdx]&(1<<(index&31)) > 0
}

func (b *BitMap) String() string {
	result := ""
	for _, v := range b.data {
		result += strconv.FormatInt(v, 10) + "(" + DecToBin(v) + "),"
	}
	return result
}

func DecToBin(n int64) string {
	result := ""
	if n == 0 {
		return "0"
	}
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.FormatInt(lsb, 10) + result
	}
	return result
}
