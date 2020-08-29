package iterator_test

import (
	"fmt"
	"iterator"
	"testing"
)

func TestIterator(t *testing.T) {
	root := iterator.NewNode(3)
	node1 := iterator.NewNode(1)
	node2 := iterator.NewNode(0)
	node3 := iterator.NewNode(2)
	node4 := iterator.NewNode(4)
	root.Left = node1
	node1.Left = node2
	node1.Right = node3
	root.Right = node4

	it := root.Iterator()
	for i := 0; i < 5; i++ {
		if !it.HasNext() {
			t.Fatal("Error ended HasNext()", it.Cur())
		}
		cur := it.Next()
		fmt.Println("Debug:", cur)
		if cur != i {
			t.Fatal(fmt.Sprintf("Expecting %d but got %d", i, cur))
		}
	}
}
