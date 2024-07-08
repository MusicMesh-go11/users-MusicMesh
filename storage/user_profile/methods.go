package userprofiles

import (
	pb "MusicMesh/users-MusicMesh/generate/user_profile"
	"context"
	"log"
)

func (u *UserProfileRepo) Create(ctx context.Context, in *pb.UserProfile) (*pb.Void, error) {
	_, err := u.DB.Exec(`INSERT INTO user_profiles(user_id, full_name, bio, role, location, avatar_url, website)
	values ($1, $2, $3, $4, $5, $6, $7)`, in.UserID, in.FullName, in.Bio, in.Role,
		in.Location, in.AvatarUrl, in.Website)
	log.Println(err)
	return &pb.Void{}, err
}

func (u *UserProfileRepo) GetByID(ctx context.Context, in *pb.UserProfileId) (*pb.UserProfile, error) {
	user := pb.UserProfile{}
	err := u.DB.QueryRow("select user_id, full_name, bio, role, location, avatar_url, website from user_profiles where user_id = $1", in.Id).
		Scan(&user.UserID, &user.FullName, &user.Bio, &user.Role, &user.Location, &user.AvatarUrl, &user.Website)
	log.Println(err)
	return &user, err
}

func (u *UserProfileRepo) Get(ctx context.Context, in *pb.FilterRequest) (*pb.UsersProfiles, error) {
	users := []*pb.UserProfile{}

	rows, err := u.DB.Query(in.Query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		user := pb.UserProfile{}
		err = rows.Scan(&user.UserID, &user.FullName, &user.Bio, &user.Role,
			&user.Location, &user.AvatarUrl, &user.Website)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	return &pb.UsersProfiles{UsersProfile: users}, err
}

func (u *UserProfileRepo) Update(ctx context.Context, in *pb.UserProfile) (*pb.Void, error) {
	_, err := u.DB.Exec(`UPDATE user_profiles SET full_name = $1, bio = $2, role = $3,
                         location = $4, avatar_url = $5, website = $6 WHERE user_id = $7`,
		in.FullName, in.Bio, in.Role, in.Location, in.AvatarUrl, in.Website, in.UserID)
	return &pb.Void{}, err
}

func (u *UserProfileRepo) Delete(ctx context.Context, in *pb.UserProfileId) (*pb.Void, error) {
	_, err := u.DB.Exec(`update user_profiles set 
                         deleted_at = date_part('epoch', current_timestamp)::INT 
                     where user_id = $1 and deleted_at = 0`, in.Id)
	return &pb.Void{}, err
}
