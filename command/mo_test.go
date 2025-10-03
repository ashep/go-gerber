package command_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ashep/go-gerber/command"
)

func TestMO_New(t *testing.T) {
	tests := []struct {
		name     string
		line     string
		expected *command.MO
		errMsg   string
	}{
		{
			name:     "valid MO command with inches",
			line:     "MOIN",
			expected: &command.MO{Unit: "IN"},
			errMsg:   "",
		},
		{
			name:     "valid MO command with millimeters",
			line:     "MOMM",
			expected: &command.MO{Unit: "MM"},
			errMsg:   "",
		},
		{
			name:     "missing MO prefix",
			line:     "XXIN",
			expected: nil,
			errMsg:   "not an MO command",
		},
		{
			name:     "invalid length - too short",
			line:     "MOI",
			expected: nil,
			errMsg:   "invalid MO command",
		},
		{
			name:     "invalid length - too long",
			line:     "MOINN",
			expected: nil,
			errMsg:   "invalid MO command",
		},
		{
			name:     "invalid unit",
			line:     "MOCM",
			expected: nil,
			errMsg:   "invalid MO command",
		},
		{
			name:     "invalid unit - lowercase",
			line:     "MOin",
			expected: nil,
			errMsg:   "invalid MO command",
		},
		{
			name:     "empty string",
			line:     "",
			expected: nil,
			errMsg:   "not an MO command",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mo, err := command.NewMO(tt.line)
			if tt.errMsg != "" {
				require.Error(t, err)
				assert.EqualError(t, err, tt.errMsg)
				assert.Nil(t, mo)
			} else {
				require.NoError(t, err)
				require.NotNil(t, mo)
				assert.Equal(t, tt.expected.Unit, mo.Unit)
			}
		})
	}
}
