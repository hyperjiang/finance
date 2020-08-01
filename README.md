# finance

[![GoDoc](https://godoc.org/github.com/hyperjiang/finance?status.svg)](https://pkg.go.dev/github.com/hyperjiang/finance?tab=doc)
[![Build Status](https://travis-ci.org/hyperjiang/finance.svg?branch=master)](https://travis-ci.org/hyperjiang/finance)
[![](https://goreportcard.com/badge/github.com/hyperjiang/finance)](https://goreportcard.com/report/github.com/hyperjiang/finance)
[![codecov](https://codecov.io/gh/hyperjiang/finance/branch/master/graph/badge.svg)](https://codecov.io/gh/hyperjiang/finance)
[![Release](https://img.shields.io/github/release/hyperjiang/finance.svg)](https://github.com/hyperjiang/finance/releases)

Financial functions with the Excel function names and parameter order, along with an installment calculator.

Require go version >= 1.13.

## Available Functions

- `PMT`: calculates the payment for a loan based on constant payments and a constant interest rate.
- `IPMT`: returns the interest payment for a given period for an investment based on periodic, constant payments and a constant interest rate.
- `PPMT`: returns the payment on the principal for a given period for an investment based on periodic, constant payments and a constant interest rate.
- `PV`: returns the present value of an investment.
- `FV`: returns the future value of an investment based on periodic, constant payments and a constant interest rate.
- `NPER`: returns the number of periods for an investment based on periodic, constant payments and a constant interest rate.
- `RATE`: calculates interest rate per period of an annuity.

## Installment Calculator

Demo usage:

```
import (
    "fmt"
    "github.com/hyperjiang/finance"
)

loan := finance.Loan{
    AnnualRate: 0.07,
    Periods:    12,
    Amount:     1000000,
    Method:     finance.EqualPayment,
}

installments := loan.CalculateInstallments()

fmt.Println(installments)
```
