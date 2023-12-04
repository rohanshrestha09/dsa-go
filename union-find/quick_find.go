package unionfind

type QuickFindUF struct {
	ID []int
}

func (qf *QuickFindUF) Init(n int) {
	qf.ID = make([]int, n)

	for i := 0; i < n; i++ {
		qf.ID[i] = i
	}
}

// O(1)
func (qf *QuickFindUF) Connected(p, q int) bool {
	return qf.ID[p] == qf.ID[q]
}

// O(n)
func (qf *QuickFindUF) Union(p, q int) {
	pid := qf.ID[p]

	qid := qf.ID[q]

	for i := 0; i < len(qf.ID); i++ {
		if qf.ID[i] == pid {
			qf.ID[i] = qid
		}
	}
}

// O(n)
func (qf *QuickFindUF) Count() int {
	count := 0

	visited := make(map[int]bool)

	for _, v := range qf.ID {
		if !visited[v] {
			visited[v] = true
			count++
		}
	}

	return count
}
