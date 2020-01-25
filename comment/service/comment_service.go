package service

import (
	"github.com/GoGroup/Movie-and-events/comment"
	"github.com/GoGroup/Movie-and-events/model"
)

// CommentService implements menu.CommentService interface
type CommentService struct {
	commentRepo comment.CommentRepository
}

// NewCommentService returns a new CommentService object
func NewCommentService(commRepo comment.CommentRepository) comment.CommentService {
	return &CommentService{commentRepo: commRepo}
}

// Comments returns all stored comments
func (cs *CommentService) Comments() ([]model.Comment, []error) {
	cmnts, errs := cs.commentRepo.Comments()
	if len(errs) > 0 {
		return nil, errs
	}
	return cmnts, errs
}

// Comment retrieves stored comment by its id
func (cs *CommentService) Comment(id uint) (*model.Comment, []error) {
	cmnt, errs := cs.commentRepo.Comment(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return cmnt, errs
}

// UpdateComment updates a given comment
func (cs *CommentService) UpdateComment(comment *model.Comment) (*model.Comment, []error) {
	cmnt, errs := cs.commentRepo.UpdateComment(comment)
	if len(errs) > 0 {
		return nil, errs
	}
	return cmnt, errs
}

// DeleteComment deletes a given comment
func (cs *CommentService) DeleteComment(id uint) (*model.Comment, []error) {
	cmnt, errs := cs.commentRepo.DeleteComment(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return cmnt, errs
}

// StoreComment stores a given comment
func (cs *CommentService) StoreComment(comment *model.Comment) (*model.Comment, []error) {
	cmnt, errs := cs.commentRepo.StoreComment(comment)
	if len(errs) > 0 {
		return nil, errs
	}
	return cmnt, errs
}
func (cs *CommentService) RetrieveComments(movieid uint) ([]model.Comment, []error) {
	cmnts, errs := cs.commentRepo.RetrieveComments(movieid)
	if len(errs) > 0 {
		return nil, errs
	}
	return cmnts, errs
}
