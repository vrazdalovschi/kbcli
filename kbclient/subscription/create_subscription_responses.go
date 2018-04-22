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

// CreateSubscriptionReader is a Reader for the CreateSubscription structure.
type CreateSubscriptionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateSubscriptionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewCreateSubscriptionCreated()
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

// NewCreateSubscriptionCreated creates a CreateSubscriptionCreated with default headers values
func NewCreateSubscriptionCreated() *CreateSubscriptionCreated {
	return &CreateSubscriptionCreated{}
}

/*CreateSubscriptionCreated handles this case with default header values.

Subscription created successfully
*/
type CreateSubscriptionCreated struct {
	Payload *kbmodel.Subscription
}

func (o *CreateSubscriptionCreated) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/subscriptions][%d] createSubscriptionCreated  %+v", 201, o.Payload)
}

func (o *CreateSubscriptionCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(kbmodel.Subscription)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
