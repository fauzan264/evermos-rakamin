package repositories

import (
	"github.com/fauzan264/evermos-rakamin/domain/model"
	"gorm.io/gorm"
)

type categoryRepository struct {
	db *gorm.DB
}

type CategoryRepository interface {
	GetListCategory() ([]model.Category, error)
	CreateCategory(category model.Category) (model.Category, error)
	GetCategoryByID(id int) (model.Category, error)
	UpdateCategory(category model.Category) (model.Category, error)
	DeleteCategory(id int) error
}

func NewCategoryRepository(db *gorm.DB) *categoryRepository {
	return &categoryRepository{db}
}

func (r *categoryRepository) GetListCategory() ([]model.Category, error) {
	var categories []model.Category

	err := r.db.Find(&categories).Error
	if err != nil {
		return categories, err
	}

	return categories, nil
}

func (r *categoryRepository) CreateCategory(category model.Category) (model.Category, error) {
	err := r.db.Create(&category).Error
	if err != nil {
		return category, err
	}

	return category, nil
}

func (r *categoryRepository) GetCategoryByID(id int) (model.Category, error) {
	var category model.Category
	err := r.db.Where("id = ?", id).First(&category).Error
	if err != nil {
		return category, err
	}

	return category, nil
}

func (r *categoryRepository) UpdateCategory(category model.Category) (model.Category, error) {
	err := r.db.Save(&category).Error
	if err != nil {
		return category, err
	}

	return category, nil
}

func (r *categoryRepository) DeleteCategory(id int) error {
	var category model.Category

	err := r.db.Delete(&category, id).Error
	if err != nil {
		return err
	}

	return nil
}