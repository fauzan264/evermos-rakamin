package services

import (
	"time"

	"github.com/fauzan264/evermos-rakamin/domain/dto/request"
	"github.com/fauzan264/evermos-rakamin/domain/dto/response"
	"github.com/fauzan264/evermos-rakamin/domain/model"
	"github.com/fauzan264/evermos-rakamin/repositories"
)

type CategoryService interface {
	GetListCategory() ([]response.CategoryResponse, error)
	GetDetailCategory(requestID request.GetByCategoryIDRequest) (response.CategoryResponse, error)
	CreateCategory(request request.CategoryRequest) (response.CategoryResponse, error)
	UpdateCategory(requestID request.GetByCategoryIDRequest, requestData request.CategoryRequest) (response.CategoryResponse, error)
	DeleteCategory(request request.GetByCategoryIDRequest) error
}

type categoryService struct {
	repository repositories.CategoryRepository
}

func NewCategoryService(repository repositories.CategoryRepository) *categoryService {
	return &categoryService{repository}
}


func (s *categoryService) GetListCategory() ([]response.CategoryResponse, error) {
	categories, err := s.repository.GetListCategory()
	if err != nil {
		return []response.CategoryResponse{}, err
	}

	listCategoryResponse := response.ListCategoryResponseFormatter(categories)
	
	return listCategoryResponse, nil
}

func (s *categoryService) GetDetailCategory(request request.GetByCategoryIDRequest) (response.CategoryResponse, error) {
	category, err := s.repository.GetCategoryByID(request.ID)
	if err != nil {
		return response.CategoryResponse{}, err
	}

	categoryResponse := response.CategoryResponseFormatter(category)

	return categoryResponse, nil
}

func (s *categoryService) CreateCategory(request request.CategoryRequest) (response.CategoryResponse, error) {
	category := model.Category{
		NamaCategory: request.NamaCategory,
		CreatedAt: time.Now(),
	}

	createCategory, err := s.repository.CreateCategory(category)
	if err != nil {
		return response.CategoryResponse{}, err
	}

	categoryResponse := response.CategoryResponseFormatter(createCategory)
	
	return categoryResponse, nil
}

func (s *categoryService) UpdateCategory(requestID request.GetByCategoryIDRequest, requestData request.CategoryRequest) (response.CategoryResponse, error) {
	category, err := s.repository.GetCategoryByID(requestID.ID)
	if err != nil {
		return response.CategoryResponse{}, err
	}

	category.ID = requestID.ID
	category.NamaCategory = requestData.NamaCategory
	category.UpdatedAt = time.Now()

	updateCategory, err := s.repository.UpdateCategory(category)
	if err != nil {
		return response.CategoryResponse{}, err
	}

	categoryResponse := response.CategoryResponseFormatter(updateCategory)
	
	return categoryResponse, nil
}

func (s *categoryService) DeleteCategory(request request.GetByCategoryIDRequest) error {
	category, err := s.repository.GetCategoryByID(request.ID)
	if err != nil {
		return err
	}
	
	err = s.repository.DeleteCategory(category.ID)
	if err != nil {
		return err
	}

	return nil
}

