package iterator

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Value int
}

func NewNode(value int) *TreeNode {
	return &TreeNode{
		Value: value,
	}
}

func (node *TreeNode) Iterator() Iterator {
	return &MorrisInOrderIterator{
		node: node,
	}
}

type Iterator interface {
	HasNext() bool
	Next() int
	Cur() int
}

// InInOrderIterator is implemented in a morris way which will actually modify the tree structure.
type MorrisInOrderIterator struct {
	node *TreeNode
}

func (it *MorrisInOrderIterator) HasNext() bool {
	return it.node != nil
}

func (it *MorrisInOrderIterator) Cur() int {
	return it.node.Value
}

func (it *MorrisInOrderIterator) Next() int {
	var res int
	cur := it.node
	for cur != nil {
		left := cur.Left
		if left == nil {
			res = cur.Value
			cur = cur.Right
			break
		} else {
			for left.Right != nil && left.Right != cur {
				left = left.Right
			}
			if left.Right == cur {
				res = cur.Value
				cur = cur.Right
				left.Right = nil
				break
			} else {
				left.Right = cur
				cur = cur.Left
			}
		}
	}

	it.node = cur

	return res
}
