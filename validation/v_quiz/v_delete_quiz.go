package v_quiz

import (
	"errors"
	"strconv"
)

func ValidateID(id string) (int, error) {
	IDQuiz, err := strconv.Atoi(id)
	if err != nil || IDQuiz < 1 {
		return 0, errors.New("invalid Quiz ID")
	}
	return IDQuiz, nil
}
