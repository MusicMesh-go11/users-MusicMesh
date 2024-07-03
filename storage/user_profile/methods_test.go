package userprofiles

import (
	pb "MusicMesh/users-MusicMesh/generate/user_profile"
	"MusicMesh/users-MusicMesh/storage/postgres"
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

//
//func TestUserProfileRepo_Create(t *testing.T) {
//	// Connect to the test database
//	db, err := postgres.Conn()
//	if err != nil {
//		t.Fatalf("Failed to open sql db, %v", err)
//	}
//	defer db.Close()
//
//	// Initialize the repository with the test database connection
//	repo := &UserProfileRepo{DB: db}
//
//	userID := "36be4369-6df1-444b-9b8c-1e52c36fbcf7" // Use a fixed UUID for consistency in testing
//	_, err = db.Exec(`INSERT INTO users(user_id, user_name, email, password, created_at, updated_at, deleted_at)
//		VALUES ($1, 'test_user', 'test@example.com', 'password', $2, $2, 0)
//		ON CONFLICT (user_id) DO NOTHING`, userID, time.Now())
//	if err != nil {
//		t.Fatalf("Failed to insert user, %v", err)
//	}
//	// Define a test user profile
//	testUserProfile := &pb.UserProfile{
//		UserID:    userID,
//		FullName:  "John Doe",
//		Bio:       "Bio",
//		Role:      "musician",
//		Location:  "Location",
//		AvatarUrl: "avatar_url",
//		Website:   "website",
//	}
//
//	// Call the Create method
//	_, err = repo.Create(context.Background(), testUserProfile)
//
//	// Assert the Create method returned no error
//	assert.NoError(t, err)
//
//	// Verify the user profile was inserted correctly by querying the database
//	var count int
//	err = db.QueryRow("SELECT COUNT(*) FROM user_profiles WHERE user_id = $1", testUserProfile.UserID).Scan(&count)
//	if err != nil {
//		t.Fatalf("Failed to query user_profiles table, %v", err)
//	}
//
//	// Assert that exactly one row was inserted
//	assert.Equal(t, 1, count)
//}

//func TestUserProfileRepo_GetById(t *testing.T) {
//	// Connect to the test database
//	db, err := postgres.Conn()
//	if err != nil {
//		t.Fatalf("Failed to open sql db, %v", err)
//	}
//	defer db.Close()
//
//	// Initialize the repository with the test database connection
//	repo := &UserProfileRepo{DB: db}
//
//	// Insert a test user profile into the user_profiles table if not already present
//	userID := "36be4369-6df1-444b-9b8c-1e52c36fbcf7"
//	_, err = db.Exec(`INSERT INTO user_profiles(user_id, full_name, bio, role, location, avatar_url, website)
//		VALUES ($1, 'John Doe', 'Bio', 'musician', 'Location', 'avatar_url', 'website')
//		ON CONFLICT (user_id) DO NOTHING`, userID)
//	if err != nil {
//		t.Fatalf("Failed to insert test data, %v", err)
//	}
//
//	// Define the user profile ID to get
//	userProfileID := &pb.UserProfileId{Id: userID}
//
//	// Call the GetById method
//	userProfile, err := repo.GetById(context.Background(), userProfileID)
//
//	// Assert the GetById method returned no error
//	assert.NoError(t, err)
//
//	// Verify the details of the user profile
//	assert.Equal(t, userID, userProfile.UserID)
//	assert.Equal(t, "John Doe", userProfile.FullName)
//	assert.Equal(t, "Bio", userProfile.Bio)
//	assert.Equal(t, "musician", userProfile.Role)
//	assert.Equal(t, "Location", userProfile.Location)
//	assert.Equal(t, "avatar_url", userProfile.AvatarUrl)
//	assert.Equal(t, "website", userProfile.Website)
//}

//func TestUserProfileRepo_Update(t *testing.T) {
//	// Connect to the test database
//	db, err := postgres.Conn()
//	if err != nil {
//		t.Fatalf("Failed to open sql db, %v", err)
//	}
//	defer db.Close()
//
//	// Initialize the repository with the test database connection
//	repo := &UserProfileRepo{DB: db}
//
//	// Insert a test user profile into the user_profiles table if not already present
//	userID := "36be4369-6df1-444b-9b8c-1e52c36fbcf7"
//	_, err = db.Exec(`INSERT INTO user_profiles(user_id, full_name, bio, role, location, avatar_url, website)
//		VALUES ($1, 'John Doe', 'Bio', 'musician', 'Location', 'avatar_url', 'website')
//		ON CONFLICT (user_id) DO NOTHING`, userID)
//	if err != nil {
//		t.Fatalf("Failed to insert test data, %v", err)
//	}
//
//	// Define the updated user profile data
//	updatedUserProfile := &pb.UserProfile{
//		UserID:    userID,
//		FullName:  "John Doe Updated",
//		Bio:       "Bio Updated",
//		Role:      "producer",
//		Location:  "New Location",
//		AvatarUrl: "new_avatar_url",
//		Website:   "new_website",
//	}
//
//	// Call the Update method
//	_, err = repo.Update(context.Background(), updatedUserProfile)
//
//	// Assert the Update method returned no error
//	assert.NoError(t, err)
//
//	// Verify the user profile details were updated correctly by querying the database
//	var fullName, bio, role, location, avatarUrl, website string
//	err = db.QueryRow("SELECT full_name, bio, role, location, avatar_url, website FROM user_profiles WHERE user_id = $1", userID).
//		Scan(&fullName, &bio, &role, &location, &avatarUrl, &website)
//	if err != nil {
//		t.Fatalf("Failed to query user_profiles table, %v", err)
//	}
//
//	// Assert that the user profile details were updated correctly
//	assert.Equal(t, "John Doe Updated", fullName)
//	assert.Equal(t, "Bio Updated", bio)
//	assert.Equal(t, "producer", role)
//	assert.Equal(t, "New Location", location)
//	assert.Equal(t, "new_avatar_url", avatarUrl)
//	assert.Equal(t, "new_website", website)
//}

func TestUserProfileRepo_Delete(t *testing.T) {
	// Connect to the test database
	db, err := postgres.Conn()
	if err != nil {
		t.Fatalf("Failed to open sql db, %v", err)
	}
	defer db.Close()

	// Initialize the repository with the test database connection
	repo := &UserProfileRepo{DB: db}

	// Insert a test user profile into the user_profiles table if not already present
	userID := "36be4369-6df1-444b-9b8c-1e52c36fbcf7"
	_, err = db.Exec(`INSERT INTO user_profiles(user_id, full_name, bio, role, location, avatar_url, website)
		VALUES ($1, 'John Doe', 'Bio', 'musician', 'Location', 'avatar_url', 'website')
		ON CONFLICT (user_id) DO NOTHING`, userID)
	if err != nil {
		t.Fatalf("Failed to insert test data, %v", err)
	}

	// Define the user profile ID to delete
	userProfileID := &pb.UserProfileId{Id: userID}

	// Call the Delete method
	_, err = repo.Delete(context.Background(), userProfileID)

	// Assert the Delete method returned no error
	assert.NoError(t, err)

	// Verify the user profile was marked as deleted by querying the database
	var deletedAt int64
	err = db.QueryRow("SELECT deleted_at FROM user_profiles WHERE user_id = $1", userID).Scan(&deletedAt)
	if err != nil {
		t.Fatalf("Failed to query user_profiles table, %v", err)
	}

	// Assert that the user profile was marked as deleted
	assert.NotEqual(t, 0, deletedAt, "Expected deleted_at to be set to a non-zero value")
}
