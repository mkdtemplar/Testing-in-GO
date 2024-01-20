package main

import (
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"
)

func TestForm_Check(t *testing.T) {
	type fields struct {
		Data   url.Values
		Errors errors
	}
	type args struct {
		ok      bool
		key     string
		message string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := Form{
				Data:   tt.fields.Data,
				Errors: tt.fields.Errors,
			}
			f.Check(tt.args.ok, tt.args.key, tt.args.message)
		})
	}
}

func TestForm_Has(t *testing.T) {
	form := NewForm(nil)

	has := form.Has("some_field")
	if has {
		t.Error("Form shows that have a field but it should not")
	}

	postedData := url.Values{}
	postedData.Add("a", "a")
	form = NewForm(postedData)

	has = form.Has("a")
	if !has {
		t.Error("The form shows that there is no fields but it should")
	}
}

func TestForm_Required(t *testing.T) {
	r := httptest.NewRequest("POST", "/testroute", nil)
	form := NewForm(r.PostForm)
	form.Required("a", "b", "c")
	if form.Valid() {
		t.Error("Form shows valid when required fields are missing")
	}

	postedData := url.Values{}
	postedData.Add("a", "a")
	postedData.Add("b", "b")
	postedData.Add("c", "c")

	r = httptest.NewRequest("POST", "/testroute", nil)
	r.PostForm = postedData
	form = NewForm(r.PostForm)
	form.Required("a", "b", "c")
	if !form.Valid() {
		t.Error("Form shows that has not the required fields but it does")
	}
}

func TestForm_Valid(t *testing.T) {
	type fields struct {
		Data   url.Values
		Errors errors
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := Form{
				Data:   tt.fields.Data,
				Errors: tt.fields.Errors,
			}
			if got := f.Valid(); got != tt.want {
				t.Errorf("Valid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewForm(t *testing.T) {
	type args struct {
		data url.Values
	}
	tests := []struct {
		name string
		args args
		want *Form
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewForm(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewForm() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_errors_Add(t *testing.T) {
	type args struct {
		field   string
		message string
	}
	tests := []struct {
		name string
		e    errors
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.e.Add(tt.args.field, tt.args.message)
		})
	}
}

func Test_errors_Get(t *testing.T) {
	type args struct {
		field string
	}
	tests := []struct {
		name string
		e    errors
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Get(tt.args.field); got != tt.want {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
