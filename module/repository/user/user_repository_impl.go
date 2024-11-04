package user

import (
	"context"

	"github.com/arudji1211/ecommerce/module/model/user"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	Db *gorm.DB
}

func (U *UserRepositoryImpl) GetAll(ctx context.Context) (UsersOut []user.User, err error) {
	// panic("not implemented") // TODO: Implement
	tx := U.Db.Model(user.User{}).Find(&UsersOut)
	if err = tx.Error; err != nil {
		return
	}
	return
}

func (U *UserRepositoryImpl) GetUserByUsername(ctx context.Context, Username string) (UserOut user.User, err error) {
	tx := U.Db.Model(user.User{}).Where("Username = ?", Username).Find(&UserOut)
	if err = tx.Error; err != nil {
		return
	}
	return
}

func (U *UserRepositoryImpl) Create(ctx context.Context, UserIn user.User) (err error) {
	//panic("not implemented") // TODO: Implement
	tx := U.Db.Model(user.User{}).Create(&UserIn)
	if err = tx.Error; err != nil {
		return
	}
	return
}

func (U *UserRepositoryImpl) Update(ctx context.Context, UserIn user.User) (err error) {
	//panic("not implemented") // TODO: Implement
	tx := U.Db.Model(user.User{}).Where("ID = ? ", UserIn.ID).Updates(&UserIn)
	if err = tx.Error; err != nil {
		return
	}
	return
}

func (U *UserRepositoryImpl) Delete(ctx context.Context, ID uint) (err error) {
	tx := U.Db.Model(user.User{}).Where("username = ? ", ID).Delete(user.User{})
	if err = tx.Error; err != nil {
		return
	}
	return
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{
		Db: db,
	}
}
