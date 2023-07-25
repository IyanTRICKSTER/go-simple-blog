package contracts

import (
	"context"
	"go-simple-blog/requests"
	"go-simple-blog/responses"
)

type IPostUsecase interface {
	FetchAll(ctx context.Context, page uint, perPage uint) responses.PostResponse
	Show(ctx context.Context, postID uint) responses.PostResponse
	Create(ctx context.Context, request requests.CreatePostRequest) responses.PostResponse
	Update(ctx context.Context, postID uint, request requests.UpdatePostRequest) responses.PostResponse
	Delete(ctx context.Context, postID uint) responses.PostResponse
	Search(ctx context.Context, keyword string, page uint, perPage uint) responses.PostResponse
}

type IUserUsecase interface {
	Login(ctx context.Context, request requests.LoginRequest) responses.UserResponse
}
