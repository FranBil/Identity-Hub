package formats

import (
	"errors"
	"regexp"
)

type PersonRequest struct {
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	PhoneNumber string `json:"phoneNumber"`
	Address     string `json:"address"`
}

func (personRequest PersonRequest) validateFirstName() (flag bool, err error) {
	if personRequest.FirstName == "" {
		flag = false
		err = errors.New("first name is required")
		return
	}
	flag = true
	return
}

func (personRequest PersonRequest) validateLastName() (flag bool, err error) {
	if personRequest.LastName == "" {
		flag = false
		err = errors.New("last name is required")
		return
	}
	flag = true
	return
}

func (personRequest PersonRequest) validatePhoneNumber() (flag bool, err error) {
	if personRequest.PhoneNumber == "" {
		flag = false
		err = errors.New("phone number is required")
		return
	}

	re := regexp.MustCompile(`^[0-9]+$`)
	if !re.MatchString(personRequest.PhoneNumber) {
		flag = false
		err = errors.New("phone number should only contain digits")
		return
	}

	flag = true
	return
}

func (personRequest PersonRequest) validateAddress() (flag bool, err error) {
	if personRequest.Address == "" {
		flag = false
		err = errors.New("address is required")
		return
	}
	flag = true
	return
}

func (personRequest PersonRequest) IsValid() (bool, []error) {
	var errorList []error

	firstNameFlag, firstNameErr := personRequest.validateFirstName()
	if firstNameErr != nil {
		errorList = append(errorList, firstNameErr)
	}

	lastNameFlag, lastNameErr := personRequest.validateLastName()
	if lastNameErr != nil {
		errorList = append(errorList, lastNameErr)
	}

	phoneNumberFlag, phoneNumberErr := personRequest.validatePhoneNumber()
	if phoneNumberErr != nil {
		errorList = append(errorList, phoneNumberErr)
	}

	addressFlag, addressErr := personRequest.validateAddress()
	if addressErr != nil {
		errorList = append(errorList, addressErr)
	}

	return (firstNameFlag && lastNameFlag && addressFlag && phoneNumberFlag), errorList
}
