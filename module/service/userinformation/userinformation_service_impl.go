package userinformation

import (
	"context"

	MDuserinformation "github.com/arudji1211/ecommerce/module/model/userinformation"
	UIRepo "github.com/arudji1211/ecommerce/module/repository/userinformation"
)

type UserInformationServiceImpl struct {
	UiRepo UIRepo.UserInformationRepository
}

func (p *UserInformationServiceImpl) GetAll(ctx context.Context) (UsersOut []MDuserinformation.UserInformation, err error) {
	//panic("not implemented") // TODO: Implement
	UsersOut, err = p.UiRepo.GetAll(ctx)
	if err != nil {
		return
	}
	return
}

func (p *UserInformationServiceImpl) GetUserInformationByUserID(ctx context.Context, UserID string) (UserOut MDuserinformation.UserInformation, err error) {
	//panic("not implemented") // TODO: Implement
	UserOut, err = p.UiRepo.GetUserInformationByUserID(ctx, UserID)
	if err != nil {
		return
	}
	return
}

func (p *UserInformationServiceImpl) Create(ctx context.Context, UserIn MDuserinformation.UserInformation) (MDuserinformation.UserInformation, error) {
	//panic("not implemented") // TODO: Implement
	UserIn, err := p.UiRepo.Create(ctx, UserIn)
	if err != nil {
		return UserIn, err
	}
	return UserIn, nil
}

func (p *UserInformationServiceImpl) Update(ctx context.Context, UserIn MDuserinformation.UserInformation) (err error) {
	//panic("not implemented") // TODO: Implement
	err = p.UiRepo.Update(ctx, UserIn)
	if err != nil {
		return
	}
	return
}

func (p *UserInformationServiceImpl) Delete(ctx context.Context, ID uint) (err error) {
	//panic("not implemented") // TODO: Implement
	err = p.UiRepo.Delete(ctx, ID)
	if err != nil {
		return
	}
	return
}

func NewUserInformationService(repo UIRepo.UserInformationRepository) UserInformationService {
	return &UserInformationServiceImpl{
		UiRepo: repo,
	}
}
