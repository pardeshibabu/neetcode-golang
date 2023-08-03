package heappriorityqueue

import "container/heap"

type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *MaxHeap) Pop() interface{} {
	ele := (*h)[0]
	*h = (*h)[1:]
	return ele
}

func lastStoneWeight(stones []int) int {
	mh := &MaxHeap{}
	*mh = stones
	heap.Init(mh)
	for mh.Len() > 1 {
		y := mh.Pop()
		heap.Init(mh)
		x := mh.Pop()
		heap.Init(mh)
		y1 := y.(int)
		x1 := x.(int)
		if y1-x1 == 0 {
			continue
		}
		if y1-x1 > 0 {
			heap.Push(mh, (y1 - x1))
			heap.Init(mh)
		}
	}
	if mh.Len() == 0 {
		return 0
	}
	return heap.Pop(mh).(int)
}
