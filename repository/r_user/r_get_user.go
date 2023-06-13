package ruser

import (
	"lms/config/db"
	"lms/model"
	"sort"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func MGetUsersOrm(limit, offset int) (users []model.User, count int64, err error) {
	if err = db.Server().Limit(limit).Offset(offset).Find(&users).Error; err != nil {
		return nil, 0, err
	}
	db.Server().Model(&model.User{}).Count(&count)
	return users, count, nil
}

func MGetUsers() (users []model.User, count int64, err error) {
	stmt := `
	SELECT u.id_user, COALESCE(u.id_role, 0), u.created_at, u.updated_at, u.deleted_at, u.username, u.email, u.password, u.last_login, COALESCE(u.profile_picture, ''), COALESCE(r.id_role, 0), COALESCE(r.rolename, ''), COALESCE(c.idcourse, 0), COALESCE(c.id_knowledge, 0), COALESCE(c.course_name, ''), COALESCE(c.course_desc, '')
	FROM users u
	LEFT JOIN user_roles ur ON u.id_user = ur.user_id_user
	LEFT JOIN roles r ON ur.role_id_role = r.id_role
	LEFT JOIN courses c ON u.id_user = c.id_user
	ORDER BY u.created_at ASC 
        `
	rows, err := db.Server().Raw(stmt).Rows()
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	userMap := make(map[uuid.UUID]*model.User)

	for rows.Next() {
		var (
			userIDStr       string
			roleID          int
			roleName        string
			userCreated     time.Time
			userUpdated     time.Time
			userDeleted     gorm.DeletedAt
			username        string
			email           string
			password        string
			lastLogin       *time.Time
			profile_picture string
			courseID        int
			courseName      string
			courseDesc      string
			idKnowledge     int
		)
		if err := rows.Scan(&userIDStr, &roleID, &userCreated, &userUpdated, &userDeleted, &username, &email, &password, &lastLogin, &profile_picture, &roleID, &roleName, &courseID, &idKnowledge, &courseName, &courseDesc); err != nil {
			return nil, 0, err
		}

		userID, err := uuid.Parse(userIDStr)
		if err != nil {
			return nil, 0, err
		}

		if user, ok := userMap[userID]; ok {
			if courseID != 0 {
				user.Course = append(user.Course, model.Course{IDcourse: courseID, IdKnowledge: idKnowledge, IDUser: userID, CourseName: courseName, CourseDesc: courseDesc})
			}
		} else {
			newUser := model.User{
				IDUser:         userID,
				IDRole:         roleID,
				CreatedAt:      userCreated,
				UpdatedAt:      userUpdated,
				DeletedAt:      userDeleted,
				Username:       username,
				Email:          email,
				Password:       password,
				LastLogin:      lastLogin,
				ProfilePicture: profile_picture,
				Roles:          []model.Role{{IDRole: roleID, Rolename: roleName}},
				Course:         []model.Course{},
			}
			if courseID != 0 {
				newUser.Course = append(newUser.Course, model.Course{IDcourse: courseID, IDUser: userID, CourseName: courseName, CourseDesc: courseDesc, IdKnowledge: idKnowledge})
			}
			userMap[userID] = &newUser
		}
	}

	// for _, user := range userMap {
	// 	users = append(users, *user)
	// }

	// Dengan ASC
	keys := make([]uuid.UUID, len(userMap))
	i := 0
	for key := range userMap {
		keys[i] = key
		i++
	}
	sort.Slice(keys, func(i, j int) bool {
		return userMap[keys[i]].CreatedAt.Before(userMap[keys[j]].CreatedAt)
	})

	for _, key := range keys {
		user := userMap[key]
		users = append(users, *user)
	}

	if err := db.Server().Model(&model.User{}).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	return users, count, nil
}

func MGetUserByIDOrm(userID string) (user model.User, err error) {
	if err = db.Server().First(&user, userID).Error; err != nil {
		return model.User{}, err
	}
	return user, nil
}

func MGetUserByID(userID uuid.UUID) (model.User, error) {
	var user model.User
	stmt := `
		SELECT id_user, id_role, created_at, updated_at, deleted_at, username, email, password, last_login, profile_picture
		FROM users
		WHERE id_user = $1
		LIMIT 1;
	`
	row := db.Server().Raw(stmt, userID).Row()

	if err := row.Scan(&user.IDUser, &user.IDRole, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt, &user.Username, &user.Email, &user.Password, &user.LastLogin, &user.ProfilePicture); err != nil {
		return model.User{}, err
	}

	return user, nil
}
