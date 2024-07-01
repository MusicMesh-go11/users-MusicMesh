package users

import (
	"MusicMesh/users-MusicMesh/generate/user"
	"database/sql"
)

type UserRepo struct {
	user.UnimplementedUserServiceServer
	DB *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{DB: db}
}
