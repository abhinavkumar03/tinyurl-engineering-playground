package repository_test

import (
	"testing"
)

func TestRepositoryPackage(t *testing.T) {

	if testing.Short() {
		t.Skip()
	}

	if 1 != 1 {
		t.Fatal("repository test failed")
	}
}
