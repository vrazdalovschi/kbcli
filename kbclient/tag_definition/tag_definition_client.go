// Code generated by go-swagger; DO NOT EDIT.

package tag_definition

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/killbill/kbcli/kbcommon"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new tag definition API client.
func New(transport runtime.ClientTransport,
	formats strfmt.Registry,
	authInfo runtime.ClientAuthInfoWriter,
	defaults KillbillDefaults) *Client {

	return &Client{transport: transport, formats: formats, authInfo: authInfo, defaults: defaults}
}

// killbill default values. When a call is made to an operation, these values are used
// if params doesn't specify them.
type KillbillDefaults interface {
	// Default CreatedBy. If not set explicitly in params, this will be used.
	XKillbillCreatedBy() *string
	// Default Comment. If not set explicitly in params, this will be used.
	XKillbillComment() *string
	// Default Reason. If not set explicitly in params, this will be used.
	XKillbillReason() *string
	// Default WithStackTrace. If not set explicitly in params, this will be used.
	KillbillWithStackTrace() *bool
}

/*
Client for tag definition API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
	authInfo  runtime.ClientAuthInfoWriter
	defaults  KillbillDefaults
}

// ITagDefinition - interface for TagDefinition client.
type ITagDefinition interface {
	/*
		CreateTagDefinition creates a tag definition
	*/
	CreateTagDefinition(ctx context.Context, params *CreateTagDefinitionParams) (*CreateTagDefinitionCreated, error)

	/*
		DeleteTagDefinition deletes a tag definition
	*/
	DeleteTagDefinition(ctx context.Context, params *DeleteTagDefinitionParams) (*DeleteTagDefinitionNoContent, error)

	/*
		GetTagDefinition retrieves a tag definition
	*/
	GetTagDefinition(ctx context.Context, params *GetTagDefinitionParams) (*GetTagDefinitionOK, error)

	/*
		GetTagDefinitionAuditLogsWithHistory retrieves tag definition audit logs with history by id
	*/
	GetTagDefinitionAuditLogsWithHistory(ctx context.Context, params *GetTagDefinitionAuditLogsWithHistoryParams) (*GetTagDefinitionAuditLogsWithHistoryOK, error)

	/*
		GetTagDefinitions lists tag definitions
	*/
	GetTagDefinitions(ctx context.Context, params *GetTagDefinitionsParams) (*GetTagDefinitionsOK, error)
}

/*
CreateTagDefinition creates a tag definition
*/
func (a *Client) CreateTagDefinition(ctx context.Context, params *CreateTagDefinitionParams) (*CreateTagDefinitionCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateTagDefinitionParams()
	}
	getParams := NewCreateTagDefinitionParams()
	getParams.Context = ctx
	params.Context = ctx
	if params.XKillbillComment == nil && a.defaults.XKillbillComment() != nil {
		params.XKillbillComment = a.defaults.XKillbillComment()
	}
	getParams.XKillbillComment = params.XKillbillComment
	if params.XKillbillCreatedBy == "" && a.defaults.XKillbillCreatedBy() != nil {
		params.XKillbillCreatedBy = *a.defaults.XKillbillCreatedBy()
	}
	getParams.XKillbillCreatedBy = params.XKillbillCreatedBy
	if params.XKillbillReason == nil && a.defaults.XKillbillReason() != nil {
		params.XKillbillReason = a.defaults.XKillbillReason()
	}
	getParams.XKillbillReason = params.XKillbillReason
	if params.WithStackTrace == nil && a.defaults.KillbillWithStackTrace() != nil {
		params.WithStackTrace = a.defaults.KillbillWithStackTrace()
	}
	getParams.WithStackTrace = params.WithStackTrace

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createTagDefinition",
		Method:             "POST",
		PathPattern:        "/1.0/kb/tagDefinitions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateTagDefinitionReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	createdResult := result.(*CreateTagDefinitionCreated)
	location := kbcommon.ParseLocationHeader(createdResult.HttpResponse.GetHeader("Location"))
	if !params.ProcessLocationHeader || location == "" {
		return createdResult, nil
	}

	getResult, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createTagDefinition",
		Method:             "GET",
		PathPattern:        location,
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             getParams,
		Reader:             &CreateTagDefinitionReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            getParams.Context,
		Client:             getParams.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return getResult.(*CreateTagDefinitionCreated), nil

}

/*
DeleteTagDefinition deletes a tag definition
*/
func (a *Client) DeleteTagDefinition(ctx context.Context, params *DeleteTagDefinitionParams) (*DeleteTagDefinitionNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteTagDefinitionParams()
	}
	params.Context = ctx
	if params.XKillbillComment == nil && a.defaults.XKillbillComment() != nil {
		params.XKillbillComment = a.defaults.XKillbillComment()
	}

	if params.XKillbillCreatedBy == "" && a.defaults.XKillbillCreatedBy() != nil {
		params.XKillbillCreatedBy = *a.defaults.XKillbillCreatedBy()
	}

	if params.XKillbillReason == nil && a.defaults.XKillbillReason() != nil {
		params.XKillbillReason = a.defaults.XKillbillReason()
	}

	if params.WithStackTrace == nil && a.defaults.KillbillWithStackTrace() != nil {
		params.WithStackTrace = a.defaults.KillbillWithStackTrace()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteTagDefinition",
		Method:             "DELETE",
		PathPattern:        "/1.0/kb/tagDefinitions/{tagDefinitionId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteTagDefinitionReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteTagDefinitionNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteTagDefinition: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)

}

/*
GetTagDefinition retrieves a tag definition
*/
func (a *Client) GetTagDefinition(ctx context.Context, params *GetTagDefinitionParams) (*GetTagDefinitionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTagDefinitionParams()
	}
	params.Context = ctx
	if params.WithStackTrace == nil && a.defaults.KillbillWithStackTrace() != nil {
		params.WithStackTrace = a.defaults.KillbillWithStackTrace()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getTagDefinition",
		Method:             "GET",
		PathPattern:        "/1.0/kb/tagDefinitions/{tagDefinitionId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetTagDefinitionReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetTagDefinitionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getTagDefinition: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)

}

/*
GetTagDefinitionAuditLogsWithHistory retrieves tag definition audit logs with history by id
*/
func (a *Client) GetTagDefinitionAuditLogsWithHistory(ctx context.Context, params *GetTagDefinitionAuditLogsWithHistoryParams) (*GetTagDefinitionAuditLogsWithHistoryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTagDefinitionAuditLogsWithHistoryParams()
	}
	params.Context = ctx
	if params.WithStackTrace == nil && a.defaults.KillbillWithStackTrace() != nil {
		params.WithStackTrace = a.defaults.KillbillWithStackTrace()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getTagDefinitionAuditLogsWithHistory",
		Method:             "GET",
		PathPattern:        "/1.0/kb/tagDefinitions/{tagDefinitionId}/auditLogsWithHistory",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetTagDefinitionAuditLogsWithHistoryReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetTagDefinitionAuditLogsWithHistoryOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getTagDefinitionAuditLogsWithHistory: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)

}

/*
GetTagDefinitions lists tag definitions
*/
func (a *Client) GetTagDefinitions(ctx context.Context, params *GetTagDefinitionsParams) (*GetTagDefinitionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTagDefinitionsParams()
	}
	params.Context = ctx
	if params.WithStackTrace == nil && a.defaults.KillbillWithStackTrace() != nil {
		params.WithStackTrace = a.defaults.KillbillWithStackTrace()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getTagDefinitions",
		Method:             "GET",
		PathPattern:        "/1.0/kb/tagDefinitions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetTagDefinitionsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetTagDefinitionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getTagDefinitions: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
