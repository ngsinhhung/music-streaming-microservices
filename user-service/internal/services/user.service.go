package services

import (
	"fmt"
	"music-streaming-microservices/common-lib/response"
	"music-streaming-microservices/user-service/internal/database"
	"music-streaming-microservices/user-service/internal/models/dto"
	"music-streaming-microservices/user-service/internal/repositories"
	"music-streaming-microservices/user-service/internal/utils"
	"music-streaming-microservices/user-service/internal/utils/hash"
	"music-streaming-microservices/user-service/internal/utils/random"
	"music-streaming-microservices/user-service/validation"
	"time"
)

type IUserService interface {
	Register(userRegisterRequest validation.UserRegisterSchema) (code int, msg string, data interface{})
	ConvertSchemaValidateToParams(avatar, email, name, password string) database.CreateUserParams
	ToDTO(user database.User) dto.UsersDTO
}

type userService struct {
	userRepository     repositories.IUserRepository
	userAuthRepository repositories.IUserAuthRepository
}

func (us *userService) Register(userRegisterRequest validation.UserRegisterSchema) (code int, msg string, data interface{}) {
	isEmailExist, _ := us.userRepository.IsEmailExist(userRegisterRequest.Email)
	if isEmailExist {
		return response.BAD_REQUEST, "Email already exists", nil
	}

	hashEmail := hash.GetHashString(userRegisterRequest.Email)
	otp := random.GenerateOTP()

	fmt.Printf("Generated OTP for email %s: %d\n", hashEmail, otp)

	key := utils.GetKeyOTP(hashEmail)

	fmt.Println(key)

	err := us.userAuthRepository.AddOTP(key, otp, int64(5*time.Minute)) // 5 minutes expiration
	if err != nil {
		return response.INTERNAL_SERVER_ERROR, "Failed to send OTP", nil
	}
	return response.OK, "OTP sent successfully", nil

	//hashedPassword, err := hash.HashPassword(userRegisterRequest.Password)
	//if err != nil {
	//	return response.INTERNAL_SERVER_ERROR, "Failed to hash password", nil
	//}

	//userParams := us.ConvertSchemaValidateToParams(userRegisterRequest.Avatar, userRegisterRequest.Email, userRegisterRequest.Name, hashedPassword)
	//newUser := us.userRepository.CreateNewUser(userParams)
	//data = us.ToDTO(newUser)
	//return response.CREATED, "User created successfully", data

}

func (us *userService) ConvertSchemaValidateToParams(avatar, email, name, password string) database.CreateUserParams {
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

func NewUserService(
	userRepository repositories.IUserRepository,
	userAuthRepository repositories.IUserAuthRepository,
) IUserService {
	return &userService{
		userRepository:     userRepository,
		userAuthRepository: userAuthRepository,
	}
}
