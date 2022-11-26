package service

import (
	"context"
	"fmt"
	"shmz_book/entity"
	"shmz_book/store"

	"golang.org/x/crypto/bcrypt"
)

type RegisterUser struct {
	DB   store.Execer
	Repo UserRegister
}

func (r *RegisterUser) RegisterUser(
	ctx context.Context, name, password, role string,
) (*entity.User, error) {

	pw, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("cannot hash password: %w", err)
	}
	user := &entity.User{
		Name:     name,
		Password: string(pw),
		Role:     role,
	}
	if err := r.Repo.RegisterUser(ctx, r.DB, user); err != nil {
		return nil, fmt.Errorf("failed to register: %w", err)
	}
	return user, nil
}
