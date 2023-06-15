package repository

import (
	"a21hc3NpZ25tZW50/model"
	

	"gorm.io/gorm"
)

type CourseRepository interface {
	FetchByID(id int) (*model.Course, error)
	Store(course *model.Course) error
	Delete(id int) error
}

type courseRepository struct {
	db *gorm.DB
}

func NewCourseRepo(db *gorm.DB) *courseRepository {
	return &courseRepository{db}
}

func (c *courseRepository) FetchByID(id int) (*model.Course, error) {
	var course model.Course
	err := c.db.Where("id = ?", id).First(&course).Error
	if err != nil {
		return nil, err
	}

	return &course, nil
}

func (c *courseRepository) Store(course *model.Course) error {
	err := c.db.Create(course).Error
	if err != nil {
		return err
	}

	return nil
}

func (c *courseRepository) Delete(id int) error {
	var course model.Course
	result := c.db.Where("id = ?", id).First(&course)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	err := c.db.Delete(&course).Error
	if err != nil {
		return err
	}

	return nil
}
