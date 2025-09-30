package gerber

// Operation consist of a block containing coordinate data before the command code.
// The operation code defines how to use the preceding coordinate.
// Code can be only CommandCodeD01, CommandCodeD02 or CommandCodeD03.
//
// Examples:
// X50001000Y12500000D02*
// X11002000Y12502000I50001000J0D01*
type Operation struct {
	X int
	Y int
	I int
	J int
}
