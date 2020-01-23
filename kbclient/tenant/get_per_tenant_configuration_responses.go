// Code generated by go-swagger; DO NOT EDIT.

package tenant

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

// GetPerTenantConfigurationReader is a Reader for the GetPerTenantConfiguration structure.
type GetPerTenantConfigurationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPerTenantConfigurationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPerTenantConfigurationOK()
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

// NewGetPerTenantConfigurationOK creates a GetPerTenantConfigurationOK with default headers values
func NewGetPerTenantConfigurationOK() *GetPerTenantConfigurationOK {
	return &GetPerTenantConfigurationOK{}
}

/*GetPerTenantConfigurationOK handles this case with default header values.

successful operation
*/
type GetPerTenantConfigurationOK struct {
	Payload *kbmodel.TenantKeyValue

	HttpResponse runtime.ClientResponse
}

func (o *GetPerTenantConfigurationOK) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/tenants/uploadPerTenantConfig][%d] getPerTenantConfigurationOK  %+v", 200, o.Payload)
}

func (o *GetPerTenantConfigurationOK) GetPayload() *kbmodel.TenantKeyValue {
	return o.Payload
}

func (o *GetPerTenantConfigurationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(kbmodel.TenantKeyValue)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPerTenantConfigurationBadRequest creates a GetPerTenantConfigurationBadRequest with default headers values
func NewGetPerTenantConfigurationBadRequest() *GetPerTenantConfigurationBadRequest {
	return &GetPerTenantConfigurationBadRequest{}
}

/*GetPerTenantConfigurationBadRequest handles this case with default header values.

Invalid tenantId supplied
*/
type GetPerTenantConfigurationBadRequest struct {
	HttpResponse runtime.ClientResponse
}

func (o *GetPerTenantConfigurationBadRequest) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/tenants/uploadPerTenantConfig][%d] getPerTenantConfigurationBadRequest ", 400)
}

func (o *GetPerTenantConfigurationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
