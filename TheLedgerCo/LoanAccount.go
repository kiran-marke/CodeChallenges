package main

type LoanAccount struct {
	BankName                     string
	BorrowerName                 string
	OriginalAmountToBePaidAmount int
	PrincipalAmount              int
	NoOfYears                    int
	RateOfInterest               int
	NoOfEMIs                     int
	EMIAmount                    int
	LumsumPayments               []LumsumPayment
}

type LumsumPayment struct {
	LumpSumAmount             int
	LumpSumPaidAfterEMINumber int
}
