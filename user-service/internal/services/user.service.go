package services

import "music-streaming-microservices/user-service/internal/repositories"

type IUserService interface {
	Register(email string, password string, name string, phone string, roles []string) (bool, error)
}

type userService struct {
	userRepository *repositories.IUserRepository
}

func (u userService) Register(email string, password string, name string, phone string, roles []string) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func NewUserService(userRepository repositories.IUserRepository) IUserService {
	return &userService{userRepository: &userRepository}
}
