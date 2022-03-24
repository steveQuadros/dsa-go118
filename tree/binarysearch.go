package tree

type Tree interface {
	Search(val comparable) bool
	Insert(val comparable)
	Min() comparable
	Max() comparable
	Delete(val comparable) bool
}

type Node struct {
	Val         comparable
	Left, Right *Node
}

// AVL is a self-balancing Binary Search Tree (BST)
// where the difference between heights of left and right subtrees cannot be more than one for all nodes.
type AVL struct{}

var _ Tree = (*AVL)(nil)

func (t *AVL) Search(val comparable) bool {
	return false
}

// Insert inserts a value into the tree
// To make sure that the given tree remains AVL after every insertion,
// we must augment the standard BST insert operation to perform some re-balancing.
// Following are two basic operations that can be performed to re-balance a BST without violating
// the BST property (keys(left) < key(root) < keys(right)).
// 1) Left Rotation
// 2) Right Rotation
func (t *AVL) Insert(val comparable) {

}

func (t *AVL) Min() comparable {
	return -1
}

func (t *AVL) Max() comparable {
	return -1
}

func (t *AVL) Delete(val comparable) bool {
	return false
}
