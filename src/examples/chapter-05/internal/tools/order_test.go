package tools

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateOrderAmount(t *testing.T) {
	// Test valid amounts
	assert.NoError(t, validateOrderAmount(500.00))
	assert.NoError(t, validateOrderAmount(0))
	assert.NoError(t, validateOrderAmount(10000.00))

	// Test amount over limit
	err := validateOrderAmount(10000.01)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "exceeds maximum limit")

	// Test negative amount
	err = validateOrderAmount(-1.00)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "cannot be negative")
}