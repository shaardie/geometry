package datastructures

import (
	"reflect"
	"testing"
)

type intBinarySearchTreeItem int

func (i intBinarySearchTreeItem) Compare(j BinarySearchTreeItem) int {
	return int(i - j.(intBinarySearchTreeItem))
}

func TestBinarySearchTree_Min(t *testing.T) {
	type fields struct {
		Item  BinarySearchTreeItem
		Left  *BinarySearchTree
		Right *BinarySearchTree
	}
	tests := []struct {
		name   string
		fields fields
		want   BinarySearchTreeItem
	}{
		{"Root", fields{Item: intBinarySearchTreeItem(0)}, intBinarySearchTreeItem(0)},
		{
			"Left",
			fields{Item: intBinarySearchTreeItem(0), Left: &BinarySearchTree{Item: intBinarySearchTreeItem(-1)}},
			intBinarySearchTreeItem(-1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bst := &BinarySearchTree{
				Item:  tt.fields.Item,
				Left:  tt.fields.Left,
				Right: tt.fields.Right,
			}
			if got := bst.Min(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BinarySearchTree.Min() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinarySearchTree_Max(t *testing.T) {
	type fields struct {
		Item  BinarySearchTreeItem
		Left  *BinarySearchTree
		Right *BinarySearchTree
	}
	tests := []struct {
		name   string
		fields fields
		want   BinarySearchTreeItem
	}{
		{"Root", fields{Item: intBinarySearchTreeItem(0)}, intBinarySearchTreeItem(0)},
		{
			"Right",
			fields{Item: intBinarySearchTreeItem(0), Right: &BinarySearchTree{Item: intBinarySearchTreeItem(1)}},
			intBinarySearchTreeItem(1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bst := &BinarySearchTree{
				Item:  tt.fields.Item,
				Left:  tt.fields.Left,
				Right: tt.fields.Right,
			}
			if got := bst.Max(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BinarySearchTree.Max() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinarySearchTree_Insert(t *testing.T) {
	type fields struct {
		Item  BinarySearchTreeItem
		Left  *BinarySearchTree
		Right *BinarySearchTree
	}
	type args struct {
		item BinarySearchTreeItem
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *BinarySearchTree
	}{
		{
			name: "Empty",
			args: args{intBinarySearchTreeItem(0)},
			want: &BinarySearchTree{Item: intBinarySearchTreeItem(0)},
		},
		{
			name:   "Same twice",
			fields: fields{Item: intBinarySearchTreeItem(0)},
			args:   args{intBinarySearchTreeItem(0)},
			want:   &BinarySearchTree{Item: intBinarySearchTreeItem(0)},
		},
		{
			name:   "Insert left",
			fields: fields{Item: intBinarySearchTreeItem(0)},
			args:   args{intBinarySearchTreeItem(-1)},
			want: &BinarySearchTree{
				Item: intBinarySearchTreeItem(0),
				Left: &BinarySearchTree{Item: intBinarySearchTreeItem(-1)},
			},
		},
		{
			name:   "Insert right",
			fields: fields{Item: intBinarySearchTreeItem(0)},
			args:   args{intBinarySearchTreeItem(1)},
			want: &BinarySearchTree{
				Item:  intBinarySearchTreeItem(0),
				Right: &BinarySearchTree{Item: intBinarySearchTreeItem(1)},
			},
		},
		{
			name: "Insert recursive left",
			fields: fields{
				Item: intBinarySearchTreeItem(0),
				Left: &BinarySearchTree{Item: intBinarySearchTreeItem(-1)},
			},
			args: args{intBinarySearchTreeItem(-2)},
			want: &BinarySearchTree{
				Item: intBinarySearchTreeItem(0),
				Left: &BinarySearchTree{
					Item: intBinarySearchTreeItem(-1),
					Left: &BinarySearchTree{Item: intBinarySearchTreeItem(-2)},
				},
			},
		},
		{
			name: "Insert recursive right",
			fields: fields{
				Item:  intBinarySearchTreeItem(0),
				Right: &BinarySearchTree{Item: intBinarySearchTreeItem(1)},
			},
			args: args{intBinarySearchTreeItem(2)},
			want: &BinarySearchTree{
				Item: intBinarySearchTreeItem(0),
				Right: &BinarySearchTree{
					Item:  intBinarySearchTreeItem(1),
					Right: &BinarySearchTree{Item: intBinarySearchTreeItem(2)},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bst := &BinarySearchTree{
				Item:  tt.fields.Item,
				Left:  tt.fields.Left,
				Right: tt.fields.Right,
			}
			bst.Insert(tt.args.item)
			if !reflect.DeepEqual(bst, tt.want) {
				t.Errorf("BinarySearchTree.Search() = %v, want %v", bst, tt.want)
			}
		})
	}
}

func TestBinarySearchTree_Search(t *testing.T) {
	type fields struct {
		Item  BinarySearchTreeItem
		Left  *BinarySearchTree
		Right *BinarySearchTree
	}
	type args struct {
		item BinarySearchTreeItem
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "Empty",
			args: args{intBinarySearchTreeItem(0)},
		},
		{
			name:   "Root",
			fields: fields{Item: intBinarySearchTreeItem(0)},
			args:   args{intBinarySearchTreeItem(0)},
			want:   true,
		},
		{
			name:   "Not present",
			fields: fields{Item: intBinarySearchTreeItem(0)},
			args:   args{intBinarySearchTreeItem(1)},
			want:   false,
		},
		{
			name: "Found left",
			fields: fields{
				Item: intBinarySearchTreeItem(0),
				Left: &BinarySearchTree{Item: intBinarySearchTreeItem(-1)},
			},
			args: args{intBinarySearchTreeItem(-1)},
			want: true,
		},
		{
			name: "Found right",
			fields: fields{
				Item:  intBinarySearchTreeItem(0),
				Right: &BinarySearchTree{Item: intBinarySearchTreeItem(1)},
			},
			args: args{intBinarySearchTreeItem(1)},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bst := &BinarySearchTree{
				Item:  tt.fields.Item,
				Left:  tt.fields.Left,
				Right: tt.fields.Right,
			}
			if got := bst.Search(tt.args.item); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BinarySearchTree.Search() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinarySearchTree_Delete(t *testing.T) {
	type fields struct {
		Item  BinarySearchTreeItem
		Left  *BinarySearchTree
		Right *BinarySearchTree
	}
	type args struct {
		item BinarySearchTreeItem
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
		after  *BinarySearchTree
	}{
		{
			name: "Empty",
			args: args{intBinarySearchTreeItem(0)},
			want: false,
		},
		{
			name:   "Root",
			fields: fields{Item: intBinarySearchTreeItem(0)},
			args:   args{intBinarySearchTreeItem(0)},
			want:   true,
			after:  &BinarySearchTree{},
		},
		{
			name:   "Not found left",
			fields: fields{Item: intBinarySearchTreeItem(0)},
			args:   args{intBinarySearchTreeItem(-1)},
			want:   false,
			after:  &BinarySearchTree{Item: intBinarySearchTreeItem(0)},
		},
		{
			name:   "Not found right",
			fields: fields{Item: intBinarySearchTreeItem(0)},
			args:   args{intBinarySearchTreeItem(1)},
			want:   false,
			after:  &BinarySearchTree{Item: intBinarySearchTreeItem(0)},
		},
		{
			name: "delete Left",
			fields: fields{
				Item: intBinarySearchTreeItem(0),
				Left: &BinarySearchTree{Item: intBinarySearchTreeItem(-1)},
			},
			args: args{intBinarySearchTreeItem(-1)},
			want: true,
			after: &BinarySearchTree{
				Item: intBinarySearchTreeItem(0),
				Left: &BinarySearchTree{},
			},
		},
		{
			name: "delete Right",
			fields: fields{
				Item:  intBinarySearchTreeItem(0),
				Right: &BinarySearchTree{Item: intBinarySearchTreeItem(1)},
			},
			args: args{intBinarySearchTreeItem(1)},
			want: true,
			after: &BinarySearchTree{
				Item:  intBinarySearchTreeItem(0),
				Right: &BinarySearchTree{},
			},
		},
		{
			name: "delete with left arm",
			fields: fields{
				Item: intBinarySearchTreeItem(0),
				Left: &BinarySearchTree{Item: intBinarySearchTreeItem(-1)},
			},
			args: args{intBinarySearchTreeItem(0)},
			want: true,
			after: &BinarySearchTree{
				Item: intBinarySearchTreeItem(-1),
			},
		},
		{
			name: "delete with right arm",
			fields: fields{
				Item:  intBinarySearchTreeItem(0),
				Right: &BinarySearchTree{Item: intBinarySearchTreeItem(1)},
			},
			args: args{intBinarySearchTreeItem(0)},
			want: true,
			after: &BinarySearchTree{
				Item: intBinarySearchTreeItem(1),
			},
		},
		{
			name: "delete with arms",
			fields: fields{
				Item:  intBinarySearchTreeItem(0),
				Left:  &BinarySearchTree{Item: intBinarySearchTreeItem(-1)},
				Right: &BinarySearchTree{Item: intBinarySearchTreeItem(1)},
			},
			args: args{intBinarySearchTreeItem(0)},
			want: true,
			after: &BinarySearchTree{
				Item: intBinarySearchTreeItem(1),
				Left: &BinarySearchTree{
					Item: intBinarySearchTreeItem(-1),
				},
				Right: &BinarySearchTree{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bst := &BinarySearchTree{
				Item:  tt.fields.Item,
				Left:  tt.fields.Left,
				Right: tt.fields.Right,
			}
			if got := bst.Delete(tt.args.item); got != tt.want {
				t.Errorf("BinarySearchTree.Delete() = %v, want %v", got, tt.want)
			}
			if tt.after != nil && !reflect.DeepEqual(bst, tt.after) {
				t.Errorf("BinarySearchTree.Search() = %v, want %v", bst, tt.after)
			}
		})
	}
}
