package core

import "testing"

func TestNewBoard(t *testing.T) {
	b, err := NewBoard(3, 3)
	if err != nil {
		t.Fatal("wasn't expecting an error")
	}
	if b.SizeX != 3 || b.SizeY != 3 {
		t.Fatal("size is wrong")
	}
}

func TestNewBoardWithNegativeSize(t *testing.T) {
	b, err := NewBoard(-3, -3)
	if err == nil {
		t.Fatal("wasn't expecting an error")
	}
	if b != nil {
		t.Fatal("board should be nil")
	}
}

func TestBoard_IsPositionValidAt00(t *testing.T) {
	b := Board{
		SizeX: 5,
		SizeY: 5,
	}

	if !b.IsPositionValid(&Position{
		X:    0,
		Y:    0,
		Face: North,
	}) {
		t.Fatal("0,0 should be a valid case")
	}
}

func TestBoard_IsPositionValidAtOutside(t *testing.T) {
	b := Board{
		SizeX: 5,
		SizeY: 5,
	}

	if b.IsPositionValid(&Position{
		X:    6,
		Y:    6,
		Face: North,
	}) {
		t.Fatal("6,6 should be an invalid case")
	}
}

func TestBoard_IsPositionValidWithNegative(t *testing.T) {
	b := Board{
		SizeX: 5,
		SizeY: 5,
	}

	if b.IsPositionValid(&Position{
		X:    -2,
		Y:    5,
		Face: North,
	}) {
		t.Fatal("position should not be valid")
	}

	if b.IsPositionValid(&Position{
		X:    2,
		Y:    -5,
		Face: North,
	}) {
		t.Fatal("position should not be valid")
	}
	if b.IsPositionValid(&Position{
		X:    -2,
		Y:    -5,
		Face: North,
	}) {
		t.Fatal("position should not be valid")
	}
}

func TestBoard_IsPositionValidWithSmallBoard(t *testing.T) {
	b := Board{
		SizeX: 1,
		SizeY: 1,
	}

	if !b.IsPositionValid(&Position{
		X:    0,
		Y:    0,
		Face: North,
	}) {
		t.Fatal("0,0 should be a valid case")
	}

	if b.IsPositionValid(&Position{
		X:    1,
		Y:    1,
		Face: North,
	}) {
		t.Fatal("0,0 should be a valid case")
	}
}

func TestFaceFromString(t *testing.T) {
	f, _ := FaceFromString("NORTH")
	if f != North {
		t.Fatal("Should be a valid face")
	}

	f, _ = FaceFromString("SOUTH")
	if f != South {
		t.Fatal("Should be a valid face")
	}
	f, _ = FaceFromString("EAST")
	if f != East {
		t.Fatal("Should be a valid face")
	}
	f, _ = FaceFromString("WEST")
	if f != West {
		t.Fatal("Should be a valid face")
	}

	f, _ = FaceFromString("ANY")
	if f != None {
		t.Fatal("Should be a valid face")
	}
}

func TestNewPosition(t *testing.T) {
	newPosition, err := NewPosition(2, 3, North)
	if err != nil {
		t.Fatal("Wasn't expecting an error here.")
	}
	if newPosition.X != 2 || newPosition.Y != 3 || newPosition.Face != North {
		t.Fatal("Values injected are not correct")
	}
}

func TestNewPositionWithNegatives(t *testing.T) {
	newPosition, err := NewPosition(-2, -3, North)
	if err == nil {
		t.Fatal("There should be an error here")
	}
	if newPosition != nil {
		t.Fatal("position should be nil")
	}
}
