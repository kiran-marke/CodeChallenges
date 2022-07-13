package main

import (
	"fmt"
	"math"
)

type BalanceCommand struct {
	BankName     string
	BorrowerName string
	EMIs         int
}

type BalanceResponse struct {
	BankName       string
	BorrowerName   string
	AmountPaid     int
	Remaining_EMIs int
}

func (b *BalanceCommand) GetLoanAccountDetails() BalanceResponse {

	loanAccountUniqueName := fmt.Sprintf("%v_%v", b.BankName, b.BorrowerName)

	res := BalanceResponse{}
	//Fetch Ban account entry from map
	if la, ok := maploanAccounts[loanAccountUniqueName]; ok {
		res.BankName = la.BankName
		res.BorrowerName = la.BorrowerName
		var totalAmountPaid int

		for _, v := range la.LumsumPayments {
			if b.EMIs >= v.LumpSumPaidAfterEMINumber {
				totalAmountPaid = totalAmountPaid + v.LumpSumAmount
			}
		}
		res.AmountPaid = totalAmountPaid + b.EMIs*la.EMIAmount
		res.Remaining_EMIs = la.NoOfEMIs - b.EMIs - int(math.Floor(float64(totalAmountPaid)/(float64(la.EMIAmount))))
	}

	return res
}
