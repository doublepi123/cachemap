package priorityqueue

type item struct {
	priority int
	value    interface{}
}

type PQueue []item

func (pq PQueue) Len() int {
	return len(pq)
}

func (pq PQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
func (pq *PQueue) Push(x interface{}) {
	*pq = append(*pq, x.(item))
}
func (pq *PQueue) Pop() interface{} {
	temp := (*pq)[pq.Len()-1]
	*pq = (*pq)[0 : pq.Len()-1]
	return temp
}
