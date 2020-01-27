package service

import (
	"github.com/GoGroup/Movie-and-events/model"
	"github.com/GoGroup/Movie-and-events/user"
)

type UserService struct {
	userRepo user.UserRepository
}

func NewUserService(userRepository user.UserRepository) user.UserService {
	return &UserService{userRepo: userRepository}
}

// func (us *UserService) Users() ([]model.User, []error) {
// 	usrs, errs := us.userRepo.Users()
// 	if len(errs) > 0 {
// 		return nil, errs
// 	}
// 	return usrs, errs
// }

// User retrieves an application user by its id
func (us *UserService) User(id uint) (*model.User, []error) {
	usr, errs := us.userRepo.User(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return usr, errs
}

// UserByEmail retrieves an application user by its email address
func (us *UserService) UserByEmail(email string) (*model.User, []error) {
	usr, errs := us.userRepo.UserByEmail(email)
	if len(errs) > 0 {
		return nil, errs
	}
	return usr, errs
}

func (us *UserService) UpdateUserAmount(user *model.User, Amount uint) *model.User {
	usr := us.userRepo.UpdateUserAmount(user, Amount)

	return usr
}

// DeleteUser deletes a given application user
// func (us *UserService) DeleteUser(id uint) (*model.User, []error) {
// 	usr, errs := us.userRepo.DeleteUser(id)
// 	if len(errs) > 0 {
// 		return nil, errs
// 	}
// 	return usr, errs
// }

// StoreUser stores a given application user
func (us *UserService) StoreUser(user *model.User) (*model.User, []error) {
	usr, errs := us.userRepo.StoreUser(user)
	if len(errs) > 0 {
		return nil, errs
	}
	return usr, errs
}
func (us *UserService) EmailExists(email string) bool {
	exists := us.userRepo.EmailExists(email)
	return exists
}
