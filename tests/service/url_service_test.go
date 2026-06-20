package service_test

import (
	"context"
	"testing"

	"github.com/abhinavkumar03/tinyurl-engineering-playground/internal/model"
	"github.com/abhinavkumar03/tinyurl-engineering-playground/pkg/shortener"
)

func TestBase62Generation(t *testing.T) {

	url := &model.URL{
		ID: 125,
	}

	code := shortener.Encode(url.ID)

	if code == "" {
		t.Fatal("short code should not be empty")
	}
}

func TestContextCreation(t *testing.T) {

	ctx := context.Background()

	if ctx == nil {
		t.Fatal("context should not be nil")
	}
}
