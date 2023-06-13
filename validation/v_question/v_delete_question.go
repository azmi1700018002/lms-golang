package v_question

import (
	"errors"
	"strconv"
)

func ValidateID(id string) (int, error) {
	IDQuestion, err := strconv.Atoi(id)
	if err != nil || IDQuestion < 1 {
		return 0, errors.New("invalid Question ID")
	}
	return IDQuestion, nil
}
