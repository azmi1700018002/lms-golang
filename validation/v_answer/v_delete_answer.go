package v_answer

import (
	"errors"
	"strconv"
)

func ValidateID(id string) (int, error) {
	IDAnswer, err := strconv.Atoi(id)
	if err != nil || IDAnswer < 1 {
		return 0, errors.New("invalid Answer ID")
	}
	return IDAnswer, nil
}
