package main

import "fmt"

type LoanCommand struct {
	BankName        string
	BorrowerName    string
	PrincipalAmount float64
	NoOfYears       int
	RateOfInterest  float64
}

func (l *LoanCommand) CreateLoanAccount() {

	loanAccountUniqueName := fmt.Sprintf("%v_%v", l.BankName, l.BorrowerName)

	interestAmount := (l.PrincipalAmount * float64(l.NoOfYears) * float64(l.NoOfYears) / 100)
	la := LoanAccount{
		BankName:                     l.BankName,
		BorrowerName:                 l.BorrowerName,
		OriginalAmountToBePaidAmount: l.PrincipalAmount + interestAmount,
		PrincipalAmount:              l.PrincipalAmount,
		NoOfYears:                    l.NoOfYears,
		RateOfInterest:               l.RateOfInterest,
		NoOfEMIs:                     l.NoOfYears * 12,
		EMIAmount:                    float64(l.PrincipalAmount+interestAmount) / float64(l.NoOfYears*12),
	}

	//Make account entry into map
	maploanAccounts[loanAccountUniqueName] = la
}
