// Code generated by go-swagger; DO NOT EDIT.

package bundle

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

// NewGetBundleByKeyParams creates a new GetBundleByKeyParams object
// with the default values initialized.
func NewGetBundleByKeyParams() *GetBundleByKeyParams {
	var (
		auditDefault           = string("NONE")
		includedDeletedDefault = bool(false)
	)
	return &GetBundleByKeyParams{
		Audit:           &auditDefault,
		IncludedDeleted: &includedDeletedDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetBundleByKeyParamsWithTimeout creates a new GetBundleByKeyParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetBundleByKeyParamsWithTimeout(timeout time.Duration) *GetBundleByKeyParams {
	var (
		auditDefault           = string("NONE")
		includedDeletedDefault = bool(false)
	)
	return &GetBundleByKeyParams{
		Audit:           &auditDefault,
		IncludedDeleted: &includedDeletedDefault,

		timeout: timeout,
	}
}

// NewGetBundleByKeyParamsWithContext creates a new GetBundleByKeyParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetBundleByKeyParamsWithContext(ctx context.Context) *GetBundleByKeyParams {
	var (
		auditDefault           = string("NONE")
		includedDeletedDefault = bool(false)
	)
	return &GetBundleByKeyParams{
		Audit:           &auditDefault,
		IncludedDeleted: &includedDeletedDefault,

		Context: ctx,
	}
}

// NewGetBundleByKeyParamsWithHTTPClient creates a new GetBundleByKeyParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetBundleByKeyParamsWithHTTPClient(client *http.Client) *GetBundleByKeyParams {
	var (
		auditDefault           = string("NONE")
		includedDeletedDefault = bool(false)
	)
	return &GetBundleByKeyParams{
		Audit:           &auditDefault,
		IncludedDeleted: &includedDeletedDefault,
		HTTPClient:      client,
	}
}

/*GetBundleByKeyParams contains all the parameters to send to the API endpoint
for the get bundle by key operation typically these are written to a http.Request
*/
type GetBundleByKeyParams struct {

	/*Audit*/
	Audit *string
	/*ExternalKey*/
	ExternalKey string
	/*IncludedDeleted*/
	IncludedDeleted *bool

	WithStackTrace        *bool // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithTimeout adds the timeout to the get bundle by key params
func (o *GetBundleByKeyParams) WithTimeout(timeout time.Duration) *GetBundleByKeyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get bundle by key params
func (o *GetBundleByKeyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get bundle by key params
func (o *GetBundleByKeyParams) WithContext(ctx context.Context) *GetBundleByKeyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get bundle by key params
func (o *GetBundleByKeyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get bundle by key params
func (o *GetBundleByKeyParams) WithHTTPClient(client *http.Client) *GetBundleByKeyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get bundle by key params
func (o *GetBundleByKeyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAudit adds the audit to the get bundle by key params
func (o *GetBundleByKeyParams) WithAudit(audit *string) *GetBundleByKeyParams {
	o.SetAudit(audit)
	return o
}

// SetAudit adds the audit to the get bundle by key params
func (o *GetBundleByKeyParams) SetAudit(audit *string) {
	o.Audit = audit
}

// WithExternalKey adds the externalKey to the get bundle by key params
func (o *GetBundleByKeyParams) WithExternalKey(externalKey string) *GetBundleByKeyParams {
	o.SetExternalKey(externalKey)
	return o
}

// SetExternalKey adds the externalKey to the get bundle by key params
func (o *GetBundleByKeyParams) SetExternalKey(externalKey string) {
	o.ExternalKey = externalKey
}

// WithIncludedDeleted adds the includedDeleted to the get bundle by key params
func (o *GetBundleByKeyParams) WithIncludedDeleted(includedDeleted *bool) *GetBundleByKeyParams {
	o.SetIncludedDeleted(includedDeleted)
	return o
}

// SetIncludedDeleted adds the includedDeleted to the get bundle by key params
func (o *GetBundleByKeyParams) SetIncludedDeleted(includedDeleted *bool) {
	o.IncludedDeleted = includedDeleted
}

// WriteToRequest writes these params to a swagger request
func (o *GetBundleByKeyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	// query param externalKey
	qrExternalKey := o.ExternalKey
	qExternalKey := qrExternalKey
	if qExternalKey != "" {
		if err := r.SetQueryParam("externalKey", qExternalKey); err != nil {
			return err
		}
	}

	if o.IncludedDeleted != nil {

		// query param includedDeleted
		var qrIncludedDeleted bool
		if o.IncludedDeleted != nil {
			qrIncludedDeleted = *o.IncludedDeleted
		}
		qIncludedDeleted := swag.FormatBool(qrIncludedDeleted)
		if qIncludedDeleted != "" {
			if err := r.SetQueryParam("includedDeleted", qIncludedDeleted); err != nil {
				return err
			}
		}

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
