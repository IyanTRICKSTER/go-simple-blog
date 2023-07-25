package usecases

import (
	"context"
	"go-simple-blog/contracts"
	"go-simple-blog/contracts/statusCodes"
	"go-simple-blog/entities"
	"go-simple-blog/requests"
	"go-simple-blog/responses"
)

type PostUsecase struct {
	postRepo     contracts.IPostRepository
	cloudStorage contracts.ICloudStorageRepo
}

func (p PostUsecase) FetchAll(ctx context.Context, page uint, perPage uint) responses.PostResponse {

	posts, totalData, statusCode, err := p.postRepo.FetchAll(ctx, page, perPage)
	if err != nil {
		return responses.New(responses.PostResponse{}, false, statusCode, err.Error(), nil)
	}
	return responses.New(responses.PostResponse{}, true, statusCode, "ok",
		map[string]interface{}{"posts": posts, "totalData": totalData, "page": page, "perPage": perPage})
}

func (p PostUsecase) Show(ctx context.Context, postID uint) responses.PostResponse {

	post, statusCode, err := p.postRepo.Find(ctx, postID)
	if err != nil {
		return responses.New(responses.PostResponse{}, false, statusCode, err.Error(), nil)
	}
	return responses.New(responses.PostResponse{}, true, statusCode, "ok", post)
}

func (p PostUsecase) Create(ctx context.Context, request requests.CreatePostRequest) responses.PostResponse {

	var newPost entities.Post
	newPost.UserID = AuthGuard(ctx).GetUserID()
	newPost.Content = request.Content
	newPost.Title = request.Title
	newPost.Slug = Slugify(request.Title + "-" + UnixMilliToStr())
	newPost.Image, _ = p.cloudStorage.UploadFile(ctx, request.Image, request.Title+UnixMilliToStr())

	post, statusCode, err := p.postRepo.Create(ctx, newPost)
	if err != nil {
		return responses.New(responses.PostResponse{}, false, statusCode, err.Error(), nil)
	}
	return responses.New(responses.PostResponse{}, true, statusCode, "ok", post)
}

func (p PostUsecase) Update(ctx context.Context, postID uint, request requests.UpdatePostRequest) responses.PostResponse {

	post, statusCode, err := p.postRepo.Find(ctx, postID)
	if err != nil {
		return responses.New(responses.PostResponse{}, false, statusCode, err.Error(), nil)
	}

	//Authorize request
	if post.UserID != AuthGuard(ctx).GetUserID() {
		return responses.New(responses.PostResponse{}, false, statusCodes.ErrForbidden, "Forbidden", nil)
	}

	//Delete old file
	go func() {
		_ = p.cloudStorage.DeleteFile(ctx, post.Image)
	}()

	post.Title = request.Title
	post.Slug = Slugify(request.Title + "-" + UnixMilliToStr())
	post.Content = request.Content
	post.Image, _ = p.cloudStorage.UploadFile(ctx, request.Image, request.Title+UnixMilliToStr())

	updatedPost, statusCode, err := p.postRepo.Update(ctx, postID, post)
	if err != nil {
		return responses.New(responses.PostResponse{}, false, statusCode, err.Error(), nil)
	}

	return responses.New(responses.PostResponse{}, true, statusCode, "ok", updatedPost)
}

func (p PostUsecase) Delete(ctx context.Context, postID uint) responses.PostResponse {

	post, statusCode, err := p.postRepo.Find(ctx, postID)
	if err != nil {
		return responses.New(responses.PostResponse{}, false, statusCode, err.Error(), nil)
	}

	//Authorize request
	if post.UserID != AuthGuard(ctx).GetUserID() {
		return responses.New(responses.PostResponse{}, false, statusCodes.ErrForbidden, "Forbidden", nil)
	}

	go func() {
		_ = p.cloudStorage.DeleteFile(ctx, post.Image)
	}()

	statusCode, err = p.postRepo.Delete(ctx, postID)
	if err != nil {
		return responses.New(responses.PostResponse{}, false, statusCode, err.Error(), nil)
	}

	return responses.New(responses.PostResponse{}, true, statusCode, "ok", nil)
}

func (p PostUsecase) Search(ctx context.Context, keyword string, page uint, perPage uint) responses.PostResponse {

	posts, totalData, statusCode, err := p.postRepo.Search(ctx, keyword, page, perPage)
	if err != nil {
		return responses.New(responses.PostResponse{}, false, statusCode, err.Error(), posts)
	}
	return responses.New(responses.PostResponse{}, true, statusCode, "ok",
		map[string]interface{}{"posts": posts, "totalData": totalData, "page": page, "perPage": perPage})
}

func NewPostUsecase(postRepo contracts.IPostRepository, cloudStorage *contracts.ICloudStorageRepo) contracts.IPostUsecase {
	return &PostUsecase{postRepo: postRepo, cloudStorage: *cloudStorage}
}
