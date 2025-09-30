package aperture

type Obround struct {
	X            float64
	Y            float64
	HoleDiameter float64
}

func (a *Obround) SVG() string {
	return ""
}

func (a *Obround) GCode() string {
	return ""
}
