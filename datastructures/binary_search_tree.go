package datastructures

// BinarySearchTreeItem is the items of the BinarySearchTree.
// Must implement Compare function.
type BinarySearchTreeItem interface {
	// Compare with the Parameter.
	// If smaller return < 0, if greater return > 0 and 0 if equal.
	Compare(BinarySearchTreeItem) int
}

// BinarySearchTree is a binary search tree, see https://en.wikipedia.org/wiki/Binary_search_tree
type BinarySearchTree struct {
	Item  BinarySearchTreeItem
	Left  *BinarySearchTree
	Right *BinarySearchTree
}

// Min returns mimimum Item in the Binary Search Tree
func (bst *BinarySearchTree) Min() BinarySearchTreeItem {
	if bst.Left == nil {
		return bst.Item
	}
	return bst.Left.Min()
}

// Max returns maximum Item in the Binary Search Tree
func (bst *BinarySearchTree) Max() BinarySearchTreeItem {
	if bst.Right == nil {
		return bst.Item
	}
	return bst.Right.Max()
}

// Insert new item to the Binary Search Tree, return Tree with item as root
func (bst *BinarySearchTree) Insert(item BinarySearchTreeItem) {

	// Tree is empty
	if bst.Item == nil {
		bst.Item = item
		return
	}

	cmp := item.Compare(bst.Item)
	if cmp < 0 {
		// Left spot is empty. Use it
		if bst.Left == nil {
			bst.Left = &BinarySearchTree{Item: item}
			return
		}
		// Call recursice on left tree
		bst.Left.Insert(item)
		return
	} else if cmp > 0 {
		// Right spot is empty. Use it
		if bst.Right == nil {
			bst.Right = &BinarySearchTree{Item: item}
			return
		}
		// Call recursice on right tree
		bst.Right.Insert(item)
		return
	}
	return
}

// Search for an item in the Binary Search Tree
func (bst *BinarySearchTree) Search(item BinarySearchTreeItem) bool {
	if bst.Item == nil {
		return false
	}
	current := bst
	for current != nil {
		cmp := item.Compare(current.Item)
		if cmp < 0 {
			current = current.Left
			continue
		} else if cmp > 0 {
			current = current.Right
			continue
		}
		return true
	}
	// Not in the tree
	return false
}

// Delete an item from the Binary Search Tree
func (bst *BinarySearchTree) Delete(item BinarySearchTreeItem) bool {
	// Empty BinarySearchTree
	if bst.Item == nil {
		return false
	}

	cmp := item.Compare(bst.Item)
	switch {
	case cmp < 0:
		if bst.Left == nil {
			return false
		}
		return bst.Left.Delete(item)
	case cmp > 0:
		if bst.Right == nil {
			return false
		}
		return bst.Right.Delete(item)
	default:
		if bst.Left == nil && bst.Right == nil {
			bst.Item = nil
			return true
		}
		if bst.Left == nil {
			bst.Item, bst.Left, bst.Right = bst.Right.Item, bst.Right.Left, bst.Right.Right
			return true
		}
		if bst.Right == nil {
			bst.Item, bst.Left, bst.Right = bst.Left.Item, bst.Left.Left, bst.Left.Right
			return true
		}
		// Find miminal Item in the Right arm, use it as new Item in Root
		// and delete it from the right arm.
		minItem := bst.Right.Min()
		bst.Item = minItem
		return bst.Right.Delete(minItem)
	}
}
