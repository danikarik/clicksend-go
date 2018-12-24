// Code generated by go-swagger; DO NOT EDIT.

package referral_account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// ReferralAccountsGetReader is a Reader for the ReferralAccountsGet structure.
type ReferralAccountsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReferralAccountsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewReferralAccountsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewReferralAccountsGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewReferralAccountsGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewReferralAccountsGetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewReferralAccountsGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewReferralAccountsGetMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewReferralAccountsGetTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewReferralAccountsGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewReferralAccountsGetOK creates a ReferralAccountsGetOK with default headers values
func NewReferralAccountsGetOK() *ReferralAccountsGetOK {
	return &ReferralAccountsGetOK{}
}

/*ReferralAccountsGetOK handles this case with default header values.

SUCCESS
*/
type ReferralAccountsGetOK struct {
	Payload interface{}
}

func (o *ReferralAccountsGetOK) Error() string {
	return fmt.Sprintf("[GET /referral/accounts][%d] referralAccountsGetOK  %+v", 200, o.Payload)
}

func (o *ReferralAccountsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReferralAccountsGetBadRequest creates a ReferralAccountsGetBadRequest with default headers values
func NewReferralAccountsGetBadRequest() *ReferralAccountsGetBadRequest {
	return &ReferralAccountsGetBadRequest{}
}

/*ReferralAccountsGetBadRequest handles this case with default header values.

BAD_REQUEST
*/
type ReferralAccountsGetBadRequest struct {
	Payload interface{}
}

func (o *ReferralAccountsGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /referral/accounts][%d] referralAccountsGetBadRequest  %+v", 400, o.Payload)
}

func (o *ReferralAccountsGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReferralAccountsGetUnauthorized creates a ReferralAccountsGetUnauthorized with default headers values
func NewReferralAccountsGetUnauthorized() *ReferralAccountsGetUnauthorized {
	return &ReferralAccountsGetUnauthorized{}
}

/*ReferralAccountsGetUnauthorized handles this case with default header values.

UNAUTHORIZED
*/
type ReferralAccountsGetUnauthorized struct {
	Payload interface{}
}

func (o *ReferralAccountsGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /referral/accounts][%d] referralAccountsGetUnauthorized  %+v", 401, o.Payload)
}

func (o *ReferralAccountsGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReferralAccountsGetForbidden creates a ReferralAccountsGetForbidden with default headers values
func NewReferralAccountsGetForbidden() *ReferralAccountsGetForbidden {
	return &ReferralAccountsGetForbidden{}
}

/*ReferralAccountsGetForbidden handles this case with default header values.

FORBIDDEN
*/
type ReferralAccountsGetForbidden struct {
	Payload interface{}
}

func (o *ReferralAccountsGetForbidden) Error() string {
	return fmt.Sprintf("[GET /referral/accounts][%d] referralAccountsGetForbidden  %+v", 403, o.Payload)
}

func (o *ReferralAccountsGetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReferralAccountsGetNotFound creates a ReferralAccountsGetNotFound with default headers values
func NewReferralAccountsGetNotFound() *ReferralAccountsGetNotFound {
	return &ReferralAccountsGetNotFound{}
}

/*ReferralAccountsGetNotFound handles this case with default header values.

NOT_FOUND
*/
type ReferralAccountsGetNotFound struct {
	Payload interface{}
}

func (o *ReferralAccountsGetNotFound) Error() string {
	return fmt.Sprintf("[GET /referral/accounts][%d] referralAccountsGetNotFound  %+v", 404, o.Payload)
}

func (o *ReferralAccountsGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReferralAccountsGetMethodNotAllowed creates a ReferralAccountsGetMethodNotAllowed with default headers values
func NewReferralAccountsGetMethodNotAllowed() *ReferralAccountsGetMethodNotAllowed {
	return &ReferralAccountsGetMethodNotAllowed{}
}

/*ReferralAccountsGetMethodNotAllowed handles this case with default header values.

METHOD_NOT_FOUND
*/
type ReferralAccountsGetMethodNotAllowed struct {
	Payload interface{}
}

func (o *ReferralAccountsGetMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /referral/accounts][%d] referralAccountsGetMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *ReferralAccountsGetMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReferralAccountsGetTooManyRequests creates a ReferralAccountsGetTooManyRequests with default headers values
func NewReferralAccountsGetTooManyRequests() *ReferralAccountsGetTooManyRequests {
	return &ReferralAccountsGetTooManyRequests{}
}

/*ReferralAccountsGetTooManyRequests handles this case with default header values.

TOO_MANY_REQUESTS
*/
type ReferralAccountsGetTooManyRequests struct {
	Payload interface{}
}

func (o *ReferralAccountsGetTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /referral/accounts][%d] referralAccountsGetTooManyRequests  %+v", 429, o.Payload)
}

func (o *ReferralAccountsGetTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReferralAccountsGetDefault creates a ReferralAccountsGetDefault with default headers values
func NewReferralAccountsGetDefault(code int) *ReferralAccountsGetDefault {
	return &ReferralAccountsGetDefault{
		_statusCode: code,
	}
}

/*ReferralAccountsGetDefault handles this case with default header values.

INTERNAL_SERVER_ERROR
*/
type ReferralAccountsGetDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the referral accounts get default response
func (o *ReferralAccountsGetDefault) Code() int {
	return o._statusCode
}

func (o *ReferralAccountsGetDefault) Error() string {
	return fmt.Sprintf("[GET /referral/accounts][%d] ReferralAccountsGet default  %+v", o._statusCode, o.Payload)
}

func (o *ReferralAccountsGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
