package command_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ashep/go-gerber/command"
)

func TestFS_New(t *testing.T) {
	tests := []struct {
		name     string
		line     string
		errMsg   string
	}{
		{
			name:     "valid FS command",
			line:     "FSLAX25Y25",
			errMsg:   "",
		},
		{
			name:     "valid FS command with different values",
			line:     "FSLAX34Y34",
			errMsg:   "",
		},
		{
			name:     "X and Y values not equal",
			line:     "FSLAX34Y12",
			errMsg:   "X and Y must equal: FSLAX34Y12",
		},
		{
			name:     "incorrect length",
			line:     "FSLAX2Y25",
			errMsg:   "not an FS command: FSLAX2Y25",
		},
		{
			name:     "missing FSLA prefix",
			line:     "X25Y25X25Y",
			errMsg:   "not an FS command: X25Y25X25Y",
		},
		{
			name:     "missing X after FSLA",
			line:     "FSLAY25Y25",
			errMsg:   "not an FS command: FSLAY25Y25",
		},
		{
			name:     "missing Y after X value",
			line:     "FSLAX25Z25",
			errMsg:   "not an FS command: FSLAX25Z25",
		},
		{
			name:     "non-numeric X value",
			line:     "FSLAXX2Y25",
			errMsg:   "invalid X value in FS command: FSLAXX2Y25",
		},
		{
			name:     "non-numeric Y value",
			line:     "FSLAX25YY2",
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
			}
		})
	}
}
