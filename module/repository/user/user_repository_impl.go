package user

import (
	"context"

	MDuser "github.com/arudji1211/ecommerce/module/model/user"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepositoryImpl struct {
	Db *gorm.DB
}

func (U *UserRepositoryImpl) GetAll(ctx context.Context, db *gorm.DB) (UsersOut []MDuser.User, err error) {
	// panic("not implemented") // TODO: Implement
	tx := U.Db.Model(MDuser.User{}).Find(&UsersOut)
	if err = tx.Error; err != nil {
		return
	}
	return
}

func (U *UserRepositoryImpl) GetUserByUsername(ctx context.Context, db *gorm.DB, Username string) (UserOut MDuser.User, err error) {
	tx := U.Db.Model(MDuser.User{}).Where("Username = ?", Username).Find(&UserOut)
	if err = tx.Error; err != nil {
		return
	}
	return
}

func (U *UserRepositoryImpl) Create(ctx context.Context, db *gorm.DB, UserIn MDuser.User) (MDuser.User, error) {
	//panic("not implemented") // TODO: Implement
	tx := U.Db.Model(MDuser.User{}).Clauses(clause.Returning{}).Create(&UserIn)
	if err := tx.Error; err != nil {
		return UserIn, err
	}
	return UserIn, nil
}

func (U *UserRepositoryImpl) Update(ctx context.Context, db *gorm.DB, UserIn MDuser.User) (err error) {
	//panic("not implemented") // TODO: Implement
	tx := U.Db.Model(MDuser.User{}).Where("ID = ? ", UserIn.ID).Updates(&UserIn)
	if err = tx.Error; err != nil {
		return
	}
	return
}

func (U *UserRepositoryImpl) Delete(ctx context.Context, db *gorm.DB, ID uint) (err error) {
	tx := U.Db.Model(MDuser.User{}).Where("username = ? ", ID).Delete(MDuser.User{})
	if err = tx.Error; err != nil {
		return
	}
	return
}

func (U *UserRepositoryImpl) IsTransaction(Db *gorm.DB) {
	if Db != nil {
		U.Db = Db
	}
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{
		Db: db,
	}
}
