package heap

import "sort"

type Interface interface {
	sort.Interface
	Push(x interface{})
	Pop() interface{}
}

func Init(h Interface) {
	n := h.Len()
	for i := n/2 - 1; i >= 0; i-- {
		down(h, i, n)
	}
}

func down(h Interface, i, n int) bool {
	t := i
	// for循环内t和i的使用容易出错
	for {
		il := 2*t + 1
		if il >= n || il < 0 {
			break
		}
		smallest := il
		ir := il + 1
		if ir < n && h.Less(ir, il) {
			smallest = ir
		}
		if !h.Less(smallest, t) {
			break
		}
		h.Swap(t, smallest)
		t = smallest
	}
	return t > i
}

func up(h Interface, i int) {
	for {
		p := (i - 1) / 2
		if p == i || h.Less(p, i) {
			break
		}
		h.Swap(p, i)
		i = p

	}
}

func Push(h Interface, x interface{}) {
	h.Push(x)
	up(h, h.Len()-1)
}

func Pop(h Interface) interface{} {
	n := h.Len() - 1
	h.Swap(0, n)
	down(h, 0, n)
	return h.Pop()
}

func Fix(h Interface, i int) {
	if !down(h, i, h.Len()-1) {
		up(h, i)
	}
}

func Remove(h Interface, i int) interface{} {
	h.Swap(i, h.Len()-1)
	if !down(h, i, h.Len()-1) {
		up(h, i)
	}
	return h.Pop()
}
