package services_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"main.go/services"
)

func TestAssign(t *testing.T) {
	handler := services.NewRecursiveAssign(200, 300, 400)
	a, b, c, err := handler.Assign(400)

	assert.NoError(t, err)
	assert.Greater(t, (a + b + c), int32(0))
}

func TestAssignWithRandomOrderInCreditTypes(t *testing.T) {
	handler := services.NewRecursiveAssign(200, 100, 400)
	a, b, c, err := handler.Assign(400)

	assert.NoError(t, err)
	assert.Greater(t, (a + b + c), int32(0))
}

func TestAssignWhenCombinationIsDuplicated(t *testing.T) {
	handler := services.NewRecursiveAssign(300, 500, 700)
	a, b, c, err := handler.Assign(3000)

	assert.NoError(t, err)
	assert.Greater(t, (a + b + c), int32(0))
}

func TestAssignWithTheMinValue(t *testing.T) {
	handler := services.NewRecursiveAssign(200, 300, 400)
	a, b, c, err := handler.Assign(200)

	assert.NoError(t, err)
	assert.Greater(t, (a + b + c), int32(0))
}

func TestAssignWithInvalidInvestment(t *testing.T) {
	handler := services.NewRecursiveAssign(200, 300, 400)
	a, b, c, err := handler.Assign(1)

	assert.Error(t, err)
	assert.Equal(t, (a + b + c), int32(0))
}

func TestAssignWhenIsNotPossible(t *testing.T) {
	handler := services.NewRecursiveAssign(200, 800, 900)
	a, b, c, err := handler.Assign(300)

	assert.Error(t, err)
	assert.Equal(t, (a + b + c), int32(0))
}

func TestAssignWhenIsNotAMultipleOf100(t *testing.T) {
	handler := services.NewRecursiveAssign(200, 800, 900)
	a, b, c, err := handler.Assign(201)

	assert.Error(t, err)
	assert.Equal(t, (a + b + c), int32(0))
}
