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

// VoiceDeliveryReceiptAutomationDeleteReader is a Reader for the VoiceDeliveryReceiptAutomationDelete structure.
type VoiceDeliveryReceiptAutomationDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VoiceDeliveryReceiptAutomationDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewVoiceDeliveryReceiptAutomationDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewVoiceDeliveryReceiptAutomationDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewVoiceDeliveryReceiptAutomationDeleteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewVoiceDeliveryReceiptAutomationDeleteForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewVoiceDeliveryReceiptAutomationDeleteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewVoiceDeliveryReceiptAutomationDeleteMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 429:
		result := NewVoiceDeliveryReceiptAutomationDeleteTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewVoiceDeliveryReceiptAutomationDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewVoiceDeliveryReceiptAutomationDeleteOK creates a VoiceDeliveryReceiptAutomationDeleteOK with default headers values
func NewVoiceDeliveryReceiptAutomationDeleteOK() *VoiceDeliveryReceiptAutomationDeleteOK {
	return &VoiceDeliveryReceiptAutomationDeleteOK{}
}

/*VoiceDeliveryReceiptAutomationDeleteOK handles this case with default header values.

SUCCESS
*/
type VoiceDeliveryReceiptAutomationDeleteOK struct {
	Payload interface{}
}

func (o *VoiceDeliveryReceiptAutomationDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /automations/voice/receipts/{receipt_rule_id}][%d] voiceDeliveryReceiptAutomationDeleteOK  %+v", 200, o.Payload)
}

func (o *VoiceDeliveryReceiptAutomationDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVoiceDeliveryReceiptAutomationDeleteBadRequest creates a VoiceDeliveryReceiptAutomationDeleteBadRequest with default headers values
func NewVoiceDeliveryReceiptAutomationDeleteBadRequest() *VoiceDeliveryReceiptAutomationDeleteBadRequest {
	return &VoiceDeliveryReceiptAutomationDeleteBadRequest{}
}

/*VoiceDeliveryReceiptAutomationDeleteBadRequest handles this case with default header values.

BAD_REQUEST
*/
type VoiceDeliveryReceiptAutomationDeleteBadRequest struct {
	Payload interface{}
}

func (o *VoiceDeliveryReceiptAutomationDeleteBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /automations/voice/receipts/{receipt_rule_id}][%d] voiceDeliveryReceiptAutomationDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *VoiceDeliveryReceiptAutomationDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVoiceDeliveryReceiptAutomationDeleteUnauthorized creates a VoiceDeliveryReceiptAutomationDeleteUnauthorized with default headers values
func NewVoiceDeliveryReceiptAutomationDeleteUnauthorized() *VoiceDeliveryReceiptAutomationDeleteUnauthorized {
	return &VoiceDeliveryReceiptAutomationDeleteUnauthorized{}
}

/*VoiceDeliveryReceiptAutomationDeleteUnauthorized handles this case with default header values.

UNAUTHORIZED
*/
type VoiceDeliveryReceiptAutomationDeleteUnauthorized struct {
	Payload interface{}
}

func (o *VoiceDeliveryReceiptAutomationDeleteUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /automations/voice/receipts/{receipt_rule_id}][%d] voiceDeliveryReceiptAutomationDeleteUnauthorized  %+v", 401, o.Payload)
}

func (o *VoiceDeliveryReceiptAutomationDeleteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVoiceDeliveryReceiptAutomationDeleteForbidden creates a VoiceDeliveryReceiptAutomationDeleteForbidden with default headers values
func NewVoiceDeliveryReceiptAutomationDeleteForbidden() *VoiceDeliveryReceiptAutomationDeleteForbidden {
	return &VoiceDeliveryReceiptAutomationDeleteForbidden{}
}

/*VoiceDeliveryReceiptAutomationDeleteForbidden handles this case with default header values.

FORBIDDEN
*/
type VoiceDeliveryReceiptAutomationDeleteForbidden struct {
	Payload interface{}
}

func (o *VoiceDeliveryReceiptAutomationDeleteForbidden) Error() string {
	return fmt.Sprintf("[DELETE /automations/voice/receipts/{receipt_rule_id}][%d] voiceDeliveryReceiptAutomationDeleteForbidden  %+v", 403, o.Payload)
}

func (o *VoiceDeliveryReceiptAutomationDeleteForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVoiceDeliveryReceiptAutomationDeleteNotFound creates a VoiceDeliveryReceiptAutomationDeleteNotFound with default headers values
func NewVoiceDeliveryReceiptAutomationDeleteNotFound() *VoiceDeliveryReceiptAutomationDeleteNotFound {
	return &VoiceDeliveryReceiptAutomationDeleteNotFound{}
}

/*VoiceDeliveryReceiptAutomationDeleteNotFound handles this case with default header values.

NOT_FOUND
*/
type VoiceDeliveryReceiptAutomationDeleteNotFound struct {
	Payload interface{}
}

func (o *VoiceDeliveryReceiptAutomationDeleteNotFound) Error() string {
	return fmt.Sprintf("[DELETE /automations/voice/receipts/{receipt_rule_id}][%d] voiceDeliveryReceiptAutomationDeleteNotFound  %+v", 404, o.Payload)
}

func (o *VoiceDeliveryReceiptAutomationDeleteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVoiceDeliveryReceiptAutomationDeleteMethodNotAllowed creates a VoiceDeliveryReceiptAutomationDeleteMethodNotAllowed with default headers values
func NewVoiceDeliveryReceiptAutomationDeleteMethodNotAllowed() *VoiceDeliveryReceiptAutomationDeleteMethodNotAllowed {
	return &VoiceDeliveryReceiptAutomationDeleteMethodNotAllowed{}
}

/*VoiceDeliveryReceiptAutomationDeleteMethodNotAllowed handles this case with default header values.

METHOD_NOT_FOUND
*/
type VoiceDeliveryReceiptAutomationDeleteMethodNotAllowed struct {
	Payload interface{}
}

func (o *VoiceDeliveryReceiptAutomationDeleteMethodNotAllowed) Error() string {
	return fmt.Sprintf("[DELETE /automations/voice/receipts/{receipt_rule_id}][%d] voiceDeliveryReceiptAutomationDeleteMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *VoiceDeliveryReceiptAutomationDeleteMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVoiceDeliveryReceiptAutomationDeleteTooManyRequests creates a VoiceDeliveryReceiptAutomationDeleteTooManyRequests with default headers values
func NewVoiceDeliveryReceiptAutomationDeleteTooManyRequests() *VoiceDeliveryReceiptAutomationDeleteTooManyRequests {
	return &VoiceDeliveryReceiptAutomationDeleteTooManyRequests{}
}

/*VoiceDeliveryReceiptAutomationDeleteTooManyRequests handles this case with default header values.

TOO_MANY_REQUESTS
*/
type VoiceDeliveryReceiptAutomationDeleteTooManyRequests struct {
	Payload interface{}
}

func (o *VoiceDeliveryReceiptAutomationDeleteTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /automations/voice/receipts/{receipt_rule_id}][%d] voiceDeliveryReceiptAutomationDeleteTooManyRequests  %+v", 429, o.Payload)
}

func (o *VoiceDeliveryReceiptAutomationDeleteTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVoiceDeliveryReceiptAutomationDeleteDefault creates a VoiceDeliveryReceiptAutomationDeleteDefault with default headers values
func NewVoiceDeliveryReceiptAutomationDeleteDefault(code int) *VoiceDeliveryReceiptAutomationDeleteDefault {
	return &VoiceDeliveryReceiptAutomationDeleteDefault{
		_statusCode: code,
	}
}

/*VoiceDeliveryReceiptAutomationDeleteDefault handles this case with default header values.

INTERNAL_SERVER_ERROR
*/
type VoiceDeliveryReceiptAutomationDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the voice delivery receipt automation delete default response
func (o *VoiceDeliveryReceiptAutomationDeleteDefault) Code() int {
	return o._statusCode
}

func (o *VoiceDeliveryReceiptAutomationDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /automations/voice/receipts/{receipt_rule_id}][%d] VoiceDeliveryReceiptAutomationDelete default  %+v", o._statusCode, o.Payload)
}

func (o *VoiceDeliveryReceiptAutomationDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
