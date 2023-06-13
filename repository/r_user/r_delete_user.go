package ruser

import (
	"lms/config/db"
	"lms/model"

	"github.com/google/uuid"
)

type UserDeleteRepository interface {
	DeleteUserByID(id uuid.UUID) error
}

type userDeleteRepository struct{}

func NewUserDeleteRepository() UserDeleteRepository {
	return &userDeleteRepository{}
}

func (*userDeleteRepository) DeleteUserByID(id uuid.UUID) (err error) {
	// Menghapus foreign key constraint pada user_roles jika sudah ada
	var exists bool
	if err = db.Server().Raw("SELECT EXISTS (SELECT 1 FROM information_schema.table_constraints WHERE constraint_name = 'fk_user_roles_users' AND table_name = 'user_roles' AND constraint_type = 'FOREIGN KEY')").Scan(&exists).Error; err != nil {
		return err
	}

	if exists {
		if err = db.Server().Exec("ALTER TABLE user_roles DROP CONSTRAINT fk_user_roles_users").Error; err != nil {
			return err
		}
	}

	// Menambahkan foreign key constraint pada user_roles jika belum ada
	if exists {
		if err = db.Server().Exec("ALTER TABLE user_roles ADD CONSTRAINT fk_user_roles_users FOREIGN KEY (user_id_user) REFERENCES users(id_user) ON DELETE CASCADE").Error; err != nil {
			return err
		}
	}

	// Menghapus data user dari tabel courses yang terkait
	if err = db.Server().Exec("DELETE FROM courses WHERE id_user = ?", id).Error; err != nil {
		return err
	}

	// Menghapus data answers yang terkait dengan questions
	if err = db.Server().Exec("DELETE FROM answers WHERE id_question IN (SELECT id_question FROM questions WHERE id_quiz IN (SELECT id_quiz FROM quizzes WHERE id_user = ?))", id).Error; err != nil {
		return err
	}

	// Menghapus data questions yang terkait dengan user
	if err = db.Server().Exec("DELETE FROM questions WHERE id_quiz IN (SELECT id_quiz FROM quizzes WHERE id_user = ?)", id).Error; err != nil {
		return err
	}

	// Menghapus data user dari tabel quizzes
	if err = db.Server().Exec("DELETE FROM quizzes WHERE id_user = ?", id).Error; err != nil {
		return err
	}

	// Menghapus data user dari tabel user_roles
	if err = db.Server().Exec("DELETE FROM user_roles WHERE user_id_user = ?", id).Error; err != nil {
		return err
	}

	// Menghapus data user dari tabel User
	if err = db.Server().Unscoped().Where("id_user = ?", id).Delete(&model.User{}).Error; err != nil {
		return err
	}

	return nil
}
