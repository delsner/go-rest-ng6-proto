package services

import (
    "sync"
    "github.com/delsner/go-rest-ng6-proto/backend/repositories"
    "github.com/delsner/go-rest-ng6-proto/backend/models"
)

// Singleton setup
var singleton *UserService
var once sync.Once

type UserService struct {
    userRepository *repositories.UserRepository
}

func GetUserServiceInstance() (*UserService) {
    once.Do(func() {
        singleton = &UserService{}
        singleton.userRepository = repositories.GetUserRepositoryInstance()
    })
    return singleton
}

// Service methods
func (us *UserService) GetByName(name string) (*models.User) {
    return us.userRepository.GetByName(name)
}

func (us *UserService) GetById(id string) (*models.User) {
    return us.userRepository.GetById(id)
}

func (us *UserService) Create(user *models.User) (*models.User) {
    return us.userRepository.Create(user)
}
func (us *UserService) GetAll() ([]*models.User) {
    return us.userRepository.GetAll()
}
