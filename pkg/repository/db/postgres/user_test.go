package postgres

import (
	"errors"
	"testing"

	"giants/pkg/domain"
)

func TestUser(t *testing.T) {
	defer DeleteAll()

	dbrepo, err := NewPostgres()
	if err != nil {
		t.Fatal(err)
	}
	defer dbrepo.Close()

	// Add a new user
	uObj, _ := domain.NewUser("u1")

	err = dbrepo.AddUser(uObj)
	if err != nil {
		t.Fatal(err)
	}

	// Check if duplicated error occurs
	err = dbrepo.AddUser(uObj)
	if !errors.Is(err, ErrDuplicated) {
		t.Fatal("Expected ErrDuplicated")
	}

	// Fetch user from db and check
	_, err = dbrepo.FindUserById(uObj.UserId)
	if err != nil {
		t.Fatal(err)
	}
}
