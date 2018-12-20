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
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Fields Contains all details for the main contact.
// swagger:discriminator Fields classType
type Fields interface {
	runtime.Validatable

	// fields
	Fields() *FieldsFields
	SetFields(*FieldsFields)
}

type fields struct {
	fieldsField *FieldsFields
}

// Fields gets the fields of this polymorphic type
func (m *fields) Fields() *FieldsFields {
	return m.fieldsField
}

// SetFields sets the fields of this polymorphic type
func (m *fields) SetFields(val *FieldsFields) {
	m.fieldsField = val
}

// UnmarshalFieldsSlice unmarshals polymorphic slices of Fields
func UnmarshalFieldsSlice(reader io.Reader, consumer runtime.Consumer) ([]Fields, error) {
	var elements []json.RawMessage
	if err := consumer.Consume(reader, &elements); err != nil {
		return nil, err
	}

	var result []Fields
	for _, element := range elements {
		obj, err := unmarshalFields(element, consumer)
		if err != nil {
			return nil, err
		}
		result = append(result, obj)
	}
	return result, nil
}

// UnmarshalFields unmarshals polymorphic Fields
func UnmarshalFields(reader io.Reader, consumer runtime.Consumer) (Fields, error) {
	// we need to read this twice, so first into a buffer
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return unmarshalFields(data, consumer)
}

func unmarshalFields(data []byte, consumer runtime.Consumer) (Fields, error) {
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
	case "Fields":
		var result fields
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	}
	return nil, errors.New(422, "invalid classType value: %q", getType.ClassType)

}

// Validate validates this fields
func (m *fields) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFields(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *fields) validateFields(formats strfmt.Registry) error {

	if swag.IsZero(m.Fields()) { // not required
		return nil
	}

	if m.Fields() != nil {
		if err := m.Fields().Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fields")
			}
			return err
		}
	}

	return nil
}