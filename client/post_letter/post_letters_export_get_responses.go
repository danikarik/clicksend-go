// Code generated by go-swagger; DO NOT EDIT.

package post_letter

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// PostLettersExportGetReader is a Reader for the PostLettersExportGet structure.
type PostLettersExportGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostLettersExportGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostLettersExportGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPostLettersExportGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewPostLettersExportGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewPostLettersExportGetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPostLettersExportGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewPostLettersExportGetMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewPostLettersExportGetTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewPostLettersExportGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostLettersExportGetOK creates a PostLettersExportGetOK with default headers values
func NewPostLettersExportGetOK() *PostLettersExportGetOK {
	return &PostLettersExportGetOK{}
}

/*PostLettersExportGetOK handles this case with default header values.

SUCCESS
*/
type PostLettersExportGetOK struct {
	Payload interface{}
}

func (o *PostLettersExportGetOK) Error() string {
	return fmt.Sprintf("[GET /post/letters/history/export][%d] postLettersExportGetOK  %+v", 200, o.Payload)
}

func (o *PostLettersExportGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLettersExportGetBadRequest creates a PostLettersExportGetBadRequest with default headers values
func NewPostLettersExportGetBadRequest() *PostLettersExportGetBadRequest {
	return &PostLettersExportGetBadRequest{}
}

/*PostLettersExportGetBadRequest handles this case with default header values.

BAD_REQUEST
*/
type PostLettersExportGetBadRequest struct {
	Payload interface{}
}

func (o *PostLettersExportGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /post/letters/history/export][%d] postLettersExportGetBadRequest  %+v", 400, o.Payload)
}

func (o *PostLettersExportGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLettersExportGetUnauthorized creates a PostLettersExportGetUnauthorized with default headers values
func NewPostLettersExportGetUnauthorized() *PostLettersExportGetUnauthorized {
	return &PostLettersExportGetUnauthorized{}
}

/*PostLettersExportGetUnauthorized handles this case with default header values.

UNAUTHORIZED
*/
type PostLettersExportGetUnauthorized struct {
	Payload interface{}
}

func (o *PostLettersExportGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /post/letters/history/export][%d] postLettersExportGetUnauthorized  %+v", 401, o.Payload)
}

func (o *PostLettersExportGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLettersExportGetForbidden creates a PostLettersExportGetForbidden with default headers values
func NewPostLettersExportGetForbidden() *PostLettersExportGetForbidden {
	return &PostLettersExportGetForbidden{}
}

/*PostLettersExportGetForbidden handles this case with default header values.

FORBIDDEN
*/
type PostLettersExportGetForbidden struct {
	Payload interface{}
}

func (o *PostLettersExportGetForbidden) Error() string {
	return fmt.Sprintf("[GET /post/letters/history/export][%d] postLettersExportGetForbidden  %+v", 403, o.Payload)
}

func (o *PostLettersExportGetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLettersExportGetNotFound creates a PostLettersExportGetNotFound with default headers values
func NewPostLettersExportGetNotFound() *PostLettersExportGetNotFound {
	return &PostLettersExportGetNotFound{}
}

/*PostLettersExportGetNotFound handles this case with default header values.

NOT_FOUND
*/
type PostLettersExportGetNotFound struct {
	Payload interface{}
}

func (o *PostLettersExportGetNotFound) Error() string {
	return fmt.Sprintf("[GET /post/letters/history/export][%d] postLettersExportGetNotFound  %+v", 404, o.Payload)
}

func (o *PostLettersExportGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLettersExportGetMethodNotAllowed creates a PostLettersExportGetMethodNotAllowed with default headers values
func NewPostLettersExportGetMethodNotAllowed() *PostLettersExportGetMethodNotAllowed {
	return &PostLettersExportGetMethodNotAllowed{}
}

/*PostLettersExportGetMethodNotAllowed handles this case with default header values.

METHOD_NOT_FOUND
*/
type PostLettersExportGetMethodNotAllowed struct {
	Payload interface{}
}

func (o *PostLettersExportGetMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /post/letters/history/export][%d] postLettersExportGetMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *PostLettersExportGetMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLettersExportGetTooManyRequests creates a PostLettersExportGetTooManyRequests with default headers values
func NewPostLettersExportGetTooManyRequests() *PostLettersExportGetTooManyRequests {
	return &PostLettersExportGetTooManyRequests{}
}

/*PostLettersExportGetTooManyRequests handles this case with default header values.

TOO_MANY_REQUESTS
*/
type PostLettersExportGetTooManyRequests struct {
	Payload interface{}
}

func (o *PostLettersExportGetTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /post/letters/history/export][%d] postLettersExportGetTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostLettersExportGetTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLettersExportGetDefault creates a PostLettersExportGetDefault with default headers values
func NewPostLettersExportGetDefault(code int) *PostLettersExportGetDefault {
	return &PostLettersExportGetDefault{
		_statusCode: code,
	}
}

/*PostLettersExportGetDefault handles this case with default header values.

INTERNAL_SERVER_ERROR
*/
type PostLettersExportGetDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the post letters export get default response
func (o *PostLettersExportGetDefault) Code() int {
	return o._statusCode
}

func (o *PostLettersExportGetDefault) Error() string {
	return fmt.Sprintf("[GET /post/letters/history/export][%d] PostLettersExportGet default  %+v", o._statusCode, o.Payload)
}

func (o *PostLettersExportGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
