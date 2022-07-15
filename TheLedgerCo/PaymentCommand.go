package main

type PaymentCommand struct {
	BankName                  string
	BorrowerName              string
	LumpSumAmount             int
	LumpSumPaidAfterEMINumber int
}

func (p *PaymentCommand) UpdateLoanAccount() {

	loanAccountUniqueName := getLoanAccountUniqueName(p.BankName, p.BorrowerName)

	//Fetch Loan account entry from map
	if la, ok := maploanAccounts[loanAccountUniqueName]; ok {

		lumsum := LumsumPayment{
			LumpSumAmount:             p.LumpSumAmount,
			LumpSumPaidAfterEMINumber: p.LumpSumPaidAfterEMINumber,
		}
		la.LumsumPayments = append(la.LumsumPayments, lumsum)
	}

}
