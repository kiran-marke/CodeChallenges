package main

type LoanAccount struct {
	UniqueName                   string //This is formed by BankName_BorrowerName
	BankName                     string
	BorrowerName                 string
	OriginalAmountToBePaidAmount float64
	PrincipalAmount              float64
	NoOfYears                    int
	RateOfInterest               float64
	NoOfEMIs                     int
	EMIAmount                    float64
}
