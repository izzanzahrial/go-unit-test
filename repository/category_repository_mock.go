package repository

import (
	entity "go-unit-test/category"

	"github.com/stretchr/testify/mock"
)

// Mock object using Testify Mock
// used for compclicated testing where you need data from third party service, API call, or database
// by using mock you dont need the data from the example above
// https://pkg.go.dev/github.com/stretchr/testify/mock

type CategoryRepositoryMock struct {
	Mock mock.Mock
}

func (repository *CategoryRepositoryMock) FindById(id string) *entity.Category {
	arguments := repository.Mock.Called(id)
	if arguments.Get(0) == nil {
		return nil
	} else {
		category := arguments.Get(0).(entity.Category)
		return &category
	}
}
