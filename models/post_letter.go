// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/validate"
)

// PostLetter PostLetter model
// swagger:discriminator PostLetter classType
type PostLetter interface {
	runtime.Validatable

	// Whether letter is in colour
	Colour() int32
	SetColour(int32)

	// Whether letter is duplex
	Duplex() int32
	SetDuplex(int32)

	// URL of file to send
	// Required: true
	FileURL() *string
	SetFileURL(*string)

	// Whether letter is priority
	PriorityPost() int32
	SetPriorityPost(int32)

	Recipients() []PostRecipient
	SetRecipients([]PostRecipient)

	// Source being sent from
	Source() *string
	SetSource(*string)

	// Whether using our template
	TemplateUsed() int32
	SetTemplateUsed(int32)
}

type postLetter struct {
	colourField int32

	duplexField int32

	fileUrlField *string

	priorityPostField int32

	recipientsField []PostRecipient

	sourceField *string

	templateUsedField int32
}

// Colour gets the colour of this polymorphic type
func (m *postLetter) Colour() int32 {
	return m.colourField
}

// SetColour sets the colour of this polymorphic type
func (m *postLetter) SetColour(val int32) {
	m.colourField = val
}

// Duplex gets the duplex of this polymorphic type
func (m *postLetter) Duplex() int32 {
	return m.duplexField
}

// SetDuplex sets the duplex of this polymorphic type
func (m *postLetter) SetDuplex(val int32) {
	m.duplexField = val
}

// FileURL gets the file url of this polymorphic type
func (m *postLetter) FileURL() *string {
	return m.fileUrlField
}

// SetFileURL sets the file url of this polymorphic type
func (m *postLetter) SetFileURL(val *string) {
	m.fileUrlField = val
}

// PriorityPost gets the priority post of this polymorphic type
func (m *postLetter) PriorityPost() int32 {
	return m.priorityPostField
}

// SetPriorityPost sets the priority post of this polymorphic type
func (m *postLetter) SetPriorityPost(val int32) {
	m.priorityPostField = val
}

// Recipients gets the recipients of this polymorphic type
func (m *postLetter) Recipients() []PostRecipient {
	return m.recipientsField
}

// SetRecipients sets the recipients of this polymorphic type
func (m *postLetter) SetRecipients(val []PostRecipient) {
	m.recipientsField = val
}

// Source gets the source of this polymorphic type
func (m *postLetter) Source() *string {
	return m.sourceField
}

// SetSource sets the source of this polymorphic type
func (m *postLetter) SetSource(val *string) {
	m.sourceField = val
}

// TemplateUsed gets the template used of this polymorphic type
func (m *postLetter) TemplateUsed() int32 {
	return m.templateUsedField
}

// SetTemplateUsed sets the template used of this polymorphic type
func (m *postLetter) SetTemplateUsed(val int32) {
	m.templateUsedField = val
}

// UnmarshalPostLetterSlice unmarshals polymorphic slices of PostLetter
func UnmarshalPostLetterSlice(reader io.Reader, consumer runtime.Consumer) ([]PostLetter, error) {
	var elements []json.RawMessage
	if err := consumer.Consume(reader, &elements); err != nil {
		return nil, err
	}

	var result []PostLetter
	for _, element := range elements {
		obj, err := unmarshalPostLetter(element, consumer)
		if err != nil {
			return nil, err
		}
		result = append(result, obj)
	}
	return result, nil
}

// UnmarshalPostLetter unmarshals polymorphic PostLetter
func UnmarshalPostLetter(reader io.Reader, consumer runtime.Consumer) (PostLetter, error) {
	// we need to read this twice, so first into a buffer
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return unmarshalPostLetter(data, consumer)
}

func unmarshalPostLetter(data []byte, consumer runtime.Consumer) (PostLetter, error) {
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
	case "PostLetter":
		var result postLetter
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	}
	return nil, errors.New(422, "invalid classType value: %q", getType.ClassType)

}

// Validate validates this post letter
func (m *postLetter) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFileURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecipients(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *postLetter) validateFileURL(formats strfmt.Registry) error {

	if err := validate.Required("file_url", "body", m.FileURL()); err != nil {
		return err
	}

	return nil
}

func (m *postLetter) validateRecipients(formats strfmt.Registry) error {

	if err := validate.Required("recipients", "body", m.Recipients()()); err != nil {
		return err
	}

	for i := 0; i < len(m.Recipients()()); i++ {

		if err := m.recipientsField[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("recipients" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}