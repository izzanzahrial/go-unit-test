package repository

import entity "go-unit-test/category"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}
