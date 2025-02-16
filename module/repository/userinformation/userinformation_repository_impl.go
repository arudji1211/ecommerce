package userinformation

import (
	"context"

	MDuserinformation "github.com/arudji1211/ecommerce/module/model/userinformation"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserInformationRepositoryImpl struct {
	Db *gorm.DB
}

func (U *UserInformationRepositoryImpl) GetAll(ctx context.Context, db *gorm.DB) (UsersOut []MDuserinformation.UserInformation, err error) {
	tx := U.Db.Model(MDuserinformation.UserInformation{}).Find(&UsersOut)
	if err = tx.Error; err != nil {
		return
	}
	return
}

func (U *UserInformationRepositoryImpl) GetUserInformationByUserID(ctx context.Context, db *gorm.DB, UserID string) (UserOut MDuserinformation.UserInformation, err error) {
	// panic("not implemented") // TODO: Implement
	tx := U.Db.Model(MDuserinformation.UserInformation{}).Where("UserID = ? ", UserID).Find(&UserOut)
	if err = tx.Error; err != nil {
		return
	}
	return
}

func (U *UserInformationRepositoryImpl) Create(ctx context.Context, db *gorm.DB, UserIn MDuserinformation.UserInformation) (MDuserinformation.UserInformation, error) {
	//panic("not implemented") // TODO: Implement
	tx := U.Db.Model(MDuserinformation.UserInformation{}).Clauses(clause.Returning{}).Create(&UserIn)
	if err := tx.Error; err != nil {
		return UserIn, err
	}
	return UserIn, nil
}

func (U *UserInformationRepositoryImpl) Update(ctx context.Context, db *gorm.DB, UserIn MDuserinformation.UserInformation) (err error) {
	tx := U.Db.Model(MDuserinformation.UserInformation{}).Where("UserID = ? ", UserIn.UserID).Updates(&UserIn)
	if err = tx.Error; err != nil {
		return
	}
	return
}

func (U *UserInformationRepositoryImpl) Delete(ctx context.Context, db *gorm.DB, ID uint) (err error) {
	tx := U.Db.Model(MDuserinformation.UserInformation{}).Where("ID = ? ", ID).Delete(&MDuserinformation.UserInformation{})
	if err = tx.Error; err != nil {
		return
	}
	return
}

func (U *UserInformationRepositoryImpl) IsTransaction(Db *gorm.DB) {
	if Db != nil {
		U.Db = Db
	}
}

func NewUserInformationRepositoryImpl(db *gorm.DB) UserInformationRepository {
	return &UserInformationRepositoryImpl{
		Db: db,
	}
}
