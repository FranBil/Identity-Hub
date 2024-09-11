package formats

import (
	"errors"
	"testing"
)

func TestValidateFirstName(t *testing.T) {
	t.Run("FirstName is empty", func(t *testing.T) {
		person := PersonRequest{
			FirstName:   "",
			LastName:    "Doe",
			PhoneNumber: "1234567890",
			Address:     "123 Street",
		}

		flag, err := person.validateFirstName()

		if flag {
			t.Errorf("expected flag to be false, got true")
		}
		if err == nil {
			t.Errorf("expected an error, got nil")
		}
		expectedErr := errors.New("first name is required")
		if err.Error() != expectedErr.Error() {
			t.Errorf("expected error: %s, got: %s", expectedErr.Error(), err.Error())
		}
	})

	t.Run("FirstName is valid", func(t *testing.T) {
		person := PersonRequest{
			FirstName:   "John",
			LastName:    "Doe",
			PhoneNumber: "1234567890",
			Address:     "123 Street",
		}

		flag, err := person.validateFirstName()

		if !flag {
			t.Errorf("expected flag to be true, got false")
		}
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
	})
}

func TestValidateLastName(t *testing.T) {
	t.Run("LastName is empty", func(t *testing.T) {
		person := PersonRequest{
			FirstName:   "John",
			LastName:    "",
			PhoneNumber: "1234567890",
			Address:     "123 Street",
		}

		flag, err := person.validateLastName()

		if flag {
			t.Errorf("expected flag to be false, got true")
		}
		if err == nil {
			t.Errorf("expected an error, got nil")
		}
		expectedErr := errors.New("last name is required")
		if err.Error() != expectedErr.Error() {
			t.Errorf("expected error: %s, got: %s", expectedErr.Error(), err.Error())
		}
	})

	t.Run("LastName is valid", func(t *testing.T) {
		person := PersonRequest{
			FirstName:   "John",
			LastName:    "Doe",
			PhoneNumber: "1234567890",
			Address:     "123 Street",
		}

		flag, err := person.validateLastName()

		if !flag {
			t.Errorf("expected flag to be true, got false")
		}
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
	})
}

func TestValidatePhoneNumber(t *testing.T) {
	t.Run("PhoneNumber is empty", func(t *testing.T) {
		person := PersonRequest{
			FirstName:   "John",
			LastName:    "Doe",
			PhoneNumber: "",
			Address:     "123 Street",
		}

		flag, err := person.validatePhoneNumber()

		if flag {
			t.Errorf("expected flag to be false, got true")
		}
		if err == nil {
			t.Errorf("expected an error, got nil")
		}
		expectedErr := errors.New("phone number is required")
		if err.Error() != expectedErr.Error() {
			t.Errorf("expected error: %s, got: %s", expectedErr.Error(), err.Error())
		}
	})

	t.Run("PhoneNumber contains invalid characters", func(t *testing.T) {
		person := PersonRequest{
			FirstName:   "John",
			LastName:    "Doe",
			PhoneNumber: "123abc456",
			Address:     "123 Street",
		}

		flag, err := person.validatePhoneNumber()

		if flag {
			t.Errorf("expected flag to be false, got true")
		}
		if err == nil {
			t.Errorf("expected an error, got nil")
		}
		expectedErr := errors.New("phone number should only contain digits")
		if err.Error() != expectedErr.Error() {
			t.Errorf("expected error: %s, got: %s", expectedErr.Error(), err.Error())
		}
	})

	t.Run("PhoneNumber is valid", func(t *testing.T) {
		person := PersonRequest{
			FirstName:   "John",
			LastName:    "Doe",
			PhoneNumber: "1234567890",
			Address:     "123 Street",
		}

		flag, err := person.validatePhoneNumber()

		if !flag {
			t.Errorf("expected flag to be true, got false")
		}
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
	})
}

func TestValidateAddress(t *testing.T) {
	t.Run("Address is empty", func(t *testing.T) {
		person := PersonRequest{
			FirstName:   "John",
			LastName:    "Doe",
			PhoneNumber: "1234567890",
			Address:     "",
		}

		flag, err := person.validateAddress()

		if flag {
			t.Errorf("expected flag to be false, got true")
		}
		if err == nil {
			t.Errorf("expected an error, got nil")
		}
		expectedErr := errors.New("address is required")
		if err.Error() != expectedErr.Error() {
			t.Errorf("expected error: %s, got: %s", expectedErr.Error(), err.Error())
		}
	})

	t.Run("Address is valid", func(t *testing.T) {
		person := PersonRequest{
			FirstName:   "John",
			LastName:    "Doe",
			PhoneNumber: "1234567890",
			Address:     "123 Main Street",
		}

		flag, err := person.validateAddress()

		if !flag {
			t.Errorf("expected flag to be true, got false")
		}
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
	})
}

func TestIsValid(t *testing.T) {
	t.Run("All fields are valid", func(t *testing.T) {
		person := PersonRequest{
			FirstName:   "John",
			LastName:    "Doe",
			PhoneNumber: "1234567890",
			Address:     "123 Main Street",
		}

		isValid, errs := person.IsValid()

		if !isValid {
			t.Errorf("expected true, got false")
		}
		if len(errs) != 0 {
			t.Errorf("expected no errors, got %v", errs)
		}
	})

	t.Run("Multiple fields are invalid", func(t *testing.T) {
		person := PersonRequest{
			FirstName:   "",
			LastName:    "",
			PhoneNumber: "abc123456",
			Address:     "",
		}

		isValid, errs := person.IsValid()

		if isValid {
			t.Errorf("expected false, got true")
		}
		if len(errs) != 4 {
			t.Errorf("expected 4 errors, got %v", len(errs))
		}

		expectedErrs := []error{
			errors.New("first name is required"),
			errors.New("last name is required"),
			errors.New("phone number should only contain digits"),
			errors.New("address is required"),
		}

		for i, err := range errs {
			if err.Error() != expectedErrs[i].Error() {
				t.Errorf("expected error: %s, got: %s", expectedErrs[i].Error(), err.Error())
			}
		}
	})
}