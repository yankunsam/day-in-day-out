package heap

import (
	"log"
	"testing"
)

func TestPriorityQueue(t *testing.T) {
	items := map[string]int{
		"banana": 3,
		"apple":  2,
		"pear":   4,
	}
	pq := make(PriorityQueue, len(items))
	i := 0
	for value, priority := range items {
		pq[i] = &Item{
			value:    value,
			priority: priority,
			index:    i,
		}
		i++
	}
	Init(&pq)
	item := &Item{
		value:    "orange",
		priority: 1,
	}
	Push(&pq, item)
	pq.update(item, item.value, 5)
	for pq.Len() > 0 {
		log.Println(Pop(&pq).(*Item))
	}
}
