package v_knowledge

import (
	"errors"
	"strconv"
)

func ValidateID(id string) (int, error) {
	knowledgeID, err := strconv.Atoi(id)
	if err != nil || knowledgeID < 1 {
		return 0, errors.New("invalid Knowledge ID")
	}
	return knowledgeID, nil
}
