package repositories

import (
	"context"
	"go-simple-blog/contracts"
	"go-simple-blog/contracts/statusCodes"
	"go-simple-blog/entities"
	"gorm.io/gorm"
)

type UserRepo struct {
	conn *gorm.DB
}

func (u UserRepo) Find(ctx context.Context, email string) (post entities.User, status statusCodes.StatusCode, err error) {
	err = u.conn.WithContext(ctx).Where("email = ?", email).First(&post).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return post, statusCodes.ModelNotFound, err
		}
		return post, statusCodes.Error, err
	}

	return post, statusCodes.Success, nil
}

func NewUserRepo(conn *gorm.DB) contracts.IUserRepository {
	return &UserRepo{conn: conn}
}
