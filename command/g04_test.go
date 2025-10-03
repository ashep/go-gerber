package command_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ashep/go-gerber/command"
)

func TestG04_New(t *testing.T) {
	tests := []struct {
		name     string
		line     string
		expected *command.G04
		errMsg   string
	}{
		{
			name:     "valid with text",
			line:     "G04 Hello world",
			expected: &command.G04{Text: "Hello world"},
			errMsg:   "",
		},
		{
			name:     "valid empty text",
			line:     "G04 ",
			expected: &command.G04{Text: ""},
			errMsg:   "",
		},
		{
			name:     "invalid missing space",
			line:     "G04",
			expected: nil,
			errMsg:   "not a G04 command",
		},
		{
			name:     "invalid different command",
			line:     "G05 Something",
			expected: nil,
			errMsg:   "not a G04 command",
		},
		{
			name:     "invalid empty",
			line:     "",
			expected: nil,
			errMsg:   "not a G04 command",
		},
		{
			name:     "invalid lowercase",
			line:     "g04 comment",
			expected: nil,
			errMsg:   "not a G04 command",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, err := command.NewG04(tt.line)
			if tt.errMsg != "" {
				require.Error(t, err)
				assert.EqualError(t, err, tt.errMsg)
				assert.Nil(t, c)
			} else {
				require.NoError(t, err)
				require.NotNil(t, c)
				assert.Equal(t, tt.expected.Text, c.Text)
			}
		})
	}
}
