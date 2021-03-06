// Code generated by go-swagger; DO NOT EDIT.

package account_recharge

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// RechargeTransactionsByTransactionIDGetReader is a Reader for the RechargeTransactionsByTransactionIDGet structure.
type RechargeTransactionsByTransactionIDGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RechargeTransactionsByTransactionIDGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewRechargeTransactionsByTransactionIDGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewRechargeTransactionsByTransactionIDGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewRechargeTransactionsByTransactionIDGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewRechargeTransactionsByTransactionIDGetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewRechargeTransactionsByTransactionIDGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewRechargeTransactionsByTransactionIDGetMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewRechargeTransactionsByTransactionIDGetTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewRechargeTransactionsByTransactionIDGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRechargeTransactionsByTransactionIDGetOK creates a RechargeTransactionsByTransactionIDGetOK with default headers values
func NewRechargeTransactionsByTransactionIDGetOK() *RechargeTransactionsByTransactionIDGetOK {
	return &RechargeTransactionsByTransactionIDGetOK{}
}

/*RechargeTransactionsByTransactionIDGetOK handles this case with default header values.

SUCCESS
*/
type RechargeTransactionsByTransactionIDGetOK struct {
	Payload interface{}
}

func (o *RechargeTransactionsByTransactionIDGetOK) Error() string {
	return fmt.Sprintf("[GET /recharge/transactions/{transaction_id}][%d] rechargeTransactionsByTransactionIdGetOK  %+v", 200, o.Payload)
}

func (o *RechargeTransactionsByTransactionIDGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRechargeTransactionsByTransactionIDGetBadRequest creates a RechargeTransactionsByTransactionIDGetBadRequest with default headers values
func NewRechargeTransactionsByTransactionIDGetBadRequest() *RechargeTransactionsByTransactionIDGetBadRequest {
	return &RechargeTransactionsByTransactionIDGetBadRequest{}
}

/*RechargeTransactionsByTransactionIDGetBadRequest handles this case with default header values.

BAD_REQUEST
*/
type RechargeTransactionsByTransactionIDGetBadRequest struct {
	Payload interface{}
}

func (o *RechargeTransactionsByTransactionIDGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /recharge/transactions/{transaction_id}][%d] rechargeTransactionsByTransactionIdGetBadRequest  %+v", 400, o.Payload)
}

func (o *RechargeTransactionsByTransactionIDGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRechargeTransactionsByTransactionIDGetUnauthorized creates a RechargeTransactionsByTransactionIDGetUnauthorized with default headers values
func NewRechargeTransactionsByTransactionIDGetUnauthorized() *RechargeTransactionsByTransactionIDGetUnauthorized {
	return &RechargeTransactionsByTransactionIDGetUnauthorized{}
}

/*RechargeTransactionsByTransactionIDGetUnauthorized handles this case with default header values.

UNAUTHORIZED
*/
type RechargeTransactionsByTransactionIDGetUnauthorized struct {
	Payload interface{}
}

func (o *RechargeTransactionsByTransactionIDGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /recharge/transactions/{transaction_id}][%d] rechargeTransactionsByTransactionIdGetUnauthorized  %+v", 401, o.Payload)
}

func (o *RechargeTransactionsByTransactionIDGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRechargeTransactionsByTransactionIDGetForbidden creates a RechargeTransactionsByTransactionIDGetForbidden with default headers values
func NewRechargeTransactionsByTransactionIDGetForbidden() *RechargeTransactionsByTransactionIDGetForbidden {
	return &RechargeTransactionsByTransactionIDGetForbidden{}
}

/*RechargeTransactionsByTransactionIDGetForbidden handles this case with default header values.

FORBIDDEN
*/
type RechargeTransactionsByTransactionIDGetForbidden struct {
	Payload interface{}
}

func (o *RechargeTransactionsByTransactionIDGetForbidden) Error() string {
	return fmt.Sprintf("[GET /recharge/transactions/{transaction_id}][%d] rechargeTransactionsByTransactionIdGetForbidden  %+v", 403, o.Payload)
}

func (o *RechargeTransactionsByTransactionIDGetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRechargeTransactionsByTransactionIDGetNotFound creates a RechargeTransactionsByTransactionIDGetNotFound with default headers values
func NewRechargeTransactionsByTransactionIDGetNotFound() *RechargeTransactionsByTransactionIDGetNotFound {
	return &RechargeTransactionsByTransactionIDGetNotFound{}
}

/*RechargeTransactionsByTransactionIDGetNotFound handles this case with default header values.

NOT_FOUND
*/
type RechargeTransactionsByTransactionIDGetNotFound struct {
	Payload interface{}
}

func (o *RechargeTransactionsByTransactionIDGetNotFound) Error() string {
	return fmt.Sprintf("[GET /recharge/transactions/{transaction_id}][%d] rechargeTransactionsByTransactionIdGetNotFound  %+v", 404, o.Payload)
}

func (o *RechargeTransactionsByTransactionIDGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRechargeTransactionsByTransactionIDGetMethodNotAllowed creates a RechargeTransactionsByTransactionIDGetMethodNotAllowed with default headers values
func NewRechargeTransactionsByTransactionIDGetMethodNotAllowed() *RechargeTransactionsByTransactionIDGetMethodNotAllowed {
	return &RechargeTransactionsByTransactionIDGetMethodNotAllowed{}
}

/*RechargeTransactionsByTransactionIDGetMethodNotAllowed handles this case with default header values.

METHOD_NOT_FOUND
*/
type RechargeTransactionsByTransactionIDGetMethodNotAllowed struct {
	Payload interface{}
}

func (o *RechargeTransactionsByTransactionIDGetMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /recharge/transactions/{transaction_id}][%d] rechargeTransactionsByTransactionIdGetMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *RechargeTransactionsByTransactionIDGetMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRechargeTransactionsByTransactionIDGetTooManyRequests creates a RechargeTransactionsByTransactionIDGetTooManyRequests with default headers values
func NewRechargeTransactionsByTransactionIDGetTooManyRequests() *RechargeTransactionsByTransactionIDGetTooManyRequests {
	return &RechargeTransactionsByTransactionIDGetTooManyRequests{}
}

/*RechargeTransactionsByTransactionIDGetTooManyRequests handles this case with default header values.

TOO_MANY_REQUESTS
*/
type RechargeTransactionsByTransactionIDGetTooManyRequests struct {
	Payload interface{}
}

func (o *RechargeTransactionsByTransactionIDGetTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /recharge/transactions/{transaction_id}][%d] rechargeTransactionsByTransactionIdGetTooManyRequests  %+v", 429, o.Payload)
}

func (o *RechargeTransactionsByTransactionIDGetTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRechargeTransactionsByTransactionIDGetDefault creates a RechargeTransactionsByTransactionIDGetDefault with default headers values
func NewRechargeTransactionsByTransactionIDGetDefault(code int) *RechargeTransactionsByTransactionIDGetDefault {
	return &RechargeTransactionsByTransactionIDGetDefault{
		_statusCode: code,
	}
}

/*RechargeTransactionsByTransactionIDGetDefault handles this case with default header values.

INTERNAL_SERVER_ERROR
*/
type RechargeTransactionsByTransactionIDGetDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the recharge transactions by transaction Id get default response
func (o *RechargeTransactionsByTransactionIDGetDefault) Code() int {
	return o._statusCode
}

func (o *RechargeTransactionsByTransactionIDGetDefault) Error() string {
	return fmt.Sprintf("[GET /recharge/transactions/{transaction_id}][%d] RechargeTransactionsByTransactionIdGet default  %+v", o._statusCode, o.Payload)
}

func (o *RechargeTransactionsByTransactionIDGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
