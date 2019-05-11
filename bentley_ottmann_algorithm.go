package geometry

import (
	"github.com/shaardie/geometry/datastructures"
)

func BentleyOttmannAlgorithm(segments []Segment) ([]Point, error) {

	// Initialize Event Queue
	startEvents := make([]datastructures.PriorityQueueItem, 2*len(segments))
	for i, s := range segments {
		s.Sort()
		startEvents[i] = event{s.P1, s, true}
		startEvents[2*i] = event{s.P2, s, false}
	}
	eventQueue := datastructures.NewPriorityQueue(startEvents...)

	// Initialize Sweep Line
	sweepLine := datastructures.BinarySearchTree{}

	// Initialize intersection points
	intersections := make([]Point, 0)

}

type event struct {
	point      Point
	segment    Segment
	firstPoint bool
}

func (e event) Compare(pqi datastructures.PriorityQueueItem) int {
	e2 := pqi.(event)

	// X differ.
	if r := int(e.point.X - e2.point.X); r != 0 {
		return r

	}

	// X equal. Order with Y.
	return int(e.point.Y - e2.point.Y)
}
