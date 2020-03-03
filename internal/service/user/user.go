package skeleton

import (
	"context"
	"fmt"

	userEntity "go-tutorial-2020/internal/entity/user"
	"go-tutorial-2020/pkg/errors"
)

// UserData ...
type UserData interface {
	GetAllUsers(ctx context.Context) ([]userEntity.User, error)
	InsertAllUsers(ctx context.Context, user userEntity.User) error
	UpdateAllUsers(ctx context.Context, user userEntity.User) error
	DeleteAllUsers(ctx context.Context, user userEntity.User) error
}

// Service ...
type Service struct {
	userData UserData
}

// New ...
func New(userData UserData) Service {
	return Service{
		userData: userData,
	}
}

// GetAllUsers ...
func (s Service) GetAllUsers(ctx context.Context) ([]userEntity.User, error) {
	// Panggil method GetAllUsers di data layer user
	users, err := s.userData.GetAllUsers(ctx)
	// Error handling
	if err != nil {
		return users, errors.Wrap(err, "[SERVICE][GetAllUsers]")
	}
	// Return users array
	return users, err
}

//insert
func (s Service) InsertAllUsers(ctx context.Context, user userEntity.User) error {
	err := s.userData.InsertAllUsers(ctx, user)
	fmt.Println(user)
	return err

}

//update
func (s Service) UpdateAllUsers(ctx context.Context, user userEntity.User) error {
	err := s.userData.UpdateAllUsers(ctx, user)
	fmt.Println(user)
	return err
}

//delete
func (s Service) DeleteAllUsers(ctx context.Context, user userEntity.User) error {
	err := s.userData.DeleteAllUsers(ctx, user)
	fmt.Println(user)
	return err
}
