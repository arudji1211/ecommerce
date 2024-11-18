package transaction

import (
	"context"

	MDtransaction "github.com/arudji1211/ecommerce/module/model/transaction"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type TransactionRepositoryImpl struct {
	Db *gorm.DB
}

func (t *TransactionRepositoryImpl) GetAllByUserID(ctx context.Context, UserID uint) (TransactionOut []MDtransaction.Transaction, err error) {
	tx := t.Db.Model(MDtransaction.Transaction{}).Where("UserID = ?", UserID).Find(&TransactionOut)
	if err = tx.Error; err != nil {
		return
	}
	return
}
func (t *TransactionRepositoryImpl) GetByID(ctx context.Context, ID uint) (TransactionOut []MDtransaction.Transaction, err error) {
	tx := t.Db.Model(MDtransaction.Transaction{}).Where("ID = ?", ID).Find(&TransactionOut)
	if err = tx.Error; err != nil {
		return
	}
	return
}

func (t *TransactionRepositoryImpl) GetAll(ctx context.Context) (TransactionOut []MDtransaction.Transaction, err error) {
	tx := t.Db.Model(MDtransaction.Transaction{}).Find(&TransactionOut)
	if err = tx.Error; err != nil {
		return
	}
	return
}

func (t *TransactionRepositoryImpl) UpdateByID(ctx context.Context, ID uint, TransactionIn MDtransaction.Transaction) (err error) {
	tx := t.Db.Model(MDtransaction.Transaction{}).Where("ID = ?", ID).Updates(&TransactionIn)
	if err = tx.Error; err != nil {
		return
	}
	return
}

func (t *TransactionRepositoryImpl) Create(ctx context.Context, TransactionIn MDtransaction.Transaction) (MDtransaction.Transaction, error) {
	tx := t.Db.Model(MDtransaction.Transaction{}).Clauses(clause.Returning{}).Create(&TransactionIn)
	if err := tx.Error; err != nil {
		return TransactionIn, err
	}
	return TransactionIn, nil
}

func (t *TransactionRepositoryImpl) DeleteByID(ctx context.Context, ID uint) (err error) {
	tx := t.Db.Model(MDtransaction.Transaction{}).Where("ID = ?", ID).Delete(MDtransaction.Transaction{})
	if err = tx.Error; err != nil {
		return
	}
	return
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &TransactionRepositoryImpl{
		Db: db,
	}
}
