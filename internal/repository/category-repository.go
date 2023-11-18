package repository

import (
	"final-project-03/internal/database"
	"final-project-03/internal/helper"
	"final-project-03/internal/model"
)

type categoryModelRepo interface {
	CreateCategory(*model.Category) (*model.Category, helper.Error)
	UpdateCategory(*model.CategoryUpdate, uint) (*model.Category, helper.Error)
	GetAllCategories() ([]*model.Category, helper.Error)
	DeleteCategory(uint) helper.Error
}

type categoryModel struct{}

var CategoryModel categoryModelRepo = &categoryModel{}

func (t *categoryModel) CreateCategory(category *model.Category) (*model.Category, helper.Error) {
	db := database.GetDB()

	err := db.Create(&category).Error

	if err != nil {
		return nil, helper.ParseError(err)
	}

	return category, nil
}


func (t *categoryModel) UpdateCategory(update *model.CategoryUpdate, categoryId uint) (*model.Category, helper.Error) {
	db := database.GetDB()

	var category model.Category
	err := db.First(&category, categoryId).Error

	if err != nil {
		return nil, helper.ParseError(err)
	}

	db.Model(&category).Updates(update)

	return &category, nil
}

func (c *categoryModel) GetAllCategories() ([]*model.Category, helper.Error) {
	db := database.GetDB()
	var categories []*model.Category

	err := db.Preload("Tasks").Find(&categories).Error

	if err != nil {
		return nil, helper.ParseError(err)
	}

	return categories, nil
}

func (c *categoryModel) DeleteCategory(categoryId uint) helper.Error {
	db := database.GetDB()
	var category model.Category

	err := db.First(&category, categoryId).Error

	if err != nil {
		return helper.ParseError(err)
	}

	db.Delete(&category)

	return nil
}
