package maximumsubsequencescore

import (
	"container/heap"
	"sort"
)

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *MinHeap) Pop() any {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}

func maxScore(nums1 []int, nums2 []int, k int) int64 {
	n := len(nums1)
	pairs := make([][2]int, n)
	for i := 0; i < n; i++ {
		pairs[i] = [2]int{nums2[i], nums1[i]}
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][0] > pairs[j][0]
	})

	h := &MinHeap{}
	heap.Init(h)
	var sum int64 = 0
	var result int64 = 0

	for _, p := range pairs {
		heap.Push(h, p[1])
		sum += int64(p[1])

		if h.Len() > k {
			removed := heap.Pop(h).(int)
			sum -= int64(removed)
		}

		if h.Len() == k {
			score := sum * int64(p[0])
			if score > result {
				result = score
			}
		}
	}

	return result
}
