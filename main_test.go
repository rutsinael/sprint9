package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMaximumNegative(t *testing.T) {
	cases := []struct {
		input    []int
		expected int
	}{
		{[]int{}, 0},
		{nil, 0},
	}
	for _, c := range cases {
		require.Equal(t, c.expected, maximum(c.input))
	}
}

func TestMaximumPositive(t *testing.T) {
	cases := []struct {
		input    []int
		expected int
	}{
		{[]int{5, 5, 5, 5}, 5},
		{[]int{100}, 100},
		{[]int{8, 99, 88, 7, 0, 1, 100, 110, 120, 130}, 130},
	}
	for _, c := range cases {
		require.Equal(t, c.expected, maximum(c.input))
	}
}

func TestGenerateRandomElementsNegative(t *testing.T) {
	input := []int{
		0, -100,
	}
	for _, i := range input {
		require.Empty(t, generateRandomElements(i))
	}
}

func TestGenerateRandomElementsPositive(t *testing.T) {
	input := generateRandomElements(10)
	require.Equal(t, 10, len(input))
	for i := range input {
		require.Positive(t, input[i])
	}
}
