package aperture

type Circle struct {
	Diameter     float64
	HoleDiameter float64
}

func (a *Circle) SVG() string {
	return ""
}

func (a *Circle) GCode() string {
	return ""
}
