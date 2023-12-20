package repositiories

import (
	"fmt"

	domain "github.com/neutrin/expense_sharing_app/internal/core/domian"
	"github.com/neutrin/expense_sharing_app/internal/core/repositiory"
)

type TransactionRepoMap struct {
	mp map[string]map[string]domain.Transaction
}

func NewTransactionRepoMap() repositiory.TransactionRepo {
	return &TransactionRepoMap{
		mp: make(map[string]map[string]domain.Transaction),
	}
}

func (repo *TransactionRepoMap) Save(txn domain.Transaction) {
	var (
		owedBy = txn.OwedBy
		owedTo = txn.OwedTo
	)

	_, exists := repo.mp[owedBy]
	if !exists {
		repo.mp[owedBy] = make(map[string]domain.Transaction)
	}
	repo.mp[owedBy][owedTo] = txn

}

func (repo *TransactionRepoMap) Delete(owedBy, owedTo string) error {
	owedByTxn, exists := repo.mp[owedBy]
	if !exists {
		return fmt.Errorf(" not exists")
	}
	delete(owedByTxn, owedTo)
	repo.mp[owedBy] = owedByTxn
	return nil
}

func (repo *TransactionRepoMap) Get(owedBy string) []*domain.Transaction {
	txn := make([]*domain.Transaction, 0)
	for key := range repo.mp[owedBy] {
		txn = append(txn, domain.NewTransaction(repo.mp[owedBy][key].OwedBy, repo.mp[owedBy][key].OwedTo, repo.mp[owedBy][key].Amount))

	}
	return txn
}

func (repo *TransactionRepoMap) GetAll() []*domain.Transaction {
	var txn = make([]*domain.Transaction, 0)
	//ß	fmt.Println(" ******** GetAll ********* len becomes =", len(repo.mp))
	for _, owedByTxns := range repo.mp {
		for _, curTxn := range owedByTxns {
			txn = append(txn, domain.NewTransaction(curTxn.OwedBy, curTxn.OwedTo, curTxn.Amount))
		}
	}
	return txn
}
