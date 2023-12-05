package unionfind

type Percolation struct {
	grid   [][]bool
	len    int
	top    int
	bottom int
	uf     WeightedQuickUnionUF
	count  int
}

func (p *Percolation) Init(n int) {
	p.len = n

	p.grid = make([][]bool, n)

	for i := 0; i < n; i++ {
		p.grid[i] = make([]bool, n)
	}

	p.bottom = n*n + 1

	p.uf.Init(n*n + 2)
}

// Open opens the site (row, col) if it is not open already
func (p *Percolation) Open(row, col int) {
	if row < 1 || col < 1 || row > p.len || col > p.len {
		panic("IllegalArgumentException")
	}

	i := row - 1
	j := col - 1

	if !p.IsOpen(row, col) {
		p.count++

		p.grid[i][j] = true

		if row == 1 {
			p.uf.Union(p.GetIndex(row, col), p.top)
		}

		if row == p.len {
			p.uf.Union(p.GetIndex(row, col), p.bottom)
		}

		// union the neighbors
		// down
		if row < p.len && p.IsOpen(row+1, col) {
			p.uf.Union(p.GetIndex(row, col), p.GetIndex(row+1, col))
		}
		// up
		if row > 1 && p.IsOpen(row-1, col) {
			p.uf.Union(p.GetIndex(row, col), p.GetIndex(row-1, col))
		}
		// left
		if col > 1 && p.IsOpen(row, col-1) {
			p.uf.Union(p.GetIndex(row, col), p.GetIndex(row, col-1))
		}
		// right
		if col < p.len && p.IsOpen(row, col+1) {
			p.uf.Union(p.GetIndex(row, col), p.GetIndex(row, col+1))
		}
	}
}

// IsOpen is the site (row, col) open?
func (p *Percolation) IsOpen(row, col int) bool {
	if row < 1 || col < 1 || row > p.len || col > p.len {
		panic("IllegalArgumentException")
	}

	return p.grid[row-1][col-1]
}

// IsFull is the site (row, col) full?
func (p *Percolation) IsFull(row, col int) bool {
	if row < 1 || col < 1 || row > p.len || col > p.len {
		panic("IllegalArgumentException")
	}

	return (p.uf.Root(p.GetIndex(row, col)) == p.uf.Root(p.top))
}

// NumberOfOpenSites returns the number of open sites
func (p *Percolation) NumberOfOpenSites() int {
	return p.count
}

// Percolates does the system percolate?
func (p *Percolation) Percolates() bool {
	return (p.uf.Root(p.top) == p.uf.Root(p.bottom))
}

// GetIndex get the index of row,column
func (p *Percolation) GetIndex(row, col int) int {
	return p.len*(row-1) + (col)
}
