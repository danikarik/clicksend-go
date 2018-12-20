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

// PostLettersSendPostReader is a Reader for the PostLettersSendPost structure.
type PostLettersSendPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostLettersSendPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostLettersSendPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPostLettersSendPostBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewPostLettersSendPostUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewPostLettersSendPostForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPostLettersSendPostNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewPostLettersSendPostMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewPostLettersSendPostTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewPostLettersSendPostDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostLettersSendPostOK creates a PostLettersSendPostOK with default headers values
func NewPostLettersSendPostOK() *PostLettersSendPostOK {
	return &PostLettersSendPostOK{}
}

/*PostLettersSendPostOK handles this case with default header values.

SUCCESS
*/
type PostLettersSendPostOK struct {
	Payload string
}

func (o *PostLettersSendPostOK) Error() string {
	return fmt.Sprintf("[POST /post/letters/send][%d] postLettersSendPostOK  %+v", 200, o.Payload)
}

func (o *PostLettersSendPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLettersSendPostBadRequest creates a PostLettersSendPostBadRequest with default headers values
func NewPostLettersSendPostBadRequest() *PostLettersSendPostBadRequest {
	return &PostLettersSendPostBadRequest{}
}

/*PostLettersSendPostBadRequest handles this case with default header values.

BAD_REQUEST
*/
type PostLettersSendPostBadRequest struct {
	Payload string
}

func (o *PostLettersSendPostBadRequest) Error() string {
	return fmt.Sprintf("[POST /post/letters/send][%d] postLettersSendPostBadRequest  %+v", 400, o.Payload)
}

func (o *PostLettersSendPostBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLettersSendPostUnauthorized creates a PostLettersSendPostUnauthorized with default headers values
func NewPostLettersSendPostUnauthorized() *PostLettersSendPostUnauthorized {
	return &PostLettersSendPostUnauthorized{}
}

/*PostLettersSendPostUnauthorized handles this case with default header values.

UNAUTHORIZED
*/
type PostLettersSendPostUnauthorized struct {
	Payload string
}

func (o *PostLettersSendPostUnauthorized) Error() string {
	return fmt.Sprintf("[POST /post/letters/send][%d] postLettersSendPostUnauthorized  %+v", 401, o.Payload)
}

func (o *PostLettersSendPostUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLettersSendPostForbidden creates a PostLettersSendPostForbidden with default headers values
func NewPostLettersSendPostForbidden() *PostLettersSendPostForbidden {
	return &PostLettersSendPostForbidden{}
}

/*PostLettersSendPostForbidden handles this case with default header values.

FORBIDDEN
*/
type PostLettersSendPostForbidden struct {
	Payload string
}

func (o *PostLettersSendPostForbidden) Error() string {
	return fmt.Sprintf("[POST /post/letters/send][%d] postLettersSendPostForbidden  %+v", 403, o.Payload)
}

func (o *PostLettersSendPostForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLettersSendPostNotFound creates a PostLettersSendPostNotFound with default headers values
func NewPostLettersSendPostNotFound() *PostLettersSendPostNotFound {
	return &PostLettersSendPostNotFound{}
}

/*PostLettersSendPostNotFound handles this case with default header values.

NOT_FOUND
*/
type PostLettersSendPostNotFound struct {
	Payload string
}

func (o *PostLettersSendPostNotFound) Error() string {
	return fmt.Sprintf("[POST /post/letters/send][%d] postLettersSendPostNotFound  %+v", 404, o.Payload)
}

func (o *PostLettersSendPostNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLettersSendPostMethodNotAllowed creates a PostLettersSendPostMethodNotAllowed with default headers values
func NewPostLettersSendPostMethodNotAllowed() *PostLettersSendPostMethodNotAllowed {
	return &PostLettersSendPostMethodNotAllowed{}
}

/*PostLettersSendPostMethodNotAllowed handles this case with default header values.

METHOD_NOT_FOUND
*/
type PostLettersSendPostMethodNotAllowed struct {
	Payload string
}

func (o *PostLettersSendPostMethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /post/letters/send][%d] postLettersSendPostMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *PostLettersSendPostMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLettersSendPostTooManyRequests creates a PostLettersSendPostTooManyRequests with default headers values
func NewPostLettersSendPostTooManyRequests() *PostLettersSendPostTooManyRequests {
	return &PostLettersSendPostTooManyRequests{}
}

/*PostLettersSendPostTooManyRequests handles this case with default header values.

TOO_MANY_REQUESTS
*/
type PostLettersSendPostTooManyRequests struct {
	Payload string
}

func (o *PostLettersSendPostTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /post/letters/send][%d] postLettersSendPostTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostLettersSendPostTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLettersSendPostDefault creates a PostLettersSendPostDefault with default headers values
func NewPostLettersSendPostDefault(code int) *PostLettersSendPostDefault {
	return &PostLettersSendPostDefault{
		_statusCode: code,
	}
}

/*PostLettersSendPostDefault handles this case with default header values.

INTERNAL_SERVER_ERROR
*/
type PostLettersSendPostDefault struct {
	_statusCode int

	Payload string
}

// Code gets the status code for the post letters send post default response
func (o *PostLettersSendPostDefault) Code() int {
	return o._statusCode
}

func (o *PostLettersSendPostDefault) Error() string {
	return fmt.Sprintf("[POST /post/letters/send][%d] PostLettersSendPost default  %+v", o._statusCode, o.Payload)
}

func (o *PostLettersSendPostDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
