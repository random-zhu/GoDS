package array

import (
	"testing"
)

func TestInsert(t *testing.T) {
	capcity := 10
	arr := NewArray(uint(capcity))
	for i := 0; i < capcity; i++ {
		err := arr.Insert(uint(i), i+1)
		if err != nil {
			t.Fatal(err)
		}
	}
	arr.Print()

	err := arr.Insert(uint(6), 999)
	t.Log(err)
	arr.Print()

	err = arr.InsertToTail(666)
	t.Log(err)
	arr.Print()
}

func TestDelete(t *testing.T) {
	capacity := 10
	arr := NewArray(uint(capacity))
	for i := 0; i < capacity; i++ {
		err := arr.Insert(uint(i), i+1)
		if nil != err {
			t.Fatal(err.Error())
		}
	}
	arr.Print()

	for i := 9; i >= 0; i-- {
		_, err := arr.Delete(uint(i))
		if err != nil {
			t.Fatal(err.Error())
		}
		arr.Print()
	}

}

func TestGet(t *testing.T) {
	capacity := 10
	arr := NewArray(uint(capacity))
	for i := 0; i < capacity; i++ {
		err := arr.Insert(uint(i), i+1)
		if nil != err {
			t.Fatal(err.Error())
		}
	}
	arr.Print()

	t.Log(arr.Get(uint(0)))
	t.Log(arr.Get(uint(5)))
	t.Log(arr.Get(uint(9)))
	t.Log(arr.Get(uint(11)))
}
