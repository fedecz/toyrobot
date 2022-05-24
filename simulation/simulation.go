package simulation

import (
	"errors"
	"log"
	"toyrobot/core"
)

type Simulation struct {
	board           *core.Board
	currentPosition *core.Position
}

func NewSimulation(b *core.Board) (*Simulation, error) {
	if b == nil {
		return nil, errors.New("board can't be nil")
	}
	return &Simulation{
		board: b,
	}, nil
}

func (s *Simulation) Move() error {
	if err := s.hasSimulationStarted(); err != nil {
		return err
	}
	var futureX int = int(s.currentPosition.X)
	var futureY int = int(s.currentPosition.Y)

	switch s.currentPosition.Face {
	case core.North:
		futureY = futureY + 1
	case core.South:
		futureY = futureY - 1
	case core.East:
		futureX = futureX + 1
	case core.West:
		futureX = futureX - 1
	}
	if futureX < 0 || futureY < 0 {
		return errors.New("Invalid move")
	}
	newPosition := &core.Position{
		X:    futureX,
		Y:    futureY,
		Face: s.currentPosition.Face,
	}
	if s.board.IsPositionValid(newPosition) == false {
		return errors.New("Invalid move")
	}
	s.currentPosition = newPosition
	return nil
}

func (s *Simulation) Left() error {
	if err := s.hasSimulationStarted(); err != nil {
		return err
	}
	switch s.currentPosition.Face {
	case core.North:
		s.currentPosition.Face = core.West
	case core.East:
		s.currentPosition.Face = core.North
	case core.South:
		s.currentPosition.Face = core.East
	case core.West:
		s.currentPosition.Face = core.South
	}
	return nil
}

func (s *Simulation) Right() error {
	if err := s.hasSimulationStarted(); err != nil {
		return err
	}
	switch s.currentPosition.Face {
	case core.North:
		s.currentPosition.Face = core.East
	case core.East:
		s.currentPosition.Face = core.South
	case core.South:
		s.currentPosition.Face = core.West
	case core.West:
		s.currentPosition.Face = core.North
	}
	return nil
}

func (s *Simulation) Report() error {
	if err := s.hasSimulationStarted(); err != nil {
		return err
	}
	log.Printf("Current position of robot is %d, %d, %s\n", s.currentPosition.X, s.currentPosition.Y, s.currentPosition.Face)
	return nil
}

func (s *Simulation) Place(position *core.Position) error {
	if s.board.IsPositionValid(position) {
		s.currentPosition = position
		return nil
	} else {
		return errors.New("Position is not valid for that board")
	}
}

func (s *Simulation) hasSimulationStarted() error {
	if s.currentPosition == nil {
		return errors.New("simulation doesn't seem to have started")
	}
	return nil
}
