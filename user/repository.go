package user

import (
	"github.com/GoGroup/Movie-and-events/model"
)

type UserRepository interface {
	// Users() ([]model.User, []error)
	User(id uint) (*model.User, []error)
	UserByEmail(email string) (*model.User, []error)
	UpdateUserAmount(user *model.User, Amount uint) (*model.User)
	// DeleteUser(id uint) (*model.User, []error)
	StoreUser(user *model.User) (*model.User, []error)

	EmailExists(email string) bool
}
type RoleRepository interface {
	Roles() ([]model.Role, []error)
	Role(id uint) (*model.Role, []error)
	RoleByName(name string) (*model.Role, []error)
	UpdateRole(role *model.Role) (*model.Role, []error)
	DeleteRole(id uint) (*model.Role, []error)
	StoreRole(role *model.Role) (*model.Role, []error)
}
type SessionRepository interface {
	Session(sessionId string) (*model.Session, []error)
	Sessions() ([]model.Session, []error)
	StoreSession(session *model.Session) (*model.Session, []error)
	DeleteSession(sessionId string) (*model.Session, []error)
}
