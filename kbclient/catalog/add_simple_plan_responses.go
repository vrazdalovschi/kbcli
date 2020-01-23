// Code generated by go-swagger; DO NOT EDIT.

package catalog

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/killbill/kbcli/kbcommon"

	strfmt "github.com/go-openapi/strfmt"
)

// AddSimplePlanReader is a Reader for the AddSimplePlan structure.
type AddSimplePlanReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddSimplePlanReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201, 200:
		result := NewAddSimplePlanCreated()
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

// NewAddSimplePlanCreated creates a AddSimplePlanCreated with default headers values
func NewAddSimplePlanCreated() *AddSimplePlanCreated {
	return &AddSimplePlanCreated{}
}

/*AddSimplePlanCreated handles this case with default header values.

Created new plan successfully
*/
type AddSimplePlanCreated struct {
	Payload string

	HttpResponse runtime.ClientResponse
}

func (o *AddSimplePlanCreated) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/catalog/simplePlan][%d] addSimplePlanCreated  %+v", 201, o.Payload)
}

func (o *AddSimplePlanCreated) GetPayload() string {
	return o.Payload
}

func (o *AddSimplePlanCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
