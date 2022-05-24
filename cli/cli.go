package cli

import (
	"errors"
	"fmt"
	"strings"
	"toyrobot/cmd"
	"toyrobot/core"
)

type Cli struct {
}

func (c *Cli) ParseLine(line string) (cmd.Command, error) {
	if line == "MOVE" {
		return cmd.MoveCmd{}, nil
	}
	if line == "LEFT" {
		return cmd.LeftCmd{}, nil
	}
	if line == "RIGHT" {
		return cmd.RightCmd{}, nil
	}
	if line == "REPORT" {
		return cmd.ReportCmd{}, nil
	}
	if strings.HasPrefix(line, "PLACE") {
		var faceStr string = ""
		var x, y int = 0, 0

		n, err := fmt.Sscanf(line, "PLACE %d,%d,%s", &x, &y, &faceStr)
		if n != 3 || err != nil {
			return nil, errors.New("Could not parse the command correctly")
		}
		face, err := core.FaceFromString(faceStr)
		if err != nil {
			return nil, err
		}

		return cmd.PlaceCmd{
			Position: core.Position{
				X:    x,
				Y:    y,
				Face: face,
			},
		}, nil
	}
	return nil, fmt.Errorf("Can't parse the command: '%s'", line)
}
