package command

type Unit int

const (
	UnitIN Unit = iota
	UnitMM
)

// MO (Mode) command sets the units used for coordinates and for parameters indicating
// sizes or coordinates. The units can be either inches or millimeters.
type MO struct {
	Unit Unit
}

func (c *MO) SVG() string {
	return ""
}

func (c *MO) GCode() string {
	return ""
}
