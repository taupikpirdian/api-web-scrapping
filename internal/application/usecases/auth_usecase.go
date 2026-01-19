package usecases

import (
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"

	"api-web-scrapping/internal/application/dto"
	"api-web-scrapping/internal/domain/entities"
	"api-web-scrapping/internal/domain/repositories"
	"api-web-scrapping/pkg/auth"
)

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrUserNotFound       = errors.New("user not found")
)

type AuthUseCase interface {
	Login(ctx context.Context, req dto.LoginRequest) (*dto.LoginResponse, error)
	Register(ctx context.Context, req dto.RegisterRequest) (*dto.RegisterResponse, error)
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

func (uc *authUseCase) Register(ctx context.Context, req dto.RegisterRequest) (*dto.RegisterResponse, error) {
	// Check if user already exists
	existingUser, _ := uc.userRepo.FindByEmail(ctx, req.Email)
	if existingUser != nil {
		return nil, errors.New("user already exists")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Create new user
	user, err := entities.NewUser(req.Email, string(hashedPassword), req.FullName)
	if err != nil {
		return nil, err
	}

	// Save user
	if err := uc.userRepo.Create(ctx, user); err != nil {
		return nil, err
	}

	return &dto.RegisterResponse{
		ID:       user.ID.String(),
		Email:    user.Email,
		FullName: user.FullName,
	}, nil
}
