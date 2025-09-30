package aperture

type Polygon struct {
	OuterDiameter float64
	InnerDiameter float64
	Vertices      int
	RotationAngle float64
	HoleDiameter  float64
}

func (a *Polygon) SVG() string {
	return ""
}

func (a *Polygon) GCode() string {
	return ""
}
