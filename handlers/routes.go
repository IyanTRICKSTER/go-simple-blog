package handlers

import (
	"github.com/gofiber/fiber/v2"
	"go-simple-blog/contracts"
	"go-simple-blog/handlers/middleware"
	jwtUtils "go-simple-blog/utils/jwt"
)

func NewRoutes(
	app *fiber.App,
	userUsecase contracts.IUserUsecase,
	postUsecase contracts.IPostUsecase) {

	userController := NewUserController(userUsecase)
	postController := NewPostController(postUsecase)

	jwt := jwtUtils.New()

	//User Controller
	v1 := app.Group("v1")
	v1.Post("/login", userController.Login)

	//Post Controller
	v1.Get("/posts/s/", postController.SearchPost)
	v1.Get("/posts/:postID", postController.ShowPost)
	v1.Get("/posts", postController.Fetch)
	v1.Post("/posts", middleware.AuthGuardMiddleware(jwt), postController.CreatePost)
	v1.Patch("/posts/:postID", middleware.AuthGuardMiddleware(jwt), postController.UpdatePost)
	v1.Delete("/posts/:postID", middleware.AuthGuardMiddleware(jwt), postController.DeletePost)
}
