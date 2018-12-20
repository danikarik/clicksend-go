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

// PostLettersPricePostReader is a Reader for the PostLettersPricePost structure.
type PostLettersPricePostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostLettersPricePostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostLettersPricePostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPostLettersPricePostBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewPostLettersPricePostUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewPostLettersPricePostForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPostLettersPricePostNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewPostLettersPricePostMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewPostLettersPricePostTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewPostLettersPricePostDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostLettersPricePostOK creates a PostLettersPricePostOK with default headers values
func NewPostLettersPricePostOK() *PostLettersPricePostOK {
	return &PostLettersPricePostOK{}
}

/*PostLettersPricePostOK handles this case with default header values.

SUCCESS
*/
type PostLettersPricePostOK struct {
	Payload string
}

func (o *PostLettersPricePostOK) Error() string {
	return fmt.Sprintf("[POST /post/letters/price][%d] postLettersPricePostOK  %+v", 200, o.Payload)
}

func (o *PostLettersPricePostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLettersPricePostBadRequest creates a PostLettersPricePostBadRequest with default headers values
func NewPostLettersPricePostBadRequest() *PostLettersPricePostBadRequest {
	return &PostLettersPricePostBadRequest{}
}

/*PostLettersPricePostBadRequest handles this case with default header values.

BAD_REQUEST
*/
type PostLettersPricePostBadRequest struct {
	Payload string
}

func (o *PostLettersPricePostBadRequest) Error() string {
	return fmt.Sprintf("[POST /post/letters/price][%d] postLettersPricePostBadRequest  %+v", 400, o.Payload)
}

func (o *PostLettersPricePostBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLettersPricePostUnauthorized creates a PostLettersPricePostUnauthorized with default headers values
func NewPostLettersPricePostUnauthorized() *PostLettersPricePostUnauthorized {
	return &PostLettersPricePostUnauthorized{}
}

/*PostLettersPricePostUnauthorized handles this case with default header values.

UNAUTHORIZED
*/
type PostLettersPricePostUnauthorized struct {
	Payload string
}

func (o *PostLettersPricePostUnauthorized) Error() string {
	return fmt.Sprintf("[POST /post/letters/price][%d] postLettersPricePostUnauthorized  %+v", 401, o.Payload)
}

func (o *PostLettersPricePostUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLettersPricePostForbidden creates a PostLettersPricePostForbidden with default headers values
func NewPostLettersPricePostForbidden() *PostLettersPricePostForbidden {
	return &PostLettersPricePostForbidden{}
}

/*PostLettersPricePostForbidden handles this case with default header values.

FORBIDDEN
*/
type PostLettersPricePostForbidden struct {
	Payload string
}

func (o *PostLettersPricePostForbidden) Error() string {
	return fmt.Sprintf("[POST /post/letters/price][%d] postLettersPricePostForbidden  %+v", 403, o.Payload)
}

func (o *PostLettersPricePostForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLettersPricePostNotFound creates a PostLettersPricePostNotFound with default headers values
func NewPostLettersPricePostNotFound() *PostLettersPricePostNotFound {
	return &PostLettersPricePostNotFound{}
}

/*PostLettersPricePostNotFound handles this case with default header values.

NOT_FOUND
*/
type PostLettersPricePostNotFound struct {
	Payload string
}

func (o *PostLettersPricePostNotFound) Error() string {
	return fmt.Sprintf("[POST /post/letters/price][%d] postLettersPricePostNotFound  %+v", 404, o.Payload)
}

func (o *PostLettersPricePostNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLettersPricePostMethodNotAllowed creates a PostLettersPricePostMethodNotAllowed with default headers values
func NewPostLettersPricePostMethodNotAllowed() *PostLettersPricePostMethodNotAllowed {
	return &PostLettersPricePostMethodNotAllowed{}
}

/*PostLettersPricePostMethodNotAllowed handles this case with default header values.

METHOD_NOT_FOUND
*/
type PostLettersPricePostMethodNotAllowed struct {
	Payload string
}

func (o *PostLettersPricePostMethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /post/letters/price][%d] postLettersPricePostMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *PostLettersPricePostMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLettersPricePostTooManyRequests creates a PostLettersPricePostTooManyRequests with default headers values
func NewPostLettersPricePostTooManyRequests() *PostLettersPricePostTooManyRequests {
	return &PostLettersPricePostTooManyRequests{}
}

/*PostLettersPricePostTooManyRequests handles this case with default header values.

TOO_MANY_REQUESTS
*/
type PostLettersPricePostTooManyRequests struct {
	Payload string
}

func (o *PostLettersPricePostTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /post/letters/price][%d] postLettersPricePostTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostLettersPricePostTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLettersPricePostDefault creates a PostLettersPricePostDefault with default headers values
func NewPostLettersPricePostDefault(code int) *PostLettersPricePostDefault {
	return &PostLettersPricePostDefault{
		_statusCode: code,
	}
}

/*PostLettersPricePostDefault handles this case with default header values.

INTERNAL_SERVER_ERROR
*/
type PostLettersPricePostDefault struct {
	_statusCode int

	Payload string
}

// Code gets the status code for the post letters price post default response
func (o *PostLettersPricePostDefault) Code() int {
	return o._statusCode
}

func (o *PostLettersPricePostDefault) Error() string {
	return fmt.Sprintf("[POST /post/letters/price][%d] PostLettersPricePost default  %+v", o._statusCode, o.Payload)
}

func (o *PostLettersPricePostDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}