package user

import (
	"github.com/GoGroup/Movie-and-events/model"
)

type UserService interface {
	// Users() ([]model.User, []error)
	User(id uint) (*model.User, []error)
	UserByEmail(email string) (*model.User, []error)
	// UpdateUser(user *model.User) (*model.User, []error)
	// DeleteUser(id uint) (*model.User, []error)
	StoreUser(user *model.User) (*model.User, []error)

	EmailExists(email string) bool
}
type RoleService interface {
	Roles() ([]model.Role, []error)
	Role(id uint) (*model.Role, []error)
	RoleByName(name string) (*model.Role, []error)
	UpdateRole(role *model.Role) (*model.Role, []error)
	DeleteRole(id uint) (*model.Role, []error)
	StoreRole(role *model.Role) (*model.Role, []error)
}

// SessionService specifies logged in user session related service
type SessionService interface {
	Session(sessionId string) (*model.Session, []error)
	Sessions() ([]model.Session, []error)
	StoreSession(session *model.Session) (*model.Session, []error)
	DeleteSession(sessionId string) (*model.Session, []error)
}
