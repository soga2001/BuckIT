package users

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/supabase-community/gotrue-go/types"
)

type User struct {
	types.User
}

type UserProfile struct {
	ID         uuid.UUID `json:"id" doc:"The unique identifier for the user" required:"true"`
	Created_At string    `json:"created_at" doc:"The date and time when the user was created" required:"true"`
	Updated_At string    `json:"updated_at" doc:"The date and time when the user was last updated" required:"true"`
	Username   string    `json:"username" doc:"The username of the user" required:"true"`
	Full_Name  string    `json:"full_name" doc:"The full name of the user" required:"true"`
	Avatar_URL string    `json:"avatar_url" doc:"The URL of the user's avatar image" required:"true"`
	Website    string    `json:"website" doc:"The user's personal or professional website URL" required:"true"`
	Bio        string    `json:"bio" doc:"A short biography or description of the user" required:"true"`
	Location   string    `json:"location" doc:"The user's location, such as city or country" required:"true"`
	Dob        string    `json:"dob" doc:"The date of birth of the user in ISO 8601 format" required:"true"`
	Private    bool      `json:"private" doc:"Indicates whether the user's profile is private (true) or public (false)" required:"true"`
	Verified   bool      `json:"verified" doc:"Indicates whether the user is verified (true) or not (false)" required:"true"`
	Followers  int       `json:"followers" doc:"The number of followers the user has" required:"true"`
	Following  int       `json:"following" doc:"The number of users the user is following" required:"true"`
}

// UserOutput is user to return a singular user profile.
type UserOutput struct {
	Body struct {
		User UserProfile `json:"user" doc:"The user object" required:"true"`
	}
}

// UsersOutput is the output type of handlers that return multiple UserProfiles.
type UsersOutput struct {
	Body struct {
		User []UserProfile `json:"users" doc:"A list of users" required:"true"`
	}
}

// UserLoginInput is the input the Login handler digests.
type UserLoginInput struct {
	Body struct {
		Email    string `json:"email" doc:"The email address of the user" required:"true"`
		Password string `json:"password" doc:"The password of the user" required:"true"`
	}
}

// UserLoginOutput is the output type of the user handler.
type UserLoginOutput struct {
	SetCookie []*http.Cookie `header:"Set-Cookie"`
	Body      struct {
		User UserProfile `json:"user" doc:"The user object" required:"true"`
	}
}

type UserRegisterInput struct {
	Body struct {
		Email    string               `json:"email" doc:"The email address of the user" required:"false"`
		Phone    string               `json:"phone" doc:"The phone number of the user" required:"false"`
		Password string               `json:"password" doc:"The password of the user" required:"true"`
		Data     UserRegisterMetaData `json:"data" doc:"The user profile information" required:"true"`
	}
}

type UserRegisterMetaData struct {
	Username   string `json:"username" doc:"The username of the user" required:"true"`
	Full_Name  string `json:"full_name" doc:"The full name of the user" required:"true"`
	Avatar_URL string `json:"avatar_url,omitempty" doc:"The URL of the user's avatar image" required:"false"`
	Website    string `json:"website,omitempty" doc:"The user's personal or professional website URL" required:"false"`
	Bio        string `json:"bio,omitempty" doc:"A short biography or description of the user" required:"false"`
	Location   string `json:"location,omitempty" doc:"The user's location, such as city or country" required:"false"`
	Dob        string `json:"dob" doc:"The date of birth of the user in ISO 8601 format" required:"true"`
}

type UserRegisterOutput struct {
	Body struct {
		Message string `json:"message" doc:"A message of whether registration was succuessful or not"`
	}
}

type UserLogoutInput struct {
	AccessToken  http.Cookie `cookie:"access_token"`
	RefreshToken http.Cookie `cookie:"refresh_token"`
}

type UserLogoutOutput struct {
	SetCookie []*http.Cookie `header:"Set-Cookie"`
}
