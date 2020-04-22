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

// GetPaymentByExternalKeyReader is a Reader for the GetPaymentByExternalKey structure.
type GetPaymentByExternalKeyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPaymentByExternalKeyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPaymentByExternalKeyOK()
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

// NewGetPaymentByExternalKeyOK creates a GetPaymentByExternalKeyOK with default headers values
func NewGetPaymentByExternalKeyOK() *GetPaymentByExternalKeyOK {
	return &GetPaymentByExternalKeyOK{}
}

/*GetPaymentByExternalKeyOK handles this case with default header values.

successful operation
*/
type GetPaymentByExternalKeyOK struct {
	Payload *kbmodel.Payment

	HttpResponse runtime.ClientResponse
}

func (o *GetPaymentByExternalKeyOK) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/payments][%d] getPaymentByExternalKeyOK  %+v", 200, o.Payload)
}

func (o *GetPaymentByExternalKeyOK) GetPayload() *kbmodel.Payment {
	return o.Payload
}

func (o *GetPaymentByExternalKeyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(kbmodel.Payment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPaymentByExternalKeyNotFound creates a GetPaymentByExternalKeyNotFound with default headers values
func NewGetPaymentByExternalKeyNotFound() *GetPaymentByExternalKeyNotFound {
	return &GetPaymentByExternalKeyNotFound{}
}

/*GetPaymentByExternalKeyNotFound handles this case with default header values.

Payment not found
*/
type GetPaymentByExternalKeyNotFound struct {
	HttpResponse runtime.ClientResponse
}

func (o *GetPaymentByExternalKeyNotFound) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/payments][%d] getPaymentByExternalKeyNotFound ", 404)
}

func (o *GetPaymentByExternalKeyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
