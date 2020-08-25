package user_test

import (
	"testing"
	"user"
)

func TestNew(t *testing.T) {
	var result user.User
	expected := user.User{
		Email: "test@example.com",
	}
	result = user.New("test@example.com")

	if expected != result {
		t.Errorf("We want %v but got %v", expected, result)
	}
}
