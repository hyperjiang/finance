package finance

import "github.com/hyperjiang/php"

// Loan is a loan
type Loan struct {
	Amount     float64
	Periods    int
	AnnualRate float64
}

// Installment is an installment
type Installment struct {
	Period          int
	Payment         float64
	Principal       float64
	Interest        float64
	RemainingAmount float64
}

// CalculatePayment calculates payment in each period
func (loan Loan) CalculatePayment() float64 {
	monthlyRate := loan.AnnualRate / 12

	return PMT(monthlyRate, loan.Periods, -loan.Amount, 0, 0)
}

// CalculatePrincipal calculates principal in given period
func (loan Loan) CalculatePrincipal(period int) float64 {
	monthlyRate := loan.AnnualRate / 12

	return PPMT(monthlyRate, period, loan.Periods, -loan.Amount, 0, 0)
}

// CalculateInterest calculates interest in given period
func (loan Loan) CalculateInterest(period int) float64 {
	monthlyRate := loan.AnnualRate / 12

	return IPMT(monthlyRate, period, loan.Periods, -loan.Amount, 0, 0)
}

// CalculateInstallments calculates installments
func (loan Loan) CalculateInstallments() []Installment {
	payment := php.Round(loan.CalculatePayment(), 2)

	var installments []Installment
	for p := 1; p < loan.Periods; p++ {
		var installment Installment
		installment.Period = p
		installment.Payment = payment
		installment.Principal = php.Round(loan.CalculatePrincipal(p), 2)
		installment.Interest = php.Round(loan.CalculateInterest(p), 2)
		installment.RemainingAmount = php.Round(loan.Amount-installment.Payment, 2)
	}

	return installments
}
