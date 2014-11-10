package main

import "testing"

func TestMove(t *testing.T) {
	type test struct {
		pos       Position
		expected1 int
		expected2 int
	}
	tests := []test{
		test{pos: Position{x: 2, y: 4, dx: 1, dy: 3}, expected1: 3, expected2: 7},
		test{pos: Position{x: 20, y: -45, dx: -7, dy: 40}, expected1: 13, expected2: -5},
	}
	for _, tt := range tests {
		result1, result2 := Move(tt.pos)
		if result1 != tt.expected1 || result2 != tt.expected2 {
			t.Errorf("Move(%v) = %v, %v; want %v, %v", tt.pos, result1, result2, tt.expected1, tt.expected2)
		}
	}
}
