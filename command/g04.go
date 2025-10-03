package command

import (
	"errors"
	"strings"

	"github.com/ashep/go-gerber/gerber"
)

// G04 command is used for human-readable comments. It does not affect the image.
// Gerber readers must ignore the command when generating the image.
type G04 struct {
	Text string
}

func NewG04(line string) (*G04, error) {
	if !strings.HasPrefix(line, "G04 ") {
		return nil, errors.New("not a G04 command")
	}
	return &G04{Text: line[4:]}, nil
}

func (*G04) Apply(_ *gerber.Gerber) {
}
