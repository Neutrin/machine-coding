package main

import (
	"fmt"
	"os"

	domain "github.com/neutrin/expense_sharing_app/internal/core/domian"
	"github.com/neutrin/expense_sharing_app/internal/repositiories"
	"github.com/neutrin/expense_sharing_app/internal/service"
)

func main() {
	repo := repositiories.NewTransactionRepoMap()
	expenseService := service.NewExpenseService(repo)
	//SHOW
	showCommand(expenseService)
	//SHOW u1
	showCommandById(expenseService, "u1")
	//EXPENSE u1 1000 4 u1 u2 u3 u4 EQUAL
	equalExpense := domain.NewExpense("u1", []string{"u1", "u2", "u3", "u4"}, float64(1000), domain.Equal, []float64{})
	err := expenseService.AddExpense(equalExpense)
	if err != nil {
		fmt.Printf(" error = %+v\n", err)
		os.Exit(1)
	}
	//showCommand(expenseService)
	//SHOW u4
	showCommandById(expenseService, "u4")
	//SHOW u1
	showCommandById(expenseService, "u1")
	//EXPENSE u1 1250 2 u2 u3 EXACT 370 880
	exactExpense := domain.NewExpense("u1", []string{"u2", "u3"}, float64(1250), domain.Exact,
		[]float64{float64(370), float64(880)})
	err = expenseService.AddExpense(exactExpense)
	if err != nil {
		fmt.Printf(" erroor = %+v\n", err)
		os.Exit(1)
	}
	//SHOW
	showCommand(expenseService)
	//EXPENSE u4 1200 4 u1 u2 u3 u4 PERCENT 40 20 20 20
	percentageExpense := domain.NewExpense("u4", []string{"u1", "u2", "u3", "u4"}, float64(1200), domain.Percentage, []float64{float64(40), float64(20), float64(20), float64(20)})
	err = expenseService.AddExpense(percentageExpense)
	if err != nil {
		fmt.Printf(" error = %+v \n", err)
		os.Exit(1)
	}
	//SHOW u1
	showCommandById(expenseService, "u1")
	//SHOW
	showCommand(expenseService)

}

func showCommand(service service.ExpenseService) {
	fmt.Println("--------------------------------------------------------------")
	for _, curResp := range service.Show() {
		fmt.Println(curResp)
	}
	fmt.Println("--------------------------------------------------------------")
}

func showCommandById(service service.ExpenseService, userId string) {
	fmt.Println("-------------------------------------------------------------")
	for _, curResp := range service.ShouwId(userId) {
		fmt.Println(curResp)
	}
	fmt.Println("-------------------------------------------------------------")
}
