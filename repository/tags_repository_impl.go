package repository

import (
	"crud_store_gin/dto/request"
	"crud_store_gin/helper"
	"crud_store_gin/models"
	"errors"

	"gorm.io/gorm"
)

type TagsRepositoryImpl struct {
	Db *gorm.DB
}

func NewTagsRepositoryImpl(Db *gorm.DB) TagsRepository {
	return &TagsRepositoryImpl{Db: Db}
}

func (t *TagsRepositoryImpl) FindAll() []models.Tags {
	var tags []models.Tags
	result := t.Db.Find(&tags)
	helper.ErrorPanic(result.Error)
	return tags
}

func (t *TagsRepositoryImpl) FindById(tagsId int) (tags models.Tags, err error) {
	var tag models.Tags
	result := t.Db.First(&tag, tagsId)
	if result != nil {
		return tag, nil
	} else {
		return tag, errors.New("tag is not found")
	}
}

func (t *TagsRepositoryImpl) Save(tags models.Tags) {
	result := t.Db.Create(&tags)
	helper.ErrorPanic(result.Error)
}

func (t *TagsRepositoryImpl) Update(tags models.Tags) {

	var updateTag = request.UpdateTagsRequest{
		Id:   tags.Id,
		Name: tags.Name,
	}
	result := t.Db.Model(&tags).Updates(updateTag)
	helper.ErrorPanic(result.Error)

}

func (t *TagsRepositoryImpl) Delete(tagsId int) {
	var tags models.Tags
	result := t.Db.Where("id = ?", tagsId).Delete(&tags)
	helper.ErrorPanic(result.Error)
}
