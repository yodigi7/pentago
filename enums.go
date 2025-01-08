package pentago

type Space int

const (
	Empty Space = iota
	White
	Black
)

func (s Space) String() string {
	switch s {
	case Empty:
		return "Empty"
	case White:
		return "White"
	case Black:
		return "Black"
	default:
		return "Invalid"
	}
}

type Quadrant int

const (
	TopLeft     Quadrant = 1
	TopRight    Quadrant = 2
	BottomLeft  Quadrant = 3
	BottomRight Quadrant = 4
)

type RotationDirection int

const (
	Clockwise RotationDirection = iota
	CounterClockwise
)
