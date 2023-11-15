package service

import (
	"final-project-03/internal/helper"
	"final-project-03/internal/model"
	"final-project-03/internal/repository"

	"github.com/asaskevich/govalidator"
)

type categoryServiceRepo interface {
	CreateCategory(*model.Category) (*model.Category, helper.Error)
	UpdateCategory(*model.CategoryUpdate, uint) (*model.Category, helper.Error)
}

type categoryService struct{}

var CategoryService categoryServiceRepo = &categoryService{}

func (t *categoryService) CreateCategory(category *model.Category) (*model.Category, helper.Error) {
	if _, err := govalidator.ValidateStruct(category); err != nil {
		return nil, helper.BadRequest(err.Error())
	}

	result, err := repository.CategoryModel.CreateCategory(category)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (t *categoryService) UpdateCategory(category *model.CategoryUpdate, categoryId uint) (*model.Category, helper.Error) {
	if _, err := govalidator.ValidateStruct(category); err != nil {
		return nil, helper.BadRequest(err.Error())
	}

	result, err := repository.CategoryModel.UpdateCategory(category, categoryId)
	if err != nil {
		return nil, err
	}

	return result, nil
}
