package comment

import "github.com/GoGroup/Movie-and-events/model"

// CommentService specifies customer comment related service
type CommentService interface {
	Comments() ([]model.Comment, []error)
	Comment(id uint) (*model.Comment, []error)
	UpdateComment(comment *model.Comment) (*model.Comment, []error)
	DeleteComment(id uint) (*model.Comment, []error)
	StoreComment(comment *model.Comment) (*model.Comment, []error)
}
