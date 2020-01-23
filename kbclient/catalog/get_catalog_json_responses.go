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

	kbmodel "github.com/killbill/kbcli/kbmodel"
)

// GetCatalogJSONReader is a Reader for the GetCatalogJSON structure.
type GetCatalogJSONReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCatalogJSONReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCatalogJSONOK()
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

// NewGetCatalogJSONOK creates a GetCatalogJSONOK with default headers values
func NewGetCatalogJSONOK() *GetCatalogJSONOK {
	return &GetCatalogJSONOK{}
}

/*GetCatalogJSONOK handles this case with default header values.

successful operation
*/
type GetCatalogJSONOK struct {
	Payload []*kbmodel.Catalog

	HttpResponse runtime.ClientResponse
}

func (o *GetCatalogJSONOK) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/catalog][%d] getCatalogJsonOK  %+v", 200, o.Payload)
}

func (o *GetCatalogJSONOK) GetPayload() []*kbmodel.Catalog {
	return o.Payload
}

func (o *GetCatalogJSONOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
