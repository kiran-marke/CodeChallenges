package main

import "testing"

func Test_To_Check_UniqueAccountName(t *testing.T) {
	//Act
	expectedName := "ICICI_Kiran"

	//Arrange
	actualName := getLoanAccountUniqueName("ICICI", "Kiran")

	//Assert
	if actualName != expectedName {
		t.Error("UniqueName is not created in correct pattern")
	}
}
