package util

import "testing"

func TestNewStoreImpl1(t *testing.T) {
	var s Store = NewStoreImpl1("myUser", 12.12)

	if user, err := s.GetUser(); err != nil {
		t.Error("get user")
	} else if user != "myUser" {
		t.Error("incorrect user")
	}

	if value, err := s.GetValue(); err != nil {
		t.Error("get value")
	} else if value != 12.12 {
		t.Error("incorrect value")
	}
}

func TestStoreImpl2(t *testing.T) {
	var sTrue Store = StoreImpl2(true)
	var sFalse Store = StoreImpl2(false)

	if user, err := sTrue.GetUser(); err != nil {
		t.Error("get user")
	} else if user != "yes" {
		t.Error("incorrect user")
	}

	if value, err := sTrue.GetValue(); err != nil {
		t.Error("get value")
	} else if value != 1 {
		t.Error("incorrect value")
	}

	if user, err := sFalse.GetUser(); err != nil {
		t.Error("get user")
	} else if user != "n0" {
		t.Error("incorrect user")
	}

	if value, err := sFalse.GetValue(); err != nil {
		t.Error("get value")
	} else if value != -1 {
		t.Error("incorrect value")
	}
}

func TestStore(t *testing.T) {
	var stores = []Store{
		NewStoreImpl1("asdf", 1337),
		StoreImpl2(true),
		StoreImpl2(false),
	}

	var userResults = []string{
		"asdf",
		"yes",
		"n0",
	}

	for i := range stores {
		if user, err := stores[i].GetUser(); err != nil {
			t.Error("get user")
		} else if user != userResults[i] {
			t.Error("incorrect user")
		}
	}
}
