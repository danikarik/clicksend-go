// Code generated by go-swagger; DO NOT EDIT.

package voice_delivery_receipt_rules

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// VoiceDeliveryReceiptAutomationsGetReader is a Reader for the VoiceDeliveryReceiptAutomationsGet structure.
type VoiceDeliveryReceiptAutomationsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VoiceDeliveryReceiptAutomationsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewVoiceDeliveryReceiptAutomationsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewVoiceDeliveryReceiptAutomationsGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewVoiceDeliveryReceiptAutomationsGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewVoiceDeliveryReceiptAutomationsGetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewVoiceDeliveryReceiptAutomationsGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewVoiceDeliveryReceiptAutomationsGetMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewVoiceDeliveryReceiptAutomationsGetTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewVoiceDeliveryReceiptAutomationsGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewVoiceDeliveryReceiptAutomationsGetOK creates a VoiceDeliveryReceiptAutomationsGetOK with default headers values
func NewVoiceDeliveryReceiptAutomationsGetOK() *VoiceDeliveryReceiptAutomationsGetOK {
	return &VoiceDeliveryReceiptAutomationsGetOK{}
}

/*VoiceDeliveryReceiptAutomationsGetOK handles this case with default header values.

SUCCESS
*/
type VoiceDeliveryReceiptAutomationsGetOK struct {
	Payload string
}

func (o *VoiceDeliveryReceiptAutomationsGetOK) Error() string {
	return fmt.Sprintf("[GET /automations/voice/receipts][%d] voiceDeliveryReceiptAutomationsGetOK  %+v", 200, o.Payload)
}

func (o *VoiceDeliveryReceiptAutomationsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVoiceDeliveryReceiptAutomationsGetBadRequest creates a VoiceDeliveryReceiptAutomationsGetBadRequest with default headers values
func NewVoiceDeliveryReceiptAutomationsGetBadRequest() *VoiceDeliveryReceiptAutomationsGetBadRequest {
	return &VoiceDeliveryReceiptAutomationsGetBadRequest{}
}

/*VoiceDeliveryReceiptAutomationsGetBadRequest handles this case with default header values.

BAD_REQUEST
*/
type VoiceDeliveryReceiptAutomationsGetBadRequest struct {
	Payload string
}

func (o *VoiceDeliveryReceiptAutomationsGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /automations/voice/receipts][%d] voiceDeliveryReceiptAutomationsGetBadRequest  %+v", 400, o.Payload)
}

func (o *VoiceDeliveryReceiptAutomationsGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVoiceDeliveryReceiptAutomationsGetUnauthorized creates a VoiceDeliveryReceiptAutomationsGetUnauthorized with default headers values
func NewVoiceDeliveryReceiptAutomationsGetUnauthorized() *VoiceDeliveryReceiptAutomationsGetUnauthorized {
	return &VoiceDeliveryReceiptAutomationsGetUnauthorized{}
}

/*VoiceDeliveryReceiptAutomationsGetUnauthorized handles this case with default header values.

UNAUTHORIZED
*/
type VoiceDeliveryReceiptAutomationsGetUnauthorized struct {
	Payload string
}

func (o *VoiceDeliveryReceiptAutomationsGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /automations/voice/receipts][%d] voiceDeliveryReceiptAutomationsGetUnauthorized  %+v", 401, o.Payload)
}

func (o *VoiceDeliveryReceiptAutomationsGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVoiceDeliveryReceiptAutomationsGetForbidden creates a VoiceDeliveryReceiptAutomationsGetForbidden with default headers values
func NewVoiceDeliveryReceiptAutomationsGetForbidden() *VoiceDeliveryReceiptAutomationsGetForbidden {
	return &VoiceDeliveryReceiptAutomationsGetForbidden{}
}

/*VoiceDeliveryReceiptAutomationsGetForbidden handles this case with default header values.

FORBIDDEN
*/
type VoiceDeliveryReceiptAutomationsGetForbidden struct {
	Payload string
}

func (o *VoiceDeliveryReceiptAutomationsGetForbidden) Error() string {
	return fmt.Sprintf("[GET /automations/voice/receipts][%d] voiceDeliveryReceiptAutomationsGetForbidden  %+v", 403, o.Payload)
}

func (o *VoiceDeliveryReceiptAutomationsGetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVoiceDeliveryReceiptAutomationsGetNotFound creates a VoiceDeliveryReceiptAutomationsGetNotFound with default headers values
func NewVoiceDeliveryReceiptAutomationsGetNotFound() *VoiceDeliveryReceiptAutomationsGetNotFound {
	return &VoiceDeliveryReceiptAutomationsGetNotFound{}
}

/*VoiceDeliveryReceiptAutomationsGetNotFound handles this case with default header values.

NOT_FOUND
*/
type VoiceDeliveryReceiptAutomationsGetNotFound struct {
	Payload string
}

func (o *VoiceDeliveryReceiptAutomationsGetNotFound) Error() string {
	return fmt.Sprintf("[GET /automations/voice/receipts][%d] voiceDeliveryReceiptAutomationsGetNotFound  %+v", 404, o.Payload)
}

func (o *VoiceDeliveryReceiptAutomationsGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVoiceDeliveryReceiptAutomationsGetMethodNotAllowed creates a VoiceDeliveryReceiptAutomationsGetMethodNotAllowed with default headers values
func NewVoiceDeliveryReceiptAutomationsGetMethodNotAllowed() *VoiceDeliveryReceiptAutomationsGetMethodNotAllowed {
	return &VoiceDeliveryReceiptAutomationsGetMethodNotAllowed{}
}

/*VoiceDeliveryReceiptAutomationsGetMethodNotAllowed handles this case with default header values.

METHOD_NOT_FOUND
*/
type VoiceDeliveryReceiptAutomationsGetMethodNotAllowed struct {
	Payload string
}

func (o *VoiceDeliveryReceiptAutomationsGetMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /automations/voice/receipts][%d] voiceDeliveryReceiptAutomationsGetMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *VoiceDeliveryReceiptAutomationsGetMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVoiceDeliveryReceiptAutomationsGetTooManyRequests creates a VoiceDeliveryReceiptAutomationsGetTooManyRequests with default headers values
func NewVoiceDeliveryReceiptAutomationsGetTooManyRequests() *VoiceDeliveryReceiptAutomationsGetTooManyRequests {
	return &VoiceDeliveryReceiptAutomationsGetTooManyRequests{}
}

/*VoiceDeliveryReceiptAutomationsGetTooManyRequests handles this case with default header values.

TOO_MANY_REQUESTS
*/
type VoiceDeliveryReceiptAutomationsGetTooManyRequests struct {
	Payload string
}

func (o *VoiceDeliveryReceiptAutomationsGetTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /automations/voice/receipts][%d] voiceDeliveryReceiptAutomationsGetTooManyRequests  %+v", 429, o.Payload)
}

func (o *VoiceDeliveryReceiptAutomationsGetTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVoiceDeliveryReceiptAutomationsGetDefault creates a VoiceDeliveryReceiptAutomationsGetDefault with default headers values
func NewVoiceDeliveryReceiptAutomationsGetDefault(code int) *VoiceDeliveryReceiptAutomationsGetDefault {
	return &VoiceDeliveryReceiptAutomationsGetDefault{
		_statusCode: code,
	}
}

/*VoiceDeliveryReceiptAutomationsGetDefault handles this case with default header values.

INTERNAL_SERVER_ERROR
*/
type VoiceDeliveryReceiptAutomationsGetDefault struct {
	_statusCode int

	Payload string
}

// Code gets the status code for the voice delivery receipt automations get default response
func (o *VoiceDeliveryReceiptAutomationsGetDefault) Code() int {
	return o._statusCode
}

func (o *VoiceDeliveryReceiptAutomationsGetDefault) Error() string {
	return fmt.Sprintf("[GET /automations/voice/receipts][%d] VoiceDeliveryReceiptAutomationsGet default  %+v", o._statusCode, o.Payload)
}

func (o *VoiceDeliveryReceiptAutomationsGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
