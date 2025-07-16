package dto

import "github.com/jackc/pgx/v5/pgtype"

type UsersDTO struct {
	ID        int32            `json:"id"`
	Avatar    string           `json:"avatar"`
	Email     string           `json:"email"`
	Name      string           `json:"name"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
	UpdatedAt pgtype.Timestamp `json:"updated_at"`
}
