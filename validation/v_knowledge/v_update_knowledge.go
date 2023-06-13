package v_knowledge

import (
	"errors"
	"lms/model"
)

func ValidateUpdateKnowledge(knowledge model.Knowledge) error {
	if knowledge.IDcategory == 0 {
		return errors.New("IDcategory is required")
	}
	if knowledge.KnowledgeName == "" {
		return errors.New("KnowledgeName is required")
	}
	if knowledge.Description == "" {
		return errors.New("description is required")
	}
	return nil
}
