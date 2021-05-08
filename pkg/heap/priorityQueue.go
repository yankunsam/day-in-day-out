package heap

import (
	"log"
	"reflect"
)

type Item struct {
	value    string
	priority int
	index    int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	if i == j {
		return
	}
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = j
	pq[j].index = i
}

func (pq *PriorityQueue) Push(x interface{}) {
	// 类型判断
	if reflect.TypeOf(x).Kind() != reflect.TypeOf(pq).Kind() {
		log.Println("the type of x not support")
	}
	n := pq.Len()
	x.(*Item).index = n
	*pq = append(*pq, x.(*Item))
}

func (pq *PriorityQueue) Pop() interface{} {
	res := (*pq)[(*pq).Len()-1]
	// avoid memory leak
	(*pq)[(*pq).Len()-1] = nil
	*pq = (*pq)[0 : (*pq).Len()-1]
	return res
}

func (pq *PriorityQueue) update(item *Item, value string, priority int) {
	item.value = value
	item.priority = priority
	Fix(pq, item.index)
}
