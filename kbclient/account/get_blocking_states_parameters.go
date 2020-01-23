// Code generated by go-swagger; DO NOT EDIT.

package account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetBlockingStatesParams creates a new GetBlockingStatesParams object
// with the default values initialized.
func NewGetBlockingStatesParams() *GetBlockingStatesParams {
	var (
		auditDefault = string("NONE")
	)
	return &GetBlockingStatesParams{
		Audit: &auditDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetBlockingStatesParamsWithTimeout creates a new GetBlockingStatesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetBlockingStatesParamsWithTimeout(timeout time.Duration) *GetBlockingStatesParams {
	var (
		auditDefault = string("NONE")
	)
	return &GetBlockingStatesParams{
		Audit: &auditDefault,

		timeout: timeout,
	}
}

// NewGetBlockingStatesParamsWithContext creates a new GetBlockingStatesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetBlockingStatesParamsWithContext(ctx context.Context) *GetBlockingStatesParams {
	var (
		auditDefault = string("NONE")
	)
	return &GetBlockingStatesParams{
		Audit: &auditDefault,

		Context: ctx,
	}
}

// NewGetBlockingStatesParamsWithHTTPClient creates a new GetBlockingStatesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetBlockingStatesParamsWithHTTPClient(client *http.Client) *GetBlockingStatesParams {
	var (
		auditDefault = string("NONE")
	)
	return &GetBlockingStatesParams{
		Audit:      &auditDefault,
		HTTPClient: client,
	}
}

/*GetBlockingStatesParams contains all the parameters to send to the API endpoint
for the get blocking states operation typically these are written to a http.Request
*/
type GetBlockingStatesParams struct {

	/*AccountID*/
	AccountID strfmt.UUID
	/*Audit*/
	Audit *string
	/*BlockingStateSvcs*/
	BlockingStateSvcs []string
	/*BlockingStateTypes*/
	BlockingStateTypes []string

	WithStackTrace        *bool // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithTimeout adds the timeout to the get blocking states params
func (o *GetBlockingStatesParams) WithTimeout(timeout time.Duration) *GetBlockingStatesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get blocking states params
func (o *GetBlockingStatesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get blocking states params
func (o *GetBlockingStatesParams) WithContext(ctx context.Context) *GetBlockingStatesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get blocking states params
func (o *GetBlockingStatesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get blocking states params
func (o *GetBlockingStatesParams) WithHTTPClient(client *http.Client) *GetBlockingStatesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get blocking states params
func (o *GetBlockingStatesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccountID adds the accountID to the get blocking states params
func (o *GetBlockingStatesParams) WithAccountID(accountID strfmt.UUID) *GetBlockingStatesParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the get blocking states params
func (o *GetBlockingStatesParams) SetAccountID(accountID strfmt.UUID) {
	o.AccountID = accountID
}

// WithAudit adds the audit to the get blocking states params
func (o *GetBlockingStatesParams) WithAudit(audit *string) *GetBlockingStatesParams {
	o.SetAudit(audit)
	return o
}

// SetAudit adds the audit to the get blocking states params
func (o *GetBlockingStatesParams) SetAudit(audit *string) {
	o.Audit = audit
}

// WithBlockingStateSvcs adds the blockingStateSvcs to the get blocking states params
func (o *GetBlockingStatesParams) WithBlockingStateSvcs(blockingStateSvcs []string) *GetBlockingStatesParams {
	o.SetBlockingStateSvcs(blockingStateSvcs)
	return o
}

// SetBlockingStateSvcs adds the blockingStateSvcs to the get blocking states params
func (o *GetBlockingStatesParams) SetBlockingStateSvcs(blockingStateSvcs []string) {
	o.BlockingStateSvcs = blockingStateSvcs
}

// WithBlockingStateTypes adds the blockingStateTypes to the get blocking states params
func (o *GetBlockingStatesParams) WithBlockingStateTypes(blockingStateTypes []string) *GetBlockingStatesParams {
	o.SetBlockingStateTypes(blockingStateTypes)
	return o
}

// SetBlockingStateTypes adds the blockingStateTypes to the get blocking states params
func (o *GetBlockingStatesParams) SetBlockingStateTypes(blockingStateTypes []string) {
	o.BlockingStateTypes = blockingStateTypes
}

// WriteToRequest writes these params to a swagger request
func (o *GetBlockingStatesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param accountId
	if err := r.SetPathParam("accountId", o.AccountID.String()); err != nil {
		return err
	}

	if o.Audit != nil {

		// query param audit
		var qrAudit string
		if o.Audit != nil {
			qrAudit = *o.Audit
		}
		qAudit := qrAudit
		if qAudit != "" {
			if err := r.SetQueryParam("audit", qAudit); err != nil {
				return err
			}
		}

	}

	valuesBlockingStateSvcs := o.BlockingStateSvcs

	joinedBlockingStateSvcs := swag.JoinByFormat(valuesBlockingStateSvcs, "multi")
	// query array param blockingStateSvcs
	if err := r.SetQueryParam("blockingStateSvcs", joinedBlockingStateSvcs...); err != nil {
		return err
	}

	valuesBlockingStateTypes := o.BlockingStateTypes

	joinedBlockingStateTypes := swag.JoinByFormat(valuesBlockingStateTypes, "multi")
	// query array param blockingStateTypes
	if err := r.SetQueryParam("blockingStateTypes", joinedBlockingStateTypes...); err != nil {
		return err
	}

	// header param withStackTrace
	if o.WithStackTrace != nil && *o.WithStackTrace {
		if err := r.SetQueryParam("withStackTrace", "true"); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
