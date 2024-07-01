package main

import (
	"MusicMesh/users-MusicMesh/generate/user"
	"MusicMesh/users-MusicMesh/generate/user_profile"
	"MusicMesh/users-MusicMesh/storage/postgres"
	"MusicMesh/users-MusicMesh/storage/user"
	userprofiles "MusicMesh/users-MusicMesh/storage/user_profile"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	db, err := postgres.Conn()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	listner, err := net.Listen("tcp", ":5051")
	if err != nil {
		panic(err)
	}
	defer listner.Close()

	server := grpc.NewServer()

	u := users.NewUserRepo(db)
	up := userprofiles.NewUserProfile(db)

	user.RegisterUserServiceServer(server, u)
	user_profile.RegisterUserServiceServer(server, up)

	log.Println("Starting server no :5051...")

	panic(server.Serve(listner))
}
