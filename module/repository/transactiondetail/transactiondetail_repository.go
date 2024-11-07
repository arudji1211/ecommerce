package transactiondetail

import (
	"context"

	MDtransactiondetail "github.com/arudji1211/ecommerce/module/model/transactionDetail"
)

type TransactionDetailRepository interface {
	GetAllByTransactionID(ctx context.Context, TransactionID uint) (TransactionDetailOut []MDtransactiondetail.TransactionDetail, err error)
	GetByTransactionDetailID(ctx context.Context, TransactionID uint) (TransactionDetailOut MDtransactiondetail.TransactionDetail, err error)
	Create(ctx context.Context, TransactionDetailIn MDtransactiondetail.TransactionDetail) (MDtransactiondetail.TransactionDetail, error)
	UpdateByTransactionDetailID(ctx context.Context, TransactionDetailIn MDtransactiondetail.TransactionDetail) (err error)
	DeleteByTransactionDetailID(ctx context.Context, TransactionID uint) (err error)
}
