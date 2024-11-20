package userinformation

import (
	"context"

	MDuserinformation "github.com/arudji1211/ecommerce/module/model/userinformation"
	UIRepo "github.com/arudji1211/ecommerce/module/repository/userinformation"
	"gorm.io/gorm"
)

type UserInformationServiceImpl struct {
	UiRepo UIRepo.UserInformationRepository
}

func (p *UserInformationServiceImpl) GetAll(ctx context.Context, db *gorm.DB) (UsersOut []MDuserinformation.UserInformation, err error) {
	//panic("not implemented") // TODO: Implement
	UsersOut, err = p.UiRepo.GetAll(ctx, db)
	if err != nil {
		return
	}
	return
}

func (p *UserInformationServiceImpl) GetUserInformationByUserID(ctx context.Context, db *gorm.DB, UserID string) (UserOut MDuserinformation.UserInformation, err error) {
	//panic("not implemented") // TODO: Implement
	UserOut, err = p.UiRepo.GetUserInformationByUserID(ctx, db, UserID)
	if err != nil {
		return
	}
	return
}

func (p *UserInformationServiceImpl) Create(ctx context.Context, db *gorm.DB, UserIn MDuserinformation.UserInformation) (MDuserinformation.UserInformation, error) {
	//panic("not implemented") // TODO: Implement
	UserIn, err := p.UiRepo.Create(ctx, db, UserIn)
	if err != nil {
		return UserIn, err
	}
	return UserIn, nil
}

func (p *UserInformationServiceImpl) Update(ctx context.Context, db *gorm.DB, UserIn MDuserinformation.UserInformation) (err error) {
	//panic("not implemented") // TODO: Implement
	err = p.UiRepo.Update(ctx, db, UserIn)
	if err != nil {
		return
	}
	return
}

func (p *UserInformationServiceImpl) Delete(ctx context.Context, db *gorm.DB, ID uint) (err error) {
	//panic("not implemented") // TODO: Implement
	err = p.UiRepo.Delete(ctx, db, ID)
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
