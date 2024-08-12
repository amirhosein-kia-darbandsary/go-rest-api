package comment

import (
	"context"
	"errors"
	"fmt"
)

var (
	fetchCommentError = errors.New("Can't Fetch Comments from the database")
)

type Store interface {
	GetComment(context.Context, string) (Comment, error)
}

// Comment - a reperesentation of a structure
// for our service
type Comment struct {
	ID     string
	Slug   string
	Body   string
	Author string
}

// Service - is the struct on which all our
// logic will be built on top of it
type Service struct {
	Store Store
}

// NewService - is the constructure and a
// pointer to a new service
func NewService(store Store) *Service {
	return &Service{
		Store: store,
	}
}

func (service *Service) GetComment(ctx context.Context, id string) (Comment, error) {
	fmt.Println("Retrieving the Comment from the database")
	cmt, err := service.Store.GetComment(ctx, id)
	if err != nil {
		fmt.Println("Error has Accured :", fetchCommentError)
	}
	return cmt, nil
}
