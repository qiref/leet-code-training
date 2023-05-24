package sort

import "fmt"

type maxHeap struct {
	Data  []int
	Count int
}

func NewMaxHeap(size int) *maxHeap {
	return &maxHeap{
		Data:  make([]int, size),
		Count: 0,
	}
}

// insert 向堆中插入元素
func (heap *maxHeap) Insert(val int) {
	heap.Data[heap.Count] = val
	heap.Count++
	// heap.shiftUp(heap.Count)
	heap.shiftUp1(heap.Count - 1) // 最后一个元素的下标
}

// shiftUp 向堆中插入元素时，叶子节点可能需要向上移动
func (heap *maxHeap) shiftUp(count int) {
	for count > 1 && heap.Data[count-1] > heap.Data[(count/2)-1] {
		heap.Data[count-1], heap.Data[(count/2)-1] = heap.Data[(count/2)-1], heap.Data[count-1]
		count = (count / 2)
	}
	fmt.Printf("shiftUp heap %v \n", heap)
}

// shiftUp 向堆中插入元素时，叶子节点可能需要向上移动
func (heap *maxHeap) shiftUp1(index int) {
	// 父节点 = (i-1)/2
	for index > 0 && heap.Data[index] > heap.Data[(index-1)/2] {
		heap.Data[index], heap.Data[(index-1)/2] = heap.Data[(index-1)/2], heap.Data[index]
		index = (index - 1) / 2
	}
	fmt.Printf("shiftUp heap %v \n", heap)
}

// GetMax 大根堆，最大值就是跟节点
func (heap *maxHeap) GetMax() int {
	return heap.Data[0]
}

// ExtraMax 提取大根堆的最大值，也就是根节点，大根堆整体性质不变
func (heap *maxHeap) ExtraMax() int {
	heap.Count--
	result := heap.Data[0]
	heap.Data[0], heap.Data[heap.Count] = heap.Data[heap.Count], heap.Data[0]
	heap.shiftDown(0)
	return result
}

// shiftDown 移除堆中元素，节点向下移动
func (heap *maxHeap) shiftDown(index int) {
	for index <= heap.Count-1 {
		i := heap.getMaxChildNode(index)
		if i == -1 {
			break
		}
		if heap.Data[index] < heap.Data[i] {
			heap.Data[index], heap.Data[i] = heap.Data[i], heap.Data[index]
		}
		index = i
	}
}

// getMaxChildNode 获取最大子节点下标，如果没有子节点，return -1
func (heap *maxHeap) getMaxChildNode(index int) int {
	if index <= heap.Count-1 {
		if index*2+2 > heap.Count-1 {
			// 当右节点下标越界时
			if index*2+1 <= heap.Count-1 {
				// 如果左下标没有越界，则返回左节点下标
				return index*2 + 1
			}
		} else {
			// 当右节点下标没有越界时，此时左下标一定没有越界
			if heap.Data[index*2+1] >= heap.Data[index*2+2] {
				return index*2 + 1
			} else {
				return index*2 + 2
			}
		}
	}
	// 没有子节点时，返回-1
	return -1
}

// MaxHeapipy 最大堆初始化
func MaxHeapipy(arr []int) *maxHeap {
	mh := NewMaxHeap(len(arr))
	mh.Data = arr
	mh.Count = len(arr)
	// 堆是一棵完全二叉树，最后一个非叶子节点是：(n-1)/2，n 是数组大小
	// 基于这个性质，可以依次对每一个非叶子节点进行 shiftDown，相当于每一个子树都完成 shiftDown，最后完成根节点的 shiftDown
	// shiftDown 之后，依然保持大根堆的性质
	for i := (mh.Count - 1) / 2; i >= 0; i-- {
		mh.shiftDown(i)
	}
	return mh
}
