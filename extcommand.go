package gerber

// ExtCommand consists of a two-character command code followed by command data.
// The command code identifies the command. The command data depend on the command code.
// An extended command is enclosed into a separate pair of ‘%’ delimiter characters. A command
// consists of a single data block, except the AM command, which normally has more than one
// data block . The ‘%’ must immediately follow the ‘*’ of the last data block without intervening line
// separators. It is allowed to put a line feed between data blocks of a single command.
//
// Examples:
// %FSLAX24Y24*%
//
// %AMDONUTFIX*1,1,0.100,0,0*1,0,0.080,0,0*%
//
// %AMDONUTFIX*
// 1,1,0.100,0,0*
// 1,0,0.080,0,0*%
type ExtCommand struct {
	Code CommandCode
}
