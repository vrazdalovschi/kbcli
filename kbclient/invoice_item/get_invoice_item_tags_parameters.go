// Code generated by go-swagger; DO NOT EDIT.

package invoice_item

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetInvoiceItemTagsParams creates a new GetInvoiceItemTagsParams object
// with the default values initialized.
func NewGetInvoiceItemTagsParams() *GetInvoiceItemTagsParams {
	var (
		auditDefault           = string("NONE")
		includedDeletedDefault = bool(false)
	)
	return &GetInvoiceItemTagsParams{
		Audit:           &auditDefault,
		IncludedDeleted: &includedDeletedDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetInvoiceItemTagsParamsWithTimeout creates a new GetInvoiceItemTagsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetInvoiceItemTagsParamsWithTimeout(timeout time.Duration) *GetInvoiceItemTagsParams {
	var (
		auditDefault           = string("NONE")
		includedDeletedDefault = bool(false)
	)
	return &GetInvoiceItemTagsParams{
		Audit:           &auditDefault,
		IncludedDeleted: &includedDeletedDefault,

		timeout: timeout,
	}
}

// NewGetInvoiceItemTagsParamsWithContext creates a new GetInvoiceItemTagsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetInvoiceItemTagsParamsWithContext(ctx context.Context) *GetInvoiceItemTagsParams {
	var (
		auditDefault           = string("NONE")
		includedDeletedDefault = bool(false)
	)
	return &GetInvoiceItemTagsParams{
		Audit:           &auditDefault,
		IncludedDeleted: &includedDeletedDefault,

		Context: ctx,
	}
}

// NewGetInvoiceItemTagsParamsWithHTTPClient creates a new GetInvoiceItemTagsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetInvoiceItemTagsParamsWithHTTPClient(client *http.Client) *GetInvoiceItemTagsParams {
	var (
		auditDefault           = string("NONE")
		includedDeletedDefault = bool(false)
	)
	return &GetInvoiceItemTagsParams{
		Audit:           &auditDefault,
		IncludedDeleted: &includedDeletedDefault,
		HTTPClient:      client,
	}
}

/*GetInvoiceItemTagsParams contains all the parameters to send to the API endpoint
for the get invoice item tags operation typically these are written to a http.Request
*/
type GetInvoiceItemTagsParams struct {

	/*XKillbillAPIKey*/
	XKillbillAPIKey string
	/*XKillbillAPISecret*/
	XKillbillAPISecret string
	/*AccountID*/
	AccountID strfmt.UUID
	/*Audit*/
	Audit *string
	/*IncludedDeleted*/
	IncludedDeleted *bool
	/*InvoiceItemID*/
	InvoiceItemID strfmt.UUID

	WithStackTrace        *bool // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithTimeout adds the timeout to the get invoice item tags params
func (o *GetInvoiceItemTagsParams) WithTimeout(timeout time.Duration) *GetInvoiceItemTagsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get invoice item tags params
func (o *GetInvoiceItemTagsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get invoice item tags params
func (o *GetInvoiceItemTagsParams) WithContext(ctx context.Context) *GetInvoiceItemTagsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get invoice item tags params
func (o *GetInvoiceItemTagsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get invoice item tags params
func (o *GetInvoiceItemTagsParams) WithHTTPClient(client *http.Client) *GetInvoiceItemTagsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get invoice item tags params
func (o *GetInvoiceItemTagsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillAPIKey adds the xKillbillAPIKey to the get invoice item tags params
func (o *GetInvoiceItemTagsParams) WithXKillbillAPIKey(xKillbillAPIKey string) *GetInvoiceItemTagsParams {
	o.SetXKillbillAPIKey(xKillbillAPIKey)
	return o
}

// SetXKillbillAPIKey adds the xKillbillApiKey to the get invoice item tags params
func (o *GetInvoiceItemTagsParams) SetXKillbillAPIKey(xKillbillAPIKey string) {
	o.XKillbillAPIKey = xKillbillAPIKey
}

// WithXKillbillAPISecret adds the xKillbillAPISecret to the get invoice item tags params
func (o *GetInvoiceItemTagsParams) WithXKillbillAPISecret(xKillbillAPISecret string) *GetInvoiceItemTagsParams {
	o.SetXKillbillAPISecret(xKillbillAPISecret)
	return o
}

// SetXKillbillAPISecret adds the xKillbillApiSecret to the get invoice item tags params
func (o *GetInvoiceItemTagsParams) SetXKillbillAPISecret(xKillbillAPISecret string) {
	o.XKillbillAPISecret = xKillbillAPISecret
}

// WithAccountID adds the accountID to the get invoice item tags params
func (o *GetInvoiceItemTagsParams) WithAccountID(accountID strfmt.UUID) *GetInvoiceItemTagsParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the get invoice item tags params
func (o *GetInvoiceItemTagsParams) SetAccountID(accountID strfmt.UUID) {
	o.AccountID = accountID
}

// WithAudit adds the audit to the get invoice item tags params
func (o *GetInvoiceItemTagsParams) WithAudit(audit *string) *GetInvoiceItemTagsParams {
	o.SetAudit(audit)
	return o
}

// SetAudit adds the audit to the get invoice item tags params
func (o *GetInvoiceItemTagsParams) SetAudit(audit *string) {
	o.Audit = audit
}

// WithIncludedDeleted adds the includedDeleted to the get invoice item tags params
func (o *GetInvoiceItemTagsParams) WithIncludedDeleted(includedDeleted *bool) *GetInvoiceItemTagsParams {
	o.SetIncludedDeleted(includedDeleted)
	return o
}

// SetIncludedDeleted adds the includedDeleted to the get invoice item tags params
func (o *GetInvoiceItemTagsParams) SetIncludedDeleted(includedDeleted *bool) {
	o.IncludedDeleted = includedDeleted
}

// WithInvoiceItemID adds the invoiceItemID to the get invoice item tags params
func (o *GetInvoiceItemTagsParams) WithInvoiceItemID(invoiceItemID strfmt.UUID) *GetInvoiceItemTagsParams {
	o.SetInvoiceItemID(invoiceItemID)
	return o
}

// SetInvoiceItemID adds the invoiceItemId to the get invoice item tags params
func (o *GetInvoiceItemTagsParams) SetInvoiceItemID(invoiceItemID strfmt.UUID) {
	o.InvoiceItemID = invoiceItemID
}

// WriteToRequest writes these params to a swagger request
func (o *GetInvoiceItemTagsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param X-Killbill-ApiKey
	if err := r.SetHeaderParam("X-Killbill-ApiKey", o.XKillbillAPIKey); err != nil {
		return err
	}

	// header param X-Killbill-ApiSecret
	if err := r.SetHeaderParam("X-Killbill-ApiSecret", o.XKillbillAPISecret); err != nil {
		return err
	}

	// query param accountId
	qrAccountID := o.AccountID
	qAccountID := qrAccountID.String()
	if qAccountID != "" {
		if err := r.SetQueryParam("accountId", qAccountID); err != nil {
			return err
		}
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

	// path param invoiceItemId
	if err := r.SetPathParam("invoiceItemId", o.InvoiceItemID.String()); err != nil {
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
