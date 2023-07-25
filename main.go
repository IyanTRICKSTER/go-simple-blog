package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go-simple-blog/database"
	"go-simple-blog/entities"
	"go-simple-blog/handlers"
	"go-simple-blog/repositories"
	"go-simple-blog/usecases"
	bcryptUtils "go-simple-blog/utils/bcrypt"
	jwtUtils "go-simple-blog/utils/jwt"
	"log"
	"os"
)

func main() {

	//Load .env file
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file")
	}

	db := database.Database{
		Host:     os.Getenv("DB_HOST"),
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DbName:   os.Getenv("DB_NAME"),
		DbPort:   os.Getenv("DB_PORT"),
	}

	err := db.Connect()
	if err != nil {
		log.Fatalf("DB connection error: %v", err.Error())
	}

	//Run Migrations
	_ = db.DropTable(entities.User{})
	_ = db.DropTable(entities.Post{})
	_ = db.MigrateTable(entities.Post{})
	_ = db.MigrateTable(entities.User{})

	//Run Seeders
	_ = db.Seed(database.UserSeedData())

	//Setup Cloud Storage Service
	wd, _ := os.Getwd()
	cloudStorage := repositories.NewGCStorageService("ayocode1-bucker", wd+"/keys.json")
	err = cloudStorage.Connect()
	if err != nil {
		panic(err)
	}

	//Setup User Service
	userRepo := repositories.NewUserRepo(db.GetConnection())
	userUsecase := usecases.NewUserUsecase(userRepo, bcryptUtils.NewHashFunction(), jwtUtils.New())

	//Setup Post Service
	postRepo := repositories.NewPostRepo(db.GetConnection())
	postUsecase := usecases.NewPostUsecase(postRepo, &cloudStorage)

	//Setup Handler
	app := fiber.New()
	handlers.NewRoutes(app, userUsecase, postUsecase)

	err = app.Listen("127.0.0.1:8097")
	if err != nil {
		panic(err)
	}
}
