package gerber

type Block interface {
	SVG() string
	GCode() string
}

type Gerber struct {
	Blocks []Block
}

func (g *Gerber) SVG() string {
	return ""
}

func (g *Gerber) GCode() string {
	return ""
}
