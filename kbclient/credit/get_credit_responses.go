// Code generated by go-swagger; DO NOT EDIT.

package credit

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/killbill/kbcli/kbcommon"

	strfmt "github.com/go-openapi/strfmt"

	kbmodel "github.com/killbill/kbcli/kbmodel"
)

// GetCreditReader is a Reader for the GetCredit structure.
type GetCreditReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCreditReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCreditOK()
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

// NewGetCreditOK creates a GetCreditOK with default headers values
func NewGetCreditOK() *GetCreditOK {
	return &GetCreditOK{}
}

/*GetCreditOK handles this case with default header values.

successful operation
*/
type GetCreditOK struct {
	Payload *kbmodel.InvoiceItem

	HttpResponse runtime.ClientResponse
}

func (o *GetCreditOK) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/credits/{creditId}][%d] getCreditOK  %+v", 200, o.Payload)
}

func (o *GetCreditOK) GetPayload() *kbmodel.InvoiceItem {
	return o.Payload
}

func (o *GetCreditOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(kbmodel.InvoiceItem)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCreditBadRequest creates a GetCreditBadRequest with default headers values
func NewGetCreditBadRequest() *GetCreditBadRequest {
	return &GetCreditBadRequest{}
}

/*GetCreditBadRequest handles this case with default header values.

Invalid credit id supplied
*/
type GetCreditBadRequest struct {
	HttpResponse runtime.ClientResponse
}

func (o *GetCreditBadRequest) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/credits/{creditId}][%d] getCreditBadRequest ", 400)
}

func (o *GetCreditBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetCreditNotFound creates a GetCreditNotFound with default headers values
func NewGetCreditNotFound() *GetCreditNotFound {
	return &GetCreditNotFound{}
}

/*GetCreditNotFound handles this case with default header values.

Credit not found
*/
type GetCreditNotFound struct {
	HttpResponse runtime.ClientResponse
}

func (o *GetCreditNotFound) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/credits/{creditId}][%d] getCreditNotFound ", 404)
}

func (o *GetCreditNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
