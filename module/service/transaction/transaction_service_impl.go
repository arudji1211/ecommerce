package transaction

import (
	"context"

	MDtransaction "github.com/arudji1211/ecommerce/module/model/transaction"
	Rtransaction "github.com/arudji1211/ecommerce/module/repository/transaction"
)

type TransactionServiceImpl struct {
	TRepo Rtransaction.TransactionRepository
}

func (t *TransactionServiceImpl) GetAllByUserID(ctx context.Context, UserID uint) (TransactionOut []MDtransaction.Transaction, err error) {
	//panic("not implemented") // TODO: Implement
	TransactionOut, err = t.TRepo.GetAllByUserID(ctx, UserID)
	if err != nil {
		return
	}
	return
}

func (t *TransactionServiceImpl) GetByID(ctx context.Context, ID uint) (TransactionOut []MDtransaction.Transaction, err error) {
	//panic("not implemented") // TODO: Implement
	TransactionOut, err = t.TRepo.GetByID(ctx, ID)
	if err != nil {
		return
	}
	return
}

func (t *TransactionServiceImpl) GetAll(ctx context.Context, UserID uint) (TransactionOut []MDtransaction.Transaction, err error) {
	//panic("not implemented") // TODO: Implement
	TransactionOut, err = t.TRepo.GetAll(ctx)
	if err != nil {
		return
	}
	return
}

func (t *TransactionServiceImpl) Create(ctx context.Context, ID uint, TransactionIn MDtransaction.Transaction) (MDtransaction.Transaction, error) {
	//panic("not implemented") // TODO: Implement
	TransactionIn, err := t.TRepo.Create(ctx, TransactionIn)
	if err != nil {
		return TransactionIn, err
	}
	return TransactionIn, nil
}

func (t *TransactionServiceImpl) UpdateByID(ctx context.Context, ID uint, TransactionIn MDtransaction.Transaction) (err error) {
	//panic("not implemented") // TODO: Implement
	err = t.TRepo.UpdateByID(ctx, ID, TransactionIn)
	if err != nil {
		return
	}
	return
}

func (t *TransactionServiceImpl) DeleteByID(ctx context.Context, ID uint) (err error) {
	//panic("not implemented") // TODO: Implement
	err = t.TRepo.DeleteByID(ctx, ID)
	if err != nil {
		return
	}
	return
}

func NewTransactionService(repo Rtransaction.TransactionRepository) TransactionService {
	return &TransactionServiceImpl{
		TRepo: repo,
	}
}
