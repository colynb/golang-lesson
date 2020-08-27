package user_test

import (
	"testing"
	"user"

	"github.com/google/go-cmp/cmp"
)

func TestNew(t *testing.T) {
	t.Parallel()
	want := user.User{
		Email: "test@example.com",
		Name:  "Mr Test",
	}
	got := user.New(want)
	want.ID = 1
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
	got2 := user.New(want)
	if got.ID == got2.ID {
		t.Errorf("want different IDS, got both %v", got.ID)
	}
}

func TestGet(t *testing.T) {
	t.Parallel()
	// If you create a user and know its ID
	// you should be able to call Get(ID) and
	// get the same user.
	want := user.User{
		Email: "test@example.com",
		Name:  "Mr Test",
	}
	u := user.New(want)
	got := user.Get(u.ID)
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
	// what about multiple users?
}
