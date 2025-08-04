package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"music-streaming-microservices/common-lib/consts"
	"music-streaming-microservices/common-lib/response"
	"music-streaming-microservices/common-lib/types"
	"music-streaming-microservices/user-service/global"
	"music-streaming-microservices/user-service/internal/database"
	"music-streaming-microservices/user-service/internal/helper"
	"music-streaming-microservices/user-service/internal/models/dto"
	"music-streaming-microservices/user-service/internal/repositories"
	"music-streaming-microservices/user-service/internal/utils"
	"music-streaming-microservices/user-service/internal/utils/hash"
	"music-streaming-microservices/user-service/internal/utils/random"
	"music-streaming-microservices/user-service/internal/validation"
)

type IUserService interface {
	Login(schema validation.UserLoginSchema) (code int, msg string, data interface{})
	Register(userRegisterRequest validation.UserRegisterSchema) (code int, msg string, data interface{})
	VerifyOTPRequest(otpRequest validation.VerifyOTPRequest) (code int, msg string, data interface{})
	ConvertSchemaValidateToParams(avatar, email, name, password string) database.CreateUserParams
	ToDTO(user database.User) dto.UsersDTO
}

type userService struct {
	userRepository             repositories.IUserRepository
	userAuthRepository         repositories.IUserAuthRepository
	userLoginSessionRepository repositories.IUserLoginSessionRepository
}

func (us *userService) Login(userLogin validation.UserLoginSchema) (code int, msg string, data interface{}) {
	user, err := us.userRepository.GetUserByEmail(userLogin.Email)
	if err != nil {
		return response.NOT_FOUND, "User not found", nil
	}

	if !hash.MatchingWithHashPassword(user.Password, userLogin.Password) {
		return response.UNAUTHORIZED, "Invalid password", nil
	}

	key := helper.GenerateKey()
	publicKey, err := helper.GetPublicKeyString(key)
	if err != nil {
		log.Printf("Failed to get public key string: %v", err)
		return
	}

	uuid := helper.GenerateJTI()

	token, err := helper.CreateTokenPair(uuid.String(), user, key)
	if err != nil {
		return response.INTERNAL_SERVER_ERROR, "Failed to create token", nil
	}

	sessionParams := database.CreateUserLoginSessionParams{
		Uuid:        uuid,
		UserID:      int64(user.ID),
		PublicKey:   publicKey,
		RfToken:     token.RfToken,
		RfTokenUsed: []string{},
	}

	session, err := us.userLoginSessionRepository.CreateLoginSession(sessionParams)
	if err != nil {
		log.Printf("Failed to create login : %v", err)
		return
	}
	log.Printf("Created login session for user: %v", session.UserID)
	return response.OK, "Login successful", token
}

func (us *userService) Register(userRegisterRequest validation.UserRegisterSchema) (code int, msg string, data interface{}) {
	isEmailExist, _ := us.userRepository.IsEmailExist(userRegisterRequest.Email)
	if isEmailExist {
		return response.BAD_REQUEST, "From already exists", nil
	}

	hashEmail := hash.GetHashString(userRegisterRequest.Email)
	key := utils.GetKeyOTP(hashEmail)

	otpFound, err := us.userAuthRepository.GetData(key)
	switch {
	case errors.Is(err, redis.Nil):
		log.Println("Key does not exist")
	case err != nil:
		return response.INTERNAL_SERVER_ERROR, "Failed to retrieve OTP data", nil
	case otpFound != "":
		return response.BAD_REQUEST, "OTP already exist", nil
	}

	hashedPassword, err := hash.HashPassword(userRegisterRequest.Password)
	if err != nil {
		return response.INTERNAL_SERVER_ERROR, "Failed to hash password", nil
	}

	userRegisterRequest.Password = hashedPassword

	var otpWithMetadata types.OTPWithMetadata[validation.UserRegisterSchema]
	otpWithMetadata.OTP = random.GenerateOTP()
	otpWithMetadata.Metadata = userRegisterRequest

	jsonBytes, _ := json.Marshal(otpWithMetadata)

	err = us.userAuthRepository.AddData(key, jsonBytes)
	if err != nil {
		return response.INTERNAL_SERVER_ERROR, "Failed to send OTP", nil
	}

	subject := "EMAIL.VerifyOTP"
	messageData := types.SendEmail[types.SendEmailOTPRegistry]{
		Type:      consts.VERIFY_OTP_USER_REGISTER,
		Recipient: userRegisterRequest.Email,
		Message: types.SendEmailOTPRegistry{
			Key: key,
		},
	}
	message, err := json.Marshal(messageData)
	if err != nil {
		return response.INTERNAL_SERVER_ERROR, "Failed to marshal message", nil
	}

	go func(subject string, message []byte) {
		ack, err := global.NatsJetStream.Publish(subject, message)
		if err != nil {
			log.Printf("Error publishing message to subject %s: %v", subject, err)
			return
		}

		if ack == nil {
			log.Printf("Publish returned nil ack for subject %s", subject)
			return
		}

		log.Printf("Message published to stream %s with sequence %d", ack.Stream, ack.Sequence)
	}(subject, message)

	return response.OK, "OTP sent successfully", nil

}

func (us *userService) VerifyOTPRequest(otpRequest validation.VerifyOTPRequest) (code int, msg string, data interface{}) {
	email := otpRequest.Email
	hashEmail := hash.GetHashString(email)
	key := utils.GetKeyOTP(hashEmail)

	data, err := us.userAuthRepository.GetData(key)
	if err != nil {
		return response.INTERNAL_SERVER_ERROR, "Failed to retrieve OTP data", nil
	}

	var otpWithMetadata types.OTPWithMetadata[validation.UserRegisterSchema]
	if err := json.Unmarshal([]byte(data.(string)), &otpWithMetadata); err != nil {
		return response.INTERNAL_SERVER_ERROR, "Failed to unmarshal OTP data", nil
	}

	if otpWithMetadata.OTP != otpRequest.OTP {
		return response.BAD_REQUEST, "Invalid OTP", nil
	}

	go func() {
		if err := us.userAuthRepository.DeleteData(key); err != nil {
			fmt.Printf("Failed to delete OTP Data: %v", err)
		}
	}()

	userParams := us.ConvertSchemaValidateToParams(otpWithMetadata.Metadata.Avatar, otpWithMetadata.Metadata.Email, otpWithMetadata.Metadata.Name, otpWithMetadata.Metadata.Password)

	newUser := us.userRepository.CreateNewUser(userParams)

	data = us.ToDTO(newUser)
	return response.CREATED, "User created successfully", data
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
	userLoginSessionRepository repositories.IUserLoginSessionRepository,
) IUserService {
	return &userService{
		userRepository:             userRepository,
		userAuthRepository:         userAuthRepository,
		userLoginSessionRepository: userLoginSessionRepository,
	}
}
