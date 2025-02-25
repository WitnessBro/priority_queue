package priority_queue

import (
	"container/heap"
	"sync"
)

type Item struct {
	Value    string
	Priority int
	Index    int
}

type PriorityQueue struct {
	sync.Mutex
	items []*Item
}

func New() *PriorityQueue {
	pq := &PriorityQueue{
		items: make([]*Item, 0),
	}
	heap.Init(pq)

	return pq
}

func NewItem(value string, priority int) *Item {
	item := Item{
		Value:    value,
		Priority: priority,
	}
	return &item
}
func (pq *PriorityQueue) Len() int {
	return len(pq.items)
}
func (pq *PriorityQueue) Less(i, j int) bool {
	return pq.items[i].Priority > pq.items[j].Priority
}

func (pq *PriorityQueue) Push(x any) {
	n := len(pq.items)
	item := x.(*Item)
	item.Index = n
	pq.items = append(pq.items, item)
}
func (pq *PriorityQueue) Swap(i, j int) {
	pq.items[i], pq.items[j] = pq.items[j], pq.items[i]
	pq.items[i].Index = i
	pq.items[j].Index = j
}
func (pq *PriorityQueue) Update(item *Item, value string, priority int) {
	pq.Lock()
	defer pq.Unlock()
	item.Value = value
	item.Priority = priority
	heap.Fix(pq, item.Index)
}
func (pq *PriorityQueue) IsEmpty() bool {
	pq.Lock()
	defer pq.Unlock()
	return pq.Len() == 0
}
func (pq *PriorityQueue) Pop() any {
	old := pq.items
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.Index = -1
	pq.items = old[0 : n-1]
	return item
}
func (pq *PriorityQueue) Put(value string, priority int) {
	pq.Lock()
	defer pq.Unlock()
	item := &Item{
		Value:    value,
		Priority: priority,
	}
	heap.Push(pq, item)
}

func (pq *PriorityQueue) Get() *Item {
	pq.Lock()
	defer pq.Unlock()
	item := heap.Pop(pq).(*Item)
	return item
}
func (pq *PriorityQueue) GetMax() *Item {
	pq.Lock()
	defer pq.Unlock()
	max := &Item{}
	for _, item := range pq.items {
		if item.Priority > max.Priority {
			max = item
		}
	}
	return max
}
