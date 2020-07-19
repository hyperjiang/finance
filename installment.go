package finance

import "github.com/hyperjiang/php"

// Loan is a loan, default method is EqualPayment
type Loan struct {
	Amount     float64
	Periods    int
	AnnualRate float64
	Method     int
}

// Installment is an installment
type Installment struct {
	Period          int
	Payment         float64
	Principal       float64
	Interest        float64
	RemainingAmount float64
}

// CalculatePayment calculates payment in given period
func (loan Loan) CalculatePayment(period int) float64 {
	return loan.CalculatePrincipal(period) + loan.CalculateInterest(period)
}

// CalculatePrincipal calculates principal in given period
func (loan Loan) CalculatePrincipal(period int) float64 {
	if loan.Method == EqualPrincipal {
		return loan.Amount / float64(loan.Periods)
	}

	monthlyRate := loan.AnnualRate / 12

	return PPMT(monthlyRate, period, loan.Periods, -loan.Amount, 0, 0)
}

// CalculateInterest calculates interest in given period
func (loan Loan) CalculateInterest(period int) float64 {
	monthlyRate := loan.AnnualRate / 12

	if loan.Method == EqualPrincipal {
		remainingAmount := loan.Amount * float64(loan.Periods-period+1) / float64(loan.Periods)
		return remainingAmount * monthlyRate
	}

	return IPMT(monthlyRate, period, loan.Periods, -loan.Amount, 0, 0)
}

// CalculateTotalPayment calculates total payment
func (loan Loan) CalculateTotalPayment() float64 {
	monthlyRate := loan.AnnualRate / 12

	if loan.Method == EqualPrincipal {
		return php.Round(loan.Amount*(1+monthlyRate*float64(1+loan.Periods)/2), Precision)
	}

	return php.Round(PMT(monthlyRate, loan.Periods, -loan.Amount, 0, 0)*float64(loan.Periods), Precision)
}

// CalculateTotalInterest calculates total interest
func (loan Loan) CalculateTotalInterest() float64 {
	return php.Round(loan.CalculateTotalPayment()-loan.Amount, Precision)
}

// CalculateInstallments calculates installments
func (loan Loan) CalculateInstallments() []Installment {
	var installments []Installment
	for p := 1; p < loan.Periods; p++ {
		var installment Installment
		installment.Period = p
		installment.Payment = php.Round(loan.CalculatePayment(p), Precision)
		installment.Principal = php.Round(loan.CalculatePrincipal(p), Precision)
		installment.Interest = php.Round(loan.CalculateInterest(p), Precision)
		installment.RemainingAmount = php.Round(loan.Amount-installment.Principal, Precision)
	}

	return installments
}
