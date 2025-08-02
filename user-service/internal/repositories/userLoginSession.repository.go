package repositories

import (
	"music-streaming-microservices/user-service/global"
	"music-streaming-microservices/user-service/internal/database"
)

type IUserLoginSessionRepository interface {
	CreateLoginSession(sessionParams database.CreateLoginSessionParams) (database.UserLoginSession, error)
}

type UserLoginSessionRepository struct {
	sqlc *database.Queries
}

func (u *UserLoginSessionRepository) CreateLoginSession(sessionParams database.CreateLoginSessionParams) (database.UserLoginSession, error) {
	userLoginSession, err := u.sqlc.CreateLoginSession(ctx, sessionParams)
	return userLoginSession, err
}

func NewUserLoginSessionRepository() IUserLoginSessionRepository {
	return &UserLoginSessionRepository{
		sqlc: database.New(global.PostgresConn),
	}
}
