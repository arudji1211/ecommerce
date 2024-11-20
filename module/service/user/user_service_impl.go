package user

import (
	"context"

	USmodel "github.com/arudji1211/ecommerce/module/model/user"
	USrepo "github.com/arudji1211/ecommerce/module/repository/user"
	"gorm.io/gorm"
)

type UserServiceImpl struct {
	UserRepo USrepo.UserRepository
}

func (u *UserServiceImpl) GetAll(ctx context.Context, db *gorm.DB) (userOut []USmodel.User, err error) {
	//panic("not implemented") // TODO: Implement
	userOut, err = u.UserRepo.GetAll(ctx, db)
	if err != nil {
		return
	}
	return
}

func (u *UserServiceImpl) GetByUsername(ctx context.Context, db *gorm.DB, username string) (userOut USmodel.User, err error) {
	//panic("not implemented") // TODO: Implement
	userOut, err = u.UserRepo.GetUserByUsername(ctx, db, username)
	if err != nil {
		return
	}
	return
}

func (u *UserServiceImpl) Create(ctx context.Context, db *gorm.DB, userIn USmodel.User) (userOut USmodel.User, err error) {
	//panic("not implemented") // TODO: Implement
	userOut, err = u.UserRepo.Create(ctx, db, userIn)
	if err != nil {
		return
	}
	return
}

func (u *UserServiceImpl) Update(ctx context.Context, db *gorm.DB, userIn USmodel.User) (err error) {
	//panic("not implemented") // TODO: Implement
	err = u.UserRepo.Update(ctx, db, userIn)
	if err != nil {
		return
	}
	return
}

func (u *UserServiceImpl) Delete(ctx context.Context, db *gorm.DB, ID uint) (err error) {
	//panic("not implemented") // TODO: Implement
	err = u.UserRepo.Delete(ctx, db, ID)
	if err != nil {
		return
	}
	return
}

func NewUserService(repo USrepo.UserRepository) UserService {
	return &UserServiceImpl{
		UserRepo: repo,
	}
}
