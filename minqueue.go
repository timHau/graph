package gograph

import "container/heap"

type Item[T any] struct {
	node  *Node[T]
	prio  int
	index int
}

type MinQueue[T any] []*Item[T]

func (mq MinQueue[T]) Len() int { return len(mq) }

func (mq MinQueue[T]) Less(i, j int) bool {
	return mq[i].prio > mq[j].prio
}

func (mq MinQueue[T]) Swap(i, j int) {
	mq[i], mq[j] = mq[j], mq[i]
	mq[i].index = i
	mq[j].index = j
}

func (mq *MinQueue[T]) Push(x any) {
	n := len(*mq)
	item := x.(*Item[T])
	item.index = n
	*mq = append(*mq, item)
}

func (mq *MinQueue[T]) Pop() any {
	old := *mq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*mq = old[0 : n-1]
	return item
}

func (mq *MinQueue[T]) update(item *Item[T], node *Node[T], prio int) {
	item.node = node
	item.prio = prio
	heap.Fix(mq, item.index)
}
