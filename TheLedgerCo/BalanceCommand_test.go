package main

import "testing"

func Test_should_provide_correctbalance_entry_when_received_balancecommand(t *testing.T) {
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

	b := BalanceCommand{
		BankName:     "IDIDI",
		BorrowerName: "Dale",
		EMIs:         3,
	}

	//Act
	br := b.GetLoanAccountDetails()

	//Assert

	if br.AmountPaid != 1326 {
		t.Error("AmountPaid calculation is wrong")
	}

	if br.Remaining_EMIs != 9 {
		t.Error("Remaining_EMIs calculation is wrong")
	}
}

func Test_should_provide_correctbalance_entry_when_received_balancecommand_afterpaymentcommand(t *testing.T) {
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
	p.UpdateLoanAccount()

	b := BalanceCommand{
		BankName:     "IDIDI",
		BorrowerName: "Dale",
		EMIs:         6,
	}

	//Act
	br := b.GetLoanAccountDetails()

	//Assert

	if br.AmountPaid != 3652 {
		t.Error("AmountPaid calculation is wrong")
	}

	if br.Remaining_EMIs != 4 {
		t.Error("Remaining_EMIs calculation is wrong")
	}
}
