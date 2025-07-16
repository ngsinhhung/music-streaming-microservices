package repositories

import (
	"fmt"
	"music-streaming-microservices/user-service/global"
	"music-streaming-microservices/user-service/internal/database"
)

type IUserRepository interface {
	GetUserByEmail(email string) (bool, error)
	GetUserById(id int) (bool, error)
	IsEmailExist(email string) (bool, error)
	CreateNewUser(userParams database.CreateUserParams) database.User
}

type userRepository struct {
	sqlc *database.Queries
}

func (ur *userRepository) CreateNewUser(userParams database.CreateUserParams) database.User {
	newUser, err := ur.sqlc.CreateUser(ctx, userParams)
	if err != nil {
		fmt.Println(err)
		return database.User{}
	}
	return newUser
}

func (ur *userRepository) IsEmailExist(email string) (bool, error) {
	isUserExist, err := ur.sqlc.IsUserExists(ctx, email)
	return isUserExist, err

}

func (ur *userRepository) GetUserByEmail(email string) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (ur *userRepository) GetUserById(id int) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func NewUserRepository() IUserRepository {
	return &userRepository{
		sqlc: database.New(global.PostgresConn),
	}
}
