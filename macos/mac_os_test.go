package macos

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockMacOS struct {
	mock.Mock
}

func (m *MockMacOS) GetName() (string, error) {
	args := m.Called()
	return args.Get(0).(string), args.Error(1)
}

func Test_GetName(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockMac := new(MockMacOS)
		mockMac.On("GetName", mock.Anything).Return("Mac OS User", nil)
		nameMock, errMock := mockMac.GetName()
		assert.Nil(t, errMock)
		assert.NotNil(t, nameMock)
		mockMac.AssertExpectations(t)
	})
	t.Run("error after getting name", func(t *testing.T) {
		mockMac := new(MockMacOS)
		mockMac.On("GetName").Return("", errors.New("error trying to get the macos user"))
		name, err := mockMac.GetName()
		assert.Equal(t, "", name, "name should be empty")
		assert.Error(t, err)
		mockMac.AssertExpectations(t)
	})
	t.Run("", func(t *testing.T) {

	})

}