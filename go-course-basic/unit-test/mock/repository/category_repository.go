package repository

import "go-course-basic/unit-test/mock/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}
