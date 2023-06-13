package rcourse

import (
	"lms/config/db"
	"lms/model"
)

func MDeleteCourseByID(idcourse int) (err error) {
	// Menghapus data contents yang terkait dengan sections yang terkait dengan course
	if err = db.Server().Exec("DELETE FROM contents WHERE idsection IN (SELECT idsection FROM sections WHERE idcourse = ?)", idcourse).Error; err != nil {
		return err
	}

	// Menghapus data sections yang terkait dengan course
	if err = db.Server().Exec("DELETE FROM sections WHERE idcourse = ?", idcourse).Error; err != nil {
		return err
	}

	// Menghapus data course dari tabel Course
	if err = db.Server().Unscoped().Where("idcourse = ?", idcourse).Delete(&model.Course{}).Error; err != nil {
		return err
	}

	return nil
}
