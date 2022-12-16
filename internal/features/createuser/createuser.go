package createuser

import (
	"fmt"
	"github.com/branebrvl/medium-users-service/internal"
)

type Storage interface {
	Create(id, name, email string) (*internal.User, error)
}
type Messaging interface {
	Publish(msg string) error
}

type CreateUser struct {
	store Storage
	msg   Messaging
}

func (s *CreateUser) Create(id, name, email string) error {
	if _, err := s.store.Create(id, name, email); err != nil {
		return fmt.Errorf("store.Create failed: %w", err)
	}

	if err := s.msg.Publish(fmt.Sprintf("created user, id: %s", id)); err != nil {
		return fmt.Errorf("msg.Publish failed: %w", err)
	}

	return nil
}
