package main

import (
	"net/url"
	"testing"
)

func TestForm_Has(t *testing.T) {
	postedData := url.Values{}
	postedData.Add("email", "me@here.com")
	form := NewForm(postedData)

	if !form.Has("email") {
		t.Error("Expected form to have 'email' field, but it does not")
	}

	if form.Has("password") {
		t.Error("Expected form not to have 'password' field, but it does")
	}
}

func TestForm_Required(t *testing.T) {
	form := NewForm(nil)
	form.Required("email")

	if form.Errors.Get("email") != "This field cannot be blank" {
		t.Errorf("Expected error message for 'email' field, but got: %s", form.Errors.Get("email"))
	}

	postedData := url.Values{}
	postedData.Add("email", "me@here.com")

	form = NewForm(postedData)
	form.Required("email")
	if form.Errors.Get("email") != "" {
		t.Errorf("Expected no error message for 'email' field, but got: %s", form.Errors.Get("email"))
	}
}

func TestForm_Check(t *testing.T) {
	form := NewForm(nil)
	form.Check(false, "password", "password is required")

	if form.Errors.Get("password") != "password is required" {
		t.Errorf("Expected error message for 'password' field, but got: %s", form.Errors.Get("password"))
	}
}

func TestForm_Valid(t *testing.T) {
	postedData := url.Values{}
	postedData.Add("email", "me@here.com")
	form := NewForm(postedData)

	if !form.Valid() {
		t.Error("Expected form to be valid initially")
	}

	form.Errors.Add("password", "password is invalid")
	if form.Valid() {
		t.Error("Expected form to be invalid after adding error")
	}
}
