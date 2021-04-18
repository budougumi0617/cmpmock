package _example

import (
	"context"
	"time"
)

type User struct {
	Name, Address string
	CreateAt      time.Time
}

type UserRepo interface {
	Save(context.Context, *User) error
}

type UserUsecase struct {
	Repo UserRepo
}

func (uu *UserUsecase) Register(ctx context.Context, name, address string) error {
	u := &User{
		Name:     name,
		Address:  address,
		CreateAt: time.Now(),
	}
	return uu.Repo.Save(ctx, u)
}
