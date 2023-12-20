package repositiory

import domain "github.com/neutrin/expense_sharing_app/internal/core/domian"

type TransactionRepo interface {
	Save(domain.Transaction)
	Delete(owedBy, owedTo string) error
	Get(owedBy string) []*domain.Transaction
	GetAll() []*domain.Transaction
}
