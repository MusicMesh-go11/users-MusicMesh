package userprofiles

import (
	"MusicMesh/users-MusicMesh/generate/user_profile"
	"database/sql"
)

type UserProfileRepo struct {
	DB *sql.DB
	user_profile.UnimplementedUserServiceServer
}

func NewUserProfile(db *sql.DB) *UserProfileRepo {
	return &UserProfileRepo{DB: db}
}
