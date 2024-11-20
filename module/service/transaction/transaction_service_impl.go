package transaction

import (
	"context"

	MDtransaction "github.com/arudji1211/ecommerce/module/model/transaction"
	Rtransaction "github.com/arudji1211/ecommerce/module/repository/transaction"
	"gorm.io/gorm"
)

type TransactionServiceImpl struct {
	TRepo Rtransaction.TransactionRepository
}

func (t *TransactionServiceImpl) GetAllByUserID(ctx context.Context, db *gorm.DB, UserID uint) (TransactionOut []MDtransaction.Transaction, err error) {
	//panic("not implemented") // TODO: Implement
	TransactionOut, err = t.TRepo.GetAllByUserID(ctx, db, UserID)
	if err != nil {
		return
	}
	return
}

func (t *TransactionServiceImpl) GetByID(ctx context.Context, db *gorm.DB, ID uint) (TransactionOut []MDtransaction.Transaction, err error) {
	//panic("not implemented") // TODO: Implement
	TransactionOut, err = t.TRepo.GetByID(ctx, db, ID)
	if err != nil {
		return
	}
	return
}

func (t *TransactionServiceImpl) GetAll(ctx context.Context, db *gorm.DB, UserID uint) (TransactionOut []MDtransaction.Transaction, err error) {
	//panic("not implemented") // TODO: Implement
	TransactionOut, err = t.TRepo.GetAll(ctx, db)
	if err != nil {
		return
	}
	return
}

func (t *TransactionServiceImpl) Create(ctx context.Context, db *gorm.DB, ID uint, TransactionIn MDtransaction.Transaction) (MDtransaction.Transaction, error) {
	//panic("not implemented") // TODO: Implement
	TransactionIn, err := t.TRepo.Create(ctx, db, TransactionIn)
	if err != nil {
		return TransactionIn, err
	}
	return TransactionIn, nil
}

func (t *TransactionServiceImpl) UpdateByID(ctx context.Context, db *gorm.DB, ID uint, TransactionIn MDtransaction.Transaction) (err error) {
	//panic("not implemented") // TODO: Implement
	err = t.TRepo.UpdateByID(ctx, db, ID, TransactionIn)
	if err != nil {
		return
	}
	return
}

func (t *TransactionServiceImpl) DeleteByID(ctx context.Context, db *gorm.DB, ID uint) (err error) {
	//panic("not implemented") // TODO: Implement
	err = t.TRepo.DeleteByID(ctx, db, ID)
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
