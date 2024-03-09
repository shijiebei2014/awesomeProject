package heaps

import (
	"container/heap"
	"log"
	"testing"
)

type Ints []int

func (ints Ints) Len() int {
	return len(ints)
}

func (ints Ints) Less(i, j int) bool {
	return ints[i] < ints[j]
}

func (ints Ints) Swap(i, j int) {
	ints[i], ints[j] = ints[j], ints[i]
}

func (ints *Ints) Push(x any) {
	*ints = append(*ints, x.(int))
}

func (ints *Ints) Pop() any {
	last := (*ints)[len(*ints)-1]
	*ints = (*ints)[0 : len(*ints)-1]
	return last
}

func TestHeap(_ *testing.T) {
	arr := Ints{5, 4, 3, 2, 1}
	heap.Init(&arr)

	for arr.Len() > 0 {
		log.Println(heap.Pop(&arr))
	}
}
