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

// ModifySubscriptionCustomFieldsReader is a Reader for the ModifySubscriptionCustomFields structure.
type ModifySubscriptionCustomFieldsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ModifySubscriptionCustomFieldsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewModifySubscriptionCustomFieldsNoContent()
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

// NewModifySubscriptionCustomFieldsNoContent creates a ModifySubscriptionCustomFieldsNoContent with default headers values
func NewModifySubscriptionCustomFieldsNoContent() *ModifySubscriptionCustomFieldsNoContent {
	return &ModifySubscriptionCustomFieldsNoContent{}
}

/*ModifySubscriptionCustomFieldsNoContent handles this case with default header values.

Successful operation
*/
type ModifySubscriptionCustomFieldsNoContent struct {
	HttpResponse runtime.ClientResponse
}

func (o *ModifySubscriptionCustomFieldsNoContent) Error() string {
	return fmt.Sprintf("[PUT /1.0/kb/subscriptions/{subscriptionId}/customFields][%d] modifySubscriptionCustomFieldsNoContent ", 204)
}

func (o *ModifySubscriptionCustomFieldsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewModifySubscriptionCustomFieldsBadRequest creates a ModifySubscriptionCustomFieldsBadRequest with default headers values
func NewModifySubscriptionCustomFieldsBadRequest() *ModifySubscriptionCustomFieldsBadRequest {
	return &ModifySubscriptionCustomFieldsBadRequest{}
}

/*ModifySubscriptionCustomFieldsBadRequest handles this case with default header values.

Invalid subscription id supplied
*/
type ModifySubscriptionCustomFieldsBadRequest struct {
	HttpResponse runtime.ClientResponse
}

func (o *ModifySubscriptionCustomFieldsBadRequest) Error() string {
	return fmt.Sprintf("[PUT /1.0/kb/subscriptions/{subscriptionId}/customFields][%d] modifySubscriptionCustomFieldsBadRequest ", 400)
}

func (o *ModifySubscriptionCustomFieldsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
