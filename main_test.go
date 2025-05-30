package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMaximum(t *testing.T) {
	cases := []struct {
		input    []int
		expected int
	}{
		{[]int{}, 0},
		{nil, 0},
		{[]int{5, 5, 5, 5}, 5},
		{[]int{100}, 100},
		{[]int{8, 99, 88, 7, 0, 1, 100, 110, 120, 130}, 130},
	}
	for _, c := range cases {
		require.Equal(t, c.expected, maximum(c.input))
	}
}

func TestGenerateRandomElements(t *testing.T) {
	cases := []struct {
		size        int
		expectedLen int
	}{
		{0, 0},
		{-100, 0},
		{10, 10},
	}

	for _, c := range cases {
		input := generateRandomElements(c.size)
		require.Equal(t, c.expectedLen, len(input))
		for i := range input {
			require.Positive(t, input[i])
		}
	}
}
