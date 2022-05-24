package cli

import (
	"testing"
	"toyrobot/cmd"
)

func TestCli_ParseLineMove(t *testing.T) {
	cli := Cli{}
	moveCmd, _ := cli.ParseLine("MOVE")
	_, ok := moveCmd.(cmd.MoveCmd)
	if !ok {
		t.Fatal("is not the right type")
	}
}

func TestCli_ParseLineShouldGetAnErrorIfIncorrect(t *testing.T) {
	cli := Cli{}
	_, err := cli.ParseLine("SOMETHINGELSE")
	if err == nil {
		t.Fatal("err should not be nil")
	}
}

func TestCli_ParseLineRight(t *testing.T) {
	cli := Cli{}
	command, _ := cli.ParseLine("RIGHT")
	_, ok := command.(cmd.RightCmd)
	if !ok {
		t.Fatal("is not the right type")
	}
}

func TestCli_ParseLineLeft(t *testing.T) {
	cli := Cli{}
	command, _ := cli.ParseLine("LEFT")
	_, ok := command.(cmd.LeftCmd)
	if !ok {
		t.Fatal("is not the right type")
	}
}

func TestCli_ParseLineReport(t *testing.T) {
	cli := Cli{}
	command, _ := cli.ParseLine("REPORT")
	_, ok := command.(cmd.ReportCmd)
	if !ok {
		t.Fatal("is not the right type")
	}
}

func TestCli_ParseLinePlace(t *testing.T) {
	cli := Cli{}
	command, err := cli.ParseLine("PLACE 0,0,NORTH")
	if err != nil {
		t.Fatal("An error was not expected here")
	}
	_, ok := command.(cmd.PlaceCmd)
	if !ok {
		t.Fatal("is not the right type")
	}
}

func TestCli_ParseLinePlaceWithWrongFace(t *testing.T) {
	cli := Cli{}
	_, err := cli.ParseLine("PLACE 0,0,SOMETHING")
	if err == nil {
		t.Fatal("I was expecting an error")
	}
}

func TestCli_ParseLinePlaceWithWrongFormat(t *testing.T) {
	cli := Cli{}
	_, err := cli.ParseLine("PLACE 0,0,NORTH,0")
	if err == nil {
		t.Fatal("I was expecting an error")
	}

	_, err = cli.ParseLine("PLACE")
	if err == nil {
		t.Fatal("I was expecting an error")
	}

	_, err = cli.ParseLine("PLACE 0")
	if err == nil {
		t.Fatal("I was expecting an error")
	}

	_, err = cli.ParseLine("PLACE 1,b,c")
	if err == nil {
		t.Fatal("I was expecting an error")
	}
}
