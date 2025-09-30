package command

import (
	"fmt"
	"strconv"
	"strings"
)

// FS (Format Specification) specifies the format of the coordinate data. The FS command is mandatory.
// It must be used once and only once, in the header, before the first use of coordinate data.
// It is recommended to put it as the very first non-comment line.
// The format of X and Y coordinate must be the same.
type FS struct {
	X int
	Y int
}

func NewFS(line string) (*FS, error) {
	if len(line) != 10 {
		return nil, fmt.Errorf("not an FS command: %s", line)
	}

	if !strings.HasPrefix(line, "FSLA") {
		return nil, fmt.Errorf("not an FS command: %s", line)
	}

	if line[4] != 'X' {
		return nil, fmt.Errorf("not an FS command: %s", line)
	}

	if line[7] != 'Y' {
		return nil, fmt.Errorf("not an FS command: %s", line)
	}

	x, err := strconv.Atoi(line[5:7])
	if err != nil {
		return nil, fmt.Errorf("invalid X value in FS command: %s", line)
	}

	y, err := strconv.Atoi(line[8:10])
	if err != nil {
		return nil, fmt.Errorf("invalid Y value in FS command: %s", line)
	}

	if x != y {
		return nil, fmt.Errorf("X and Y must equal: %s", line)
	}

	return &FS{
		X: x,
		Y: y,
	}, nil
}

func (c *FS) SVG() string {
	return ""
}

func (c *FS) GCode() string {
	return ""
}
