package array

import (
	"errors"
	"fmt"
)

/*
数据的插入，删除，下标的随机访问
*/

type Array struct {
	data   []int
	length uint
}

// 初始化数据内存
func NewArray(cap uint) *Array {
	if cap == 0 {
		return nil
	}
	return &Array{make([]int, cap), 0}
}

// len 返回数组长度
func (a *Array) Len() uint {
	return a.length
}

// 实现下标的随机访问 [0, n-1]
func (a *Array) Get(index uint) (int, error) {
	if a.isOutOfIndex(index) {
		return 0, errors.New("out of index")
	}
	return a.data[index], nil
}

// 索引越界判断
func (a *Array) isOutOfIndex(index uint) bool {
	return index >= uint(cap(a.data))
}

/**
1 3 4
1 2 3 4
*/
// 实现指定插入
func (a *Array) Insert(index uint, value int) error {
	if a.Len() == uint(cap(a.data)) {
		return errors.New("array is full")
	}
	if a.isOutOfIndex(index) {
		return errors.New("out of index")
	}
	for i := a.length; i > index; i-- {
		a.data[i] = a.data[i-1]
	}
	a.data[index] = value
	a.length++
	return nil
}

// 向后插入 append方法追加
func (a *Array) InsertToTail(value int) error {
	return a.Insert(a.Len(), value)
}

// 删除索引上的值
func (a *Array) Delete(index uint) (int, error) {
	if a.isOutOfIndex(index) {
		return 0, errors.New("out of index")
	}
	value := a.data[index]
	for i := index; i < a.Len()-1; i++ {
		a.data[i] = a.data[i+1]
	}
	a.length--
	return value, nil
}

// 打印数列
func (a *Array) Print() {
	var format string
	for i := uint(0); i < a.Len(); i++ {
		format += fmt.Sprintf("| %+v", a.data[i])
	}
	fmt.Println(format)
}
