// Code generated by go-swagger; DO NOT EDIT.

package tag_definition

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

// GetTagDefinitionsReader is a Reader for the GetTagDefinitions structure.
type GetTagDefinitionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTagDefinitionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetTagDefinitionsOK()
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

// NewGetTagDefinitionsOK creates a GetTagDefinitionsOK with default headers values
func NewGetTagDefinitionsOK() *GetTagDefinitionsOK {
	return &GetTagDefinitionsOK{}
}

/*GetTagDefinitionsOK handles this case with default header values.

successful operation
*/
type GetTagDefinitionsOK struct {
	Payload []*kbmodel.TagDefinition

	HttpResponse runtime.ClientResponse
}

func (o *GetTagDefinitionsOK) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/tagDefinitions][%d] getTagDefinitionsOK  %+v", 200, o.Payload)
}

func (o *GetTagDefinitionsOK) GetPayload() []*kbmodel.TagDefinition {
	return o.Payload
}

func (o *GetTagDefinitionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
