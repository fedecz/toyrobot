package core

import (
	"errors"
	"strings"
)

type Board struct {
	SizeX int
	SizeY int
}

func NewBoard(sizeX, sizeY int) (*Board, error) {
	if sizeX < 0 || sizeY < 0 {
		return nil, errors.New("Can't create a board with negative size")
	}
	return &Board{
		SizeX: sizeX,
		SizeY: sizeY,
	}, nil
}

func (b *Board) IsPositionValid(p *Position) bool {
	return (b.SizeX-1 >= p.X) && (b.SizeY-1 >= p.Y) &&
		(0 <= p.X) && (0 <= p.Y)
}

type Face string

const (
	North Face = "NORTH"
	East  Face = "EAST"
	South Face = "SOUTH"
	West  Face = "WEST"
	None  Face = "NONE"
)

func FaceFromString(s string) (Face, error) {
	if strings.Compare(s, "NORTH") == 0 {
		return North, nil
	}
	if strings.Compare(s, "SOUTH") == 0 {
		return South, nil
	}
	if strings.Compare(s, "EAST") == 0 {
		return East, nil
	}
	if strings.Compare(s, "WEST") == 0 {
		return West, nil
	}
	return None, errors.New("Can't parse that Face")
}

type Position struct {
	X    int
	Y    int
	Face Face
}

func NewPosition(x, y int, face Face) (*Position, error) {
	if x < 0 || y < 0 {
		return nil, errors.New("position can't be negative")
	}
	if face == None {
		return nil, errors.New("face can't be 'none'")
	}
	return &Position{
		X:    x,
		Y:    y,
		Face: face,
	}, nil
}
