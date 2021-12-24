package day23

type PQ []interface{}

func (pq PQ) Len() int { return len(pq) }
func (pq PQ) Less(i, j int) bool {
	return pq[i].(Move).H() < pq[j].(Move).H()
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].(Move).SetIndex(i)
	pq[j].(Move).SetIndex(j)
}

func (pq *PQ) Push(x interface{}) {
	n := len(*pq)
	item := x.(Move)
	item.SetIndex(n)
	*pq = append(*pq, item)
}

func (pq *PQ) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.(Move).SetIndex(-1)
	*pq = old[0 : n-1]
	return item
}
