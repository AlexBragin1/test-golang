package services

import (
	"context"
	"fmt"
	"time"

	"test/domain"
	"test/domain/groups"
	"test/dto"
	errors2 "test/errors"
)

type UsersRepo interface {
	Save(context.Context, domain.User) error
	FindByLogin(context.Context, domain.Login) (*domain.User, error)
	CountByLogin(context.Context, domain.Login) (int, error)
	Update(context.Context, domain.User) error
	FindByID(context.Context, domain.UUID) (*domain.User, error)
}

type TokenService interface {
	GenerateAuthToken(context.Context, *domain.User) (string, error)
	ReadFromToken(ctx context.Context, tokenString string, keys ...string) (map[string]string, error)
}

type AuthService struct {
	usersRepo    UsersRepo
	tokenService TokenService
}

func NewAuthService(usersRepo UsersRepo, gen TokenService) *AuthService {
	return &AuthService{
		usersRepo:    usersRepo,
		tokenService: gen,
	}
}

func (s *AuthService) Register(ctx context.Context, req dto.RegisterReq) (*dto.RegisterRes, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	count, err := s.usersRepo.CountByLogin(ctx, req.Login)
	if err != nil {
		return nil, err
	}

	if count != 0 {
		return nil, fmt.Errorf("login exists")
	}

	user, err := domain.NewUser(req.Login, req.Password, groups.User)
	if err != nil {
		return nil, fmt.Errorf("could not create new user: %w", err)
	}

	if err = s.usersRepo.Save(ctx, *user); err != nil {
		return nil, fmt.Errorf("could not save user: %w", err)
	}

	tokenString, err := s.tokenService.GenerateAuthToken(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("could not create jwt token: %w", err)
	}

	return &dto.RegisterRes{
		Token: tokenString,
	}, nil
}

func (s *AuthService) Login(ctx context.Context, req dto.LoginReq) (*dto.LoginRes, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	count, err := s.usersRepo.CountByLogin(ctx, req.Login)
	if err != nil {
		return nil, err
	} else if count == 0 {
		return nil, errors2.NewNotFoundError("user not found")
	}

	user, err := s.usersRepo.FindByLogin(ctx, req.Login)
	if err != nil {
		return nil, errors2.NewNotFoundError("user not found")
	}

	if user.Auth == true {
		return nil, errors2.NewAuthenticationError("your authentication")
	}

	if user.Password.CheckWithPlain(req.Password) {
		token, err := s.tokenService.GenerateAuthToken(ctx, user)
		if err != nil {
			return nil, fmt.Errorf("could not generate auth token: %w", err)
		}

		user.StartSessionAt = time.Now()
		user.Auth = true
		if err := s.usersRepo.Update(ctx, *user); err != nil {
			return nil, err
		}

		return &dto.LoginRes{Token: token}, nil
	} else {
		return nil, errors2.NewAuthenticationError("wrong password")
	}
}

func (s *AuthService) LoginOut(ctx context.Context, req dto.LoginOutReq) (*dto.LoginOutRes, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	user, err := s.usersRepo.FindByID(ctx, req.UserID)
	if err != nil {
		return nil, err
	}

	user.EndSessionAt = time.Now()
	user.Auth = false
	if err := s.usersRepo.Update(ctx, *user); err != nil {
		return nil, err
	}

	return &dto.LoginOutRes{}, nil
}
