package simulation

import (
	"testing"
	"toyrobot/core"
)

func TestAllActionsShouldReturnErrorIfThereIsNoPlacement(t *testing.T) {
	b, _ := core.NewBoard(5, 5)
	s, _ := NewSimulation(b)

	if err := s.Move(); err == nil {
		t.Fatal("there should be an error")
	}
	if err := s.Left(); err == nil {
		t.Fatal("there should be an error")
	}
	if err := s.Right(); err == nil {
		t.Fatal("there should be an error")
	}
	if err := s.Report(); err == nil {
		t.Fatal("there should be an error")
	}
}

func errorCheckCommandWithPosition(cmd func(position *core.Position) error, p *core.Position, t *testing.T) {
	if err := cmd(p); err != nil {
		t.Fatal(err.Error())
	}
}
func errorCheckCommand(cmd func() error, t *testing.T) {
	if err := cmd(); err != nil {
		t.Fatal(err.Error())
	}
}

func TestMoveOutOfBounds(t *testing.T) {
	b, _ := core.NewBoard(2, 2)
	s, _ := NewSimulation(b)
	position, _ := core.NewPosition(0, 1, core.North)
	errorCheckCommandWithPosition(s.Place, position, t)
	err := s.Move()
	if err == nil {
		t.Fatal("Should've received an error")
	}

	position, _ = core.NewPosition(0, 0, core.South)
	errorCheckCommandWithPosition(s.Place, position, t)
	err = s.Move()
	if err == nil {
		t.Fatal("Should've received an error")
	}

	position, _ = core.NewPosition(0, 0, core.West)
	errorCheckCommandWithPosition(s.Place, position, t)
	err = s.Move()
	if err == nil {
		t.Fatal("Should've received an error")
	}
}

func TestNewSimulationWithBoard(t *testing.T) {
	b, _ := core.NewBoard(2, 3)
	s, err := NewSimulation(b)
	if err != nil || s == nil {
		t.Fatal("something went wrong creating the simulation")
	}
	if s.board != b {
		t.Fatal("board should be equal")
	}
}

func TestNewSimulationWithNilBoard(t *testing.T) {
	s, err := NewSimulation(nil)
	if err == nil || s != nil {
		t.Fatal("something went wrong creating the simulation")
	}
}

func TestTurningLeft(t *testing.T) {
	b, _ := core.NewBoard(2, 2)
	s, _ := NewSimulation(b)
	position, _ := core.NewPosition(1, 1, core.North)
	errorCheckCommandWithPosition(s.Place, position, t)
	errorCheckCommand(s.Left, t) //west
	errorCheckCommand(s.Left, t) //south
	errorCheckCommand(s.Left, t) //east
	currentPosition := s.currentPosition
	if currentPosition.Face != core.East {
		t.Fatal("Left is not turning left")
	}
	if currentPosition.X != 1 || currentPosition.Y != 1 {
		t.Fatal("robot shouldn't have moved")
	}
}

func TestTurningRight(t *testing.T) {
	b, _ := core.NewBoard(2, 2)
	s, _ := NewSimulation(b)
	position, _ := core.NewPosition(1, 1, core.North)
	errorCheckCommandWithPosition(s.Place, position, t)
	errorCheckCommand(s.Right, t) //East
	errorCheckCommand(s.Right, t) //south
	errorCheckCommand(s.Right, t) //West
	currentPosition := s.currentPosition
	if currentPosition.Face != core.West {
		t.Fatal("Left is not turning left")
	}
	if currentPosition.X != 1 || currentPosition.Y != 1 {
		t.Fatal("robot shouldn't have moved")
	}
}

func TestReportShouldntMoveTheRobot(t *testing.T) {
	b, _ := core.NewBoard(2, 2)
	s, _ := NewSimulation(b)
	position, _ := core.NewPosition(1, 1, core.North)
	errorCheckCommandWithPosition(s.Place, position, t)
	errorCheckCommand(s.Report, t)
	currentPosition := s.currentPosition
	if currentPosition.Face != core.North {
		t.Fatal("Report shouldn't move the robot")
	}
	if currentPosition.X != 1 || currentPosition.Y != 1 {
		t.Fatal("Report shouldn't move the robot")
	}
}

func TestPlaceShouldFailWhenPositionIsIncorrect(t *testing.T) {
	b, _ := core.NewBoard(2, 2)
	s, _ := NewSimulation(b)
	position, _ := core.NewPosition(2, 2, core.North)
	err := s.Place(position)
	if err == nil {
		t.Fatal("there should be an error")
	}
	if err.Error() != "Position is not valid for that board" {
		t.Fatal("Got another error")
	}
}

func TestValidRun(t *testing.T) {
	b, _ := core.NewBoard(3, 3)
	s, _ := NewSimulation(b)
	position, _ := core.NewPosition(1, 1, core.North)
	errorCheckCommandWithPosition(s.Place, position, t)
	errorCheckCommand(s.Move, t)  //1,2,N
	errorCheckCommand(s.Right, t) //1,2,E
	errorCheckCommand(s.Move, t)  //2,2,E
	errorCheckCommand(s.Right, t) //2,2,S
	errorCheckCommand(s.Move, t)  //2,1,S
	errorCheckCommand(s.Move, t)  //2,0,S
	err := s.Move()               // invalid move - 2,0,S
	if err == nil {
		t.Fatal("Should've got an error")
	}
	errorCheckCommand(s.Right, t) //2,0,W
	errorCheckCommand(s.Move, t)  //1,0,W
	currentPosition := s.currentPosition
	if currentPosition.Face != core.West {
		t.Fatal("Face should be West")
	}
	if currentPosition.X != 1 || currentPosition.Y != 0 {
		t.Fatal("position should be (1,0)")
	}
}
