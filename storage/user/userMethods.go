package users

import (
	pb "MusicMesh/users-MusicMesh/generate/user"
	"context"
)

func (u *UserRepo) Create(ctx context.Context, in *pb.User) (*pb.Void, error) {
	_, err := u.DB.Exec("INSERT INTO users (user_name, email, password) VALUES ($1, $2, $3)",
		in.UserName, in.Email, in.Password)
	return &pb.Void{}, err
}

func (u *UserRepo) Get(ctx context.Context, in *pb.FilterRequest) (*pb.UsersRes, error) {
	rows, err := u.DB.Query(in.Query, in.Arr)
	if err != nil {
		return nil, err
	}
	users := &pb.UsersRes{}
	for rows.Next() {
		user := &pb.UserRes{}
		err = rows.Scan(&user.UserID, &user.UserName, &user.Email, &user.Password, &user.CreatedAt)
		if err != nil {
			return nil, err
		}
		users.Users = append(users.Users, user)
	}
	return users, nil
}

func (u *UserRepo) GetByID(ctx context.Context, in *pb.UserId) (*pb.UserRes, error) {
	row := u.DB.QueryRow("SELECT user_name, email, password, created_at FROM users WHERE user_id = $1", in.Id)

	user := &pb.UserRes{}
	err := row.Scan(&user.UserName, &user.Email, &user.Password, &user.CreatedAt)
	user.UserID = in.Id
	return user, err
}

func (u *UserRepo) Update(ctx context.Context, in *pb.UserRes) (*pb.Void, error) {
	_, err := u.DB.Exec("UPDATE users SET user_name = $1, email = $2, password = $3 WHERE user_id = $4",
		in.UserName, in.Email, in.Password, in.UserID)
	return &pb.Void{}, err
}

func (u *UserRepo) Delete(context.Context, *pb.UserId) (*pb.Void, error) {
	_, err := u.DB.Exec(`update users set
 deleted_at = date_part('epoch', current_timestamp)::INT
where user_id = $1 and deleted_at = 0`)
	return &pb.Void{}, err
}
