package main

import "testing"

func Test_should_create_new_lumpsumPaymentEntryIntoExistingLoanAccount_when_received_paymentcommand(t *testing.T) {
	//Arrange
	Init()
	l := LoanCommand{
		BankName:        "IDIDI",
		BorrowerName:    "Dale",
		PrincipalAmount: 5000,
		NoOfYears:       1,
		RateOfInterest:  6,
	}
	l.CreateLoanAccount()
	p := PaymentCommand{
		BankName:                  "IDIDI",
		BorrowerName:              "Dale",
		LumpSumAmount:             1000,
		LumpSumPaidAfterEMINumber: 5,
	}

	//Act
	p.UpdateLoanAccount()

	//Assert
	if _, ok := maploanAccounts["IDIDI_Dale"]; !ok {
		t.Error("Loan account map entry not found")
	}

	if len(maploanAccounts["IDIDI_Dale"].LumsumPayments) != 1 &&
		maploanAccounts["IDIDI_Dale"].LumsumPayments[0].LumpSumAmount != 1000 {
		t.Error("Lumsum entry is not recorded")
	}
}
