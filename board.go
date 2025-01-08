package pentago

type Game struct {
	Board  [6][6]Space
	Turn   Space
	Winner Space
}

func NewGame() *Game {
	return &Game{
		Board:  [6][6]Space{},
		Turn:   White,
		Winner: Empty,
	}
}

func (g *Game) checkForWinner() Space {
	// Check horizontal and vertical lines
	for i := 0; i < 6; i++ {
		for j := 0; j < 2; j++ { // Only check up to the 2nd-to-last row and column
			// Check horizontally
			if g.Board[i][j] != Empty && g.Board[i][j] == g.Board[i][j+1] && g.Board[i][j+1] == g.Board[i][j+2] &&
				g.Board[i][j+2] == g.Board[i][j+3] && g.Board[i][j+3] == g.Board[i][j+4] {
				g.Winner = g.Board[i][j]
				return g.Winner
			}

			// Check vertically
			if g.Board[j][i] != Empty && g.Board[j][i] == g.Board[j+1][i] && g.Board[j+1][i] == g.Board[j+2][i] &&
				g.Board[j+2][i] == g.Board[j+3][i] && g.Board[j+3][i] == g.Board[j+4][i] {
				g.Winner = g.Board[j][i]
				return g.Winner
			}
		}
	}

	// Check diagonals (both directions)
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			// Check diagonal from top-left to bottom-right
			if g.Board[i][j] != Empty && g.Board[i][j] == g.Board[i+1][j+1] && g.Board[i+1][j+1] == g.Board[i+2][j+2] &&
				g.Board[i+2][j+2] == g.Board[i+3][j+3] && g.Board[i+3][j+3] == g.Board[i+4][j+4] {
				g.Winner = g.Board[i][j]
				return g.Winner
			}

			// Check diagonal from top-right to bottom-left
			if g.Board[i][5-j] != Empty && g.Board[i][5-j] == g.Board[i+1][4-j] && g.Board[i+1][4-j] == g.Board[i+2][3-j] &&
				g.Board[i+2][3-j] == g.Board[i+3][2-j] && g.Board[i+3][2-j] == g.Board[i+4][1-j] {
				g.Winner = g.Board[i][5-j]
				return g.Winner
			}
		}
	}

	// No winner found
	return Empty
}

// Handle placing a marble
func (g *Game) PlaceMarble(row, col int) error {
	if row > 5 || col > 5 || row < 0 || col < 0 {
		return InvalidCoordinates
	}
	if g.Board[row][col] != Empty {
		return SpaceOccupied
	}
	g.Board[row][col] = g.Turn
	return nil
}

func (g *Game) switchTurn() {
	if g.Turn == White {
		g.Turn = Black
	} else {
		g.Turn = White
	}
}

// Quadrants:
// 1 | 2
// ------
// 3 | 4
func getStarts(quadrant Quadrant) (int, int, error) {
	switch quadrant {
	case TopLeft:
		return 0, 0, nil
	case TopRight:
		return 0, 3, nil
	case BottomLeft:
		return 3, 0, nil
	case BottomRight:
		return 3, 3, nil
	default:
		return -1, -1, InvalidQuadrant
	}
}

func (g *Game) RotateQuadrant(quadrant Quadrant, direction RotationDirection) error {
	startRow, startCol, err := getStarts(quadrant)
	if err != nil {
		return err
	}

	// Extract the 3x3 quadrant
	subGrid := [3][3]Space{}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			subGrid[i][j] = g.Board[startRow+i][startCol+j]
		}
	}

	// Rotate the quadrant
	switch direction {
	case Clockwise:
		subGrid = rotateClockwise(subGrid)
	case CounterClockwise:
		subGrid = rotateCounterClockwise(subGrid)
	}

	// Place the rotated quadrant back into the board
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			g.Board[startRow+i][startCol+j] = subGrid[i][j]
		}
	}

	g.checkForWinner()
	g.switchTurn()
	return nil
}

func rotateClockwise(grid [3][3]Space) [3][3]Space {
	var newGrid [3][3]Space
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			newGrid[j][2-i] = grid[i][j]
		}
	}
	return newGrid
}

func rotateCounterClockwise(grid [3][3]Space) [3][3]Space {
	var newGrid [3][3]Space
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			newGrid[2-j][i] = grid[i][j]
		}
	}
	return newGrid
}
