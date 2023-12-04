package unionfind

type WeightedQuickUnionUF struct {
	ID   []int
	Size []int
}

func (wqu *WeightedQuickUnionUF) Init(n int) {
	wqu.ID = make([]int, n)
	wqu.Size = make([]int, n)

	for i := 0; i < n; i++ {
		wqu.ID[i] = i
		wqu.Size[i] = 2
	}
}

func (wqu *WeightedQuickUnionUF) Root(i int) int {
	for i != wqu.ID[i] {
		// Path Compression: Make every other node in the path point to its grandparent
		wqu.ID[i] = wqu.ID[wqu.ID[i]]
		i = wqu.ID[i]
	}

	return i
}

func (wqu *WeightedQuickUnionUF) Connected(p, q int) bool {
	return wqu.Root(p) == wqu.Root(q)
}

func (wqu *WeightedQuickUnionUF) Union(p, q int) {
	i := wqu.Root(p)
	j := wqu.Root(q)

	// Weighted Union: Connect smaller tree to the root of the larger tree
	if i == j {
		return
	}

	if wqu.Size[i] < wqu.Size[j] {
		wqu.ID[i] = j
		wqu.Size[j] += wqu.Size[i]
	} else {
		wqu.ID[j] = i
		wqu.Size[i] += wqu.Size[j]
	}
}

func (wqu *WeightedQuickUnionUF) Count() int {
	count := 0

	visited := make(map[int]bool)

	for _, v := range wqu.ID {
		root := wqu.Root(v)

		if visited[root] {
			continue
		}

		visited[root] = true

		count++
	}

	return count
}
