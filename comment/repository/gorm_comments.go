package repository

import (
	"github.com/GoGroup/Movie-and-events/comment"
	"github.com/GoGroup/Movie-and-events/model"
	"github.com/jinzhu/gorm"
)

// CommentGormRepo implements menu.CommentRepository interface
type CommentGormRepo struct {
	conn *gorm.DB
}

// NewCommentGormRepo returns new object of CommentGormRepo
func NewCommentGormRepo(db *gorm.DB) comment.CommentRepository {
	return &CommentGormRepo{conn: db}
}

// Comments returns all customer comments stored in the database
func (cmntRepo *CommentGormRepo) Comments() ([]model.Comment, []error) {
	cmnts := []model.Comment{}
	errs := cmntRepo.conn.Find(&cmnts).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return cmnts, errs
}

// Comment retrieves a customer comment from the database by its id
func (cmntRepo *CommentGormRepo) Comment(id uint) (*model.Comment, []error) {
	cmnt := model.Comment{}
	errs := cmntRepo.conn.First(&cmnt, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &cmnt, errs
}

// UpdateComment updates a given customer comment in the database
func (cmntRepo *CommentGormRepo) UpdateComment(comment *model.Comment) (*model.Comment, []error) {
	cmnt := comment
	errs := cmntRepo.conn.Save(cmnt).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return cmnt, errs
}

// DeleteComment deletes a given customer comment from the database
func (cmntRepo *CommentGormRepo) DeleteComment(id uint) (*model.Comment, []error) {
	cmnt, errs := cmntRepo.Comment(id)

	if len(errs) > 0 {
		return nil, errs
	}

	errs = cmntRepo.conn.Delete(cmnt, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return cmnt, errs
}

// StoreComment stores a given customer comment in the database
func (cmntRepo *CommentGormRepo) StoreComment(comment *model.Comment) (*model.Comment, []error) {
	cmnt := comment
	errs := cmntRepo.conn.Create(cmnt).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return cmnt, errs
}
