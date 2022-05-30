package forms

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestForm_Valid(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	isValid := form.Valid()
	if !isValid {
		t.Error("got invalid when should have been valid")
	}
}

func TestForm_Required(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	form.Required("a", "b", "c")
	if form.Valid() {
		t.Error("form shows valid when required fields are missing")
	}

	postedData := url.Values{}
	postedData.Add("a", "a")
	postedData.Add("b", "a")
	postedData.Add("c", "a")

	r, _ = http.NewRequest("POST", "/whatever", nil)

	r.PostForm = postedData
	form = New(r.PostForm)
	form.Required("a", "b", "c")
	if !form.Valid() {
		t.Error("shows does not have required fields when it does")
	}
}

func TestForm_Has(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	has := form.Has("whatever")
	if has {
		t.Error("form shows it has field when it does not")
	}

	postedData := url.Values{}
	postedData.Add("a", "a")
	form = New(postedData)

	has = form.Has("a")
	if !has {
		t.Error("Shows form does not have field when it does")
	}
}

func TestForm_MinLength(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	form.MinLength("x", 10)
	if form.Valid() {
		t.Error("form shows min length for non-existent field")
	}

	isError := form.Errors.Get("x")
	if isError == "" {
		t.Error("should have an error, but did not get one")
	}

	postedValues := url.Values{}
	postedValues.Add("a", "aaaa")
	form = New(postedValues)

	form.MinLength("a", 100)
	if form.Valid() {
		t.Error("Form shows min length being long enough when it should be too short")
	}

	postedValues = url.Values{}
	postedValues.Add("b", "bbbb")
	form = New(postedValues)

	form.MinLength("b", 3)
	if !form.Valid() {
		t.Error("Form shows min length is too short when it is long enough")
	}

	isError = form.Errors.Get("b")
	if isError != "" {
		t.Error("should not have an error, but got one")
	}
}

func TestForm_IsEmail(t *testing.T) {
	postedEmails := url.Values{}
	form := New(postedEmails)

	form.IsEmail("x")
	if form.Valid() {
		t.Error("form shows valid email for non-existent field")
	}

	postedEmails = url.Values{}
	postedEmails.Add("email_valid", "me@here.com")
	postedEmails.Add("email_invalid", "mehere.com")
	form = New(postedEmails)

	form.IsEmail("email_valid")
	if !form.Valid() {
		t.Error("form shows email is invalid when it should be valid")
	}

	form.IsEmail("email_invalid")
	if form.Valid() {
		t.Error("form shows a valid email when it should be invalid")
	}
}