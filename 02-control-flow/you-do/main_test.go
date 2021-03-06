package main

import "testing"

func Test_wordMatch(t *testing.T) {
	type test struct {
		name          string
		first, second string
		expected      bool
	}
	tests := []test{
		test{name: "Same words", first: "dog", second: "dog", expected: true},
	}
	for _, tt := range tests {
		result := wordMatch(tt.first, tt.second)
		if result != tt.expected {
			t.Errorf("Test %q failed. got: %d, expected %v", tt.name, result, tt.expected)
		}
	}
}

func Test_doubleUp(t *testing.T) {
	type test struct {
		name     string
		a, b     int
		expected string
	}
	tests := []test{
		test{name: "a double of b", a: 4, b: 2, expected: "double up!"},
		test{name: "a equal to b", a: 4, b: 4, expected: "no double..."},
		test{name: "a not double, not equal to b", a: 1, b: 7, expected: "no double..."},
		test{name: "a half of b", a: 4, b: 8, expected: "double up!"},
	}
	for _, tt := range tests {
		result := doubleUp(tt.a, tt.b)
		if result != tt.expected {
			t.Errorf("Test %q failed. got :%v, expected: %v", tt.name, result, tt.expected)
		}
	}
}

func Test_runeToNum(t *testing.T) {
	type test struct {
		r        rune
		expected bool
	}
	tests := []test{
		test{r: 'F', expected: true},
		test{r: '/', expected: false},
	}
	for _, tt := range tests {
		result := runeToNum(tt.r)
		if result != tt.expected {
			t.Errorf("runeToNum(%d) = %d, want %v", tt.r, result, tt.expected)
		}
	}
}

func Test_multipleByte(t *testing.T) {
	type test struct {
		i        int
		expected string
	}
	tests := []test{
		test{i: 1024, expected: "kibibyte"},
		test{i: 1025, expected: "inexact"},
		test{i: 1048576, expected: "mebibyte"},
		test{i: 1073741824, expected: "gibibyte"},
	}
	for _, tt := range tests {
		result := multipleByte(tt.i)
		if result != tt.expected {
			t.Errorf("multipleByte(%d) = %d, want %v", tt.i, result, tt.expected)
		}
	}
}

func Test_correctByte(t *testing.T) {
	type test struct {
		i        int
		expected int
	}
	tests := []test{
		test{i: 1000, expected: 1024},
		test{i: 1000000, expected: 1048576},
		test{i: 1000000000, expected: 1073741824},
		test{i: 5, expected: 5},
	}
	for _, tt := range tests {
		result := correctByte(tt.i)
		if result != tt.expected {
			t.Errorf("correctByte(%d) = %d, want %v", tt.i, result, tt.expected)
		}
	}
}

func Test_threeString(t *testing.T) {
	type test struct {
		s1       string
		s2       string
		s3       string
		expected string
	}
	tests := []test{
		test{s1: "stop", s2: "second string", s3: "third string", expected: "second string"},
	}
	for _, tt := range tests {
		result := threeString(tt.s1, tt.s2, tt.s3)
		if result != tt.expected {
			t.Errorf("threeString(%v, %v, %v) = %d, want %v", tt.s1, tt.s2, tt.s3, result, tt.expected)
		}
	}
}

func Test_topThree(t *testing.T) {
	type test struct {
		b        bool
		i        int
		expected string
	}
	tests := []test{
		test{b: true, i: 1, expected: "1."},
		test{b: true, i: 2, expected: "2."},
		test{b: true, i: 3, expected: "3."},
		test{b: true, i: 10, expected: "not good enough"},
		test{b: false, i: 1, expected: "One:"},
		test{b: false, i: 2, expected: "Two:"},
		test{b: false, i: 3, expected: "Three:"},
		test{b: false, i: 45, expected: "not good enough"},
	}
	for _, tt := range tests {
		result := topThree(tt.b, tt.i)
		if result != tt.expected {
			t.Errorf("topThree(%v, %v) = %d, want %v", tt.b, tt.i, result, tt.expected)
		}
	}
}
