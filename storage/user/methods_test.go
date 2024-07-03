package users

import (
	pb "MusicMesh/users-MusicMesh/generate/user"
	"MusicMesh/users-MusicMesh/storage/postgres"
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserRepo_Create(t *testing.T) {
	db, err := postgres.Conn()
	if err != nil {
		t.Fatalf("Failed to open sql db, %v", err)
	}
	defer db.Close()

	repo := &UserRepo{DB: db}

	testUser := &pb.User{
		UserName: "test_user",
		Email:    "test@example.com",
		Password: "password",
	}

	_, err = repo.Create(context.Background(), testUser)

	assert.NoError(t, err)

	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM users WHERE user_name = $1 AND email = $2", testUser.UserName, testUser.Email).Scan(&count)
	if err != nil {
		t.Fatalf("Failed to query users table, %v", err)
	}

	assert.Equal(t, 1, count)
}

//func TestUserRepo_Get(t *testing.T) {
//	// Connect to the test database
//	db, err := postgres.Conn()
//	if err != nil {
//		t.Fatalf("Failed to open sql db, %v", err)
//	}
//	defer db.Close()
//
//	// Initialize the repository with the test database connection
//	repo := &UserRepo{DB: db}
//
//	//// Insert a test user into the users table if not already present
//	//userID := "df004dcb-4207-43d6-a028-abb48991b613"
//	//_, err = db.Exec(`INSERT INTO users (user_id, user_name, email, password, created_at, updated_at, deleted_at)
//	//	VALUES ($1, 'test_user', 'test@example.com', 'password', now(), now(), 0)
//	//	ON CONFLICT (user_id) DO NOTHING`, userID)
//	//if err != nil {
//	//	t.Fatalf("Failed to insert test data, %v", err)
//	//}
//
//	// Define the filter request to fetch all users
//	filterRequest := &pb.FilterRequest{
//		Query: "SELECT user_id, user_name, email, password, created_at FROM users",
//		Arr:   nil, // No parameters needed for this query
//	}
//
//	// Call the Get method
//	usersRes, err := repo.Get(context.Background(), filterRequest)
//
//	// Assert the Get method returned no error
//	assert.NoError(t, err)
//
//	// Verify the number of users returned
//	assert.Equal(t, 1, len(usersRes.Users))
//
//	// Verify the details of the user
//	assert.Equal(t, "df004dcb-4207-43d6-a028-abb48991b613", usersRes.Users[0].UserID)
//	assert.Equal(t, "test_user", usersRes.Users[0].UserName)
//	assert.Equal(t, "test@example.com", usersRes.Users[0].Email)
//}
//
//func TestUserRepo_GetByID(t *testing.T) {
//	// Connect to the test database
//	db, err := postgres.Conn()
//	if err != nil {
//		t.Fatalf("Failed to open sql db, %v", err)
//	}
//	defer db.Close()
//
//	// Initialize the repository with the test database connection
//	repo := &UserRepo{DB: db}
//
//	// Insert a test user into the users table if not already present
//	userID := "df004dcb-4207-43d6-a028-abb48991b613"
//	//_, err = db.Exec(`INSERT INTO users (user_id, user_name, email, password, created_at, updated_at, deleted_at)
//	//	VALUES ($1, 'test_user', 'test@example.com', 'password', now(), now(), 0)
//	//	ON CONFLICT (user_id) DO NOTHING`, userID)
//	//if err != nil {
//	//	t.Fatalf("Failed to insert test data, %v", err)
//	//}
//
//	// Create a UserId request
//	userId := pb.UserId{Id: userID}
//
//	// Call the GetByID method
//	userRes, err := repo.GetByID(context.Background(), &userId)
//
//	// Assert the GetByID method returned no error
//	assert.NoError(t, err)
//
//	// Verify the details of the user
//	assert.Equal(t, "df004dcb-4207-43d6-a028-abb48991b613", userRes.UserID)
//	assert.Equal(t, "updated_user", userRes.UserName)
//	assert.Equal(t, "updated@example.com", userRes.Email)
//}
//
//func TestUserRepo_Update(t *testing.T) {
//	// Connect to the test database
//	db, err := postgres.Conn()
//	if err != nil {
//		t.Fatalf("Failed to open sql db, %v", err)
//	}
//	defer db.Close()
//
//	// Initialize the repository with the test database connection
//	repo := &UserRepo{DB: db}
//
//	// Insert a test user into the users table if not already present
//	userID := "df004dcb-4207-43d6-a028-abb48991b613"
//	_, err = db.Exec(`INSERT INTO users (user_id, user_name, email, password, created_at, updated_at, deleted_at)
//		VALUES ($1, 'test_user', 'test@example.com', 'password', now(), now(), 0)
//		ON CONFLICT (user_id) DO NOTHING`, userID)
//	if err != nil {
//		t.Fatalf("Failed to insert test data, %v", err)
//	}
//
//	// Define the updated user data
//	updatedUser := &pb.UserRes{
//		UserID:   userID,
//		UserName: "updated_user",
//		Email:    "updated@example.com",
//		Password: "new_password",
//	}
//
//	// Call the Update method
//	_, err = repo.Update(context.Background(), updatedUser)
//
//	// Assert the Update method returned no error
//	assert.NoError(t, err)
//
//	// Verify the user details were updated correctly by querying the database
//	var userName, email, password string
//	err = db.QueryRow("SELECT user_name, email, password FROM users WHERE user_id = $1", userID).Scan(&userName, &email, &password)
//	if err != nil {
//		t.Fatalf("Failed to query users table, %v", err)
//	}
//
//	// Assert that the user details were updated correctly
//	assert.Equal(t, "updated_user", userName)
//	assert.Equal(t, "updated@example.com", email)
//	assert.Equal(t, "new_password", password)
//}
//
//func TestUserRepo_Delete(t *testing.T) {
//	// Connect to the test database
//	db, err := postgres.Conn()
//	if err != nil {
//		t.Fatalf("Failed to open sql db, %v", err)
//	}
//	defer db.Close()
//
//	// Initialize the repository with the test database connection
//	repo := &UserRepo{DB: db}
//
//	// Insert a test user into the users table if not already present
//	userID := "df004dcb-4207-43d6-a028-abb48991b613"
//	_, err = db.Exec(`INSERT INTO users (user_id, user_name, email, password, created_at, updated_at, deleted_at)
//		VALUES ($1, 'test_user', 'test@example.com', 'password', now(), now(), 0)
//		ON CONFLICT (user_id) DO NOTHING`, userID)
//	if err != nil {
//		t.Fatalf("Failed to insert test data, %v", err)
//	}
//
//	// Define the user ID to delete
//	userToDelete := &pb.UserId{Id: userID}
//
//	// Call the Delete method
//	_, err = repo.Delete(context.Background(), userToDelete)
//
//	// Assert the Delete method returned no error
//	assert.NoError(t, err)
//
//	// Verify the user details were updated correctly by querying the database
//	var deletedAt int64
//	err = db.QueryRow("SELECT deleted_at FROM users WHERE user_id = $1", userID).Scan(&deletedAt)
//	if err != nil {
//		t.Fatalf("Failed to query users table, %v", err)
//	}
//
//	// Assert that the user was marked as deleted
//	assert.NotEqual(t, 0, deletedAt, "Expected deleted_at to be set to a non-zero value")
//}
