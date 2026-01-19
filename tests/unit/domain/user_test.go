package domain

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"api-web-scrapping/internal/domain/entities"
)

func TestNewUser_Success(t *testing.T) {
	// Execute
	user, err := entities.NewUser("test@example.com", "hashedpassword", "Test User")

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.NotEqual(t, uuid.Nil, user.ID)
	assert.Equal(t, "test@example.com", user.Email)
	assert.Equal(t, "hashedpassword", user.Password)
	assert.Equal(t, "Test User", user.FullName)
	assert.False(t, user.CreatedAt.IsZero())
	assert.False(t, user.UpdatedAt.IsZero())
}

func TestNewUser_EmptyEmail(t *testing.T) {
	// Execute
	user, err := entities.NewUser("", "hashedpassword", "Test User")

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, "", user.Email)
}

func TestNewUser_EmptyPassword(t *testing.T) {
	// Execute
	user, err := entities.NewUser("test@example.com", "", "Test User")

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, "", user.Password)
}

func TestNewUser_EmptyFullName(t *testing.T) {
	// Execute
	user, err := entities.NewUser("test@example.com", "hashedpassword", "")

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, "", user.FullName)
}

func TestUser_IDGeneration(t *testing.T) {
	// Execute
	user1, _ := entities.NewUser("user1@example.com", "pass", "User 1")
	user2, _ := entities.NewUser("user2@example.com", "pass", "User 2")

	// Assert - Each user should have a unique ID
	assert.NotEqual(t, user1.ID, user2.ID)
}

func TestUser_Timestamps(t *testing.T) {
	// Execute
	user, err := entities.NewUser("test@example.com", "hashedpassword", "Test User")

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.False(t, user.CreatedAt.IsZero(), "CreatedAt should be set")
	assert.False(t, user.UpdatedAt.IsZero(), "UpdatedAt should be set")
}

// Table-driven test
func TestNewUser_TableDriven(t *testing.T) {
	tests := []struct {
		name      string
		email     string
		password  string
		fullName  string
		wantError bool
	}{
		{
			name:      "valid user",
			email:     "test@example.com",
			password:  "hashedpassword",
			fullName:  "Test User",
			wantError: false,
		},
		{
			name:      "empty email",
			email:     "",
			password:  "hashedpassword",
			fullName:  "Test User",
			wantError: false,
		},
		{
			name:      "empty password",
			email:     "test@example.com",
			password:  "",
			fullName:  "Test User",
			wantError: false,
		},
		{
			name:      "empty full name",
			email:     "test@example.com",
			password:  "hashedpassword",
			fullName:  "",
			wantError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user, err := entities.NewUser(tt.email, tt.password, tt.fullName)

			if tt.wantError {
				assert.Error(t, err)
				assert.Nil(t, user)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, user)
				assert.Equal(t, tt.email, user.Email)
				assert.Equal(t, tt.password, user.Password)
				assert.Equal(t, tt.fullName, user.FullName)
			}
		})
	}
}

// Benchmark test
func BenchmarkNewUser(b *testing.B) {
	for i := 0; i < b.N; i++ {
		entities.NewUser("test@example.com", "hashedpassword", "Test User")
	}
}
