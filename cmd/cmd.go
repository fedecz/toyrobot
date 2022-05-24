package cmd

import (
	"log"
	"toyrobot/core"
	"toyrobot/simulation"
)

type Command interface {
	Execute(s *simulation.Simulation) error
}

type MoveCmd struct {
}

func (m MoveCmd) Execute(s *simulation.Simulation) error {
	log.Println("Executing MOVE command")
	return s.Move()
}

type LeftCmd struct {
}

func (m LeftCmd) Execute(s *simulation.Simulation) error {
	log.Println("Executing LEFT command")
	return s.Left()
}

type RightCmd struct {
}

func (m RightCmd) Execute(s *simulation.Simulation) error {
	log.Println("Executing RIGHT command")
	return s.Right()
}

type ReportCmd struct {
}

func (m ReportCmd) Execute(s *simulation.Simulation) error {
	log.Println("Executing REPORT command")
	return s.Report()
}

type PlaceCmd struct {
	Position core.Position
}

func (m PlaceCmd) Execute(s *simulation.Simulation) error {
	log.Println("Executing PLACE command")
	return s.Place(&m.Position)
}
