package repository

import "user-management/internal/domain"

type UserRepository struct {
	users  []domain.User
	nextID int
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		users:  make([]domain.User, 0),
		nextID: 1,
	}
}

func (r *UserRepository) Save(user domain.User) domain.User {
	user.ID = r.nextID
	r.nextID++
	r.users = append(r.users, user)

	return user
}

func (r *UserRepository) FindAll() []domain.User {
	return r.users
}
