package service

import (
	"fmt"
	"math"
	"strings"

	domain "github.com/neutrin/expense_sharing_app/internal/core/domian"
	"github.com/neutrin/expense_sharing_app/internal/core/repositiory"
)

type ExpenseService struct {
	repo repositiory.TransactionRepo
}

func NewExpenseService(repo repositiory.TransactionRepo) ExpenseService {
	return ExpenseService{
		repo: repo,
	}
}

/*
 */
func (service ExpenseService) AddExpense(expense domain.ExpenseModel) error {
	expenseHandler := domain.ExpenseByType(expense.Type)
	if expenseHandler == nil {
		return fmt.Errorf(" transaction type not registered")
	}

	txn, err := expenseHandler.Transactions(expense)
	if err != nil {
		return err
	}

	for index, _ := range txn {
		curTxn := txn[index]
		originalTxn := service.originalTxn(curTxn.OwedBy, curTxn.OwedTo)
		counterTxn := service.counterTxn(curTxn.OwedBy, curTxn.OwedTo)
		//fmt.Printf(" owedBy = %s and owedto = %s  cur txn amount = %f, original txn amount = %f counter txn amount = %f\n",
		//curTxn.OwedBy, curTxn.OwedTo, curTxn.Amount, originalTxn.Amount, counterTxn.Amount)
		originalTxn.Amount = math.Round(max(0.0, curTxn.Amount+originalTxn.Amount-counterTxn.Amount))
		counterTxn.Amount = math.Round(max(0.0, counterTxn.Amount-(curTxn.Amount+originalTxn.Amount)))
		service.saveTxn(originalTxn)
		service.saveTxn(counterTxn)
		// log.Printf(" original txn becomes = %+v\n", *originalTxn)
		// log.Printf(" counter txn becomes = %+v\n", *counterTxn)
	}

	return err

}

func (service ExpenseService) Show() []string {
	resp := make([]string, 0)
	for _, curTxn := range service.repo.GetAll() {
		resp = append(resp, fmt.Sprintf("%s Owes %s : %.2f", curTxn.OwedBy, curTxn.OwedTo, curTxn.Amount))
	}
	return formatResp(resp)
}

func (service ExpenseService) ShouwId(userId string) []string {
	resp := make([]string, 0)
	allTxn := service.repo.GetAll()
	for _, curTxn := range allTxn {
		if strings.Compare(curTxn.OwedBy, userId) == 0 {
			resp = append(resp, fmt.Sprintf("%s Owes %s : %.2f", curTxn.OwedBy, curTxn.OwedTo, math.Round(curTxn.Amount)))
		}
	}
	for _, curTxn := range allTxn {
		//fmt.Printf(" ****** owed by = %s , owed to = %s amount = %f \n", curTxn.OwedBy, curTxn.OwedTo, curTxn.Amount)
		if strings.Compare(curTxn.OwedTo, userId) == 0 {
			//fmt.Printf(" owed by = %s , owed to = %s amount = %f \n", curTxn.OwedBy, curTxn.OwedTo, curTxn.Amount)
			resp = append(resp, fmt.Sprintf("%s Owes %s : %.2f", curTxn.OwedBy, curTxn.OwedTo, math.Round(curTxn.Amount)))
		}
	}
	return formatResp(resp)
}

func formatResp(resp []string) []string {
	if len(resp) == 0 {
		resp = append(resp, fmt.Sprintf(" No Balances"))
	}
	return resp
}
func max(a, b float64) float64 {
	if a <= b {
		return b
	}
	return a
}

func (service ExpenseService) counterTxn(owedBy string, owedTo string) *domain.Transaction {
	var txn *domain.Transaction
	txn = domain.NewTransaction(owedTo, owedBy, 0.0)
	for _, curTxn := range service.repo.Get(owedTo) {
		if strings.Compare(curTxn.OwedTo, owedBy) == 0 {
			txn = curTxn
			break
		}
	}
	return txn
}

func (service ExpenseService) originalTxn(owedBy, owedTo string) *domain.Transaction {
	var txn *domain.Transaction
	txn = domain.NewTransaction(owedBy, owedTo, 0.0)
	for _, curTxn := range service.repo.Get(owedBy) {
		if strings.Compare(owedTo, curTxn.OwedTo) == 0 {
			txn = curTxn
			break
		}
	}
	return txn
}

// Made unexported so that cannot be used by other packages
func (service ExpenseService) saveTxn(txn *domain.Transaction) {

	if txn.Amount <= 0.0 {
		//fmt.Printf(" **************txn sent for delete owedby = %s owedto = %s amount = %v\n", txn.OwedBy, txn.OwedTo, txn.Amount)
		service.repo.Delete(txn.OwedBy, txn.OwedTo)
		return
	}
	//fmt.Printf(" ********txn sent for update owedby = %s owedto = %s amount = %v\n", txn.OwedBy, txn.OwedTo, txn.Amount)
	service.repo.Save(*txn)
	return
}
