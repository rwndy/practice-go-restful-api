package service

import (
	"Users/riwandi/Documents/practice/go-restful-api/helper"
	"Users/riwandi/Documents/practice/go-restful-api/model/domain"
	"Users/riwandi/Documents/practice/go-restful-api/model/web"
	"Users/riwandi/Documents/practice/go-restful-api/repository"
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
)

type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository
	DB                 *sql.DB //connection to DB
	Validate           *validator.Validate
}

func NewCategoryService(categoryRepository repository.CategoryRepository, DB *sql.DB, validate *validator.Validate) CategoryService {
	return &CategoryServiceImpl{
		CategoryRepository: categoryRepository,
		DB:                 DB,
		Validate:           validate,
	}
}

func (service CategoryServiceImpl) Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse {
	//validation
	err := service.Validate.Struct(request)
	helper.HandlePanic(err)

	tx, err := service.DB.Begin()
	helper.HandlePanic(err)
	defer helper.HandleTx(tx)

	// body request
	category := domain.Category{
		Name: request.Name,
	}

	category = service.CategoryRepository.Save(ctx, tx, category)

	return helper.ToCategoryResponse(category)

}

func (service CategoryServiceImpl) Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse {
	//validation
	err := service.Validate.Struct(request)
	helper.HandlePanic(err)

	tx, err := service.DB.Begin()
	helper.HandlePanic(err)
	defer helper.HandleTx(tx)

	// check category is not empty
	category, err := service.CategoryRepository.FindById(ctx, tx, request.Id)
	helper.HandlePanic(err)

	category.Name = request.Name
	category = service.CategoryRepository.Update(ctx, tx, category)

	return helper.ToCategoryResponse(category)
}

func (service CategoryServiceImpl) Delete(ctx context.Context, categoryId int) {
	tx, err := service.DB.Begin()
	helper.HandlePanic(err)
	defer helper.HandleTx(tx)

	// check category is not empty
	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	helper.HandlePanic(err)

	service.CategoryRepository.Delete(ctx, tx, category)
}

func (service CategoryServiceImpl) FindById(ctx context.Context, categoryId int) web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.HandlePanic(err)
	defer helper.HandleTx(tx)

	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	helper.HandlePanic(err)

	return helper.ToCategoryResponse(category)
}

func (service CategoryServiceImpl) FindAll(ctx context.Context) []web.CategoryResponse {

	tx, err := service.DB.Begin()
	helper.HandlePanic(err)
	defer helper.HandleTx(tx)

	categories := service.CategoryRepository.FindAll(ctx, tx)

	return helper.ToCategoriesResponse(categories)
}
