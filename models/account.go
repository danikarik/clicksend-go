// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/validate"
)

// Account Complete account details needed for the user.
// swagger:discriminator Account classType
type Account interface {
	runtime.Validatable

	// Your delivery to value.
	// Required: true
	AccountName() *string
	SetAccountName(*string)

	// Your country
	// Required: true
	Country() *string
	SetCountry(*string)

	// Your password
	// Required: true
	Password() *string
	SetPassword(*string)

	// Your email
	// Required: true
	UserEmail() *string
	SetUserEmail(*string)

	// Your first name
	// Required: true
	UserFirstName() *string
	SetUserFirstName(*string)

	// Your last name
	// Required: true
	UserLastName() *string
	SetUserLastName(*string)

	// Your phone number in E.164 format.
	// Required: true
	UserPhone() *string
	SetUserPhone(*string)

	// Your username
	// Required: true
	Username() *string
	SetUsername(*string)
}

type account struct {
	accountNameField *string

	countryField *string

	passwordField *string

	userEmailField *string

	userFirstNameField *string

	userLastNameField *string

	userPhoneField *string

	usernameField *string
}

// AccountName gets the account name of this polymorphic type
func (m *account) AccountName() *string {
	return m.accountNameField
}

// SetAccountName sets the account name of this polymorphic type
func (m *account) SetAccountName(val *string) {
	m.accountNameField = val
}

// Country gets the country of this polymorphic type
func (m *account) Country() *string {
	return m.countryField
}

// SetCountry sets the country of this polymorphic type
func (m *account) SetCountry(val *string) {
	m.countryField = val
}

// Password gets the password of this polymorphic type
func (m *account) Password() *string {
	return m.passwordField
}

// SetPassword sets the password of this polymorphic type
func (m *account) SetPassword(val *string) {
	m.passwordField = val
}

// UserEmail gets the user email of this polymorphic type
func (m *account) UserEmail() *string {
	return m.userEmailField
}

// SetUserEmail sets the user email of this polymorphic type
func (m *account) SetUserEmail(val *string) {
	m.userEmailField = val
}

// UserFirstName gets the user first name of this polymorphic type
func (m *account) UserFirstName() *string {
	return m.userFirstNameField
}

// SetUserFirstName sets the user first name of this polymorphic type
func (m *account) SetUserFirstName(val *string) {
	m.userFirstNameField = val
}

// UserLastName gets the user last name of this polymorphic type
func (m *account) UserLastName() *string {
	return m.userLastNameField
}

// SetUserLastName sets the user last name of this polymorphic type
func (m *account) SetUserLastName(val *string) {
	m.userLastNameField = val
}

// UserPhone gets the user phone of this polymorphic type
func (m *account) UserPhone() *string {
	return m.userPhoneField
}

// SetUserPhone sets the user phone of this polymorphic type
func (m *account) SetUserPhone(val *string) {
	m.userPhoneField = val
}

// Username gets the username of this polymorphic type
func (m *account) Username() *string {
	return m.usernameField
}

// SetUsername sets the username of this polymorphic type
func (m *account) SetUsername(val *string) {
	m.usernameField = val
}

// UnmarshalAccountSlice unmarshals polymorphic slices of Account
func UnmarshalAccountSlice(reader io.Reader, consumer runtime.Consumer) ([]Account, error) {
	var elements []json.RawMessage
	if err := consumer.Consume(reader, &elements); err != nil {
		return nil, err
	}

	var result []Account
	for _, element := range elements {
		obj, err := unmarshalAccount(element, consumer)
		if err != nil {
			return nil, err
		}
		result = append(result, obj)
	}
	return result, nil
}

// UnmarshalAccount unmarshals polymorphic Account
func UnmarshalAccount(reader io.Reader, consumer runtime.Consumer) (Account, error) {
	// we need to read this twice, so first into a buffer
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return unmarshalAccount(data, consumer)
}

func unmarshalAccount(data []byte, consumer runtime.Consumer) (Account, error) {
	buf := bytes.NewBuffer(data)
	buf2 := bytes.NewBuffer(data)

	// the first time this is read is to fetch the value of the classType property.
	var getType struct {
		ClassType string `json:"classType"`
	}
	if err := consumer.Consume(buf, &getType); err != nil {
		return nil, err
	}

	if err := validate.RequiredString("classType", "body", getType.ClassType); err != nil {
		return nil, err
	}

	// The value of classType is used to determine which type to create and unmarshal the data into
	switch getType.ClassType {
	case "Account":
		var result account
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	}
	return nil, errors.New(422, "invalid classType value: %q", getType.ClassType)

}

// Validate validates this account
func (m *account) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccountName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCountry(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePassword(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserEmail(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserFirstName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserLastName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserPhone(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsername(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *account) validateAccountName(formats strfmt.Registry) error {

	if err := validate.Required("account_name", "body", m.AccountName()); err != nil {
		return err
	}

	return nil
}

func (m *account) validateCountry(formats strfmt.Registry) error {

	if err := validate.Required("country", "body", m.Country()); err != nil {
		return err
	}

	return nil
}

func (m *account) validatePassword(formats strfmt.Registry) error {

	if err := validate.Required("password", "body", m.Password()); err != nil {
		return err
	}

	return nil
}

func (m *account) validateUserEmail(formats strfmt.Registry) error {

	if err := validate.Required("user_email", "body", m.UserEmail()); err != nil {
		return err
	}

	return nil
}

func (m *account) validateUserFirstName(formats strfmt.Registry) error {

	if err := validate.Required("user_first_name", "body", m.UserFirstName()); err != nil {
		return err
	}

	return nil
}

func (m *account) validateUserLastName(formats strfmt.Registry) error {

	if err := validate.Required("user_last_name", "body", m.UserLastName()); err != nil {
		return err
	}

	return nil
}

func (m *account) validateUserPhone(formats strfmt.Registry) error {

	if err := validate.Required("user_phone", "body", m.UserPhone()); err != nil {
		return err
	}

	return nil
}

func (m *account) validateUsername(formats strfmt.Registry) error {

	if err := validate.Required("username", "body", m.Username()); err != nil {
		return err
	}

	return nil
}
