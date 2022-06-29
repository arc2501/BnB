package forms

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/asaskevich/govalidator"
)

// creates a custom form with passed URL values
type Form struct {
	url.Values
	Errors errors
}

func (f *Form) Valid() bool {
	// if form Error has any value then return false
	return len(f.Errors) == 0

}

// initializes a form struct
func New(data url.Values) *Form {
	return &Form{
		data,                          // the url value we passed
		errors(map[string][]string{}), // this one is initialized empty
	}
}

func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "This field cannot be blank")
		}
	}
}

// checks does the form has that field?
func (f *Form) Has(field string, r *http.Request) bool {
	x := r.Form.Get(field)
	if x == "" {
		f.Errors.Add(field, "this field cannot be blank")
		return false
	}
	return true
}

// Checks for minimum length of fields
func (f *Form) MinLength(field string, length int, r *http.Request) bool {
	x := r.Form.Get(field)
	if len(x) < length {
		f.Errors.Add(field, fmt.Sprintf("This field must be at least %d characters long", length))
		return false
	}
	return true
}

// Checks for valid Email
func (f *Form) IsEmail(field string) {
	if !govalidator.IsEmail(f.Get(field)) {
		f.Errors.Add(field, "Invalid Email Address")
	}
}
