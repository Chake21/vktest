package internal

import "sync"

type IntArray struct {
	data  []int32
	mutex sync.Mutex
}

func NewIntArray() *IntArray {
	return &IntArray{
		data:  make([]int32, 0),
		mutex: sync.Mutex{},
	}
}

func (arr *IntArray) Append(value int32) {
	arr.mutex.Lock()
	defer arr.mutex.Unlock()

	arr.data = append(arr.data, value)
}

func (arr *IntArray) Sum() int32 {
	arr.mutex.Lock()
	defer arr.mutex.Unlock()

	var sum int32
	for _, value := range arr.data {
		sum += value
	}
	return sum
}

type BoolArray struct {
	data  []bool
	mutex sync.Mutex
}

func NewBoolArray() *BoolArray {
	return &BoolArray{
		data:  make([]bool, 0),
		mutex: sync.Mutex{},
	}
}

func (arr *BoolArray) Append(value bool) {
	arr.mutex.Lock()
	defer arr.mutex.Unlock()

	arr.data = append(arr.data, value)
}

func (arr *BoolArray) AnyTrue() bool {
	arr.mutex.Lock()
	defer arr.mutex.Unlock()

	for _, value := range arr.data {
		if value {
			return true
		}
	}
	return false
}
