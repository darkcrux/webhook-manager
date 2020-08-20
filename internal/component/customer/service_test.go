package customer

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {
	c := &Customer{}

	mockRepo := new(MockRepository)
	mockRepo.On("Save", c).Return(1, nil)

	s := NewDefaultService(mockRepo)
	id, err := s.Register(c)
	assert.Nil(t, err)
	assert.Equal(t, 1, id)
	assert.True(t, mockRepo.AssertCalled(t, "Save", c))

	// TODO add negative tests
}

func TestGet(t *testing.T) {
	successCustomer := &Customer{}

	mockRepo := new(MockRepository)
	mockRepo.On("GetByID", 1).Return(successCustomer, nil)
	mockRepo.On("GetByID", 2).Return(nil, errors.New("something went wrong"))

	s := NewDefaultService(mockRepo)

	c, err := s.Get(1)
	assert.Equal(t, c, successCustomer)
	assert.Nil(t, err)

	c2, err2 := s.Get(2)
	assert.Nil(t, c2)
	assert.NotNil(t, err2)

	assert.True(t, mockRepo.AssertNumberOfCalls(t, "GetByID", 2))
}
