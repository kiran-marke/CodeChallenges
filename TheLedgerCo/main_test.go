package main

import (
	"os"
	"strings"
	"testing"
)

func Test_Should_Read_SampleFileData_Returns_MatchedData(t *testing.T) {
	//Act
	testInputFilePath1 := "SampleData/SampleInput1.txt"
	testOutputFilePath1 := "SampleData/SampleOutput1.txt"

	testInputFilePath2 := "SampleData/SampleInput2.txt"
	testOutputFilePath2 := "SampleData/SampleOutput2.txt"

	testInputFilePath3 := "SampleData/SampleInputMultiplePayments.txt"
	testOutputFilePath3 := "SampleData/SampleOutputMultiplePayments.txt"

	inputFileData1, _ := os.ReadFile(testInputFilePath1)
	inputFileData2, _ := os.ReadFile(testInputFilePath2)
	inputFileData3, _ := os.ReadFile(testInputFilePath3)

	outputFileData1, _ := os.ReadFile(testOutputFilePath1)
	outputFileData2, _ := os.ReadFile(testOutputFilePath2)
	outputFileData3, _ := os.ReadFile(testOutputFilePath3)

	//Action
	printStatements1 := ReadCommandsOneByOneAndTakeAction(string(inputFileData1))
	printStatements2 := ReadCommandsOneByOneAndTakeAction(string(inputFileData2))
	printStatements3 := ReadCommandsOneByOneAndTakeAction(string(inputFileData3))

	//Assert
	outputResults1 := strings.Split(string(outputFileData1), "\n")
	if len(outputResults1) != len(printStatements1) {
		t.Error("Count Results for InputTestFile1 are not matching")
	}

	outputResults2 := strings.Split(string(outputFileData2), "\n")
	if len(outputResults2) != len(printStatements2) {
		t.Error("Count Results for InputTestFile1 are not matching")
	}

	outputResults3 := strings.Split(string(outputFileData3), "\n")
	if len(outputResults3) != len(printStatements3) {
		t.Error("Count Results for InputTestFile1 are not matching")
	}

	for i1, o1 := range outputResults1 {
		result := printStatements1[i1]
		if o1 != result {
			t.Error("Results for InputTestFile1 are not matching")
		}
	}

	for i2, o2 := range outputResults2 {
		result := printStatements2[i2]
		if o2 != result {
			t.Error("Results for InputTestFile2 are not matching")
		}
	}

	for i3, o3 := range outputResults3 {
		result := printStatements3[i3]
		if o3 != result {
			t.Error("Results for InputTestFile3 are not matching")
		}
	}
}
