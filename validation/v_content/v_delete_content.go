package v_content

import (
	"errors"
	rcontent "lms/repository/r_content"
	"strconv"
)

func ValidateDeleteID(id string) (int, error) {
	contentID, err := strconv.Atoi(id)
	if err != nil || contentID < 1 {
		return 0, errors.New("invalid content ID")
	}

	// mengonversi contentID menjadi string
	contentIDStr := strconv.Itoa(contentID)

	// tambahkan validasi untuk mengecek apakah content dengan ID tersebut masih ada
	_, err = rcontent.MGetContentByID(contentIDStr)
	if err != nil {
		return 0, errors.New("content not found")
	}

	return contentID, nil
}
