package transactiondetail

import (
	"context"

	transactiondetail "github.com/arudji1211/ecommerce/module/model/transactionDetail"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type TransactionDetailRepositoryImpl struct {
	Db *gorm.DB
}

func (T *TransactionDetailRepositoryImpl) GetAllByTransactionID(ctx context.Context, TransactionID uint) (TransactionDetailOut []transactiondetail.TransactionDetail, err error) {
	tx := T.Db.Model(transactiondetail.TransactionDetail{}).Where("TransactionID = ?", TransactionID).Find(&TransactionDetailOut)
	if err = tx.Error; err != nil {
		return
	}
	return
}

func (T *TransactionDetailRepositoryImpl) GetByTransactionDetailID(ctx context.Context, TransactionID uint) (TransactionDetailOut transactiondetail.TransactionDetail, err error) {
	tx := T.Db.Model(transactiondetail.TransactionDetail{}).Where("ID = ?", TransactionID).Find(&TransactionDetailOut)
	if err = tx.Error; err != nil {
		return
	}
	return
}

func (T *TransactionDetailRepositoryImpl) Create(ctx context.Context, TransactionDetailIn transactiondetail.TransactionDetail) (transactiondetail.TransactionDetail, error) {
	tx := T.Db.Model(transactiondetail.TransactionDetail{}).Clauses(clause.Returning{}).Create(&TransactionDetailIn)
	if err := tx.Error; err != nil {
		return TransactionDetailIn, err
	}
	return TransactionDetailIn, nil
}

func (T *TransactionDetailRepositoryImpl) UpdateByTransactionDetailID(ctx context.Context, TransactionDetailIn transactiondetail.TransactionDetail) (err error) {
	tx := T.Db.Model(transactiondetail.TransactionDetail{}).Where("ID = ?", TransactionDetailIn.ID).Updates(&TransactionDetailIn)
	if err = tx.Error; err != nil {
		return
	}
	return
}

func (T *TransactionDetailRepositoryImpl) DeleteByTransactionDetailID(ctx context.Context, TransactionID uint) (err error) {
	tx := T.Db.Model(transactiondetail.TransactionDetail{}).Where("ID = ?", TransactionID).Delete(transactiondetail.TransactionDetail{})
	if err = tx.Error; err != nil {
		return
	}
	return
}

func NewTransactionDetailRepository(db *gorm.DB) TransactionDetailRepository {
	return &TransactionDetailRepositoryImpl{
		Db: db,
	}
}
