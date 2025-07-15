package services

import (
	"fmt"
	"music-streaming-microservices/common-lib/response"
	"music-streaming-microservices/user-service/internal/database"
	"music-streaming-microservices/user-service/internal/models/dto"
	"music-streaming-microservices/user-service/internal/repositories"
	"music-streaming-microservices/user-service/internal/utils/hash"
	"music-streaming-microservices/user-service/validation"
)

type IUserService interface {
	Register(userRegisterRequest validation.UserRegisterSchema) (code int, msg string, data interface{})
	FromSchemaValidateToParams(avatar, email, name, password string) database.CreateUserParams
	ToDTO(user database.User) dto.UsersDTO
}

type userService struct {
	userRepository repositories.IUserRepository
}

func (us *userService) Register(userRegisterRequest validation.UserRegisterSchema) (code int, msg string, data interface{}) {
	// hash email
	isEmailExist, _ := us.userRepository.IsEmailExist(userRegisterRequest.Email)
	if isEmailExist {
		return response.BAD_REQUEST, "Email already exists", nil
	}

	hashedPassword, err := hash.HashPassword(userRegisterRequest.Password)
	if err != nil {
		return response.INTERNAL_SERVER_ERROR, "Failed to hash password", nil
	}

	userParams := us.FromSchemaValidateToParams(userRegisterRequest.Avatar, userRegisterRequest.Email, userRegisterRequest.Name, hashedPassword)
	fmt.Println(userParams)
	newUser := us.userRepository.CreateNewUser(userParams)
	fmt.Println(newUser.Email)
	data = us.ToDTO(newUser)
	fmt.Println(data)
	return response.CREATED, "User created successfully", data

}

func (us *userService) FromSchemaValidateToParams(avatar, email, name, password string) database.CreateUserParams {
	return database.CreateUserParams{
		Avatar:   avatar,
		Email:    email,
		Name:     name,
		Password: password,
	}
}

func (us *userService) ToDTO(user database.User) dto.UsersDTO {
	return dto.UsersDTO{
		ID:        user.ID,
		Avatar:    user.Avatar,
		Email:     user.Email,
		Name:      user.Name,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func NewUserService(userRepository repositories.IUserRepository) IUserService {
	return &userService{userRepository: userRepository}
}
