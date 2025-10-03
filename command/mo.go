package command

import (
	"errors"
	"strings"

	"github.com/ashep/go-gerber/gerber"
)

// MO (Mode) command sets the units used for coordinates and for parameters indicating
// sizes or coordinates. The units can be either inches or millimeters.
type MO struct {
	Unit string
}

func NewMO(line string) (*MO, error) {
	if !strings.HasPrefix(line, "MO") {
		return nil, errors.New("not an MO command")
	}
	if len(line) != 4 {
		return nil, errors.New("invalid MO command")
	}

	unit := line[2:]
	if unit != "IN" && unit != "MM" {
		return nil, errors.New("invalid MO command")
	}

	return &MO{Unit: unit}, nil
}

func (c *MO) Apply(g *gerber.Gerber) {
	switch c.Unit {
	case "IN":
		g.SetUnit(gerber.UnitIN)
	default:
		g.SetUnit(gerber.UnitMM)
	}
}
