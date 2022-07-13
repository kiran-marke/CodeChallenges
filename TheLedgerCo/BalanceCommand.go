package main

import "fmt"

type BalanceCommand struct {
	BankName     string
	BorrowerName string
	EMIs         int
}

type BalanceResponse struct {
	BankName       string
	BorrowerName   string
	AmountPaid     float64
	Remaining_EMIs int
}

func (b *BalanceCommand) GetLoanAccountDetails() BalanceResponse {

	loanAccountUniqueName := fmt.Sprintf("%v_%v", b.BankName, b.BorrowerName)

	res := BalanceResponse{}
	//Fetch Ban account entry from map
	if la, ok := maploanAccounts[loanAccountUniqueName]; ok {
		res.BankName = la.BankName
		res.BorrowerName = la.BorrowerName
		res.AmountPaid = float64(b.EMIs * int(la.EMIAmount))
		res.Remaining_EMIs = la.NoOfEMIs - b.EMIs
	}

	return res
}
