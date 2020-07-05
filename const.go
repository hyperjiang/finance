package finance

/**
 * Constants for the day-counting
 *
 * base 0: BASIS_MSRB_30_360
 *         MSRB/NSAD 30/360 daycount method
 *         see MSRB Rule G-33
 *         http://www.msrb.org/msrb1/rules/ruleg33.htm
 *         http://www.msrb.org/msrb1/rules/interpg33.htm
 *         Number of days = (Y2 - Y1) 360 + (M2 - M1) 30 + (D2 - D1)
 *         The variables "Yl," "M1," and "D1" are defined as the year, month, and day, respectively,
 *         of the date on which the computation period begins (June 15, 1982, in your example),
 *         and "Y2,a" "M2," and "D2" as the year, month, and day of the date on which the
 *         computation period ends (July 1, 1982, in your example).
 *         For purposes of this formula, if the symbol "D2" has a value of "31," and the symbol "D1"
 *         has a value of "30" or "31," the value of the symbol "D2" shall be changed to "30."
 *         If the symbol "D1" has a value of "31," the value of the symbol "D1" shall be changed to
 *         "30." For purposes of this rule time periods shall be computed to include the day
 *         specified in the rule for the beginning of the period but not to include the day
 *         specified for the end of the period.
 *
 * base 1: BASIS_ACTACT
 *         Actual/Actual daycount method
 *         date adjustment: no change
 *         date difference: serial delta (# of days)
 *
 * base 2: BASIS_ACT_360
 *         Actual/360 daycount method
 *         date adjustment: no change
 *         date difference: serial delta
 *         360/freq for length of coupon period
 *
 * base 3: BASIS_ACT_365
 *         Actual/365 daycount method (short term and Canadian bonds only)
 *         date adjustment: no change
 *         date difference: serial delta
 *         365/freq for length of coupon period, (with decimal answer)
 *
 * base 4: BASIS_30E_360
 *         date adjustment: from_date is changed from 31st to 30th
 *                          to_date is changed from 31st to 30th
 *         date difference: each month 30 days, within a month serial delta
 *
 * base 5: BASIS_30Ep_360
 *         date adjustment: from_date is changed from 31st to 30th
 *                          to_date is changed from 31st to 1st of following month
 *         date difference: each month 30 days, within a month serial delta
 */
const (
	FinancialBasisMsrb = iota
	FinancialBasisActAct
	FinancialBasisAct360
	FinancialBasisAct365
	FinancialBasis30E
	FinancialBasis30Ep
)
