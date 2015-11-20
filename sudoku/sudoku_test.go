package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRowTester(t *testing.T) {
	tests := []struct {
		puzzle   puzzle
		expected bool
		reason   string
	}{
		{
			puzzle{
				{1, 2, 3, 4, 5, 6, 7, 8, 9},
				{1, 2, 3, 4, 5, 6, 7, 8, 9},
				{1, 2, 3, 4, 5, 6, 7, 8, 9},
				{1, 2, 3, 4, 5, 6, 7, 8, 9},
				{1, 2, 3, 4, 5, 6, 7, 8, 9},
				{1, 2, 3, 4, 5, 6, 7, 8, 9},
				{1, 2, 3, 4, 5, 6, 7, 8, 9},
				{1, 2, 3, 4, 5, 6, 7, 8, 9},
				{1, 2, 3, 4, 5, 6, 7, 8, 9},
			},
			true,
			"Should pass because all rows are valid",
		},
		{
			puzzle{
				{1, 2, 3, 4, 5, 6, 7, 8, 9},
				{1, 2, 3, 4, 5, 6, 7, 8, 9},
				{1, 2, 3, 4, 5, 6, 7, 8, 9},
				{1, 2, 3, 4, 5, 6, 7, 8, 9},
				{1, 2, 3, 4, 5, 6, 7, 8, 9},
				{1, 2, 3, 4, 5, 6, 7, 8, 9},
				{1, 2, 3, 4, 5, 6, 7, 8, 9},
				{1, 2, 3, 4, 5, 6, 7, 8, 9},
				{1, 2, 3, 4, 5, 6, 7, 8, 1},
			},
			false,
			"Should fail because last row has 1 twice",
		},
		{
			puzzle{
				{1, 2, 3, 4, 5, 6, 7, 8, 9},
				{1, 2, 3, 4, 5, 6, 7, 8, 9},
			},
			true,
			"Should pass because when not all rows have been filled yet",
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.puzzle.testRows(), test.expected, test.reason)
	}
}

func TestColumnTester(t *testing.T) {
	tests := []struct {
		puzzle   puzzle
		expected bool
		reason   string
	}{
		{
			puzzle{
				{1, 1, 1, 1, 1, 1, 1, 1, 1},
				{2, 2, 2, 2, 2, 2, 2, 2, 2},
				{3, 3, 3, 3, 3, 3, 3, 3, 3},
				{4, 4, 4, 4, 4, 4, 4, 4, 4},
				{5, 5, 5, 5, 5, 5, 5, 5, 5},
				{6, 6, 6, 6, 6, 6, 6, 6, 6},
				{7, 7, 7, 7, 7, 7, 7, 7, 7},
				{8, 8, 8, 8, 8, 8, 8, 8, 8},
				{9, 9, 9, 9, 9, 9, 9, 9, 9},
			},
			true,
			"Should pass because all cols are valid",
		},
		{
			puzzle{
				{1, 1, 1, 1, 1, 1, 1, 1, 1},
				{2, 2, 2, 2, 2, 2, 2, 2, 2},
				{3, 3, 3, 3, 3, 3, 3, 3, 3},
				{4, 4, 4, 4, 4, 4, 4, 4, 4},
				{5, 5, 5, 5, 5, 5, 5, 5, 5},
				{6, 6, 6, 6, 6, 6, 6, 6, 6},
				{7, 7, 7, 7, 7, 7, 7, 7, 7},
				{8, 8, 8, 8, 8, 8, 8, 8, 8},
				{9, 9, 9, 9, 9, 9, 9, 9, 2},
			},
			false,
			"Should fail because last column has 2 twice",
		}, {
			puzzle{
				{1, 1, 1, 1, 1, 1, 1, 1, 1},
				{2, 2, 2, 2, 2, 2, 2, 2, 2},
				{3, 3, 3, 3, 3, 3, 3, 3, 3},
			},
			true,
			"Should pass when not completed yet but has no conflicts",
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.puzzle.testColumns(), test.expected, test.reason)
	}
}

func TestBoxTester(t *testing.T) {
	tests := []struct {
		puzzle   puzzle
		expected bool
		reason   string
	}{
		{
			puzzle{
				{1, 2, 3, 1, 2, 3, 1, 2, 3},
				{4, 5, 6, 4, 5, 6, 4, 5, 6},
				{7, 8, 9, 7, 8, 9, 7, 8, 9},
				{1, 2, 3, 1, 2, 3, 1, 2, 3},
				{4, 5, 6, 4, 5, 6, 4, 5, 6},
				{7, 8, 9, 7, 8, 9, 7, 8, 9},
				{1, 2, 3, 1, 2, 3, 1, 2, 3},
				{4, 5, 6, 4, 5, 6, 4, 5, 6},
				{7, 8, 9, 7, 8, 9, 7, 8, 9},
			},
			true,
			"Should pass because all boxes are valid",
		},
		{
			puzzle{
				{1, 2, 3, 1, 2, 3, 1, 2, 3},
				{4, 5, 6, 4, 5, 6, 4, 5, 6},
				{7, 8, 9, 7, 8, 9, 7, 8, 9},
				{1, 2, 3, 1, 2, 3, 1, 2, 3},
				{4, 5, 6, 4, 5, 6, 4, 5, 6},
				{7, 8, 9, 7, 8, 9, 7, 8, 9},
				{1, 2, 3, 1, 2, 3, 1, 2, 3},
				{4, 5, 6, 4, 5, 6, 4, 5, 6},
				{7, 8, 9, 7, 8, 9, 7, 8, 3},
			},
			false,
			"Should fail because last box has 3 twice",
		},
		{
			puzzle{
				{1, 2, 3, 1, 2, 3, 1, 2, 3},
				{4, 5, 6, 4, 5, 6, 4, 5, 6},
				{7, 8, 9, 7, 8, 9, 7, 8, 9},
			},
			true,
			"Should pass when not completed but there are no conflicts",
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.puzzle.testBoxes(), test.expected, test.reason)
	}
}

func TestPuzzle(t *testing.T) {
	tests := []struct {
		puzzle   puzzle
		expected bool
		reason   string
	}{
		{
			puzzle{
				{1, 2, 3, 4, 5, 6, 7, 8, 9},
				{4, 5, 6, 7, 8, 9, 1, 2, 3},
				{7, 8, 9, 1, 2, 3, 4, 5, 6},
				{2, 3, 4, 5, 6, 7, 8, 9, 1},
				{5, 6, 7, 8, 9, 1, 2, 3, 4},
				{8, 9, 1, 2, 3, 4, 5, 6, 7},
				{3, 4, 5, 6, 7, 8, 9, 1, 2},
				{6, 7, 8, 9, 1, 2, 3, 4, 5},
				{9, 1, 2, 3, 4, 5, 6, 7, 8},
			},
			true,
			"Should pass because puzzle is valid",
		},
		{
			puzzle{
				{1, 2, 3, 4, 5, 6, 7, 8, 9},
				{4, 5, 6, 7, 8, 9, 1, 2, 3},
				{7, 8, 9, 1, 2, 3, 4, 5, 6},
				{2, 3, 4, 5, 6, 7, 8, 9, 1},
				{5, 6, 7, 8, 9, 1, 2, 3, 4},
				{8, 9, 1, 2, 0, 4, 5, 6, 7},
				{3, 4, 5, 6, 7, 8, 9, 1, 2},
				{6, 7, 8, 9, 1, 2, 3, 4, 5},
				{9, 1, 2, 3, 4, 5, 6, 7, 8},
			},
			true,
			"Should pass when puzzle has no conflicts but is not completed",
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.puzzle.test(), test.expected, test.reason)
	}
}
