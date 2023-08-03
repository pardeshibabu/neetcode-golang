package heappriorityqueue

// type MinHeap []int

// func (h MinHeap) Len() int {
// 	return len(h)
// }

// func (h MinHeap) Less(i, j int) bool {
// 	return h[i] > h[j]
// }

// func (h MinHeap) Swap(i, j int) {
// 	h[i], h[j] = h[j], h[i]
// }

// func (h *MinHeap) Push(x interface{}) {
// 	*h = append(*h, x.(int))
// }

// func (h *MinHeap) Pop() interface{} {
// 	old := *h
// 	element := old[0]
// 	*h = old[1:]
// 	return element
// }

// func buildHeapByInit(array []int) *MinHeap {
// 	minHeap := &MinHeap{}
// 	// for _, val := range array {
// 	// 	*minHeap = append(*minHeap, val)
// 	// }
// 	*minHeap = array
// 	heap.Init(minHeap)
// 	return minHeap
// }

// // Time: O(nlogn)
// func BuildHeapByPush(array []int) *MinHeap {
// 	// initialize the MinHeap that has implement the heap.Interface
// 	minHeap := &MinHeap{}
// 	for _, num := range array {
// 		heap.Push(minHeap, num)
// 	}
// 	fmt.Println("buildHeapByPush: ", *minHeap)
// 	return minHeap
// }
