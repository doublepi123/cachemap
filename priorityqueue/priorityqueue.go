package priorityqueue

type item struct {
	priority int
	value    interface{}
}

type priorityqueue []item

func (pq priorityqueue) Len() int {
	return len(pq)
}

func (pq priorityqueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq priorityqueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
func (pq *priorityqueue) Push(x interface{}) {
	*pq = append(*pq, x.(item))
}
func (pq *priorityqueue) Pop() interface{} {
	temp := (*pq)[pq.Len()-1]
	*pq = (*pq)[0 : pq.Len()-1]
	return temp
}
