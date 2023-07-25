package repositories

import (
	"context"
	"go-simple-blog/contracts"
	"go-simple-blog/contracts/statusCodes"
	"go-simple-blog/entities"
	"gorm.io/gorm"
)

type PostRepo struct {
	conn *gorm.DB
}

func (p *PostRepo) FetchAll(ctx context.Context, page uint, perPage uint) (posts []entities.Post, totalData uint, status statusCodes.StatusCode, err error) {

	offset := (page - 1) * perPage
	var post []entities.Post
	p.conn.WithContext(ctx).Limit(int(perPage)).Offset(int(offset)).Find(&post)

	var counted int64
	p.conn.Model(entities.Post{}).Count(&counted)

	return post, uint(counted), statusCodes.Success, nil
}

func (p *PostRepo) Search(ctx context.Context, keyword string, page uint, perPage uint) (posts []entities.Post, totalData uint, status statusCodes.StatusCode, err error) {

	offset := (page - 1) * perPage

	var post []entities.Post
	p.conn.WithContext(ctx).Where("title LIKE ?", "%"+keyword+"%").Limit(int(perPage)).Offset(int(offset)).Find(&post)

	var counted int64
	p.conn.Model(entities.Post{}).Where("title LIKE ?", "%"+keyword+"%").Count(&counted)

	return post, uint(counted), statusCodes.Success, nil
}

func (p *PostRepo) Find(ctx context.Context, postID uint) (post entities.Post, status statusCodes.StatusCode, err error) {

	err = p.conn.WithContext(ctx).Where("id = ?", postID).First(&post).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return post, statusCodes.ModelNotFound, err
		}
		return post, statusCodes.Error, err
	}

	return post, statusCodes.Success, nil
}

func (p *PostRepo) Create(ctx context.Context, model entities.Post) (post entities.Post, status statusCodes.StatusCode, err error) {

	err = p.conn.WithContext(ctx).Create(&model).Error
	if err != nil {
		if err == gorm.ErrDuplicatedKey {
			return post, statusCodes.ErrDuplicatedModel, err
		}
		return post, statusCodes.Error, err
	}

	return model, statusCodes.Success, nil
}

func (p *PostRepo) Update(ctx context.Context, postID uint, model entities.Post) (post entities.Post, status statusCodes.StatusCode, err error) {

	err = p.conn.WithContext(ctx).Where("id = ?", postID).Save(&model).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return post, statusCodes.ModelNotFound, err
		}
		return post, statusCodes.Error, err
	}
	return model, statusCodes.Success, nil
}

func (p *PostRepo) Delete(ctx context.Context, postID uint) (status statusCodes.StatusCode, err error) {

	err = p.conn.WithContext(ctx).Where("id = ?", postID).Delete(&entities.Post{}).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return statusCodes.ModelNotFound, err
		}
		return statusCodes.Error, err
	}

	return statusCodes.Success, nil
}

func NewPostRepo(conn *gorm.DB) contracts.IPostRepository {
	return &PostRepo{conn: conn}
}
