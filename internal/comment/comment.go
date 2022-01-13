package comment

import "github.com/jinzhu/gorm"

//Service - struct for our comment service
type Service struct {
	DB *gorm.DB
}

//Comment - defins our comment structure
type Comment struct {
	gorm.Model
	Slug   string
	Body   string
	Author string
}

//CommentsSerice - interface for our comment service
type CommentsSerice interface {
	GetComment(ID uint) (Comment, error)
	GetCommentBySlug(Slug string) ([]Comment, error)
	PostComment(comment Comment) (Comment, error)
	UpdateComment(ID uint, comment Comment) (Comment, error)
	DeleteCommen(ID uint) error
	GetAllComments() ([]Comment, error)
}

//NewService - returns new comment service
func newService(db *gorm.DB) *Service {
	return &Service{
		DB: db,
	}
}
