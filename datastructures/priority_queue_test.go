package datastructures

import (
	"reflect"
	"testing"
)

type intItem int

func (i intItem) Compare(j PriorityQueueItem) int {
	return int(i - j.(intItem))
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
				items: []PriorityQueueItem{intItem(0)},
			},
			want: &PriorityQueue{&priorityQueue{
				priorityQueueItem{intItem(0), 0}},
			},
		},
		{
			name: "Multiple Elements",
			args: args{
				items: []PriorityQueueItem{intItem(0), intItem(-1), intItem(1)},
			},
			want: &PriorityQueue{&priorityQueue{
				priorityQueueItem{intItem(-1), 0},
				priorityQueueItem{intItem(0), 1},
				priorityQueueItem{intItem(1), 2}},
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
			items: []PriorityQueueItem{intItem(2), intItem(1), intItem(0)},
			want:  intItem(0),
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
			item:  intItem(-1),
			items: []PriorityQueueItem{intItem(0)},
			want: &PriorityQueue{
				&priorityQueue{
					priorityQueueItem{intItem(-1), 0},
					priorityQueueItem{intItem(0), 1},
				},
			},
		},
		{
			name:  "higher",
			item:  intItem(1),
			items: []PriorityQueueItem{intItem(0)},
			want: &PriorityQueue{
				&priorityQueue{
					priorityQueueItem{intItem(0), 0},
					priorityQueueItem{intItem(1), 1},
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
