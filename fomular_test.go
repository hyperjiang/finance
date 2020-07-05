package finance

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPMT(t *testing.T) {
	should := require.New(t)

	monthlyRate := 4.9 / 100 / 12
	periods := 12 * 30
	amount := 1000000.00

	should.Equal(-5307.267206228046, PMT(monthlyRate, periods, amount, 0, false))
	should.Equal(-5285.683996575363, PMT(monthlyRate, periods, amount, 0, true))
	should.Equal(-2777.777777777778, PMT(0, periods, amount, 0, false))
}
