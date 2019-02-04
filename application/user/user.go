package user

import (
	"context"

	"github.com/inari111/layered-architecture-study/domain"

	"github.com/inari111/layered-architecture-study/domain/user"
)

type Application interface {
	Register(ctx context.Context) error

	UpdateProfile(ctx context.Context, name string, age int, description string) error
}

func NewApplication(
	userRepo user.Repository,
	profileRepo user.ProfileRepository,
	now domain.CurrentTimeFunc,
) Application {
	return &userAppImpl{
		userRepo:    userRepo,
		profileRepo: profileRepo,
		now:         now,
	}
}

type userAppImpl struct {
	userRepo    user.Repository
	profileRepo user.ProfileRepository
	now         domain.CurrentTimeFunc
}

func (app *userAppImpl) Register(ctx context.Context) error {
	// TODO: ID生成
	u := user.User{
		ID:        "user1",
		CreatedAt: app.now(),
		UpdatedAt: app.now(),
	}
	if err := app.userRepo.Put(ctx, &u); err != nil {
		return err
	}
	return nil
}

func (app *userAppImpl) UpdateProfile(ctx context.Context, name string, age int, description string) error {
	p := user.Profile{
		Name:        name,
		Age:         age,
		Description: description,
	}
	if err := app.profileRepo.Put(ctx, &p); err != nil {
		return err
	}
	return nil
}
