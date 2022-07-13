package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var maploanAccounts map[string]LoanAccount

func main() {

	// 	if len(os.Args) != 2 {
	// 		fmt.Println(`
	// Program excuted with insufficient arguments.
	// Sample Execution:
	// ./TheLedgerCo.exe "<TestFilePath>"`)
	// 		return
	// 	}
	//testFilePath := os.Args[1]
	testFilePath := "SampleInput1.txt"

	inputFileData, err := os.ReadFile(testFilePath)
	if err != nil {
		fmt.Printf("Cannot read the file: %v", err)
	}

	maploanAccounts = make(map[string]LoanAccount)

	//Split Commands
	fileCommands := strings.Split(string(inputFileData), "\n")
	for _, c := range fileCommands {
		commandDetails := strings.Split(c, " ")

		switch commandDetails[0] {
		case "LOAN":
			pa, _ := strconv.ParseFloat(commandDetails[3], 64)
			noy, _ := strconv.ParseInt(commandDetails[4], 10, 64)
			roi, _ := strconv.ParseFloat(commandDetails[5], 64)
			l := LoanCommand{
				BankName:        commandDetails[1],
				BorrowerName:    commandDetails[2],
				PrincipalAmount: pa,
				NoOfYears:       int(noy),
				RateOfInterest:  roi,
			}
			l.CreateLoanAccount()

		case "PAYMENT":
			la, _ := strconv.ParseFloat(commandDetails[3], 64)
			laen, _ := strconv.ParseInt(commandDetails[4], 10, 64)
			p := PaymentCommand{
				BankName:                  commandDetails[1],
				BorrowerName:              commandDetails[2],
				LumpSumAmount:             la,
				LumpSumPaidAfterEMINumber: int(laen),
			}
			p.UpdateLoanAccount()

		case "BALANCE":
			remis, _ := strconv.ParseInt(commandDetails[3], 10, 64)
			b := BalanceCommand{
				BankName:     commandDetails[1],
				BorrowerName: commandDetails[2],
				EMIs:         int(remis),
			}
			br := b.GetLoanAccountDetails()
			fmt.Println(fmt.Sprintf("%v %v %v %v", br.BankName, br.BorrowerName, br.AmountPaid, br.Remaining_EMIs))
		}

	}
}