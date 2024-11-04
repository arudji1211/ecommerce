package userinformation

import (
	"context"

	"github.com/arudji1211/ecommerce/module/model/userinformation"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserInformationRepositoryImpl struct {
	Db *gorm.DB
}

func (U *UserInformationRepositoryImpl) GetAll(ctx context.Context) (UsersOut []userinformation.UserInformation, err error) {
	tx := U.Db.Model(userinformation.UserInformation{}).Find(&UsersOut)
	if err = tx.Error; err != nil {
		return
	}
	return
}

func (U *UserInformationRepositoryImpl) GetUserInformationByUserID(ctx context.Context, UserID string) (UserOut userinformation.UserInformation, err error) {
	// panic("not implemented") // TODO: Implement
	tx := U.Db.Model(userinformation.UserInformation{}).Where("UserID = ? ", UserID).Find(&UserOut)
	if err = tx.Error; err != nil {
		return
	}
	return
}

func (U *UserInformationRepositoryImpl) Create(ctx context.Context, UserIn userinformation.UserInformation) (userinformation.UserInformation, error) {
	//panic("not implemented") // TODO: Implement
	tx := U.Db.Model(userinformation.UserInformation{}).Clauses(clause.Returning{}).Create(&UserIn)
	if err := tx.Error; err != nil {
		return UserIn, err
	}
	return UserIn, nil
}

func (U *UserInformationRepositoryImpl) Update(ctx context.Context, UserIn userinformation.UserInformation) (err error) {
	tx := U.Db.Model(userinformation.UserInformation{}).Where("UserID = ? ", UserIn.UserID).Updates(&UserIn)
	if err = tx.Error; err != nil {
		return
	}
	return
}

func (U *UserInformationRepositoryImpl) Delete(ctx context.Context, ID uint) (err error) {
	tx := U.Db.Model(userinformation.UserInformation{}).Where("ID = ? ", ID).Delete(&userinformation.UserInformation{})
	if err = tx.Error; err != nil {
		return
	}
	return
}

func NewUserInformationRepositoryImpl(db *gorm.DB) UserInformationRepository {
	return &UserInformationRepositoryImpl{
		Db: db,
	}
}
