package users

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/danielgtaylor/huma/v2"
	"github.com/supabase-community/gotrue-go/types"
	"github.com/supabase-community/supabase-go"
)

type UserService struct {
	// Add dependencies here, e.g., a database client
	db *supabase.Client
}

// NewUserHandler creates a new UserService handler with the provided Supabase client.
func NewUserHandler(db *supabase.Client) *UserService {
	return &UserService{
		db: db,
	}
}

func (s *UserService) GetLoggedinUser(ctx context.Context, input *struct{}) (*UserOutput, error) {
	user, err := s.db.Auth.GetUser()
	if err != nil {
		return nil, err
	}

	userProfile, err := convertUserMetaDataToUserProfile(user.User)
	if err != nil {
		return nil, err
	}
	userProfile.ID = user.ID

	return &UserOutput{
		Body: struct {
			User UserProfile `json:"user" doc:"The user object" required:"true"`
		}{
			User: *userProfile,
		},
	}, nil
}

// GetUsers retrieves a list of users (user profiles) from the database.
func (s *UserService) GetUsers(ctx context.Context, input *struct{}) (*UsersOutput, error) {
	var users []UserProfile
	_, err := s.db.From("profiles").Select("*", "", false).ExecuteTo(&users)
	if err != nil {
		log.Printf("Error fetching users: %v", err)
		return nil, huma.Error500InternalServerError("Failed to fetch users")
	}

	output := &UsersOutput{
		Body: struct {
			User []UserProfile `json:"users" doc:"A list of users" required:"true"`
		}{
			User: users,
		},
	}

	return output, nil
}

// GetUserByID gets a user (user profile) based on their id
func (s *UserService) GetUserByID(ctx context.Context, input *struct {
	ID string `path:"id" doc:"The unique identifier for the user" required:"true"`
}) (*UserOutput, error) {
	userID := input.ID
	if userID == "" {
		return nil, huma.Error400BadRequest("User ID is required")
	}

	var user UserProfile
	_, err := s.db.From("profiles").Select("*", "", false).Eq("id", userID).Single().ExecuteTo(&user)
	if err != nil {
		log.Printf("Error fetching user: %v", err)
		return nil, huma.Error500InternalServerError("Failed to fetch user")
	}

	output := &UserOutput{
		Body: struct {
			User UserProfile `json:"user" doc:"The user object" required:"true"`
		}{
			User: user,
		},
	}

	return output, nil
}

// GetUserByUsername gets a user (user profile) by their username
func (s *UserService) GetUserByUsername(ctx context.Context, input *struct {
	Username string `path:"username" doc:"The users username" required:"true"`
}) (*UserOutput, error) {
	username := input.Username
	if username == "" {
		return nil, huma.Error400BadRequest("Username is required")
	}

	var user UserProfile
	_, err := s.db.From("profiles").Select("*", "", false).Eq("username", username).Single().ExecuteTo(&user)
	if err != nil {
		log.Printf("Error fetching user: %v", err)
		return nil, huma.Error500InternalServerError("Failed to fetch user")
	}

	output := &UserOutput{
		Body: struct {
			User UserProfile `json:"user" doc:"The user object" required:"true"`
		}{
			User: user,
		},
	}

	return output, nil

}

// Register registers the user
func (s *UserService) Register(ctx context.Context, input *UserRegisterInput) (*UserRegisterOutput, error) {
	info := types.SignupRequest{
		Email:    input.Body.Email,
		Password: input.Body.Password,
		Phone:    input.Body.Phone,
	}

	// Data:     input.Body.Data,

	// Marshal the struct into a JSON byte slice
	userMetaDataBytes, err := json.Marshal(input.Body.Data)
	if err != nil {
		log.Printf("Error marshaling struct: %v", err)
		return nil, err
	}

	// Unmarshal the JSON byte slice into a map[string]interface{}
	var userMetaData map[string]interface{}
	err = json.Unmarshal(userMetaDataBytes, &userMetaData)
	if err != nil {
		fmt.Printf("Error unmarshaling JSON to map: %v\n", err)
		return nil, err
	}

	info.Data = userMetaData

	_, err = s.db.Auth.Signup(info)
	if err != nil {
		return &UserRegisterOutput{
			Body: struct {
				Message string "json:\"message\" doc:\"A message of whether registration was successful or not\""
			}{
				Message: fmt.Sprintf("Error registering %v", err),
			},
		}, err
	}

	return &UserRegisterOutput{
		Body: struct {
			Message string "json:\"message\" doc:\"A message of whether registration was successful or not\""
		}{
			Message: "Registration was successful",
		},
	}, nil
}

// Login handles user login by validating credentials and returning a session cookie.
func (s *UserService) Login(ctx context.Context, input *UserLoginInput) (*UserLoginOutput, error) {
	if input.Body.Email == "" || input.Body.Password == "" {
		return nil, huma.Error400BadRequest("Email and password are required")
	}

	session, err := s.db.SignInWithEmailPassword(input.Body.Email, input.Body.Password)

	if err != nil {
		log.Println("Login failed:", err)
		return nil, huma.Error401Unauthorized("Invalid email or password")
	}

	userProfile, err := convertUserMetaDataToUserProfile(session.User)
	if err != nil {
		return nil, err
	}

	userProfile.ID = session.User.ID

	output := &UserLoginOutput{
		Body: struct {
			User UserProfile `json:"user" doc:"The user object" required:"true"`
		}{
			User: *userProfile,
		},
	}

	output.SetCookie = []*http.Cookie{
		{
			Name:     "access_token",
			Value:    session.AccessToken,
			MaxAge:   session.ExpiresIn,
			Expires:  time.Now().Add(time.Duration(session.ExpiresIn) * time.Second),
			Path:     "/",
			Secure:   false,
			HttpOnly: true,
			SameSite: http.SameSiteLaxMode,
		},
		{
			Name:     "refresh_token",
			Value:    session.RefreshToken,
			MaxAge:   int(session.ExpiresAt),
			Expires:  time.Now().Add(time.Duration(session.ExpiresAt)),
			Path:     "/",
			Secure:   false,
			HttpOnly: true,
			SameSite: http.SameSiteLaxMode,
		},
	}

	return output, nil
}

func (s *UserService) Logout(ctx context.Context, input *UserLogoutInput) (*UserLogoutOutput, error) {
	err := s.db.Auth.Logout()
	if err != nil {
		return nil, err
	}

	output := &UserLogoutOutput{
		SetCookie: []*http.Cookie{
			{
				Name:     input.AccessToken.Name,
				Value:    "",
				MaxAge:   -1,
				Expires:  time.Unix(0, 0),
				Path:     "/",
				Secure:   false,
				HttpOnly: true,
				SameSite: http.SameSiteLaxMode,
			},
			{
				Name:     input.RefreshToken.Name,
				Value:    "",
				MaxAge:   -1,
				Expires:  time.Unix(0, 0),
				Path:     "/",
				Secure:   false,
				HttpOnly: true,
				SameSite: http.SameSiteLaxMode,
			},
		},
	}

	return output, nil
}

/*
###### Below are Helper Functions ######
*/

// convertUserMetaDataToUserProfile digests UserMetaData from supabase to UserProfile
func convertUserMetaDataToUserProfile(user types.User) (*UserProfile, error) {
	// Marshal user.UserMetadata directly
	jsonData, err := json.Marshal(user.UserMetadata)
	if err != nil {
		return nil, fmt.Errorf("error marshalling user metadata: %w", err)
	}

	// Unmarshal into UserProfile
	var userProfile UserProfile
	if err := json.Unmarshal(jsonData, &userProfile); err != nil {
		return nil, fmt.Errorf("error unmarshalling to UserProfile: %w", err)
	}

	return &userProfile, nil
}
