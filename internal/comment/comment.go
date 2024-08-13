package comment

import (
	"context"
	"errors"
	"fmt"
)

var (
	fetchCommentError = errors.New("cannot fetch comments from the database")
)

type Store interface {
	GetComment(context.Context, string) (Comment, error)
	PostComment(context.Context, Comment) error
	UpdateComment(context.Context, string, Comment) error
	DeleteComment(context.Context, string) error
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

func (service *Service) PostComment(ctx context.Context, cmt Comment) error {
	fmt.Println("Posting the Comment to the database")
	err := service.Store.PostComment(ctx, cmt)
	if err != nil {
		return err
	}
	return nil
}

func (service *Service) UpdateComment(ctx context.Context, id string, cmt Comment) error {
	fmt.Println("Updating the Comment....")
	err := service.Store.UpdateComment(ctx, id, cmt)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}

func (service *Service) DeleteComment(ctx context.Context, id string) {
	fmt.Println("Deleteing the Comment ....")
	err := service.Store.DeleteComment(ctx, id)
	if err != nil {
		fmt.Println("Can't delete the message the uuid : ", id)
	}
}
