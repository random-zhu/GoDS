package linkedlist

import "testing"

func TestInsert(t *testing.T) {
	l := NewLindedList()
	for i := 1; i <= 10; i++ {
		l.Insert(i)
	}
	l.Print()
}

func TestInsertAfter(t *testing.T) {
	l := NewLindedList()
	for i := 1; i <= 10; i++ {
		l.Insert(i)
	}
	l.Print()
	node := l.Find(5)
	l.InsertAfter(100, node)
	l.Print()
}

func TestFind(t *testing.T) {
	l := NewLindedList()
	for i := 1; i <= 10; i++ {
		l.Insert(i)
	}
	l.Print()

	node := l.Find(5)
	t.Log(node)
	node = l.Find(1)
	t.Log(node)
	node = l.Find(10)
	t.Log(node)
}

func TestDelete(t *testing.T) {
	l := NewLindedList()
	for i := 1; i <= 10; i++ {
		l.Insert(i)
	}
	l.Print()
	l.Delete(2)
	l.Print()
	l.Delete(1)
	l.Print()
	l.Delete(9)
	l.Print()
	l.Delete(10)
	l.Print()
}
