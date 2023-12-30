package linkedlist

import "fmt"

/*
	单链表的基本实现
	value int
*/

type ListNode struct {
	next  *ListNode
	value int
}

type LindedList struct {
	head   *ListNode
	length uint
}

// 构造函数
func NewListNode(value int) *ListNode {
	return &ListNode{
		value: value,
	}
}

// 构造函数
func NewLindedList() *LindedList {
	return &LindedList{
		head:   NewListNode(0),
		length: 0,
	}
}

// 尾部插入节点
func (l *LindedList) Insert(value int) {
	newNode := NewListNode(value)
	current := l.head

	for current.next != nil {
		current = current.next
	}
	current.next = newNode
}

// 插入指定节点
func (list *LindedList) InsertAfter(value int, node *ListNode) bool {
	if node == nil {
		return false
	}
	newNode := NewListNode(value)
	newNode.next = node.next
	node.next = newNode
	return true
}

// 查找指定值节点
func (list *LindedList) Find(value int) *ListNode {
	current := list.head.next
	for current != nil {
		if current.value == value {
			return current
		}
		current = current.next
	}
	return nil
}

// 删除指定值的节点
func (list *LindedList) Delete(value int) bool {
	current := list.head
	for current != nil {
		if current.next.value == value {
			current.next = current.next.next
			return true
		}
		current = current.next
	}
	return false
}

// 打印链表
func (list *LindedList) Print() {
	start := list.head.next
	var format string
	for start != nil {
		format += fmt.Sprintf("| %+v", start.value)
		start = start.next
	}
	fmt.Println(format)
}
