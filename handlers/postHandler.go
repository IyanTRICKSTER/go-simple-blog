package handlers

import (
	"github.com/gofiber/fiber/v2"
	"go-simple-blog/contracts"
	"go-simple-blog/contracts/statusCodes"
	"go-simple-blog/requests"
	"go-simple-blog/responses"
	"log"
	"strconv"
)

type PostController struct {
	postUsecase contracts.IPostUsecase
	validator   XValidator
}

func (p PostController) Fetch(c *fiber.Ctx) error {

	page := c.Query("page")
	perPage := c.Query("perPage")

	pageInt, err := strconv.ParseUint(page, 10, 64)
	if err != nil {
		c.SendStatus(500)
		return c.JSON(
			responses.New(
				responses.UserResponse{},
				false,
				statusCodes.Error,
				err.Error(),
				nil).ToMap())
	}

	perPageInt, err := strconv.ParseUint(perPage, 10, 64)
	if err != nil {
		c.SendStatus(500)
		return c.JSON(
			responses.New(
				responses.UserResponse{},
				false,
				statusCodes.Error,
				err.Error(),
				nil).ToMap())
	}

	res := p.postUsecase.FetchAll(c.Context(), uint(pageInt), uint(perPageInt))

	if res.IsFailed() {
		c.SendStatus(400)
		return c.JSON(res.ToMap())
	}

	c.SendStatus(200)
	return c.JSON(res.ToMap())
}

func (p PostController) SearchPost(c *fiber.Ctx) error {

	page := c.Query("page")
	perPage := c.Query("perPage")
	keyword := c.Query("keyword")
	log.Println(keyword)
	pageInt, err := strconv.ParseUint(page, 10, 64)
	if err != nil {
		c.SendStatus(500)
		return c.JSON(
			responses.New(
				responses.UserResponse{},
				false,
				statusCodes.Error,
				err.Error(),
				nil).ToMap())
	}

	perPageInt, err := strconv.ParseUint(perPage, 10, 64)
	if err != nil {
		c.SendStatus(500)
		return c.JSON(
			responses.New(
				responses.UserResponse{},
				false,
				statusCodes.Error,
				err.Error(),
				nil).ToMap())
	}

	res := p.postUsecase.Search(c.Context(), keyword, uint(pageInt), uint(perPageInt))

	if res.IsFailed() {
		c.SendStatus(400)
		return c.JSON(res.ToMap())
	}

	c.SendStatus(200)
	return c.JSON(res.ToMap())
}

func (p PostController) ShowPost(c *fiber.Ctx) error {

	postID, err := strconv.ParseUint(c.Params("postID"), 10, 64)
	if err != nil {
		c.SendStatus(500)
		return c.JSON(
			responses.New(
				responses.UserResponse{},
				false,
				statusCodes.Error,
				err.Error(),
				nil).ToMap())
	}

	res := p.postUsecase.Show(c.Context(), uint(postID))

	if res.IsFailed() {
		c.SendStatus(400)
		return c.JSON(res.ToMap())
	}

	c.SendStatus(200)
	return c.JSON(res.ToMap())
}

func (p PostController) CreatePost(c *fiber.Ctx) error {

	var req requests.CreatePostRequest
	req.Image, _ = c.FormFile("image")
	err := c.BodyParser(&req)
	if err != nil {
		c.SendStatus(500)
		return c.JSON(
			responses.New(
				responses.UserResponse{},
				false,
				statusCodes.Error,
				err.Error(),
				nil).ToMap())
	}

	//Validation
	if msg, ok := p.validator.ValidateWithMessage(req); !ok {
		c.SendStatus(400)
		return c.JSON(
			responses.New(
				responses.UserResponse{},
				false,
				statusCodes.Error,
				msg,
				nil).ToMap())
	}

	//Dispatch usecase
	res := p.postUsecase.Create(c.Context(), req)
	if res.IsFailed() {
		c.SendStatus(400)
		return c.JSON(res.ToMap())
	}

	c.SendStatus(200)
	return c.JSON(res.ToMap())
}

func (p PostController) UpdatePost(c *fiber.Ctx) error {

	postID, err := strconv.ParseUint(c.Params("postID"), 10, 64)
	if err != nil {
		c.SendStatus(500)
		return c.JSON(
			responses.New(
				responses.UserResponse{},
				false,
				statusCodes.Error,
				err.Error(),
				nil).ToMap())
	}

	var req requests.UpdatePostRequest
	req.Image, _ = c.FormFile("image")
	err = c.BodyParser(&req)
	if err != nil {
		c.SendStatus(500)
		return c.JSON(
			responses.New(
				responses.PostResponse{},
				false,
				statusCodes.Error,
				err.Error(),
				nil).ToMap())
	}

	//Validation
	if msg, ok := p.validator.ValidateWithMessage(req); !ok {
		c.SendStatus(400)
		return c.JSON(
			responses.New(
				responses.PostResponse{},
				false,
				statusCodes.Error,
				msg,
				nil).ToMap())
	}

	//Dispatch usecase
	res := p.postUsecase.Update(c.Context(), uint(postID), req)
	if res.IsFailed() {
		c.SendStatus(400)
		return c.JSON(res.ToMap())
	}

	c.SendStatus(200)
	return c.JSON(res.ToMap())

}

func (p PostController) DeletePost(c *fiber.Ctx) error {

	postID, err := strconv.ParseUint(c.Params("postID"), 10, 64)
	if err != nil {
		c.SendStatus(500)
		return c.JSON(
			responses.New(
				responses.UserResponse{},
				false,
				statusCodes.Error,
				err.Error(),
				nil).ToMap())
	}

	res := p.postUsecase.Delete(c.Context(), uint(postID))
	if res.IsFailed() {
		c.SendStatus(400)
		return c.JSON(res.ToMap())
	}

	c.SendStatus(200)
	return c.JSON(res.ToMap())
}

func NewPostController(postUsecase contracts.IPostUsecase) *PostController {
	return &PostController{postUsecase: postUsecase}
}
