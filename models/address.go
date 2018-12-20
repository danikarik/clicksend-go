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

// Address Base model for all address-related objects.
// swagger:discriminator Address classType
type Address interface {
	runtime.Validatable

	// Your city
	// Required: true
	AddressCity() *string
	SetAddressCity(*string)

	// Your country
	// Required: true
	AddressCountry() *string
	SetAddressCountry(*string)

	// Your address line 1
	// Required: true
	AddressLine1() *string
	SetAddressLine1(*string)

	// Your address line 2
	AddressLine2() string
	SetAddressLine2(string)

	// Your address name.
	// Required: true
	AddressName() *string
	SetAddressName(*string)

	// Your postal code
	// Required: true
	AddressPostalCode() *string
	SetAddressPostalCode(*string)

	// Your state
	AddressState() string
	SetAddressState(string)
}

type address struct {
	addressCityField *string

	addressCountryField *string

	addressLine1Field *string

	addressLine2Field string

	addressNameField *string

	addressPostalCodeField *string

	addressStateField string
}

// AddressCity gets the address city of this polymorphic type
func (m *address) AddressCity() *string {
	return m.addressCityField
}

// SetAddressCity sets the address city of this polymorphic type
func (m *address) SetAddressCity(val *string) {
	m.addressCityField = val
}

// AddressCountry gets the address country of this polymorphic type
func (m *address) AddressCountry() *string {
	return m.addressCountryField
}

// SetAddressCountry sets the address country of this polymorphic type
func (m *address) SetAddressCountry(val *string) {
	m.addressCountryField = val
}

// AddressLine1 gets the address line 1 of this polymorphic type
func (m *address) AddressLine1() *string {
	return m.addressLine1Field
}

// SetAddressLine1 sets the address line 1 of this polymorphic type
func (m *address) SetAddressLine1(val *string) {
	m.addressLine1Field = val
}

// AddressLine2 gets the address line 2 of this polymorphic type
func (m *address) AddressLine2() string {
	return m.addressLine2Field
}

// SetAddressLine2 sets the address line 2 of this polymorphic type
func (m *address) SetAddressLine2(val string) {
	m.addressLine2Field = val
}

// AddressName gets the address name of this polymorphic type
func (m *address) AddressName() *string {
	return m.addressNameField
}

// SetAddressName sets the address name of this polymorphic type
func (m *address) SetAddressName(val *string) {
	m.addressNameField = val
}

// AddressPostalCode gets the address postal code of this polymorphic type
func (m *address) AddressPostalCode() *string {
	return m.addressPostalCodeField
}

// SetAddressPostalCode sets the address postal code of this polymorphic type
func (m *address) SetAddressPostalCode(val *string) {
	m.addressPostalCodeField = val
}

// AddressState gets the address state of this polymorphic type
func (m *address) AddressState() string {
	return m.addressStateField
}

// SetAddressState sets the address state of this polymorphic type
func (m *address) SetAddressState(val string) {
	m.addressStateField = val
}

// UnmarshalAddressSlice unmarshals polymorphic slices of Address
func UnmarshalAddressSlice(reader io.Reader, consumer runtime.Consumer) ([]Address, error) {
	var elements []json.RawMessage
	if err := consumer.Consume(reader, &elements); err != nil {
		return nil, err
	}

	var result []Address
	for _, element := range elements {
		obj, err := unmarshalAddress(element, consumer)
		if err != nil {
			return nil, err
		}
		result = append(result, obj)
	}
	return result, nil
}

// UnmarshalAddress unmarshals polymorphic Address
func UnmarshalAddress(reader io.Reader, consumer runtime.Consumer) (Address, error) {
	// we need to read this twice, so first into a buffer
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return unmarshalAddress(data, consumer)
}

func unmarshalAddress(data []byte, consumer runtime.Consumer) (Address, error) {
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
	case "Address":
		var result address
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	}
	return nil, errors.New(422, "invalid classType value: %q", getType.ClassType)

}

// Validate validates this address
func (m *address) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddressCity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAddressCountry(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAddressLine1(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAddressName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAddressPostalCode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *address) validateAddressCity(formats strfmt.Registry) error {

	if err := validate.Required("address_city", "body", m.AddressCity()); err != nil {
		return err
	}

	return nil
}

func (m *address) validateAddressCountry(formats strfmt.Registry) error {

	if err := validate.Required("address_country", "body", m.AddressCountry()); err != nil {
		return err
	}

	return nil
}

func (m *address) validateAddressLine1(formats strfmt.Registry) error {

	if err := validate.Required("address_line_1", "body", m.AddressLine1()); err != nil {
		return err
	}

	return nil
}

func (m *address) validateAddressName(formats strfmt.Registry) error {

	if err := validate.Required("address_name", "body", m.AddressName()); err != nil {
		return err
	}

	return nil
}

func (m *address) validateAddressPostalCode(formats strfmt.Registry) error {

	if err := validate.Required("address_postal_code", "body", m.AddressPostalCode()); err != nil {
		return err
	}

	return nil
}
