package s_user

import (
	"lms/model"
	vuser "lms/validation/v_user"
)

type UserUpdateRepository interface {
	UpdateUser(user *model.User) error
}

type UserUpdateService struct {
	userUpdateRepository UserUpdateRepository
}

func (s *UserUpdateService) UpdateUser(user *model.User) error {
	if err := vuser.ValidateUser(user); err != nil {
		return err
	}

	return s.userUpdateRepository.UpdateUser(user)
}
