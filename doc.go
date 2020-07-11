// Package finance implements financial functions with the Excel function names and parameter order.
//
// Terms Definition
//
// rate: the interest rate for the loan. (monthly rate)
// nper: the total number of payments for the loan.
// pv: the present value, or the total amount that a series of future payments is worth now; also known as the principal.
// fv: the future value, or a cash balance you want to attain after the last payment is made.
// pmt: the payment made each period and cannot change over the life of the annuity.
// due: adjust the payment to the beginning or end of the period. 0, At the end of the period; 1, At the beginning of the period.
package finance
