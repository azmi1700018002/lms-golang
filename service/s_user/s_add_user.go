package s_user

import (
	"lms/model"
	vuser "lms/validation/v_user"
)

type UserRepository interface {
	RegisterUser(user *model.User) error
}

type UserService struct {
	userRepository UserRepository
}

func (s *UserService) RegisterUser(user *model.User) error {
	if err := vuser.ValidateUser(user); err != nil {
		return err
	}

	return s.userRepository.RegisterUser(user)
}
