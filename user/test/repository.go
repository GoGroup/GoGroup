package user

import (
	"github.com/GoGroup/Movie-and-events/model"
)

type UserRepository interface {
	// Users() ([]model.User, []error)
	User(id uint) (*model.User, []error)
	UserByEmail(email string) (*model.User, []error)
	// UpdateUser(user *model.User) (*model.User, []error)
	// DeleteUser(id uint) (*model.User, []error)
	StoreUser(user *model.User) (*model.User, []error)

	EmailExists(email string) bool
}
