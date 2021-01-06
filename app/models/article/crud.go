package article

import (
	"goblog2/pkg/logger"
	"goblog2/pkg/model"
	"goblog2/pkg/types"
)

func Get(idstr string) (Article, error) {
	var article Article

	id := types.StringToInt(idstr)
	if err := model.DB.First(&article, id).Error; err != nil {
		return article, err
	}
	return article, nil
}

func GetAll() ([]Article, error) {
	var articles []Article
	if err := model.DB.Find(&articles).Error; err != nil {
		return articles, err
	}
	return articles, nil
}

func (article *Article) Create() (err error) {
	if err = model.DB.Create(&article).Error; err != nil {
		logger.LogError(err)
		return err
	}
	return nil
}
