package s_user

import (
	ruser "lms/repository/r_user"

	"github.com/google/uuid"
)

type UserDeleteService interface {
	DeleteUserByID(id uuid.UUID) error
}

type userDeleteService struct {
	userDeleteRepo ruser.UserDeleteRepository
}

func NewUserDeleteService(userDeleteRepo ruser.UserDeleteRepository) UserDeleteService {
	return &userDeleteService{userDeleteRepo}
}

func (s *userDeleteService) DeleteUserByID(id uuid.UUID) error {
	err := s.userDeleteRepo.DeleteUserByID(id)
	if err != nil {
		return err
	}
	return nil
}
