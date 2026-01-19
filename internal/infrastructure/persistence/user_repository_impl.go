package persistence

import (
	"context"

	"github.com/google/uuid"
	"api-web-scrapping/internal/domain/entities"
	"api-web-scrapping/internal/domain/repositories"
)

// InMemoryUserRepository is an in-memory implementation for demonstration
// In production, replace this with a real database implementation
type InMemoryUserRepository struct {
	users map[uuid.UUID]*entities.User
}

func NewInMemoryUserRepository() repositories.UserRepository {
	return &InMemoryUserRepository{
		users: make(map[uuid.UUID]*entities.User),
	}
}

func (r *InMemoryUserRepository) FindByEmail(ctx context.Context, email string) (*entities.User, error) {
	for _, user := range r.users {
		if user.Email == email {
			return user, nil
		}
	}
	return nil, nil
}

func (r *InMemoryUserRepository) FindByID(ctx context.Context, id uuid.UUID) (*entities.User, error) {
	user, exists := r.users[id]
	if !exists {
		return nil, nil
	}
	return user, nil
}

func (r *InMemoryUserRepository) Create(ctx context.Context, user *entities.User) error {
	r.users[user.ID] = user
	return nil
}

func (r *InMemoryUserRepository) Update(ctx context.Context, user *entities.User) error {
	r.users[user.ID] = user
	return nil
}
