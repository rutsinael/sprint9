package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMaximumShouldReturnZeroWhenEmptyArray(t *testing.T) {
	input := []int{}
	require.Equal(t, 0, maximum(input))
}

func TestMaximumShouldReturnZeroWhenNilArray(t *testing.T) {
	require.Equal(t, 0, maximum(nil))
}

func TestMaximumShouldCountCorrectlyWhenOneElementArray(t *testing.T) {
	input := []int{100}
	require.Equal(t, 100, maximum(input))
}

func TestMaximumShouldCountCorrectly(t *testing.T) {
	input := []int{100, 99, 88, 7, 0, 1, 100}
	require.Equal(t, 100, maximum(input))
}

func TestGenerateRandomElements(t *testing.T) {
	input := generateRandomElements(10)
	require.Equal(t, 10, len(input))
	for i := range input {
		require.Positive(t, input[i])
	}
}

func TestGenerateRandomElementsWhenZero(t *testing.T) {
	input := generateRandomElements(0)
	require.Empty(t, input)
}

func TestGenerateRandomElementsWhenNegative(t *testing.T) {
	input := generateRandomElements(-100)
	require.Empty(t, input)
}
