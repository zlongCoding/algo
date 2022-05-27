package array

import (
	"errors"
	"fmt"
)

//数组find append delete

type Array struct {
	data   []int
	length int
}

//初始化数组
func NewArray(cap int) *Array {
	if cap == 0 {
		return nil
	}
	return &Array{
		data:   make([]int, cap, cap),
		length: 0,
	}
}

//获取其length
func (this *Array) Len() int {
	return this.length
}

//判断数组是否越界了
func (this *Array) indexOutOfRanger(index int) bool {
	//cap --> 返回指定类型的容量
	return index > cap(this.data)
}

//find
func (this *Array) Find(index int) (int, error) {
	if this.indexOutOfRanger(index) {
		return 0, errors.New("index is out of range")
	}
	return this.data[index], nil
}

func (this *Array) capacityArray() error {
	len := cap(this.data) * 2
	arr := make([]int, len, len)
	for i := 0; i < len; i++ {
		arr[i] = this.data[i]
	}
	this.data = arr
	return nil
}

//Append
func (this *Array) Append(index int, val int) error {
	if this.Len() == cap(this.data) {
		//数组满了，可以进行扩容
		// capacityArray
		return errors.New("array is full")
	}
	if this.indexOutOfRanger(index) {
		return errors.New("index is out of range")
	}
	for i := this.length; i > index; i-- {
		this.data[i] = this.data[i-1]
	}
	this.data[index] = val
	this.length++
	return nil
}

//Delete
func (this *Array) Delete(index int) (int, error) {
	if this.indexOutOfRanger(index) {
		return 0, errors.New("index is out of range")
	}
	val := this.data[index]
	for i := index; i < this.length-1; i++ {
		this.data[i] = this.data[i+1]
	}
	this.length--
	return val, nil
}

//print array
func (this *Array) Print() {
	var format string
	for _, v := range this.data {
		format += fmt.Sprintf("%+v\n", v)
	}
	fmt.Println(format)
}
