package usecase

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"

	"api-web-scrapping/internal/application/dto"
	"api-web-scrapping/internal/application/usecases"
	"api-web-scrapping/internal/domain/entities"
	"api-web-scrapping/pkg/auth"
)

// Mock implementations
type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) FindByEmail(ctx context.Context, email string) (*entities.User, error) {
	args := m.Called(ctx, email)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.User), args.Error(1)
}

func (m *MockUserRepository) FindByID(ctx context.Context, id uuid.UUID) (*entities.User, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.User), args.Error(1)
}

func (m *MockUserRepository) Create(ctx context.Context, user *entities.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

func (m *MockUserRepository) Update(ctx context.Context, user *entities.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

type MockJWTManager struct {
	mock.Mock
}

func (m *MockJWTManager) GenerateToken(userID, email string) (string, error) {
	args := m.Called(userID, email)
	return args.String(0), args.Error(1)
}

func (m *MockJWTManager) ValidateToken(tokenString string) (*auth.Claims, error) {
	args := m.Called(tokenString)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*auth.Claims), args.Error(1)
}

func TestAuthUseCase_Login_Success(t *testing.T) {
	// Setup
	mockRepo := new(MockUserRepository)
	mockJWT := new(MockJWTManager)

	ctx := context.Background()
	userID := uuid.New()
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)

	testUser := &entities.User{
		ID:       userID,
		Email:    "test@example.com",
		Password: string(hashedPassword),
		FullName: "Test User",
	}

	mockRepo.On("FindByEmail", ctx, "test@example.com").Return(testUser, nil)
	mockJWT.On("GenerateToken", userID.String(), "test@example.com").Return("valid-jwt-token", nil)

	useCase := usecases.NewAuthUseCase(mockRepo, mockJWT)

	// Execute
	req := dto.LoginRequest{
		Email:    "test@example.com",
		Password: "password123",
	}

	resp, err := useCase.Login(ctx, req)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "valid-jwt-token", resp.Token)
	assert.Equal(t, userID.String(), resp.User.ID)
	assert.Equal(t, "test@example.com", resp.User.Email)
	assert.Equal(t, "Test User", resp.User.FullName)

	mockRepo.AssertExpectations(t)
	mockJWT.AssertExpectations(t)
}

func TestAuthUseCase_Login_InvalidCredentials(t *testing.T) {
	// Setup
	mockRepo := new(MockUserRepository)
	mockJWT := new(MockJWTManager)

	ctx := context.Background()

	mockRepo.On("FindByEmail", ctx, "test@example.com").Return(nil, nil)

	useCase := usecases.NewAuthUseCase(mockRepo, mockJWT)

	// Execute
	req := dto.LoginRequest{
		Email:    "test@example.com",
		Password: "wrongpassword",
	}

	resp, err := useCase.Login(ctx, req)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Equal(t, usecases.ErrInvalidCredentials, err)

	mockRepo.AssertExpectations(t)
}

func TestAuthUseCase_Login_WrongPassword(t *testing.T) {
	// Setup
	mockRepo := new(MockUserRepository)
	mockJWT := new(MockJWTManager)

	ctx := context.Background()
	userID := uuid.New()
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("correctpassword"), bcrypt.DefaultCost)

	testUser := &entities.User{
		ID:       userID,
		Email:    "test@example.com",
		Password: string(hashedPassword),
		FullName: "Test User",
	}

	mockRepo.On("FindByEmail", ctx, "test@example.com").Return(testUser, nil)

	useCase := usecases.NewAuthUseCase(mockRepo, mockJWT)

	// Execute
	req := dto.LoginRequest{
		Email:    "test@example.com",
		Password: "wrongpassword",
	}

	resp, err := useCase.Login(ctx, req)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Equal(t, usecases.ErrInvalidCredentials, err)

	mockRepo.AssertExpectations(t)
}

func TestAuthUseCase_Register_Success(t *testing.T) {
	// Setup
	mockRepo := new(MockUserRepository)
	mockJWT := new(MockJWTManager)

	ctx := context.Background()

	mockRepo.On("FindByEmail", ctx, "newuser@example.com").Return(nil, nil)
	mockRepo.On("Create", ctx, mock.AnythingOfType("*entities.User")).Return(nil)

	useCase := usecases.NewAuthUseCase(mockRepo, mockJWT)

	// Execute
	req := dto.RegisterRequest{
		Email:    "newuser@example.com",
		Password: "password123",
		FullName: "New User",
	}

	resp, err := useCase.Register(ctx, req)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp.ID)
	assert.Equal(t, "newuser@example.com", resp.Email)
	assert.Equal(t, "New User", resp.FullName)

	mockRepo.AssertExpectations(t)
}

func TestAuthUseCase_Register_UserAlreadyExists(t *testing.T) {
	// Setup
	mockRepo := new(MockUserRepository)
	mockJWT := new(MockJWTManager)

	ctx := context.Background()
	userID := uuid.New()

	existingUser := &entities.User{
		ID:       userID,
		Email:    "existing@example.com",
		Password: "hashedpassword",
		FullName: "Existing User",
	}

	mockRepo.On("FindByEmail", ctx, "existing@example.com").Return(existingUser, nil)

	useCase := usecases.NewAuthUseCase(mockRepo, mockJWT)

	// Execute
	req := dto.RegisterRequest{
		Email:    "existing@example.com",
		Password: "password123",
		FullName: "Existing User",
	}

	resp, err := useCase.Register(ctx, req)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "user already exists")

	mockRepo.AssertExpectations(t)
}

// Benchmark test
func BenchmarkAuthUseCase_Login(b *testing.B) {
	mockRepo := new(MockUserRepository)
	mockJWT := new(MockJWTManager)

	ctx := context.Background()
	userID := uuid.New()
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)

	testUser := &entities.User{
		ID:       userID,
		Email:    "test@example.com",
		Password: string(hashedPassword),
		FullName: "Test User",
	}

	mockRepo.On("FindByEmail", ctx, "test@example.com").Return(testUser, nil)
	mockJWT.On("GenerateToken", userID.String(), "test@example.com").Return("valid-jwt-token", nil)

	useCase := usecases.NewAuthUseCase(mockRepo, mockJWT)

	req := dto.LoginRequest{
		Email:    "test@example.com",
		Password: "password123",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		useCase.Login(ctx, req)
	}
}
