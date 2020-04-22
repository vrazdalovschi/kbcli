// Code generated by go-swagger; DO NOT EDIT.

package payment_transaction

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

// GetTransactionCustomFieldsReader is a Reader for the GetTransactionCustomFields structure.
type GetTransactionCustomFieldsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTransactionCustomFieldsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetTransactionCustomFieldsOK()
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

// NewGetTransactionCustomFieldsOK creates a GetTransactionCustomFieldsOK with default headers values
func NewGetTransactionCustomFieldsOK() *GetTransactionCustomFieldsOK {
	return &GetTransactionCustomFieldsOK{}
}

/*GetTransactionCustomFieldsOK handles this case with default header values.

successful operation
*/
type GetTransactionCustomFieldsOK struct {
	Payload []*kbmodel.CustomField

	HttpResponse runtime.ClientResponse
}

func (o *GetTransactionCustomFieldsOK) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/paymentTransactions/{transactionId}/customFields][%d] getTransactionCustomFieldsOK  %+v", 200, o.Payload)
}

func (o *GetTransactionCustomFieldsOK) GetPayload() []*kbmodel.CustomField {
	return o.Payload
}

func (o *GetTransactionCustomFieldsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTransactionCustomFieldsBadRequest creates a GetTransactionCustomFieldsBadRequest with default headers values
func NewGetTransactionCustomFieldsBadRequest() *GetTransactionCustomFieldsBadRequest {
	return &GetTransactionCustomFieldsBadRequest{}
}

/*GetTransactionCustomFieldsBadRequest handles this case with default header values.

Invalid transaction id supplied
*/
type GetTransactionCustomFieldsBadRequest struct {
	HttpResponse runtime.ClientResponse
}

func (o *GetTransactionCustomFieldsBadRequest) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/paymentTransactions/{transactionId}/customFields][%d] getTransactionCustomFieldsBadRequest ", 400)
}

func (o *GetTransactionCustomFieldsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
