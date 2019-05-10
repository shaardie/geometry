package datastructures

import (
	"reflect"
	"testing"
)

type intPriorityQueueItem int

func (i intPriorityQueueItem) Compare(j PriorityQueueItem) int {
	return int(i - j.(intPriorityQueueItem))
}

func TestNewPriorityQueue(t *testing.T) {
	type args struct {
		items []PriorityQueueItem
	}
	tests := []struct {
		name string
		args args
		want *PriorityQueue
	}{
		{
			name: "Empty",
			want: &PriorityQueue{&priorityQueue{}},
		},
		{
			name: "Single Element",
			args: args{
				items: []PriorityQueueItem{intPriorityQueueItem(0)},
			},
			want: &PriorityQueue{&priorityQueue{
				priorityQueueItem{intPriorityQueueItem(0), 0}},
			},
		},
		{
			name: "Multiple Elements",
			args: args{
				items: []PriorityQueueItem{intPriorityQueueItem(0), intPriorityQueueItem(-1), intPriorityQueueItem(1)},
			},
			want: &PriorityQueue{&priorityQueue{
				priorityQueueItem{intPriorityQueueItem(-1), 0},
				priorityQueueItem{intPriorityQueueItem(0), 1},
				priorityQueueItem{intPriorityQueueItem(1), 2}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPriorityQueue(tt.args.items...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPriorityQueue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPriorityQueue_Pop(t *testing.T) {
	tests := []struct {
		name  string
		items []PriorityQueueItem
		want  PriorityQueueItem
	}{
		{
			name:  "simple",
			items: []PriorityQueueItem{intPriorityQueueItem(2), intPriorityQueueItem(1), intPriorityQueueItem(0)},
			want:  intPriorityQueueItem(0),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pq := NewPriorityQueue(tt.items...)
			if got := pq.Pop(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PriorityQueue.Pop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPriorityQueue_Push(t *testing.T) {
	tests := []struct {
		name  string
		item  PriorityQueueItem
		items []PriorityQueueItem
		want  *PriorityQueue
	}{
		{
			name:  "lower",
			item:  intPriorityQueueItem(-1),
			items: []PriorityQueueItem{intPriorityQueueItem(0)},
			want: &PriorityQueue{
				&priorityQueue{
					priorityQueueItem{intPriorityQueueItem(-1), 0},
					priorityQueueItem{intPriorityQueueItem(0), 1},
				},
			},
		},
		{
			name:  "higher",
			item:  intPriorityQueueItem(1),
			items: []PriorityQueueItem{intPriorityQueueItem(0)},
			want: &PriorityQueue{
				&priorityQueue{
					priorityQueueItem{intPriorityQueueItem(0), 0},
					priorityQueueItem{intPriorityQueueItem(1), 1},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pq := NewPriorityQueue(tt.items...)
			pq.Push(tt.item)
			if !reflect.DeepEqual(pq, tt.want) {
				t.Errorf("PriorityQueue.Push() = %v, want %v", pq, tt.want)
			}
		})
	}
}
