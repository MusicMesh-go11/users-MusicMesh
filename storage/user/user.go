package user

import (
	pb"MusicMesh/users-MusicMesh/generate/user"
	"database/sql"
)

type UserRepo struct {
	pb.UnimplementedUserServiceServer
	DB *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{DB: db}
}
