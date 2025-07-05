package users

// User represents a user.
type User struct {
	ID    string `json:"id" example:"12345" doc:"The unique identifier for the user"`
	Name  string `json:"name" example:"John Doe" doc:"The name of the user"`
	Email string `json:"email" example:"john.doe@example.com" doc:"The email address of the user"`
}

// UserOutput represents the user operation response.
type UserOutput struct {
	Body struct {
		User []User `json:"users" doc:"A list of users" required:"true"`
	}
}
