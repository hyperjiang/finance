package finance

import "math"

// PMT calculates the payment for a loan based on constant payments
// and a constant interest rate.
//
// If rate = 0:
//        -(FV + PV)
// PMT = ------------
//           nper
//
// Else
//                                      nper
//                   FV + PV * (1 + rate)
// PMT = --------------------------------------------
//                             /             nper \
//                            | 1 - (1 + rate)     |
//         (1 + rate * due) * | ------------------ |
//                             \        rate      /
//
func PMT(rate float64, nper int, pv float64, fv float64, due float64) float64 {
	pvif := pvif(rate, nper)
	fvifa := fvifa(rate, nper)

	return -(pv*pvif + fv) / ((1.0 + rate*due) * fvifa)
}

// IPMT returns the interest payment for a given period for an investment based
// on periodic, constant payments and a constant interest rate.
func IPMT(rate float64, per int, nper int, pv float64, fv float64, due float64) float64 {
	if (per < 1) || (per >= (nper + 1)) {
		return 0
	}

	pmt := PMT(rate, nper, pv, fv, due)
	return calculateInterest(pv, pmt, rate, per-1)
}

// PPMT returns the payment on the principal for a given period for an investment
// based on periodic, constant payments and a constant interest rate.
func PPMT(rate float64, per int, nper int, pv float64, fv float64, due float64) float64 {
	if (per < 1) || (per >= (nper + 1)) {
		return 0
	}

	pmt := PMT(rate, nper, pv, fv, due)
	ipmt := calculateInterest(pv, pmt, rate, per-1)
	return pmt - ipmt
}

// PV returns the present value of an investment. The present value is
// the total amount that a series of future payments is worth now.
// For example, when you borrow money, the loan amount is the present
// value to the lender.
//
// If rate = 0:
// PV = -(FV + PMT * nper)
//
// Else
//                                 /              nper \
//                                 | 1 - (1 + rate)    |
//        PMT * (1 + rate * due) * | ----------------- | - FV
//                                 \        rate       /
// PV = ------------------------------------------------------
//                                nper
//                       (1 + rate)
//
func PV(rate float64, nper int, pmt float64, fv float64, due float64) float64 {
	pvif := pvif(rate, nper)
	fvifa := fvifa(rate, nper)

	return (-pmt*(1.0+rate*due)*fvifa - fv) / pvif
}

// FV returns the future value of an investment based on periodic,
// constant payments and a constant interest rate.
//
// For a more complete description of the arguments in FV, see the PV function.
//
// If rate = 0:
// FV = -(PV + PMT * nper)
//
// Else
//                               /             nper \
//                               | 1 - (1 + rate)     |                 nper
// FV = PMT * (1 + rate * due) * | ------------------ | - PV * (1 + rate)
//                               \        rate      /
//
func FV(rate float64, nper int, pmt float64, pv float64, due float64) float64 {
	pvif := pvif(rate, nper)
	fvifa := fvifa(rate, nper)

	return -pmt*(1.0+rate*due)*fvifa - pv*pvif
}

// NPER returns the number of periods for an investment based on periodic, constant payments and a constant interest rate.
//
// If rate = 0:
//        -(FV + PV)
// nper = -----------
//           PMT
//
// Else
//              / PMT * (1 + rate * due) - FV * rate \
//         log | ------------------------------------ |
//              \ PMT * (1 + rate * due) + PV * rate /
// nper = -----------------------------------------------
//                          log (1 + rate)
//
func NPER(rate float64, pmt float64, pv float64, fv float64, due float64) float64 {
	if rate == 0 && pmt != 0 {
		return -fv - pv/pmt
	}

	if rate <= 0.0 {
		return 0
	}

	initial := pmt * (1.0 + rate*due)

	tmp := (initial - fv*rate) / (pv*rate + initial)
	if tmp <= 0.0 {
		return 0
	}

	return math.Log(tmp) / math.Log(1.0+rate)
}

// RATE calculates interest rate per period of an annuity.
//
// The basic rate formula derivation is to solve for the future value
// taking into account the present value:
// https://en.wikipedia.org/wiki/Future_value
func RATE(nper int, pmt float64, pv float64, fv float64, due float64, guess float64) float64 {
	rate := guess
	i := 0
	var x0, x1, y, f float64

	x1 = rate

	lamda := func(rate float64) float64 {
		if math.Abs(rate) < Accuracy {
			return pv*(1+float64(nper)*rate) + pmt*(1+rate*due)*float64(nper) + fv
		}

		f = math.Exp(float64(nper) * math.Log(1+rate))
		return pv*f + pmt*(1/rate+due)*(f-1) + fv
	}

	y = lamda(rate)

	y0 := pv + pmt*float64(nper) + fv
	y1 := pv*f + pmt*(1/rate+due)*(f-1) + fv

	// find root by secant method
	for (math.Abs(y0-y1) > Accuracy) && (i < MaxIterations) {
		rate = (y1*x0 - y0*x1) / (y1 - y0)
		x0 = x1
		x1 = rate

		y = lamda(rate)

		y0 = y1
		y1 = y
		i++
	}

	return rate
}

// Present value interest factor
//
//                 nper
// PVIF = (1 + rate)
//
func pvif(rate float64, nper int) float64 {
	return math.Pow(1+rate, float64(nper))
}

// Future value interest factor of annuities
//
//                   nper
//          (1 + rate)    - 1
// FVIFA = -------------------
//               rate
//
func fvifa(rate float64, nper int) float64 {
	// Removable singularity at rate == 0
	if rate == 0 {
		return float64(nper)
	}

	return (math.Pow(1+rate, float64(nper)) - 1) / rate
}

func calculateInterest(pv float64, pmt float64, rate float64, per int) float64 {
	return -(pv*math.Pow(1+rate, float64(per))*rate +
		pmt*(math.Pow(1+rate, float64(per))-1))
}
