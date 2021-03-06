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

// Attachment Email attachment
// swagger:discriminator Attachment classType
type Attachment interface {
	runtime.Validatable

	// The base64-encoded contents of the file.
	// Required: true
	Content() *string
	SetContent(*string)

	// An ID for the content.
	// Required: true
	ContentID() *string
	SetContentID(*string)

	// Inline for content that can be displayed within the email, or attachment for any other files.
	// Required: true
	Disposition() *string
	SetDisposition(*string)

	// The name of the file being attached.
	// Required: true
	Filename() *string
	SetFilename(*string)

	// The type of file being attached.
	// Required: true
	Type() *string
	SetType(*string)
}

type attachment struct {
	contentField *string

	contentIdField *string

	dispositionField *string

	filenameField *string

	typeField *string
}

// Content gets the content of this polymorphic type
func (m *attachment) Content() *string {
	return m.contentField
}

// SetContent sets the content of this polymorphic type
func (m *attachment) SetContent(val *string) {
	m.contentField = val
}

// ContentID gets the content id of this polymorphic type
func (m *attachment) ContentID() *string {
	return m.contentIdField
}

// SetContentID sets the content id of this polymorphic type
func (m *attachment) SetContentID(val *string) {
	m.contentIdField = val
}

// Disposition gets the disposition of this polymorphic type
func (m *attachment) Disposition() *string {
	return m.dispositionField
}

// SetDisposition sets the disposition of this polymorphic type
func (m *attachment) SetDisposition(val *string) {
	m.dispositionField = val
}

// Filename gets the filename of this polymorphic type
func (m *attachment) Filename() *string {
	return m.filenameField
}

// SetFilename sets the filename of this polymorphic type
func (m *attachment) SetFilename(val *string) {
	m.filenameField = val
}

// Type gets the type of this polymorphic type
func (m *attachment) Type() *string {
	return m.typeField
}

// SetType sets the type of this polymorphic type
func (m *attachment) SetType(val *string) {
	m.typeField = val
}

// UnmarshalAttachmentSlice unmarshals polymorphic slices of Attachment
func UnmarshalAttachmentSlice(reader io.Reader, consumer runtime.Consumer) ([]Attachment, error) {
	var elements []json.RawMessage
	if err := consumer.Consume(reader, &elements); err != nil {
		return nil, err
	}

	var result []Attachment
	for _, element := range elements {
		obj, err := unmarshalAttachment(element, consumer)
		if err != nil {
			return nil, err
		}
		result = append(result, obj)
	}
	return result, nil
}

// UnmarshalAttachment unmarshals polymorphic Attachment
func UnmarshalAttachment(reader io.Reader, consumer runtime.Consumer) (Attachment, error) {
	// we need to read this twice, so first into a buffer
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return unmarshalAttachment(data, consumer)
}

func unmarshalAttachment(data []byte, consumer runtime.Consumer) (Attachment, error) {
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
	case "Attachment":
		var result attachment
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	}
	return nil, errors.New(422, "invalid classType value: %q", getType.ClassType)

}

// Validate validates this attachment
func (m *attachment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContentID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDisposition(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFilename(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *attachment) validateContent(formats strfmt.Registry) error {

	if err := validate.Required("content", "body", m.Content()); err != nil {
		return err
	}

	return nil
}

func (m *attachment) validateContentID(formats strfmt.Registry) error {

	if err := validate.Required("content_id", "body", m.ContentID()); err != nil {
		return err
	}

	return nil
}

func (m *attachment) validateDisposition(formats strfmt.Registry) error {

	if err := validate.Required("disposition", "body", m.Disposition()); err != nil {
		return err
	}

	return nil
}

func (m *attachment) validateFilename(formats strfmt.Registry) error {

	if err := validate.Required("filename", "body", m.Filename()); err != nil {
		return err
	}

	return nil
}

func (m *attachment) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type()); err != nil {
		return err
	}

	return nil
}
