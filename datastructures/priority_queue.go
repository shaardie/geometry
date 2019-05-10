package datastructures

import (
	"container/heap"
	"fmt"
)

// PriorityQueueItem is an Item in the Priority Queue
type PriorityQueueItem interface {
	// Compare with the Parameter.
	// If smaller return < 0, if greater return > 0 and 0 if equal.
	Compare(PriorityQueueItem) int
}

type priorityQueueItem struct {
	item  PriorityQueueItem
	index int
}

type priorityQueue []priorityQueueItem

// PriorityQueue is a priority queue as defined in https://en.wikipedia.org/wiki/Priority_queue
type PriorityQueue struct {
	pq *priorityQueue
}

// NewPriorityQueue with items as optional initial PriorityQueueItems
func NewPriorityQueue(items ...PriorityQueueItem) *PriorityQueue {
	l := len(items)
	pq := make(priorityQueue, l)
	for i, item := range items {
		pq[i] = priorityQueueItem{item, i}
	}
	heap.Init(&pq)
	return &PriorityQueue{&pq}

}

func (pq PriorityQueue) String() string {
	return fmt.Sprintf("%+v", pq.pq)
}

// Len returns the length of the PriorityQueue
func (pq PriorityQueue) Len() int {
	return pq.pq.Len()
}

// Pop the PriorityQueueItem with the highest priority
func (pq PriorityQueue) Pop() PriorityQueueItem {
	return heap.Pop(pq.pq).(priorityQueueItem).item
}

// Push a new PriorityQueueItem to the PriorityQueue
func (pq PriorityQueue) Push(pi PriorityQueueItem) {
	heap.Push(pq.pq, priorityQueueItem{item: pi})
}

func (pq priorityQueue) Len() int { return len(pq) }

func (pq priorityQueue) Less(i, j int) bool {
	return pq[i].item.Compare(pq[j].item) < 0
}

func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *priorityQueue) Push(x interface{}) {
	n := pq.Len()
	item := x.(priorityQueueItem)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *priorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}
