package finance

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPMT(t *testing.T) {
	should := require.New(t)

	monthlyRate := 4.9 / 100 / 12
	periods := 12 * 30
	amount := 1000000.00

	should.Equal(-5307.267206228046, PMT(monthlyRate, periods, amount, 0, 0))
	should.Equal(-5285.683996575363, PMT(monthlyRate, periods, amount, 0, 1))
	should.Equal(-2777.777777777778, PMT(0, periods, amount, 0, 0))
}

func TestIPMT(t *testing.T) {
	should := require.New(t)

	monthlyRate := 4.9 / 100 / 12
	periods := 12 * 30
	amount := 1000000.00

	should.Equal(0.0, IPMT(monthlyRate, 0, periods, amount, 0, 0))
	should.Equal(-4083.333333333333, IPMT(monthlyRate, 1, periods, amount, 0, 0))
}

func TestPPMT(t *testing.T) {
	should := require.New(t)

	monthlyRate := 4.9 / 100 / 12
	periods := 12 * 30
	amount := 1000000.00

	should.Equal(0.0, PPMT(monthlyRate, 0, periods, amount, 0, 0))
	should.Equal(-1223.9338728947132, PPMT(monthlyRate, 1, periods, amount, 0, 0))
}

func TestPV(t *testing.T) {
	should := require.New(t)

	monthlyRate := 4.9 / 100 / 12
	periods := 12 * 30
	pmt := 1000.00

	should.Equal(-188420.8880281185, PV(monthlyRate, periods, pmt, 0, 0))
}

func TestFV(t *testing.T) {
	should := require.New(t)

	monthlyRate := 4.9 / 100 / 12
	periods := 12 * 30
	pmt := 1000.00

	should.Equal(-817037.6048461755, FV(monthlyRate, periods, pmt, 0, 0))
}

func TestNPER(t *testing.T) {
	should := require.New(t)

	monthlyRate := 4.9 / 100 / 12
	pmt := 5307.00
	pv := 1000000.00

	should.Equal(0.0, NPER(0, 0, pv, 0, 0))
	should.Equal(-188.43037497644622, NPER(0, pmt, pv, 0, 0))
	should.Equal(-140.03715468730746, NPER(monthlyRate, pmt, pv, 0, 0))
	should.Equal(0.0, NPER(monthlyRate, pmt, pv, 2000000, 0))
}

func TestRATE(t *testing.T) {
	should := require.New(t)

	should.Equal(-1.4155398849824252, RATE(12, 3612.82, 41817.82, 0.0, 0.0, 0.1))
	should.True(math.IsNaN(RATE(12, 3612.82, 41817.82, 0.0, 0.0, 0.0)))
}
