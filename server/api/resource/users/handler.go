package users

import (
	"context"
)

type UserService struct {
	// Add dependencies here, e.g., a database client
}

func NewUserHandler() *UserService {
	return &UserService{}
}

func (s *UserService) GetUsers(ctx context.Context, input *struct{}) (*UserOutput, error) {
	return &UserOutput{
		Body: struct {
			User []User `json:"users" doc:"A list of users" required:"true"`
		}{
			User: []User{
				{
					ID:    "12345",
					Name:  "John Doe",
					Email: "john.doe@example.com",
				},
				{
					ID:    "12345",
					Name:  "John Smith",
					Email: "john.smith@example.com",
				},
			},
		},
	}, nil
}
