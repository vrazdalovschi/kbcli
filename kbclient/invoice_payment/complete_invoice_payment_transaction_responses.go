// Code generated by go-swagger; DO NOT EDIT.

package invoice_payment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/killbill/kbcli/kbcommon"

	strfmt "github.com/go-openapi/strfmt"
)

// CompleteInvoicePaymentTransactionReader is a Reader for the CompleteInvoicePaymentTransaction structure.
type CompleteInvoicePaymentTransactionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CompleteInvoicePaymentTransactionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewCompleteInvoicePaymentTransactionNoContent()
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

// NewCompleteInvoicePaymentTransactionNoContent creates a CompleteInvoicePaymentTransactionNoContent with default headers values
func NewCompleteInvoicePaymentTransactionNoContent() *CompleteInvoicePaymentTransactionNoContent {
	return &CompleteInvoicePaymentTransactionNoContent{}
}

/*CompleteInvoicePaymentTransactionNoContent handles this case with default header values.

Successful operation
*/
type CompleteInvoicePaymentTransactionNoContent struct {
	HttpResponse runtime.ClientResponse
}

func (o *CompleteInvoicePaymentTransactionNoContent) Error() string {
	return fmt.Sprintf("[PUT /1.0/kb/invoicePayments/{paymentId}][%d] completeInvoicePaymentTransactionNoContent ", 204)
}

func (o *CompleteInvoicePaymentTransactionNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCompleteInvoicePaymentTransactionBadRequest creates a CompleteInvoicePaymentTransactionBadRequest with default headers values
func NewCompleteInvoicePaymentTransactionBadRequest() *CompleteInvoicePaymentTransactionBadRequest {
	return &CompleteInvoicePaymentTransactionBadRequest{}
}

/*CompleteInvoicePaymentTransactionBadRequest handles this case with default header values.

Invalid paymentId supplied
*/
type CompleteInvoicePaymentTransactionBadRequest struct {
	HttpResponse runtime.ClientResponse
}

func (o *CompleteInvoicePaymentTransactionBadRequest) Error() string {
	return fmt.Sprintf("[PUT /1.0/kb/invoicePayments/{paymentId}][%d] completeInvoicePaymentTransactionBadRequest ", 400)
}

func (o *CompleteInvoicePaymentTransactionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCompleteInvoicePaymentTransactionPaymentRequired creates a CompleteInvoicePaymentTransactionPaymentRequired with default headers values
func NewCompleteInvoicePaymentTransactionPaymentRequired() *CompleteInvoicePaymentTransactionPaymentRequired {
	return &CompleteInvoicePaymentTransactionPaymentRequired{}
}

/*CompleteInvoicePaymentTransactionPaymentRequired handles this case with default header values.

Transaction declined by gateway
*/
type CompleteInvoicePaymentTransactionPaymentRequired struct {
	HttpResponse runtime.ClientResponse
}

func (o *CompleteInvoicePaymentTransactionPaymentRequired) Error() string {
	return fmt.Sprintf("[PUT /1.0/kb/invoicePayments/{paymentId}][%d] completeInvoicePaymentTransactionPaymentRequired ", 402)
}

func (o *CompleteInvoicePaymentTransactionPaymentRequired) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCompleteInvoicePaymentTransactionNotFound creates a CompleteInvoicePaymentTransactionNotFound with default headers values
func NewCompleteInvoicePaymentTransactionNotFound() *CompleteInvoicePaymentTransactionNotFound {
	return &CompleteInvoicePaymentTransactionNotFound{}
}

/*CompleteInvoicePaymentTransactionNotFound handles this case with default header values.

Account or payment not found
*/
type CompleteInvoicePaymentTransactionNotFound struct {
	HttpResponse runtime.ClientResponse
}

func (o *CompleteInvoicePaymentTransactionNotFound) Error() string {
	return fmt.Sprintf("[PUT /1.0/kb/invoicePayments/{paymentId}][%d] completeInvoicePaymentTransactionNotFound ", 404)
}

func (o *CompleteInvoicePaymentTransactionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCompleteInvoicePaymentTransactionUnprocessableEntity creates a CompleteInvoicePaymentTransactionUnprocessableEntity with default headers values
func NewCompleteInvoicePaymentTransactionUnprocessableEntity() *CompleteInvoicePaymentTransactionUnprocessableEntity {
	return &CompleteInvoicePaymentTransactionUnprocessableEntity{}
}

/*CompleteInvoicePaymentTransactionUnprocessableEntity handles this case with default header values.

Payment is aborted by a control plugin
*/
type CompleteInvoicePaymentTransactionUnprocessableEntity struct {
	HttpResponse runtime.ClientResponse
}

func (o *CompleteInvoicePaymentTransactionUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /1.0/kb/invoicePayments/{paymentId}][%d] completeInvoicePaymentTransactionUnprocessableEntity ", 422)
}

func (o *CompleteInvoicePaymentTransactionUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCompleteInvoicePaymentTransactionBadGateway creates a CompleteInvoicePaymentTransactionBadGateway with default headers values
func NewCompleteInvoicePaymentTransactionBadGateway() *CompleteInvoicePaymentTransactionBadGateway {
	return &CompleteInvoicePaymentTransactionBadGateway{}
}

/*CompleteInvoicePaymentTransactionBadGateway handles this case with default header values.

Failed to submit payment transaction
*/
type CompleteInvoicePaymentTransactionBadGateway struct {
	HttpResponse runtime.ClientResponse
}

func (o *CompleteInvoicePaymentTransactionBadGateway) Error() string {
	return fmt.Sprintf("[PUT /1.0/kb/invoicePayments/{paymentId}][%d] completeInvoicePaymentTransactionBadGateway ", 502)
}

func (o *CompleteInvoicePaymentTransactionBadGateway) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCompleteInvoicePaymentTransactionServiceUnavailable creates a CompleteInvoicePaymentTransactionServiceUnavailable with default headers values
func NewCompleteInvoicePaymentTransactionServiceUnavailable() *CompleteInvoicePaymentTransactionServiceUnavailable {
	return &CompleteInvoicePaymentTransactionServiceUnavailable{}
}

/*CompleteInvoicePaymentTransactionServiceUnavailable handles this case with default header values.

Payment in unknown status, failed to receive gateway response
*/
type CompleteInvoicePaymentTransactionServiceUnavailable struct {
	HttpResponse runtime.ClientResponse
}

func (o *CompleteInvoicePaymentTransactionServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /1.0/kb/invoicePayments/{paymentId}][%d] completeInvoicePaymentTransactionServiceUnavailable ", 503)
}

func (o *CompleteInvoicePaymentTransactionServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCompleteInvoicePaymentTransactionGatewayTimeout creates a CompleteInvoicePaymentTransactionGatewayTimeout with default headers values
func NewCompleteInvoicePaymentTransactionGatewayTimeout() *CompleteInvoicePaymentTransactionGatewayTimeout {
	return &CompleteInvoicePaymentTransactionGatewayTimeout{}
}

/*CompleteInvoicePaymentTransactionGatewayTimeout handles this case with default header values.

Payment operation timeout
*/
type CompleteInvoicePaymentTransactionGatewayTimeout struct {
	HttpResponse runtime.ClientResponse
}

func (o *CompleteInvoicePaymentTransactionGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /1.0/kb/invoicePayments/{paymentId}][%d] completeInvoicePaymentTransactionGatewayTimeout ", 504)
}

func (o *CompleteInvoicePaymentTransactionGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}