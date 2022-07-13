package main

import "fmt"

type PaymentCommand struct {
	BankName                  string
	BorrowerName              string
	LumpSumAmount             float64
	LumpSumPaidAfterEMINumber int
}

func (p *PaymentCommand) UpdateLoanAccount() {

	loanAccountUniqueName := fmt.Sprintf("%v_%v", p.BankName, p.BorrowerName)

	//Fetch Ban account entry from map
	if la, ok := maploanAccounts[loanAccountUniqueName]; ok {
		la.PrincipalAmount = la.PrincipalAmount - p.LumpSumAmount
		la.NoOfEMIs = la.NoOfEMIs - p.LumpSumPaidAfterEMINumber
	}
}
