package builder

import (
	"fmt"
)

type Square struct {
	height int
	width  int
}

func (s *Square) Validate() bool {
	return s.height == s.width
}

type SquareBuilder struct {
	height int
	width  int
}

func (s *SquareBuilder) SetHeight(height int) *SquareBuilder {
	s.height = height
	return s
}

func (s *SquareBuilder) SetWidth(width int) *SquareBuilder {
	s.width = width
	return s
}

func (s *SquareBuilder) BuildSquare() (*Square, error) {
	if s.height == 0 || s.width == 0 || s.height != s.width {
		return nil, fmt.Errorf("height %d is not the same as width %d", s.height, s.width)
	}
	return &Square{height: s.height, width: s.width}, nil
}
