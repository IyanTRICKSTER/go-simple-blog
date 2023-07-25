package dataSource

import (
	"go-simple-blog/entities"
	bcryptUtils "go-simple-blog/utils/bcrypt"
)

func UserSeedData() []entities.User {
	return []entities.User{
		{
			Username:   "Iyan",
			Email:      "iyan@gmail.com",
			Password:   bcryptUtils.NewHashFunction().Hash("iyan12345"),
			ProfilePic: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcT-M4Qk9egGf5MlnRBMWtCwZ4XML7hixAff26Q6rmRm&s",
			DeletedAt:  nil,
			Posts:      nil,
		},
	}
}
