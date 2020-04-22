// Code generated by go-swagger; DO NOT EDIT.

package payment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/killbill/kbcli/v2/kbcommon"

	strfmt "github.com/go-openapi/strfmt"

	kbmodel "github.com/killbill/kbcli/v2/kbmodel"
)

// CaptureAuthorizationReader is a Reader for the CaptureAuthorization structure.
type CaptureAuthorizationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CaptureAuthorizationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201, 200:
		result := NewCaptureAuthorizationCreated()
		result.HttpResponse = response
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		errorResult := kbcommon.NewKillbillError(response.Code())
		if err := consumer.Consume(response.Body(), &errorResult); err != nil && err != io.EOF {
			return nil, err
		}
		return nil, errorResult
	}
}

// NewCaptureAuthorizationCreated creates a CaptureAuthorizationCreated with default headers values
func NewCaptureAuthorizationCreated() *CaptureAuthorizationCreated {
	return &CaptureAuthorizationCreated{}
}

/*CaptureAuthorizationCreated handles this case with default header values.

Payment transaction created successfully
*/
type CaptureAuthorizationCreated struct {
	Payload *kbmodel.Payment

	HttpResponse runtime.ClientResponse
}

func (o *CaptureAuthorizationCreated) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/payments/{paymentId}][%d] captureAuthorizationCreated  %+v", 201, o.Payload)
}

func (o *CaptureAuthorizationCreated) GetPayload() *kbmodel.Payment {
	return o.Payload
}

func (o *CaptureAuthorizationCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(kbmodel.Payment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCaptureAuthorizationBadRequest creates a CaptureAuthorizationBadRequest with default headers values
func NewCaptureAuthorizationBadRequest() *CaptureAuthorizationBadRequest {
	return &CaptureAuthorizationBadRequest{}
}

/*CaptureAuthorizationBadRequest handles this case with default header values.

Invalid paymentId supplied
*/
type CaptureAuthorizationBadRequest struct {
	HttpResponse runtime.ClientResponse
}

func (o *CaptureAuthorizationBadRequest) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/payments/{paymentId}][%d] captureAuthorizationBadRequest ", 400)
}

func (o *CaptureAuthorizationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCaptureAuthorizationPaymentRequired creates a CaptureAuthorizationPaymentRequired with default headers values
func NewCaptureAuthorizationPaymentRequired() *CaptureAuthorizationPaymentRequired {
	return &CaptureAuthorizationPaymentRequired{}
}

/*CaptureAuthorizationPaymentRequired handles this case with default header values.

Transaction declined by gateway
*/
type CaptureAuthorizationPaymentRequired struct {
	HttpResponse runtime.ClientResponse
}

func (o *CaptureAuthorizationPaymentRequired) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/payments/{paymentId}][%d] captureAuthorizationPaymentRequired ", 402)
}

func (o *CaptureAuthorizationPaymentRequired) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCaptureAuthorizationNotFound creates a CaptureAuthorizationNotFound with default headers values
func NewCaptureAuthorizationNotFound() *CaptureAuthorizationNotFound {
	return &CaptureAuthorizationNotFound{}
}

/*CaptureAuthorizationNotFound handles this case with default header values.

Account or payment not found
*/
type CaptureAuthorizationNotFound struct {
	HttpResponse runtime.ClientResponse
}

func (o *CaptureAuthorizationNotFound) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/payments/{paymentId}][%d] captureAuthorizationNotFound ", 404)
}

func (o *CaptureAuthorizationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCaptureAuthorizationUnprocessableEntity creates a CaptureAuthorizationUnprocessableEntity with default headers values
func NewCaptureAuthorizationUnprocessableEntity() *CaptureAuthorizationUnprocessableEntity {
	return &CaptureAuthorizationUnprocessableEntity{}
}

/*CaptureAuthorizationUnprocessableEntity handles this case with default header values.

Payment is aborted by a control plugin
*/
type CaptureAuthorizationUnprocessableEntity struct {
	HttpResponse runtime.ClientResponse
}

func (o *CaptureAuthorizationUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/payments/{paymentId}][%d] captureAuthorizationUnprocessableEntity ", 422)
}

func (o *CaptureAuthorizationUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCaptureAuthorizationBadGateway creates a CaptureAuthorizationBadGateway with default headers values
func NewCaptureAuthorizationBadGateway() *CaptureAuthorizationBadGateway {
	return &CaptureAuthorizationBadGateway{}
}

/*CaptureAuthorizationBadGateway handles this case with default header values.

Failed to submit payment transaction
*/
type CaptureAuthorizationBadGateway struct {
	HttpResponse runtime.ClientResponse
}

func (o *CaptureAuthorizationBadGateway) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/payments/{paymentId}][%d] captureAuthorizationBadGateway ", 502)
}

func (o *CaptureAuthorizationBadGateway) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCaptureAuthorizationServiceUnavailable creates a CaptureAuthorizationServiceUnavailable with default headers values
func NewCaptureAuthorizationServiceUnavailable() *CaptureAuthorizationServiceUnavailable {
	return &CaptureAuthorizationServiceUnavailable{}
}

/*CaptureAuthorizationServiceUnavailable handles this case with default header values.

Payment in unknown status, failed to receive gateway response
*/
type CaptureAuthorizationServiceUnavailable struct {
	HttpResponse runtime.ClientResponse
}

func (o *CaptureAuthorizationServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/payments/{paymentId}][%d] captureAuthorizationServiceUnavailable ", 503)
}

func (o *CaptureAuthorizationServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCaptureAuthorizationGatewayTimeout creates a CaptureAuthorizationGatewayTimeout with default headers values
func NewCaptureAuthorizationGatewayTimeout() *CaptureAuthorizationGatewayTimeout {
	return &CaptureAuthorizationGatewayTimeout{}
}

/*CaptureAuthorizationGatewayTimeout handles this case with default header values.

Payment operation timeout
*/
type CaptureAuthorizationGatewayTimeout struct {
	HttpResponse runtime.ClientResponse
}

func (o *CaptureAuthorizationGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/payments/{paymentId}][%d] captureAuthorizationGatewayTimeout ", 504)
}

func (o *CaptureAuthorizationGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
