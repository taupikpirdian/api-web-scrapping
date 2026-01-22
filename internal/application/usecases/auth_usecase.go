package usecases

import (
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"

	"api-web-scrapping/internal/application/dto"
	"api-web-scrapping/internal/domain/repositories"
	"api-web-scrapping/pkg/auth"
)

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrUserNotFound       = errors.New("user not found")
)

type AuthUseCase interface {
	Login(ctx context.Context, req dto.LoginRequest) (*dto.LoginResponse, error)
}

type authUseCase struct {
	userRepo   repositories.UserRepository
	jwtManager auth.JWTManager
}

func NewAuthUseCase(userRepo repositories.UserRepository, jwtManager auth.JWTManager) AuthUseCase {
	return &authUseCase{
		userRepo:   userRepo,
		jwtManager: jwtManager,
	}
}

func (uc *authUseCase) Login(ctx context.Context, req dto.LoginRequest) (*dto.LoginResponse, error) {
	// Find user by email
	user, err := uc.userRepo.FindByEmail(ctx, req.Email)
	if err != nil {
		return nil, ErrInvalidCredentials
	}

	if user == nil {
		return nil, ErrInvalidCredentials
	}

	// Verify password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, ErrInvalidCredentials
	}

	// Generate JWT token
	token, err := uc.jwtManager.GenerateToken(user.ID.String(), user.Email)
	if err != nil {
		return nil, err
	}

	return &dto.LoginResponse{
		Token: token,
		User: dto.UserResponse{
			ID:       user.ID.String(),
			Email:    user.Email,
			FullName: user.FullName,
		},
	}, nil
}
