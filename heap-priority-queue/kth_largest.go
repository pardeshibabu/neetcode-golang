package heappriorityqueue

type KthLargest struct {
	// Num     []int
	Kth     int
	MinHeap *MinHeap
}

// func Constructor(k int, nums []int) KthLargest {
// 	minHeap := &MinHeap{}
// 	*minHeap = nums
// 	heap.Init(minHeap)
// 	for minHeap.Len() > k {
// 		minHeap.Pop()
// 	}
// 	return KthLargest{
// 		Kth:     k,
// 		MinHeap: minHeap,
// 	}
// }

// func (this *KthLargest) Add(val int) int {
// 	heap.Push(this.MinHeap, val)
// 	for this.MinHeap.Len() > this.Kth {
// 		heap.Init(this.MinHeap)
// 		this.MinHeap.Pop()
// 	}
// 	heap.Init(this.MinHeap)
// 	return (*this.MinHeap)[0]
// }
