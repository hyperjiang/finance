package finance

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEqualInstallment(t *testing.T) {
	should := require.New(t)

	loan := Loan{
		AnnualRate: 0.07,
		Periods:    12,
		Amount:     1000000,
	}

	installments := loan.CalculateInstallments()

	for i, installment := range installments {
		should.Equal(86526.75, installment.Payment)
		should.Equal(i+1, installment.Period)

		if i == 0 {
			should.Equal(80693.41, installment.Principal)
			should.Equal(5833.33, installment.Interest)
			should.Equal(919306.59, installment.RemainingAmount)
		}

		if i == 12 {
			should.Equal(86024.93, installment.Principal)
			should.Equal(501.81, installment.Interest)
			should.Equal(0, installment.RemainingAmount)
		}
	}

	should.Equal(38320.95, loan.CalculateTotalInterest())
	should.Equal(1038320.95, loan.CalculateTotalPayment())
}

func TestEqualPrincipal(t *testing.T) {
	should := require.New(t)

	loan := Loan{
		AnnualRate: 0.07,
		Periods:    12,
		Amount:     1000000,
		Method:     EqualPrincipal,
	}

	installments := loan.CalculateInstallments()

	for i, installment := range installments {
		should.Equal(i+1, installment.Period)
		should.Equal(83333.33, installment.Principal)

		if i == 0 {
			should.Equal(89166.67, installment.Payment)
			should.Equal(5833.33, installment.Interest)
			should.Equal(916666.67, installment.RemainingAmount)
		}

		if i == 12 {
			should.Equal(83819.44, installment.Payment)
			should.Equal(486.11, installment.Interest)
			should.Equal(0, installment.RemainingAmount)
		}
	}

	should.Equal(37916.67, loan.CalculateTotalInterest())
	should.Equal(1037916.67, loan.CalculateTotalPayment())
}
