// Code generated by go-swagger; DO NOT EDIT.

package statistics

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// StatisticsVoiceGetReader is a Reader for the StatisticsVoiceGet structure.
type StatisticsVoiceGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StatisticsVoiceGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewStatisticsVoiceGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewStatisticsVoiceGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewStatisticsVoiceGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewStatisticsVoiceGetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewStatisticsVoiceGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewStatisticsVoiceGetMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewStatisticsVoiceGetTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewStatisticsVoiceGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStatisticsVoiceGetOK creates a StatisticsVoiceGetOK with default headers values
func NewStatisticsVoiceGetOK() *StatisticsVoiceGetOK {
	return &StatisticsVoiceGetOK{}
}

/*StatisticsVoiceGetOK handles this case with default header values.

SUCCESS
*/
type StatisticsVoiceGetOK struct {
	Payload interface{}
}

func (o *StatisticsVoiceGetOK) Error() string {
	return fmt.Sprintf("[GET /statistics/voice][%d] statisticsVoiceGetOK  %+v", 200, o.Payload)
}

func (o *StatisticsVoiceGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStatisticsVoiceGetBadRequest creates a StatisticsVoiceGetBadRequest with default headers values
func NewStatisticsVoiceGetBadRequest() *StatisticsVoiceGetBadRequest {
	return &StatisticsVoiceGetBadRequest{}
}

/*StatisticsVoiceGetBadRequest handles this case with default header values.

BAD_REQUEST
*/
type StatisticsVoiceGetBadRequest struct {
	Payload interface{}
}

func (o *StatisticsVoiceGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /statistics/voice][%d] statisticsVoiceGetBadRequest  %+v", 400, o.Payload)
}

func (o *StatisticsVoiceGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStatisticsVoiceGetUnauthorized creates a StatisticsVoiceGetUnauthorized with default headers values
func NewStatisticsVoiceGetUnauthorized() *StatisticsVoiceGetUnauthorized {
	return &StatisticsVoiceGetUnauthorized{}
}

/*StatisticsVoiceGetUnauthorized handles this case with default header values.

UNAUTHORIZED
*/
type StatisticsVoiceGetUnauthorized struct {
	Payload interface{}
}

func (o *StatisticsVoiceGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /statistics/voice][%d] statisticsVoiceGetUnauthorized  %+v", 401, o.Payload)
}

func (o *StatisticsVoiceGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStatisticsVoiceGetForbidden creates a StatisticsVoiceGetForbidden with default headers values
func NewStatisticsVoiceGetForbidden() *StatisticsVoiceGetForbidden {
	return &StatisticsVoiceGetForbidden{}
}

/*StatisticsVoiceGetForbidden handles this case with default header values.

FORBIDDEN
*/
type StatisticsVoiceGetForbidden struct {
	Payload interface{}
}

func (o *StatisticsVoiceGetForbidden) Error() string {
	return fmt.Sprintf("[GET /statistics/voice][%d] statisticsVoiceGetForbidden  %+v", 403, o.Payload)
}

func (o *StatisticsVoiceGetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStatisticsVoiceGetNotFound creates a StatisticsVoiceGetNotFound with default headers values
func NewStatisticsVoiceGetNotFound() *StatisticsVoiceGetNotFound {
	return &StatisticsVoiceGetNotFound{}
}

/*StatisticsVoiceGetNotFound handles this case with default header values.

NOT_FOUND
*/
type StatisticsVoiceGetNotFound struct {
	Payload interface{}
}

func (o *StatisticsVoiceGetNotFound) Error() string {
	return fmt.Sprintf("[GET /statistics/voice][%d] statisticsVoiceGetNotFound  %+v", 404, o.Payload)
}

func (o *StatisticsVoiceGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStatisticsVoiceGetMethodNotAllowed creates a StatisticsVoiceGetMethodNotAllowed with default headers values
func NewStatisticsVoiceGetMethodNotAllowed() *StatisticsVoiceGetMethodNotAllowed {
	return &StatisticsVoiceGetMethodNotAllowed{}
}

/*StatisticsVoiceGetMethodNotAllowed handles this case with default header values.

METHOD_NOT_FOUND
*/
type StatisticsVoiceGetMethodNotAllowed struct {
	Payload interface{}
}

func (o *StatisticsVoiceGetMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /statistics/voice][%d] statisticsVoiceGetMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *StatisticsVoiceGetMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStatisticsVoiceGetTooManyRequests creates a StatisticsVoiceGetTooManyRequests with default headers values
func NewStatisticsVoiceGetTooManyRequests() *StatisticsVoiceGetTooManyRequests {
	return &StatisticsVoiceGetTooManyRequests{}
}

/*StatisticsVoiceGetTooManyRequests handles this case with default header values.

TOO_MANY_REQUESTS
*/
type StatisticsVoiceGetTooManyRequests struct {
	Payload interface{}
}

func (o *StatisticsVoiceGetTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /statistics/voice][%d] statisticsVoiceGetTooManyRequests  %+v", 429, o.Payload)
}

func (o *StatisticsVoiceGetTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStatisticsVoiceGetDefault creates a StatisticsVoiceGetDefault with default headers values
func NewStatisticsVoiceGetDefault(code int) *StatisticsVoiceGetDefault {
	return &StatisticsVoiceGetDefault{
		_statusCode: code,
	}
}

/*StatisticsVoiceGetDefault handles this case with default header values.

INTERNAL_SERVER_ERROR
*/
type StatisticsVoiceGetDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the statistics voice get default response
func (o *StatisticsVoiceGetDefault) Code() int {
	return o._statusCode
}

func (o *StatisticsVoiceGetDefault) Error() string {
	return fmt.Sprintf("[GET /statistics/voice][%d] StatisticsVoiceGet default  %+v", o._statusCode, o.Payload)
}

func (o *StatisticsVoiceGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
