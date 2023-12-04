package unionfind

type QuickUnionUF struct {
	ID []int
}

// O(n)
func (qu *QuickUnionUF) Init(n int) {
	qu.ID = make([]int, n)

	for i := 0; i < n; i++ {
		qu.ID[i] = i
	}
}

// O(h) -> h is the height of the tree
func (qu *QuickUnionUF) Root(i int) int {
	for i != qu.ID[i] {
		i = qu.ID[i]
	}

	return i
}

// O(h)
func (qu *QuickUnionUF) Connected(p, q int) bool {
	return qu.Root(p) == qu.Root(q)
}

// O(h)
func (qu *QuickUnionUF) Union(p, q int) {
	i := qu.Root(p)

	j := qu.Root(q)

	qu.ID[i] = j
}

// O(n*h)
func (qu *QuickUnionUF) Count() int {
	count := 0

	visited := make(map[int]bool)

	for _, v := range qu.ID {
		root := qu.Root(v)

		if visited[root] {
			continue
		}

		visited[root] = true

		count++
	}

	return count
}
