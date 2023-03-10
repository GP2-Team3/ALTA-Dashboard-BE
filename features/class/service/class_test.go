package service

import (
	"alta-dashboard-be/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreate(t *testing.T) {
	table := CreateTestTable()
	for _, v := range table {
		t.Run(v.Name, func(t *testing.T) {
			//mock data
			classData := new(mocks.ClassData_)
			classData.On("Create", mock.Anything).Return(nil)

			//execute
			service := New(classData)
			err := service.Create(v.Input.UserID, v.Input.Class)
			if v.Output.IsError {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestUpdate(t *testing.T) {
	table := UpdateTestTable()
	for _, v := range table {
		t.Run(v.Name, func(t *testing.T) {
			//mock data
			classData := new(mocks.ClassData_)
			classData.On("Update", mock.AnythingOfType("int"), mock.Anything).Return(nil)

			//execute
			service := New(classData)
			err := service.Update(v.Input.UserID, v.Input.ID, v.Input.Class)
			if v.Output.IsError {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestDelete(t *testing.T) {
	table := DeleteTestTable()
	for _, v := range table {
		t.Run(v.Name, func(t *testing.T) {
			//mock data
			classData := new(mocks.ClassData_)
			classData.On("Delete", mock.AnythingOfType("int")).Return(nil)

			//execute
			service := New(classData)
			err := service.Delete(v.Input.UserID, v.Input.ID)
			if v.Output.IsError {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestGetAll(t *testing.T) {
	table := GetAllTestTable()
	for _, v := range table {
		t.Run(v.Name, func(t *testing.T) {
			//mock data
			classData := new(mocks.ClassData_)
			classData.On("GetAll", mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(v.Output.Result, nil)

			//execute
			service := New(classData)
			_, err := service.GetAll(v.Input.Page, v.Input.Limit)
			if v.Output.IsError {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestGetOne(t *testing.T) {
	table := GetOneTestTable()
	for _, v := range table {
		t.Run(v.Name, func(t *testing.T) {
			//mock data
			classData := new(mocks.ClassData_)
			classData.On("GetOne", mock.AnythingOfType("int")).Return(v.Output.Result, nil)

			//execute
			service := New(classData)
			_, err := service.GetOne(v.Input.ID)
			if v.Output.IsError {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
