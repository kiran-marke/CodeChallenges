package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var maploanAccounts map[string]*LoanAccount

func main() {

	if len(os.Args) != 2 {
		fmt.Println(`
	Program excuted with insufficient arguments.
	Sample Execution:
	./TheLedgerCo "<TestFilePath>"`)
		return
	}
	testFilePath := os.Args[1]
	//testFilePath := "SampleInputMultiplePayments.txt"

	inputFileData, err := os.ReadFile(testFilePath)
	if err != nil {
		fmt.Printf("Cannot read the file: %v", err)
	}

	maploanAccounts = make(map[string]*LoanAccount)

	printStatments := ReadCommandsOneByOneAndTakeAction(string(inputFileData))

	for _, s := range printStatments {
		fmt.Println(s)
	}

}

func ReadCommandsOneByOneAndTakeAction(fileData string) []string {
	//Split Commands
	fileCommands := strings.Split(fileData, "\n")
	var printStatements []string

	for _, c := range fileCommands {
		commandDetails := strings.Split(c, " ")

		switch commandDetails[0] {
		case "LOAN":
			pa, _ := strconv.ParseInt(commandDetails[3], 10, 64)
			noy, _ := strconv.ParseInt(commandDetails[4], 10, 64)
			roi, _ := strconv.ParseInt(commandDetails[5], 10, 64)
			l := LoanCommand{
				BankName:        commandDetails[1],
				BorrowerName:    commandDetails[2],
				PrincipalAmount: int(pa),
				NoOfYears:       int(noy),
				RateOfInterest:  int(roi),
			}
			l.CreateLoanAccount()

		case "PAYMENT":
			la, _ := strconv.ParseInt(commandDetails[3], 10, 64)
			laen, _ := strconv.ParseInt(commandDetails[4], 10, 64)
			p := PaymentCommand{
				BankName:                  commandDetails[1],
				BorrowerName:              commandDetails[2],
				LumpSumAmount:             int(la),
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
			printStatements = append(printStatements, fmt.Sprintf("%v %v %v %v", br.BankName, br.BorrowerName, br.AmountPaid, br.Remaining_EMIs))
		}
	}
	return printStatements
}
