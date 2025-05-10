package models

import (
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestUserValidate(t *testing.T) {
	tests := []struct {
		name    string
		user    User
		wantErr bool
		errMsg  string
	}{
		{
			name: "Valid user",
			user: User{
				ID:       primitive.NewObjectID(),
				Username: "JohnDoe",
				Password: "password123",
				Role:     RoleUser,
			},
			wantErr: false,
		},
		{
			name: "Empty username",
			user: User{
				ID:       primitive.NewObjectID(),
				Username: "   ",
				Password: "password123",
				Role:     RoleUser,
			},
			wantErr: true,
			errMsg:  "username cannot be empty",
		},
		{
			name: "Short password",
			user: User{
				ID:       primitive.NewObjectID(),
				Username: "JohnDoe",
				Password: "123",
				Role:     RoleUser,
			},
			wantErr: true,
			errMsg:  "password must be at least 6 characters long",
		},
		{
			name: "Invalid role",
			user: User{
				ID:       primitive.NewObjectID(),
				Username: "JohnDoe",
				Password: "password123",
				Role:     "manager",
			},
			wantErr: true,
			errMsg:  "role must be either 'admin' or 'user'",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.user.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr && err.Error() != tt.errMsg {
				t.Errorf("Validate() error message = %v, want %v", err.Error(), tt.errMsg)
			}
		})
	}
}
