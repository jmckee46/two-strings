package main

import "testing"

func TestTwoStrings1(t *testing.T) {
	s1 := "hello"
	s2 := "world"
	present := twoStrings(s1, s2)

	if present != "YES" {
		t.Errorf("got %s instead of YES", present)
	}
}

func TestTwoStrings2(t *testing.T) {
	s1 := "hi"
	s2 := "world"
	present := twoStrings(s1, s2)

	if present != "NO" {
		t.Errorf("got %s instead of NO", present)
	}
}
