package gerber

// BlockCommand is identified by a command code. For block commands,
// this is a letter G, D or M followed by a code number, e.g. G02.
// A block command with command code D01, D02 and D03 is Operation.
//
// Examples:
// G01*
// G74*
// D12*
// M02*
type BlockCommand struct {
	Code CommandCode
}
