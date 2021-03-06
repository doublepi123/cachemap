package priorityqueue

type Item struct {
	Priority int
	Value    interface{}
}

type PQueue []Item

func (pq PQueue) Len() int {
	return len(pq)
}

func (pq PQueue) Less(i, j int) bool {
	return pq[i].Priority < pq[j].Priority
}

func (pq PQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
func (pq *PQueue) Push(x interface{}) {
	*pq = append(*pq, x.(Item))
}
func (pq *PQueue) Pop() interface{} {
	temp := (*pq)[pq.Len()-1]
	*pq = (*pq)[0 : pq.Len()-1]
	return temp
}
