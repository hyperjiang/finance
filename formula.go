package finance

import "math"

// PMT calculates the payment for a loan based on constant payments
// and a constant interest rate.
//
// rate is the interest rate for the loan. (monthly rate)
// nper is the total number of payments for the loan.
// pv is the present value, or the total amount that a series of future payments
//  is worth now; also known as the principal.
// fv is the future value, or a cash balance you want to attain after the last
//  payment is made. If fv is omitted, it is assumed to be 0 (zero), that is,
//  the future value of a loan is 0.
// flag adjust the payment to the beginning or end of the period.
// 	0, At the end of the period
//  1, At the beginning of the period
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
//        (1 + rate * flag) * | ------------------ |
//                             \        rate      /
//
func PMT(rate float64, nper int, pv float64, fv float64, flag float64) float64 {
	pvif := pvif(rate, nper)
	fvifa := fvifa(rate, nper)

	return -(pv*pvif + fv) / ((1.0 + rate*flag) * fvifa)
}

// IPMT returns the interest payment for a given period for an investment based
// on periodic, constant payments and a constant interest rate.
func IPMT(rate float64, per int, nper int, pv float64, fv float64, flag float64) float64 {
	if (per < 1) || (per >= (nper + 1)) {
		return 0
	}

	pmt := PMT(rate, nper, pv, fv, flag)
	return calculateInterest(pv, pmt, rate, per-1)
}

// PPMT returns the payment on the principal for a given period for an investment
// based on periodic, constant payments and a constant interest rate.
func PPMT(rate float64, per int, nper int, pv float64, fv float64, flag float64) float64 {
	if (per < 1) || (per >= (nper + 1)) {
		return 0
	}

	pmt := PMT(rate, nper, pv, fv, flag)
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
//       PMT * (1 + rate * flag) * | ----------------- | - FV
//                                 \        rate       /
// PV = ------------------------------------------------------
//                                nper
//                       (1 + rate)
//
// rate is the interest rate per period. (monthly rate)
// nper is the total number of payment periods in an annuity.
// pmt  is the payment made each period and cannot change over the life of the annuity.
// fv   is the future value, or a cash balance you want to attain after the last payment is made.
// flag adjust the payment to the beginning or end of the period.
//  0, At the end of the period
//  1, At the beginning of the period
func PV(rate float64, nper int, pmt float64, fv float64, flag float64) float64 {
	pvif := pvif(rate, nper)
	fvifa := fvifa(rate, nper)

	return (-pmt*(1.0+rate*flag)*fvifa - fv) / pvif
}

// FV returns the future value of an investment based on periodic,
// constant payments and a constant interest rate.
//
// For a more complete description of the arguments in FV, see the PV function.
//
// rate is the interest rate per period. (monthly rate)
// nper is the total number of payment periods in an annuity.
// pmt  is the payment made each period and cannot change over the life of the annuity.
// pv   is the present value, or the lump-sum amount that a series of future payments is worth right now.
// flag adjust the payment to the beginning or end of the period.
//  0, At the end of the period
//  1, At the beginning of the period
//
// If rate = 0:
// FV = -(PV + PMT * nper)
//
// Else
//                                 /             nper \
//                                 | 1 - (1 + rate)     |                 nper
// FV =  PMT * (1 + rate * flag) * | ------------------ | - PV * (1 + rate)
//                                 \        rate      /
//
func FV(rate float64, nper int, pmt float64, pv float64, flag float64) float64 {
	pvif := pvif(rate, nper)
	fvifa := fvifa(rate, nper)

	return -pmt*(1.0+rate*flag)*fvifa - pv*pvif
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

func calculateInterest(pv float64, pmt float64, rate float64, per int) float64 {
	return -(pv*math.Pow(1+rate, float64(per))*rate +
		pmt*(math.Pow(1+rate, float64(per))-1))
}
