package main

import "fmt"

type PaymentCommand struct {
	BankName                  string
	BorrowerName              string
	LumpSumAmount             int
	LumpSumPaidAfterEMINumber int
}

func (p *PaymentCommand) UpdateLoanAccount() {

	loanAccountUniqueName := fmt.Sprintf("%v_%v", p.BankName, p.BorrowerName)

	//Fetch Ban account entry from map
	if la, ok := maploanAccounts[loanAccountUniqueName]; ok {

		lumsum := LumsumPayment{
			LumpSumAmount:             p.LumpSumAmount,
			LumpSumPaidAfterEMINumber: p.LumpSumPaidAfterEMINumber,
		}
		la.LumsumPayments = append(la.LumsumPayments, lumsum)
	}

}
