package array

import (
	"testing"
)

func TestAppend(t *testing.T) {
	cap := 10
	arr := NewArray(cap)
	for i := 0; i < cap; i++ {
		err := arr.Append(i, i+1)
		if nil != err {
			t.Fatal(err.Error())
		}
	}
}

func TestFind(t *testing.T) {
	cap := 10
	arr := NewArray(cap)
	for i := 0; i < cap; i++ {
		err := arr.Append(i, i+1)
		if err != nil {
			t.Fatal(err.Error())
		}
	}
	arr.Print()
	t.Log(arr.Find(0))
	t.Log(arr.Find(5))
	t.Log(arr.Find(9))
}

func TestDelete(t *testing.T) {
	cap := 10
	arr := NewArray(cap)
	for i := 0; i < cap; i++ {
		err := arr.Append(i, i+1)
		if err != nil {
			t.Fatal(err.Error())
		}
	}
	for i := 9; i >= 0; i-- {
		_, err := arr.Delete(i)
		if err != nil {
			t.Fatal(err.Error())
		}
	}
}
