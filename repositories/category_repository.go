package repositories

import (
	"github.com/fauzan264/evermos-rakamin/domain/model"
	"gorm.io/gorm"
)

type categoryRepository struct {
	db *gorm.DB
}

type CategoryRepository interface {
	CreateCategory(category model.Category) (model.Category, error)
	GetCategoryByID(id int) (model.Category, error)
	UpdateCategory(category model.Category) (model.Category, error)
	DeleteCategory(id int) error
}

func NewCategoryRepository(db *gorm.DB) *categoryRepository {
	return &categoryRepository{db}
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
	err := r.db.Where("id = ?", id).Find(&category).Error
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

	err := r.db.Where("id = ?", id).Delete(&category).Error
	if err != nil {
		return err
	}

	return nil
}