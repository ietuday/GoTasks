package models

import (
	"errors"
	"strings"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	RoleAdmin = "admin"
	RoleUser  = "user"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Username string             `bson:"username" json:"username"`
	Password string             `bson:"password" json:"password,omitempty"`
	Role     string             `bson:"role" json:"role"` // "admin" or "user"
}

func (u *User) Validate() error {
	u.Username = strings.TrimSpace(u.Username)
	u.Role = strings.ToLower(strings.TrimSpace(u.Role))

	if u.Username == "" {
		return errors.New("username cannot be empty")
	}
	if len(u.Password) < 6 {
		return errors.New("password must be at least 6 characters long")
	}
	if u.Role != RoleAdmin && u.Role != RoleUser {
		return errors.New("role must be either 'admin' or 'user'")
	}
	return nil
}
