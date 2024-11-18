package transactiondetail

import (
	"context"

	MDtransactiondetail "github.com/arudji1211/ecommerce/module/model/transactionDetail"
	RPtransactiondetail "github.com/arudji1211/ecommerce/module/repository/transactiondetail"
)

type TransactiondetailServiceImpl struct {
	TDRepo RPtransactiondetail.TransactionDetailRepository
}

func (t *TransactiondetailServiceImpl) GetAllByTransactionID(ctx context.Context, TransactionID uint) (TransactionDetailOut []MDtransactiondetail.TransactionDetail, err error) {
	//panic("not implemented") // TODO: Implement
	TransactionDetailOut, err = t.TDRepo.GetAllByTransactionID(ctx, TransactionID)
	if err != nil {
		return
	}
	return
}

func (t *TransactiondetailServiceImpl) GetByTransactionDetailID(ctx context.Context, TransactionID uint) (TransactionDetailOut MDtransactiondetail.TransactionDetail, err error) {
	//panic("not implemented") // TODO: Implement
	TransactionDetailOut, err = t.TDRepo.GetByTransactionDetailID(ctx, TransactionID)
	if err != nil {
		return
	}
	return
}

func (t *TransactiondetailServiceImpl) Create(ctx context.Context, TransactionDetailIn MDtransactiondetail.TransactionDetail) (MDtransactiondetail.TransactionDetail, error) {
	//panic("not implemented") // TODO: Implement
	TransactionDetailIn, err := t.TDRepo.Create(ctx, TransactionDetailIn)
	if err != nil {
		return TransactionDetailIn, err
	}
	return TransactionDetailIn, err
}

func (t *TransactiondetailServiceImpl) UpdateByTransactionDetailID(ctx context.Context, TransactionDetailIn MDtransactiondetail.TransactionDetail) (err error) {
	//panic("not implemented") // TODO: Implement
	err = t.TDRepo.UpdateByTransactionDetailID(ctx, TransactionDetailIn)
	if err != nil {
		return
	}
	return
}

func (t *TransactiondetailServiceImpl) DeleteByTransactionDetailID(ctx context.Context, TransactionID uint) (err error) {
	//panic("not implemented") // TODO: Implement
	err = t.TDRepo.DeleteByTransactionDetailID(ctx, TransactionID)
	if err != nil {
		return
	}
	return
}

func NewTransactiondetailService(repo RPtransactiondetail.TransactionDetailRepository) TransactiondetailService {
	return &TransactiondetailServiceImpl{
		TDRepo: repo,
	}
}