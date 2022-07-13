package main

import (
	"fmt"
	"math"
)

type LoanCommand struct {
	BankName        string
	BorrowerName    string
	PrincipalAmount int
	NoOfYears       int
	RateOfInterest  int
}

func (l *LoanCommand) CreateLoanAccount() {

	loanAccountUniqueName := fmt.Sprintf("%v_%v", l.BankName, l.BorrowerName)

	interestAmount := (l.PrincipalAmount * l.NoOfYears * l.RateOfInterest) / 100
	la := LoanAccount{
		BankName:                     l.BankName,
		BorrowerName:                 l.BorrowerName,
		OriginalAmountToBePaidAmount: l.PrincipalAmount,
		PrincipalAmount:              l.PrincipalAmount,
		NoOfYears:                    l.NoOfYears,
		RateOfInterest:               l.RateOfInterest,
		NoOfEMIs:                     l.NoOfYears * 12,
		EMIAmount:                    int(math.Ceil(float64(l.PrincipalAmount+interestAmount) / float64(l.NoOfYears*12))),
		LumsumPayments:               make([]LumsumPayment, 0),
	}

	//Make account entry into map
	maploanAccounts[loanAccountUniqueName] = &la
}
