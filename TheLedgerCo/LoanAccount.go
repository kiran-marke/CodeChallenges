package main

import "fmt"

type LoanAccount struct {
	BankName        string
	BorrowerName    string
	PrincipalAmount int
	NoOfYears       int
	RateOfInterest  int
	NoOfEMIs        int
	EMIAmount       int
	LumsumPayments  []LumsumPayment
}

type LumsumPayment struct {
	LumpSumAmount             int
	LumpSumPaidAfterEMINumber int
}

func getLoanAccountUniqueName(bankName, borrowerName string) string {
	return fmt.Sprintf("%v_%v", bankName, borrowerName)
}
