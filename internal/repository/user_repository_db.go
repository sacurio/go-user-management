package repository

import (
	"database/sql"
	"user-management/internal/config"
	"user-management/internal/domain"

	"github.com/sirupsen/logrus"
)

type UserRepositoryDB struct {
	db *sql.DB
}

func NewRepositoryDB(db *sql.DB) *UserRepositoryDB {
	config.Log.Info("UserRepositoryDB initialized")
	return &UserRepositoryDB{db: db}
}

func (r *UserRepositoryDB) Save(user domain.User) (domain.User, error) {
	query := `INSERT INTO users (name, email) VALUES (?, ?)`
	result, err := r.db.Exec(query, user.Name, user.Email)
	if err != nil {
		config.Log.WithFields(logrus.Fields{
			"error": err.Error(),
			"name":  user.Name,
			"email": user.Email,
		}).Error("Failed to save user")

		return domain.User{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		config.Log.WithField("error", err.Error()).Error("Failed to retrieve last insert ID")
		return domain.User{}, err
	}

	user.ID = int(id)
	config.Log.WithFields(logrus.Fields{
		"id":    user.ID,
		"name":  user.Name,
		"email": user.Email,
	}).Info("User saved successfully")

	return user, nil
}

func (r *UserRepositoryDB) FindAll() ([]domain.User, error) {
	query := `SELECT id, name, email FROM users`
	rows, err := r.db.Query(query)
	if err != nil {
		config.Log.WithField("error", err.Error()).Error("Failed to retrieve all user list")
		return nil, err
	}
	defer rows.Close()

	var users []domain.User
	for rows.Next() {
		var user domain.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			config.Log.WithFields(logrus.Fields{
				"id":    &user.ID,
				"name":  &user.Name,
				"email": &user.Email,
				"error": err.Error(),
			}).Error("Failed to read user record")
			return nil, err
		}
		users = append(users, user)
	}

	config.Log.Infof("User list retrieved successfully [%d]", len(users))

	return users, nil
}
