package service

import (
	"user-management/internal/domain"
	"user-management/internal/repository"
)

type UserService struct {
	repo *repository.UserRepositoryDB
}

func NewUserService(repo *repository.UserRepositoryDB) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(name, email string) (domain.User, error) {
	user := domain.User{Name: name, Email: email}
	return s.repo.Save(user)
}

func (s *UserService) GetAllUsers() ([]domain.User, error) {
	return s.repo.FindAll()
}
