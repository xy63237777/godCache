package algorithm

import (
	"testing"
)

func TestLinkedList_Remove(t *testing.T) {
	list := NewLinkedList()
	list.Add(3)
	list.Add(5)
	list.Show()
	list.Remove(4)
	list.Show()
	list.Remove(5)
	list.Show()
}
