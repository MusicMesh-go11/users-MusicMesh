package userprofile

import (
	pb "MusicMesh/users-MusicMesh/generate/user_profile"
	"context"
	"database/sql"
)

type UserProfileRepo struct {
	pb.UnimplementedUserServiceServer
	DB *sql.DB
}

func (u *UserProfileRepo) Create(ctx context.Context, user *pb.User) (*pb.Void, error){
	
}