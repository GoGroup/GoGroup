package comment

import "github.com/GoGroup/Movie-and-events/model"

// CommentRepository specifies customer comment related database operations
type CommentRepository interface {
	Comments() ([]model.Comment, []error)
	Comment(id uint) (*model.Comment, []error)
	UpdateComment(comment *model.Comment) (*model.Comment, []error)
	DeleteComment(id uint) (*model.Comment, []error)

	StoreComment(comment *model.Comment) (*model.Comment, []error)
	RetrieveComments(movieid uint) ([]model.Comment, []error)
}
