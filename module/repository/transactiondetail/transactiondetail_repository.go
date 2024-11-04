package transactiondetail

import (
	"context"

	transactiondetail "github.com/arudji1211/ecommerce/module/model/transactionDetail"
)

type TransactionDetailRepository interface {
	GetAllByTransactionID(ctx context.Context, TransactionID uint) (TransactionDetailOut []transactiondetail.TransactionDetail, err error)
	GetByTransactionDetailID(ctx context.Context, TransactionID uint) (TransactionDetailOut transactiondetail.TransactionDetail, err error)
	Create(ctx context.Context, TransactionDetailIn transactiondetail.TransactionDetail) (transactiondetail.TransactionDetail, error)
	UpdateByTransactionDetailID(ctx context.Context, TransactionDetailIn transactiondetail.TransactionDetail) (err error)
	DeleteByTransactionDetailID(ctx context.Context, TransactionID uint) (err error)
}
