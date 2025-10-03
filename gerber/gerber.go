package gerber

type Unit int

const (
	UnitMM Unit = iota
	UnitIN
)

type Block interface {
	Apply(*Gerber)
}

type Gerber struct {
	blocks       []Block
	unit         Unit
	fmtIntDigits int
	fmtDecDigits int
}

func (g *Gerber) AppendBlock(b Block) {
	g.blocks = append(g.blocks, b)
}

func (g *Gerber) SetUnit(unit Unit) {
	g.unit = unit
}

func (g *Gerber) SetFormat(intDigits, decDigits int) {
	g.fmtIntDigits = intDigits
	g.fmtDecDigits = decDigits
}
