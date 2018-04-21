package binarySearchTree

type BinarySearchTree struct {
	value     int
	parent    *BinarySearchTree
	leftNode  *BinarySearchTree
	rightNode *BinarySearchTree
}

func (tr *BinarySearchTree) Count() int {
	leftCount := 0
	if tr.leftNode != nil {
		leftCount = tr.leftNode.Count()
	}
	rightCount := 0
	if tr.rightNode != nil {
		leftCount = tr.rightNode.Count()
	}
	return leftCount + 1 + rightCount
}

func (tr *BinarySearchTree) Insert(value int) {
	// TODO: to implement
}
