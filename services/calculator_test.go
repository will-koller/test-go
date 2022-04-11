package services_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/williamkoller/test-go/services"
)

func TestSum(t *testing.T) {
	require.Equal(t, services.Sum(2, 3), 5)
}

func TestMultiply(t *testing.T) {
	require.Equal(t, services.Multiply(2, 3), 6)
}
