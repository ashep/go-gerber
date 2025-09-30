package gerber

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/ashep/go-gerber/command"
)

func FromFile(filename string) (*Gerber, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	b, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}

	return NewParser().Parse(b)
}

func FromBytes(b []byte) (*Gerber, error) {
	return NewParser().Parse(b)
}

type Parser struct {
}

func NewParser() *Parser {
	return &Parser{}
}

func (p *Parser) Parse(data []byte) (*Gerber, error) {
	linesLF := strings.ReplaceAll(string(data), "\r\n", "\n")
	lines := strings.Split(linesLF, "\n")
	blocks := make([]Block, 0)

	for cnt, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		switch {
		case strings.HasSuffix(line, "*"):
			line = strings.TrimSuffix(line, "*")
			blk, err := p.parseCommand(line)
			if err != nil {
				return nil, fmt.Errorf("line %d: %w", cnt+1, err)
			}
			blocks = append(blocks, blk)
		case strings.HasPrefix(line, "%") && strings.HasSuffix(line, "*%"):
			line = strings.TrimPrefix(line, "%")
			line = strings.TrimSuffix(line, "*%")
			blk, err := p.parseExtCommand(line)
			if err != nil {
				return nil, fmt.Errorf("line %d: %w", cnt+1, err)
			}
			blocks = append(blocks, blk)
		default:
			return nil, fmt.Errorf("line %d: invalid format", cnt+1)
		}
	}

	return nil, nil
}

func (p *Parser) parseCommand(line string) (Block, error) {
	switch {
	case strings.HasPrefix(line, "G04"):
		return command.NewG04(line)
	}

	return nil, errors.New("unsupported command: " + line)
}

func (p *Parser) parseExtCommand(line string) (Block, error) {
	line = strings.TrimSuffix(line, "*")

	switch {
	case strings.HasPrefix(line, "FS"):
		return command.NewFS(line)
	}

	return nil, errors.New("unsupported command: " + line)
}
