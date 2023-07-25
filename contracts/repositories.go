package contracts

import (
	"context"
	"go-simple-blog/contracts/statusCodes"
	"go-simple-blog/entities"
)

type IPostRepository interface {
	FetchAll(ctx context.Context, page uint, perPage uint) (posts []entities.Post, totalData uint, status statusCodes.StatusCode, err error)
	Search(ctx context.Context, keyword string, page uint, perPage uint) (posts []entities.Post, totalData uint, status statusCodes.StatusCode, err error)
	Find(ctx context.Context, postID uint) (post entities.Post, status statusCodes.StatusCode, err error)
	Create(ctx context.Context, model entities.Post) (post entities.Post, status statusCodes.StatusCode, err error)
	Update(ctx context.Context, postID uint, model entities.Post) (post entities.Post, status statusCodes.StatusCode, err error)
	Delete(ctx context.Context, postID uint) (status statusCodes.StatusCode, err error)
}

type IUserRepository interface {
	Find(ctx context.Context, email string) (post entities.User, status statusCodes.StatusCode, err error)
}
