package heap

import (
	"log"
	"testing"
)

type tHeap []int

func (th *tHeap) Less(i, j int) bool {
	return (*th)[i] < (*th)[j]
}

func (th *tHeap) Len() int {
	return len(*th)
}

func (th *tHeap) Swap(i, j int) {
	(*th)[i], (*th)[j] = (*th)[j], (*th)[i]
}

func (th *tHeap) Push(v interface{}) {
	*th = append(*th, v.(int))
}

func (th *tHeap) Pop() interface{} {
	res := (*th)[(*th).Len()-1]
	*th = (*th)[0 : (*th).Len()-1]
	return res
}

func TestInit(t *testing.T) {
	th := new(tHeap)
	for i := 20; i > 0; i-- {
		th.Push(i)
	}
	Init(th)
	for th.Len() > 0 {
		log.Println(Pop(th))
	}
}
