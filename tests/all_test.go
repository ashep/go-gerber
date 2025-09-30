package tests

import (
	_ "embed"
	"fmt"
	"testing"

	"github.com/ashep/go-gerber"

	"github.com/stretchr/testify/require"
)

var (
	//go:embed testdata/can2usb_COPPER-TOP.gbr
	testDataCopperTop []byte
)

func TestAll(t *testing.T) {
	gbr, err := gerber.FromBytes(testDataCopperTop)
	require.NoError(t, err)

	fmt.Println(gbr)
}
