package main

import "testing"

func Init() {
	maploanAccounts = make(map[string]*LoanAccount)
}
func Test_should_create_new_loanaccount_entry_when_received_loancommand(t *testing.T) {
	//Arrange
	Init()
	l := LoanCommand{
		BankName:        "IDIDI",
		BorrowerName:    "Dale",
		PrincipalAmount: 5000,
		NoOfYears:       1,
		RateOfInterest:  6,
	}

	//Act
	l.CreateLoanAccount()

	//Assert
	if len(maploanAccounts) == 0 {
		t.Error("Loan account entry is not made for LOAN Command")
	}

	if _, ok := maploanAccounts["IDIDI_Dale"]; !ok {
		t.Error("Loan account map entry not found")
	}

	if maploanAccounts["IDIDI_Dale"].NoOfEMIs != 12 {
		t.Error("NoOfEMIs calculation is wrong")
	}

	if maploanAccounts["IDIDI_Dale"].EMIAmount != 442 {
		t.Error("EMIAmount Calculation is wrong")
	}
}
