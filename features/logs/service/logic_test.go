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
			logDataMock := new(mocks.LogData_)
			logDataMock.On("Insert", mock.Anything).Return(v.Output.Result, v.Output.errResult)

			//execute
			logService := New(logDataMock)
			_, err := logService.Create(v.Input.logEntity, v.Input.LoggedInUserId)
			if v.Output.IsError {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestGetData(t *testing.T) {
	table := GetDataTestTable()
	for _, v := range table {
		t.Run(v.Name, func(t *testing.T) {
			//mock data
			logDataMock := new(mocks.LogData_)
			logDataMock.On("SelectData", mock.AnythingOfType("uint"), mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(v.Output.Result, v.Output.errResult)

			//execute
			logService := New(logDataMock)
			_, err := logService.GetData(v.Input.logEntity.MenteeID, v.Input.Limit, v.Input.Offset)
			if v.Output.IsError {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
