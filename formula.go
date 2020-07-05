package finance

import "math"

// PMT calculates the payment for a loan based on constant payments
// and a constant interest rate.
//
// rate: is the interest rate for the loan. (monthly rate)
// nper: is the total number of payments for the loan.
// pv: is the present value, or the total amount that a series of future payments
//  is worth now; also known as the principal.
// fv: is the future value, or a cash balance you want to attain after the last
//  payment is made. If fv is omitted, it is assumed to be 0 (zero), that is,
//  the future value of a loan is 0.
// beginning: adjust the payment to the beginning or end of the period.
//  false, At the end of the period
//  true, At the beginning of the period
//
// If rate = 0:
//        -(FV + PV)
// PMT = ------------
//           nper
//
// Else
//
//                                      nper
//                   FV + PV * (1 + rate)
// PMT = --------------------------------------------
//                             /             nper \
//                            | 1 - (1 + rate)     |
//        (1 + rate * type) * | ------------------ |
//                             \        rate      /
//
func PMT(rate float64, nper int, pv float64, fv float64, beginning bool) float64 {
	pvif := pvif(rate, nper)
	fvifa := fvifa(rate, nper)

	var f float64
	if beginning {
		f = 1
	}

	return (-pv*pvif - fv) / ((1.0 + rate*f) * fvifa)
}

// Present value interest factor
//
//                 nper
// PVIF = (1 + rate)
//
// rate is the interest rate per period.
// nper is the total number of periods.
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
// rate is the interest rate per period.
// nper is the total number of periods.
func fvifa(rate float64, nper int) float64 {
	// Removable singularity at rate == 0
	if rate == 0 {
		return float64(nper)
	}

	return (math.Pow(1+rate, float64(nper)) - 1) / rate
}
