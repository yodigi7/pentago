package pentago

import (
	"reflect"
	"testing"
)

func TestIsDraw(t *testing.T) {
	game := NewGame()
	game.Board = [6][6]Space{
		{White, Black, White, Black, White, Black},
		{Black, White, Black, White, Black, White},
		{White, White, White, White, Black, White},
		{Black, White, Black, Black, White, Black},
		{Black, Black, White, Black, White, Black},
		{Black, White, Black, White, Black, White},
	}
	if !game.IsDraw() {
		t.Error("Expected game to be draw but it wasn't")
	}
	game.Board = [6][6]Space{
		{White, Black, White, Black, White, Black},
		{Empty, White, Black, White, Black, White},
		{White, White, White, White, Black, White},
		{Black, White, Black, Black, White, Black},
		{Black, Black, White, Black, White, Black},
		{Black, White, Black, White, Black, White},
	}
	if game.IsDraw() {
		t.Error("Expected game to be not a draw but it wasn't")
	}
	game.Board = [6][6]Space{
		{White, Black, White, Black, White, Black},
		{Black, White, Black, White, Black, White},
		{Black, White, White, White, Black, White},
		{Black, White, Black, Black, White, Black},
		{Black, Black, White, Black, White, Black},
		{Black, White, Black, White, Black, White},
	}
	if game.IsDraw() {
		t.Error("Expected game to be not a draw but it wasn't")
	}
}

func TestRotateClockwise(t *testing.T) {
	grid := [3][3]Space{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	expected := [3][3]Space{
		{7, 4, 1},
		{8, 5, 2},
		{9, 6, 3},
	}

	result := rotateClockwise(grid)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("rotateClockwise() = %v, want %v", result, expected)
	}
}

func TestRotateCounterClockwise(t *testing.T) {
	grid := [3][3]Space{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	expected := [3][3]Space{
		{3, 6, 9},
		{2, 5, 8},
		{1, 4, 7},
	}

	result := rotateCounterClockwise(grid)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("rotateCounterClockwise() = %v, want %v", result, expected)
	}
}

func TestGetStarts(t *testing.T) {
	_, _, err := getStarts(5)
	if err != InvalidQuadrant {
		t.Errorf("Expected error thrown on invalid quadrant")
	}
}

func TestPlaceMarbleErrors(t *testing.T) {
	game := NewGame()
	game.Board[5][4] = White
	err := game.PlaceMarble(5, 4)
	if err == nil {
		t.Errorf("Expected error when SpaceOccupied")
	}
	err = game.PlaceMarble(6, 4)
	if err == nil {
		t.Errorf("Expected error when InvalidCoordinates")
	}
	err = game.PlaceMarble(4, 6)
	if err == nil {
		t.Errorf("Expected error when InvalidCoordinates")
	}
}

func TestPlaceMarble(t *testing.T) {
	game := NewGame()
	game.PlaceMarble(0, 1)
	game.PlaceMarble(0, 2)
	game.PlaceMarble(0, 3)
	game.PlaceMarble(0, 4)
	game.PlaceMarble(0, 5)
	for i := 1; i < 6; i++ {
		if game.Board[0][i] == Empty {
			t.Errorf("Expected marble to be placed at (0, %d)", i)
		}
	}
}

func TestRotateQuadrantSwitchesTurns(t *testing.T) {
	game := NewGame()
	firstPlayer := game.Turn
	game.RotateQuadrant(TopLeft, Clockwise)
	if game.Turn == firstPlayer {
		t.Errorf("Expected turn to change")
	}
}

func TestCheckForWinner(t *testing.T) {
	// vertically
	game := Game{
		Board: [6][6]Space{
			{0, 0, 0, 0, 0, 0},
			{Black, 0, 0, 0, 0, 0},
			{Black, 0, 0, 0, 0, 0},
			{Black, 0, 0, 0, 0, 0},
			{Black, 0, 0, 0, 0, 0},
			{Black, 0, 0, 0, 0, 0},
		},
	}
	winner := game.CheckForWinner()
	if winner != Black {
		t.Errorf("Expected Black to win")
	}
	// horizontally
	game = Game{
		Board: [6][6]Space{
			{0, 0, 0, 0, 0, 0},
			{0, Black, Black, Black, Black, Black},
			{0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0},
		},
	}
	winner = game.CheckForWinner()
	if winner != Black {
		t.Errorf("Expected Black to win")
	}
	// diagonally
	game = Game{
		Board: [6][6]Space{
			{0, 0, 0, 0, 0, 0},
			{0, White, 0, 0, 0, 0},
			{0, 0, White, 0, 0, 0},
			{0, 0, 0, White, 0, 0},
			{0, 0, 0, 0, White, 0},
			{0, 0, 0, 0, 0, White},
		},
	}
	winner = game.CheckForWinner()
	if winner != White {
		t.Errorf("Expected White to win")
	}
	// None
	game = Game{
		Board: [6][6]Space{
			{0, 0, 0, 0, 0, 0},
			{0, White, 0, 0, 0, 0},
			{0, 0, Black, 0, 0, 0},
			{0, 0, 0, White, 0, 0},
			{0, 0, 0, 0, White, 0},
			{0, 0, 0, 0, 0, White},
		},
	}
	winner = game.CheckForWinner()
	if winner != Empty {
		t.Errorf("Expected no one to win")
	}
}

func TestRotateQuadrant(t *testing.T) {
	game := Game{
		Board: [6][6]Space{
			{1, 2, 3, 4, 5, 6},
			{7, 8, 9, 10, 11, 12},
			{13, 14, 15, 16, 17, 18},
			{19, 20, 21, 22, 23, 24},
			{25, 26, 27, 28, 29, 30},
			{31, 32, 33, 34, 35, 36},
		},
	}

	tests := []struct {
		quadrant  Quadrant
		direction RotationDirection
		expected  [6][6]Space
	}{
		{
			quadrant:  TopLeft,
			direction: Clockwise,
			expected: [6][6]Space{
				{13, 7, 1, 4, 5, 6},
				{14, 8, 2, 10, 11, 12},
				{15, 9, 3, 16, 17, 18},
				{19, 20, 21, 22, 23, 24},
				{25, 26, 27, 28, 29, 30},
				{31, 32, 33, 34, 35, 36},
			},
		},
		{
			quadrant:  TopLeft,
			direction: CounterClockwise,
			expected: [6][6]Space{
				{3, 9, 15, 4, 5, 6},
				{2, 8, 14, 10, 11, 12},
				{1, 7, 13, 16, 17, 18},
				{19, 20, 21, 22, 23, 24},
				{25, 26, 27, 28, 29, 30},
				{31, 32, 33, 34, 35, 36},
			},
		},
		{
			quadrant:  TopRight,
			direction: Clockwise,
			expected: [6][6]Space{
				{1, 2, 3, 16, 10, 4},
				{7, 8, 9, 17, 11, 5},
				{13, 14, 15, 18, 12, 6},
				{19, 20, 21, 22, 23, 24},
				{25, 26, 27, 28, 29, 30},
				{31, 32, 33, 34, 35, 36},
			},
		},
		{
			quadrant:  TopRight,
			direction: CounterClockwise,
			expected: [6][6]Space{
				{1, 2, 3, 6, 12, 18},
				{7, 8, 9, 5, 11, 17},
				{13, 14, 15, 4, 10, 16},
				{19, 20, 21, 22, 23, 24},
				{25, 26, 27, 28, 29, 30},
				{31, 32, 33, 34, 35, 36},
			},
		},
		{
			quadrant:  BottomLeft,
			direction: Clockwise,
			expected: [6][6]Space{
				{1, 2, 3, 4, 5, 6},
				{7, 8, 9, 10, 11, 12},
				{13, 14, 15, 16, 17, 18},
				{31, 25, 19, 22, 23, 24},
				{32, 26, 20, 28, 29, 30},
				{33, 27, 21, 34, 35, 36},
			},
		},
		{
			quadrant:  BottomLeft,
			direction: CounterClockwise,
			expected: [6][6]Space{
				{1, 2, 3, 4, 5, 6},
				{7, 8, 9, 10, 11, 12},
				{13, 14, 15, 16, 17, 18},
				{21, 27, 33, 22, 23, 24},
				{20, 26, 32, 28, 29, 30},
				{19, 25, 31, 34, 35, 36},
			},
		},
		{
			quadrant:  BottomRight,
			direction: Clockwise,
			expected: [6][6]Space{
				{1, 2, 3, 4, 5, 6},
				{7, 8, 9, 10, 11, 12},
				{13, 14, 15, 16, 17, 18},
				{19, 20, 21, 34, 28, 22},
				{25, 26, 27, 35, 29, 23},
				{31, 32, 33, 36, 30, 24},
			},
		},
		{
			quadrant:  BottomRight,
			direction: CounterClockwise,
			expected: [6][6]Space{
				{1, 2, 3, 4, 5, 6},
				{7, 8, 9, 10, 11, 12},
				{13, 14, 15, 16, 17, 18},
				{19, 20, 21, 24, 30, 36},
				{25, 26, 27, 23, 29, 35},
				{31, 32, 33, 22, 28, 34},
			},
		},
	}

	// Run the tests
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			// Reset the board before each test
			game.Board = [6][6]Space{
				{1, 2, 3, 4, 5, 6},
				{7, 8, 9, 10, 11, 12},
				{13, 14, 15, 16, 17, 18},
				{19, 20, 21, 22, 23, 24},
				{25, 26, 27, 28, 29, 30},
				{31, 32, 33, 34, 35, 36},
			}

			// Call RotateQuadrant
			game.RotateQuadrant(tt.quadrant, tt.direction)

			// Check if the board matches the expected result
			if !reflect.DeepEqual(game.Board, tt.expected) {
				t.Errorf("RotateQuadrant() = %v, want %v", game.Board, tt.expected)
			}
		})
	}
}
