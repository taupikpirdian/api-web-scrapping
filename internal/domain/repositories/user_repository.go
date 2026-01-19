package repositories

import (
	"context"

	"github.com/google/uuid"
	"api-web-scrapping/internal/domain/entities"
)

type UserRepository interface {
	FindByEmail(ctx context.Context, email string) (*entities.User, error)
	FindByID(ctx context.Context, id uuid.UUID) (*entities.User, error)
	Create(ctx context.Context, user *entities.User) error
	Update(ctx context.Context, user *entities.User) error
}
