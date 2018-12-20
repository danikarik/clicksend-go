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

// AccountForgotPasswordVerify account forgot password verify
// swagger:discriminator AccountForgotPasswordVerify classType
type AccountForgotPasswordVerify interface {
	runtime.Validatable

	// Activation token
	// Required: true
	ActivationToken() *string
	SetActivationToken(*string)

	// Password
	// Required: true
	Password() *string
	SetPassword(*string)

	// ID of subaccount
	// Required: true
	SubaccountID() *int32
	SetSubaccountID(*int32)
}

type accountForgotPasswordVerify struct {
	activationTokenField *string

	passwordField *string

	subaccountIdField *int32
}

// ActivationToken gets the activation token of this polymorphic type
func (m *accountForgotPasswordVerify) ActivationToken() *string {
	return m.activationTokenField
}

// SetActivationToken sets the activation token of this polymorphic type
func (m *accountForgotPasswordVerify) SetActivationToken(val *string) {
	m.activationTokenField = val
}

// Password gets the password of this polymorphic type
func (m *accountForgotPasswordVerify) Password() *string {
	return m.passwordField
}

// SetPassword sets the password of this polymorphic type
func (m *accountForgotPasswordVerify) SetPassword(val *string) {
	m.passwordField = val
}

// SubaccountID gets the subaccount id of this polymorphic type
func (m *accountForgotPasswordVerify) SubaccountID() *int32 {
	return m.subaccountIdField
}

// SetSubaccountID sets the subaccount id of this polymorphic type
func (m *accountForgotPasswordVerify) SetSubaccountID(val *int32) {
	m.subaccountIdField = val
}

// UnmarshalAccountForgotPasswordVerifySlice unmarshals polymorphic slices of AccountForgotPasswordVerify
func UnmarshalAccountForgotPasswordVerifySlice(reader io.Reader, consumer runtime.Consumer) ([]AccountForgotPasswordVerify, error) {
	var elements []json.RawMessage
	if err := consumer.Consume(reader, &elements); err != nil {
		return nil, err
	}

	var result []AccountForgotPasswordVerify
	for _, element := range elements {
		obj, err := unmarshalAccountForgotPasswordVerify(element, consumer)
		if err != nil {
			return nil, err
		}
		result = append(result, obj)
	}
	return result, nil
}

// UnmarshalAccountForgotPasswordVerify unmarshals polymorphic AccountForgotPasswordVerify
func UnmarshalAccountForgotPasswordVerify(reader io.Reader, consumer runtime.Consumer) (AccountForgotPasswordVerify, error) {
	// we need to read this twice, so first into a buffer
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return unmarshalAccountForgotPasswordVerify(data, consumer)
}

func unmarshalAccountForgotPasswordVerify(data []byte, consumer runtime.Consumer) (AccountForgotPasswordVerify, error) {
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
	case "AccountForgotPasswordVerify":
		var result accountForgotPasswordVerify
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	}
	return nil, errors.New(422, "invalid classType value: %q", getType.ClassType)

}

// Validate validates this account forgot password verify
func (m *accountForgotPasswordVerify) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActivationToken(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePassword(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubaccountID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *accountForgotPasswordVerify) validateActivationToken(formats strfmt.Registry) error {

	if err := validate.Required("activation_token", "body", m.ActivationToken()); err != nil {
		return err
	}

	return nil
}

func (m *accountForgotPasswordVerify) validatePassword(formats strfmt.Registry) error {

	if err := validate.Required("password", "body", m.Password()); err != nil {
		return err
	}

	return nil
}

func (m *accountForgotPasswordVerify) validateSubaccountID(formats strfmt.Registry) error {

	if err := validate.Required("subaccount_id", "body", m.SubaccountID()); err != nil {
		return err
	}

	return nil
}