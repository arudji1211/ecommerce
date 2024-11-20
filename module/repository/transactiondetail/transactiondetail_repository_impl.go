package transactiondetail

import (
	"context"

	MDtransactiondetail "github.com/arudji1211/ecommerce/module/model/transactionDetail"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type TransactionDetailRepositoryImpl struct {
	Db *gorm.DB
}

func (T *TransactionDetailRepositoryImpl) GetAllByTransactionID(ctx context.Context, db *gorm.DB, TransactionID uint) (TransactionDetailOut []MDtransactiondetail.TransactionDetail, err error) {
	tx := T.Db.Model(MDtransactiondetail.TransactionDetail{}).Where("TransactionID = ?", TransactionID).Find(&TransactionDetailOut)
	if err = tx.Error; err != nil {
		return
	}
	return
}

func (T *TransactionDetailRepositoryImpl) GetByTransactionDetailID(ctx context.Context, db *gorm.DB, TransactionID uint) (TransactionDetailOut MDtransactiondetail.TransactionDetail, err error) {
	tx := T.Db.Model(MDtransactiondetail.TransactionDetail{}).Where("ID = ?", TransactionID).Find(&TransactionDetailOut)
	if err = tx.Error; err != nil {
		return
	}
	return
}

func (T *TransactionDetailRepositoryImpl) Create(ctx context.Context, db *gorm.DB, TransactionDetailIn MDtransactiondetail.TransactionDetail) (MDtransactiondetail.TransactionDetail, error) {
	tx := T.Db.Model(MDtransactiondetail.TransactionDetail{}).Clauses(clause.Returning{}).Create(&TransactionDetailIn)
	if err := tx.Error; err != nil {
		return TransactionDetailIn, err
	}
	return TransactionDetailIn, nil
}

func (T *TransactionDetailRepositoryImpl) UpdateByTransactionDetailID(ctx context.Context, db *gorm.DB, TransactionDetailIn MDtransactiondetail.TransactionDetail) (err error) {
	tx := T.Db.Model(MDtransactiondetail.TransactionDetail{}).Where("ID = ?", TransactionDetailIn.ID).Updates(&TransactionDetailIn)
	if err = tx.Error; err != nil {
		return
	}
	return
}

func (T *TransactionDetailRepositoryImpl) DeleteByTransactionDetailID(ctx context.Context, db *gorm.DB, TransactionID uint) (err error) {
	tx := T.Db.Model(MDtransactiondetail.TransactionDetail{}).Where("ID = ?", TransactionID).Delete(MDtransactiondetail.TransactionDetail{})
	if err = tx.Error; err != nil {
		return
	}
	return
}

func (T *TransactionDetailRepositoryImpl) IsTransaction(Db *gorm.DB) {
	if Db != nil {
		T.Db = Db
	}
}

func NewTransactionDetailRepository(db *gorm.DB) TransactionDetailRepository {
	return &TransactionDetailRepositoryImpl{
		Db: db,
	}
}
