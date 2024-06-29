package user

import (
	pb "MusicMesh/users-MusicMesh/generate/user"
	"context"
)

func (u *UserRepo) Create(ctx context.Context, user *pb.User) (*pb.Void, error){
	
}