package basemdl

import "fmt"

// MatrixN represents a table of numeric elements
type MatrixN[N NumericType] struct {
	elems [][]N
	rows  uint16
	cols  uint16
}

func (m *MatrixN[N]) Set(xPos, yPos uint16, value N) error {

	if m.rows-1 < xPos {
		return fmt.Errorf("Position greater than max rows")
	}

	if m.cols-1 < yPos {
		return fmt.Errorf("Position greater than max columns")
	}

	m.elems[xPos][yPos] = value
	return nil
}

func (m MatrixN[N]) ValueAt(xPos, yPos uint16) (N, error) {

	if m.rows-1 < xPos {
		return 0, fmt.Errorf("Position greater than max rows")
	}

	if m.cols-1 < yPos {
		return 0, fmt.Errorf("Position greater than max columns")
	}

	return m.elems[xPos][yPos], nil
}
