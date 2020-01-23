// Code generated by go-swagger; DO NOT EDIT.

package subscription

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

// GetSubscriptionByKeyReader is a Reader for the GetSubscriptionByKey structure.
type GetSubscriptionByKeyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSubscriptionByKeyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetSubscriptionByKeyOK()
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

// NewGetSubscriptionByKeyOK creates a GetSubscriptionByKeyOK with default headers values
func NewGetSubscriptionByKeyOK() *GetSubscriptionByKeyOK {
	return &GetSubscriptionByKeyOK{}
}

/*GetSubscriptionByKeyOK handles this case with default header values.

successful operation
*/
type GetSubscriptionByKeyOK struct {
	Payload *kbmodel.Subscription

	HttpResponse runtime.ClientResponse
}

func (o *GetSubscriptionByKeyOK) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/subscriptions][%d] getSubscriptionByKeyOK  %+v", 200, o.Payload)
}

func (o *GetSubscriptionByKeyOK) GetPayload() *kbmodel.Subscription {
	return o.Payload
}

func (o *GetSubscriptionByKeyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(kbmodel.Subscription)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSubscriptionByKeyNotFound creates a GetSubscriptionByKeyNotFound with default headers values
func NewGetSubscriptionByKeyNotFound() *GetSubscriptionByKeyNotFound {
	return &GetSubscriptionByKeyNotFound{}
}

/*GetSubscriptionByKeyNotFound handles this case with default header values.

Subscription not found
*/
type GetSubscriptionByKeyNotFound struct {
	HttpResponse runtime.ClientResponse
}

func (o *GetSubscriptionByKeyNotFound) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/subscriptions][%d] getSubscriptionByKeyNotFound ", 404)
}

func (o *GetSubscriptionByKeyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
