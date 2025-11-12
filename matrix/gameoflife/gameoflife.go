package gameoflife

func gameOfLife(board [][]int) {
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}

	next := make([][]int, len(board))
	for i := 0; i < len(board); i++ {
		next[i] = make([]int, len(board[0]))
	}

	newGame := &game{
		current: board,
		next:    next,
		height:  len(board),
		width:   len(board[0]),
		gen:     0,
	}
	newGame.nextGen()

	// Copy the next state back to the original board
	for i := 0; i < len(board); i++ {
		copy(board[i], newGame.current[i])
	}
}

type game struct {
	current, next      [][]int
	height, width, gen int
}

func (g *game) nextGen() {
	g.processAll()
	g.gen += 1
	g.current, g.next = g.next, g.current
}

func (g *game) processAll() {
	for i := 0; i < g.height; i++ {
		for j := 0; j < g.width; j++ {
			nCount := g.countNeighbors(i, j)
			g.next[i][j] = g.applyRule(g.current[i][j], nCount)
		}
	}
}

func (g *game) countNeighbors(x, y int) int {
	count := 0
	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			if dy == 0 && dx == 0 {
				continue
			}
			neighborX, neighborY := x+dx, y+dy
			if neighborX >= 0 && neighborX < g.height && neighborY >= 0 && neighborY < g.width {
				if g.current[neighborX][neighborY] == 1 {
					count++
				}
			}
		}
	}
	return count
}

func (g *game) applyRule(current, neighborCount int) int {
	if current == 1 {
		if neighborCount == 2 || neighborCount == 3 {
			return 1
		}
		return 0
	} else {
		if neighborCount == 3 {
			return 1
		}
		return 0
	}
}

func gameOfLifeInPlace(board [][]int) {
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}

	// Directions: 8 neighbors
	directions := [][]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	m, n := len(board), len(board[0])

	// First pass: encode next state in the second bit
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			liveNeighbors := 0
			for _, dir := range directions {
				ni, nj := i+dir[0], j+dir[1]
				if ni >= 0 && ni < m && nj >= 0 && nj < n {
					// Check if was alive in previous state
					if board[ni][nj] == 1 || board[ni][nj] == 2 {
						liveNeighbors++
					}
				}
			}

			// Apply rules
			if board[i][j] == 1 {
				if liveNeighbors < 2 || liveNeighbors > 3 {
					// Live cell dies: 1 -> 2 (was alive, will die)
					board[i][j] = 2
				}
				// Otherwise stays alive: 1 -> 1
			} else {
				if liveNeighbors == 3 {
					// Dead cell becomes alive: 0 -> 3 (was dead, will live)
					board[i][j] = 3
				}
			}
		}
	}

	// Second pass: decode the next state
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 2 {
				board[i][j] = 0
			} else if board[i][j] == 3 {
				board[i][j] = 1
			}
		}
	}
}
