package main

import (
	"testing"
)

func TestMutexUnsafeRace(t *testing.T) {
	// This test is illustrative; it will not fail unless run with -race
	RunMutexUnsafe()
	// To properly detect race, run: go test -race
}
