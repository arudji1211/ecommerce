package transactiondetail

import (
	"context"

	transactiondetail "github.com/arudji1211/ecommerce/module/model/transactionDetail"
)

type TransactionDetailRepository interface {
	GetAllByTransactionID(ctx context.Context, TransactionID uint) (TransactionDetailOut []transactiondetail.TransactionDetail, err error)
	GetByTransactionDetailID(ctx context.Context, TransactionID uint) (TransactionDetailOut transactiondetail.TransactionDetail, err error)
	UpdateByTransactionDetailID(ctx context.Context, TransactionID uint) (err error)
	DeleteByTransactionDetailID(ctx context.Context, TransactionID uint) (err error)
}
