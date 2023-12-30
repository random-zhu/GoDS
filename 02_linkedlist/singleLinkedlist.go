package linkedlist

/*
	单链表的基本实现
	value
*/

type ListNode struct {
	next  *ListNode
	value int
}

type LindedList struct {
	head   *ListNode
	length uint
}

func NewLindedList() *ListNode {

}
