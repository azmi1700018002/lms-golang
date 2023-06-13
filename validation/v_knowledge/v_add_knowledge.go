package v_knowledge

import (
	"errors"
	"lms/model"
)

func ValidateKnowledge(knowledge model.Knowledge) error {
	if knowledge.IDcategory == 0 {
		return errors.New("IDcategory is required")
	}
	if knowledge.KnowledgeName == "" {
		return errors.New("KnowledgeName is required")
	}
	// if knowledge.Status == 0 {
	// 	return errors.New("status is required")
	// }
	return nil
}
