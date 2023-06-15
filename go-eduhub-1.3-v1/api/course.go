package api

import (
	"a21hc3NpZ25tZW50/model"
	repo "a21hc3NpZ25tZW50/repository"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CourseAPI interface {
	AddCourse(c *gin.Context)
	DeleteCourse(c *gin.Context)
}

type courseAPI struct {
	courseRepo repo.CourseRepository
}

func NewCourseAPI(courseRepo repo.CourseRepository) *courseAPI {
	return &courseAPI{courseRepo}
}

func (cr *courseAPI) AddCourse(c *gin.Context) {
	var newCourse model.Course
	if err := c.ShouldBindJSON(&newCourse); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
		return
	}

	err := cr.courseRepo.Store(&newCourse)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse{Message: "add course success"})
}

func (cr *courseAPI) DeleteCourse(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "Invalid course ID"})
		return
	}

	err = cr.courseRepo.Delete(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, model.ErrorResponse{Error: "Course not found"})
		} else {
			c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse{Message: "course delete success"})
}
