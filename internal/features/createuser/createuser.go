package createuser

import (
	"context"
	"fmt"
	"github.com/fenderdigital/bv-medium-users-service/internal"
	"github.com/fenderdigital/bv-medium-users-service/internal/pkg"
)

type Storage interface {
	Create(ctx context.Context, id, name, email string) (*internal.User, error)
}
type Messaging interface {
	Publish(msg string) error
}

type CreateUser struct {
	store Storage
	msg   Messaging
}

func (s *CreateUser) Create(ctx context.Context, name, email string) error {
	id, err := pkg.GenerateUUID()
	if err != nil {
		return fmt.Errorf("pkg.GenerateUUID failed: %w", err)
	}

	if _, err := s.store.Create(ctx, id, name, email); err != nil {
		return fmt.Errorf("store.Create failed: %w", err)
	}

	if err := s.msg.Publish(fmt.Sprintf("created user, id: %s", id)); err != nil {
		return fmt.Errorf("msg.Publish failed: %w", err)
	}

	return nil
}
