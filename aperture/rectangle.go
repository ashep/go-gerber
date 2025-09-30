package aperture

type Rectangle struct {
	X            float64
	Y            float64
	HoleDiameter float64
}

func (a *Rectangle) SVG() string {
	return ""
}

func (a *Rectangle) GCode() string {
	return ""
}
