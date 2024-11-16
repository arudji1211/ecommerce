package user

import (
	"context"

	USmodel "github.com/arudji1211/ecommerce/module/model/user"
	USrepo "github.com/arudji1211/ecommerce/module/repository/user"
)

type UserServiceImpl struct {
	UserRepo USrepo.UserRepository
}

func (u *UserServiceImpl) GetAll(ctx context.Context) (userOut []USmodel.User, err error) {
	//panic("not implemented") // TODO: Implement
	userOut, err = u.UserRepo.GetAll(ctx)
	if err != nil {
		return
	}
	return
}

func (u *UserServiceImpl) GetByUsername(ctx context.Context, username string) (userOut USmodel.User, err error) {
	//panic("not implemented") // TODO: Implement
	userOut, err = u.UserRepo.GetUserByUsername(ctx, username)
	if err != nil {
		return
	}
	return
}

func (u *UserServiceImpl) Create(ctx context.Context, userIn USmodel.User) (userOut USmodel.User, err error) {
	//panic("not implemented") // TODO: Implement
	userOut, err = u.UserRepo.Create(ctx, userIn)
	if err != nil {
		return
	}
	return
}

func (u *UserServiceImpl) Update(ctx context.Context, userIn USmodel.User) (err error) {
	//panic("not implemented") // TODO: Implement
	err = u.UserRepo.Update(ctx, userIn)
	if err != nil {
		return
	}
	return
}

func (u *UserServiceImpl) Delete(ctx context.Context, ID uint) (err error) {
	//panic("not implemented") // TODO: Implement
	err = u.UserRepo.Delete(ctx, ID)
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
