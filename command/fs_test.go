package command_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ashep/go-gerber/command"
)

func TestFS(t *testing.T) {
	tests := []struct {
		name     string
		line     string
		expected *command.FS
		errMsg   string
	}{
		{
			name:     "valid FS command",
			line:     "FSLAX25Y25",
			expected: &command.FS{X: 25, Y: 25},
			errMsg:   "",
		},
		{
			name:     "valid FS command with different values",
			line:     "FSLAX34Y34",
			expected: &command.FS{X: 34, Y: 34},
			errMsg:   "",
		},
		{
			name:     "X and Y values not equal",
			line:     "FSLAX34Y12",
			expected: nil,
			errMsg:   "X and Y must equal: FSLAX34Y12",
		},
		{
			name:     "incorrect length",
			line:     "FSLAX2Y25",
			expected: nil,
			errMsg:   "not an FS command: FSLAX2Y25",
		},
		{
			name:     "missing FSLA prefix",
			line:     "X25Y25X25Y",
			expected: nil,
			errMsg:   "not an FS command: X25Y25X25Y",
		},
		{
			name:     "missing X after FSLA",
			line:     "FSLAY25Y25",
			expected: nil,
			errMsg:   "not an FS command: FSLAY25Y25",
		},
		{
			name:     "missing Y after X value",
			line:     "FSLAX25Z25",
			expected: nil,
			errMsg:   "not an FS command: FSLAX25Z25",
		},
		{
			name:     "non-numeric X value",
			line:     "FSLAXX2Y25",
			expected: nil,
			errMsg:   "invalid X value in FS command: FSLAXX2Y25",
		},
		{
			name:     "non-numeric Y value",
			line:     "FSLAX25YY2",
			expected: nil,
			errMsg:   "invalid Y value in FS command: FSLAX25YY2",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs, err := command.NewFS(tt.line)
			if tt.errMsg != "" {
				require.Error(t, err)
				assert.EqualError(t, err, tt.errMsg)
				assert.Nil(t, fs)
			} else {
				require.NoError(t, err)
				require.NotNil(t, fs)
				assert.Equal(t, tt.expected.X, fs.X)
				assert.Equal(t, tt.expected.Y, fs.Y)
			}
		})
	}
}

func TestFS_SVG(t *testing.T) {
	fs := &command.FS{X: 25, Y: 25}
	assert.Equal(t, "", fs.SVG(), "SVG() should return an empty string")
}

func TestFS_GCode(t *testing.T) {
	fs := &command.FS{X: 25, Y: 25}
	assert.Equal(t, "", fs.GCode(), "GCode() should return an empty string")
}
