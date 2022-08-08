package gograph

import "container/heap"

type Item struct {
	prio  float64
	index int
}

type MinQueue []*Item

func (mq MinQueue) Len() int { return len(mq) }

func (mq MinQueue) Less(i, j int) bool {
	return mq[i].prio > mq[j].prio
}

func (mq MinQueue) Swap(i, j int) {
	mq[i], mq[j] = mq[j], mq[i]
	mq[i].index = i
	mq[j].index = j
}

func (mq *MinQueue) Push(x any) {
	n := len(*mq)
	item := x.(*Item)
	item.index = n
	*mq = append(*mq, item)
}

func (mq *MinQueue) Pop() any {
	old := *mq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*mq = old[0 : n-1]
	return item
}

func (mq *MinQueue) UpdatePrio(item *Item, prio float64) {
	item.prio = prio
	heap.Fix(mq, item.index)
}
