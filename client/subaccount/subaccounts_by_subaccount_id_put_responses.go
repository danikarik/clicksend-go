// Code generated by go-swagger; DO NOT EDIT.

package subaccount

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// SubaccountsBySubaccountIDPutReader is a Reader for the SubaccountsBySubaccountIDPut structure.
type SubaccountsBySubaccountIDPutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SubaccountsBySubaccountIDPutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSubaccountsBySubaccountIDPutOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewSubaccountsBySubaccountIDPutBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewSubaccountsBySubaccountIDPutUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewSubaccountsBySubaccountIDPutForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewSubaccountsBySubaccountIDPutNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewSubaccountsBySubaccountIDPutMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewSubaccountsBySubaccountIDPutTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewSubaccountsBySubaccountIDPutDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSubaccountsBySubaccountIDPutOK creates a SubaccountsBySubaccountIDPutOK with default headers values
func NewSubaccountsBySubaccountIDPutOK() *SubaccountsBySubaccountIDPutOK {
	return &SubaccountsBySubaccountIDPutOK{}
}

/*SubaccountsBySubaccountIDPutOK handles this case with default header values.

SUCCESS
*/
type SubaccountsBySubaccountIDPutOK struct {
	Payload string
}

func (o *SubaccountsBySubaccountIDPutOK) Error() string {
	return fmt.Sprintf("[PUT /subaccounts/{subaccount_id}][%d] subaccountsBySubaccountIdPutOK  %+v", 200, o.Payload)
}

func (o *SubaccountsBySubaccountIDPutOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSubaccountsBySubaccountIDPutBadRequest creates a SubaccountsBySubaccountIDPutBadRequest with default headers values
func NewSubaccountsBySubaccountIDPutBadRequest() *SubaccountsBySubaccountIDPutBadRequest {
	return &SubaccountsBySubaccountIDPutBadRequest{}
}

/*SubaccountsBySubaccountIDPutBadRequest handles this case with default header values.

BAD_REQUEST
*/
type SubaccountsBySubaccountIDPutBadRequest struct {
	Payload string
}

func (o *SubaccountsBySubaccountIDPutBadRequest) Error() string {
	return fmt.Sprintf("[PUT /subaccounts/{subaccount_id}][%d] subaccountsBySubaccountIdPutBadRequest  %+v", 400, o.Payload)
}

func (o *SubaccountsBySubaccountIDPutBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSubaccountsBySubaccountIDPutUnauthorized creates a SubaccountsBySubaccountIDPutUnauthorized with default headers values
func NewSubaccountsBySubaccountIDPutUnauthorized() *SubaccountsBySubaccountIDPutUnauthorized {
	return &SubaccountsBySubaccountIDPutUnauthorized{}
}

/*SubaccountsBySubaccountIDPutUnauthorized handles this case with default header values.

UNAUTHORIZED
*/
type SubaccountsBySubaccountIDPutUnauthorized struct {
	Payload string
}

func (o *SubaccountsBySubaccountIDPutUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /subaccounts/{subaccount_id}][%d] subaccountsBySubaccountIdPutUnauthorized  %+v", 401, o.Payload)
}

func (o *SubaccountsBySubaccountIDPutUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSubaccountsBySubaccountIDPutForbidden creates a SubaccountsBySubaccountIDPutForbidden with default headers values
func NewSubaccountsBySubaccountIDPutForbidden() *SubaccountsBySubaccountIDPutForbidden {
	return &SubaccountsBySubaccountIDPutForbidden{}
}

/*SubaccountsBySubaccountIDPutForbidden handles this case with default header values.

FORBIDDEN
*/
type SubaccountsBySubaccountIDPutForbidden struct {
	Payload string
}

func (o *SubaccountsBySubaccountIDPutForbidden) Error() string {
	return fmt.Sprintf("[PUT /subaccounts/{subaccount_id}][%d] subaccountsBySubaccountIdPutForbidden  %+v", 403, o.Payload)
}

func (o *SubaccountsBySubaccountIDPutForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSubaccountsBySubaccountIDPutNotFound creates a SubaccountsBySubaccountIDPutNotFound with default headers values
func NewSubaccountsBySubaccountIDPutNotFound() *SubaccountsBySubaccountIDPutNotFound {
	return &SubaccountsBySubaccountIDPutNotFound{}
}

/*SubaccountsBySubaccountIDPutNotFound handles this case with default header values.

NOT_FOUND
*/
type SubaccountsBySubaccountIDPutNotFound struct {
	Payload string
}

func (o *SubaccountsBySubaccountIDPutNotFound) Error() string {
	return fmt.Sprintf("[PUT /subaccounts/{subaccount_id}][%d] subaccountsBySubaccountIdPutNotFound  %+v", 404, o.Payload)
}

func (o *SubaccountsBySubaccountIDPutNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSubaccountsBySubaccountIDPutMethodNotAllowed creates a SubaccountsBySubaccountIDPutMethodNotAllowed with default headers values
func NewSubaccountsBySubaccountIDPutMethodNotAllowed() *SubaccountsBySubaccountIDPutMethodNotAllowed {
	return &SubaccountsBySubaccountIDPutMethodNotAllowed{}
}

/*SubaccountsBySubaccountIDPutMethodNotAllowed handles this case with default header values.

METHOD_NOT_FOUND
*/
type SubaccountsBySubaccountIDPutMethodNotAllowed struct {
	Payload string
}

func (o *SubaccountsBySubaccountIDPutMethodNotAllowed) Error() string {
	return fmt.Sprintf("[PUT /subaccounts/{subaccount_id}][%d] subaccountsBySubaccountIdPutMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *SubaccountsBySubaccountIDPutMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSubaccountsBySubaccountIDPutTooManyRequests creates a SubaccountsBySubaccountIDPutTooManyRequests with default headers values
func NewSubaccountsBySubaccountIDPutTooManyRequests() *SubaccountsBySubaccountIDPutTooManyRequests {
	return &SubaccountsBySubaccountIDPutTooManyRequests{}
}

/*SubaccountsBySubaccountIDPutTooManyRequests handles this case with default header values.

TOO_MANY_REQUESTS
*/
type SubaccountsBySubaccountIDPutTooManyRequests struct {
	Payload string
}

func (o *SubaccountsBySubaccountIDPutTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /subaccounts/{subaccount_id}][%d] subaccountsBySubaccountIdPutTooManyRequests  %+v", 429, o.Payload)
}

func (o *SubaccountsBySubaccountIDPutTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSubaccountsBySubaccountIDPutDefault creates a SubaccountsBySubaccountIDPutDefault with default headers values
func NewSubaccountsBySubaccountIDPutDefault(code int) *SubaccountsBySubaccountIDPutDefault {
	return &SubaccountsBySubaccountIDPutDefault{
		_statusCode: code,
	}
}

/*SubaccountsBySubaccountIDPutDefault handles this case with default header values.

INTERNAL_SERVER_ERROR
*/
type SubaccountsBySubaccountIDPutDefault struct {
	_statusCode int

	Payload string
}

// Code gets the status code for the subaccounts by subaccount Id put default response
func (o *SubaccountsBySubaccountIDPutDefault) Code() int {
	return o._statusCode
}

func (o *SubaccountsBySubaccountIDPutDefault) Error() string {
	return fmt.Sprintf("[PUT /subaccounts/{subaccount_id}][%d] SubaccountsBySubaccountIdPut default  %+v", o._statusCode, o.Payload)
}

func (o *SubaccountsBySubaccountIDPutDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}