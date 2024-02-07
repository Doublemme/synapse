package validation

import (
	"fmt"
	"testing"
)

func TestValidator(t *testing.T) {

	data := struct {
		FirstName string
		LastName  string
	}{
		FirstName: "Michael",
		LastName:  "Martin",
	}

	var errs any

	isValid := NewValidator(data, Fields{
		"FirstName": BindingRules(Required),
	}).Validate(errs)

	if !isValid {
		fmt.Println(errs)
	}
}
