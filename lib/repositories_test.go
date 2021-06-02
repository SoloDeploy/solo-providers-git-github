package main

import (
	"testing"
)

func TestGetRepositoryNames(t *testing.T) {
	names := GetRepositoryNames()

	if len(names) == 0 {
		t.Errorf("Expected repositories to be returned")
	}
}
