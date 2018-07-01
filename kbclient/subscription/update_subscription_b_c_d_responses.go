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
)

// UpdateSubscriptionBCDReader is a Reader for the UpdateSubscriptionBCD structure.
type UpdateSubscriptionBCDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateSubscriptionBCDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewUpdateSubscriptionBCDNoContent()
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

// NewUpdateSubscriptionBCDNoContent creates a UpdateSubscriptionBCDNoContent with default headers values
func NewUpdateSubscriptionBCDNoContent() *UpdateSubscriptionBCDNoContent {
	return &UpdateSubscriptionBCDNoContent{}
}

/*UpdateSubscriptionBCDNoContent handles this case with default header values.

Successful operation
*/
type UpdateSubscriptionBCDNoContent struct {
	HttpResponse runtime.ClientResponse
}

func (o *UpdateSubscriptionBCDNoContent) Error() string {
	return fmt.Sprintf("[PUT /1.0/kb/subscriptions/{subscriptionId}/bcd][%d] updateSubscriptionBCDNoContent ", 204)
}

func (o *UpdateSubscriptionBCDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateSubscriptionBCDBadRequest creates a UpdateSubscriptionBCDBadRequest with default headers values
func NewUpdateSubscriptionBCDBadRequest() *UpdateSubscriptionBCDBadRequest {
	return &UpdateSubscriptionBCDBadRequest{}
}

/*UpdateSubscriptionBCDBadRequest handles this case with default header values.

Invalid entitlement supplied
*/
type UpdateSubscriptionBCDBadRequest struct {
	HttpResponse runtime.ClientResponse
}

func (o *UpdateSubscriptionBCDBadRequest) Error() string {
	return fmt.Sprintf("[PUT /1.0/kb/subscriptions/{subscriptionId}/bcd][%d] updateSubscriptionBCDBadRequest ", 400)
}

func (o *UpdateSubscriptionBCDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
