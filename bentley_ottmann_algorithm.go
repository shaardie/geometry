package geometry

import (
	"container/heap"
	"fmt"
)

func bentleyOttmannAlgorithm(segments []Segment) (points []Point, err error) {

	// Initiliaze sweep line
	sl := sweepLine{}

	// Initialize eventQueue
	eq := make(eventQueue, 2*len(segments))

	// Sort Endpoints in Segments and add them to event queue
	for i, s := range segments {
		s.Sort()
		eq[2*i] = &event{s.P1, s, true, i}
		eq[2*i+1] = &event{s.P2, s, false, 2 * i}
	}
	heap.Init(&eq)

	for eq.Len() > 0 {
		e := heap.Pop(&eq).(*event)
		if e.firstPoint {
			fmt.Printf("%v\n", e)
			sl.insert(e.segment)
			slTmp := sl.search(e.segment)

		}
	}

	return
}

type event struct {
	point      Point
	segment    Segment
	firstPoint bool
	index      int
}

func (e event) String() string {
	return fmt.Sprintf("{%v, %v, %v, %v}", e.point, e.segment, e.firstPoint, e.index)
}

type eventQueue []*event

func (eq eventQueue) Len() int { return len(eq) }

func (eq eventQueue) Less(i, j int) bool {
	// Order after X.
	if eq[i].point.X < eq[j].point.X {
		return true
	} else if eq[i].point.X > eq[j].point.X {
		return false
	}

	// X is equal. Order after Y.
	if eq[i].point.Y < eq[j].point.Y {
		return true
	}

	// Definitly not less
	return false
}

func (eq *eventQueue) Push(x interface{}) {
	n := len(*eq)
	item := x.(*event)
	item.index = n
	*eq = append(*eq, item)
}

func (eq *eventQueue) Pop() interface{} {
	old := *eq
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*eq = old[0 : n-1]
	return item
}

func (eq eventQueue) Swap(i, j int) {
	eq[i], eq[j] = eq[j], eq[i]
	eq[i].index = i
	eq[j].index = j
}
