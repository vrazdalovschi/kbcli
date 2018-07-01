// Code generated by go-swagger; DO NOT EDIT.

package invoice_payment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetInvoicePaymentCustomFieldsParams creates a new GetInvoicePaymentCustomFieldsParams object
// with the default values initialized.
func NewGetInvoicePaymentCustomFieldsParams() *GetInvoicePaymentCustomFieldsParams {
	var (
		auditDefault = string("NONE")
	)
	return &GetInvoicePaymentCustomFieldsParams{
		Audit: &auditDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetInvoicePaymentCustomFieldsParamsWithTimeout creates a new GetInvoicePaymentCustomFieldsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetInvoicePaymentCustomFieldsParamsWithTimeout(timeout time.Duration) *GetInvoicePaymentCustomFieldsParams {
	var (
		auditDefault = string("NONE")
	)
	return &GetInvoicePaymentCustomFieldsParams{
		Audit: &auditDefault,

		timeout: timeout,
	}
}

// NewGetInvoicePaymentCustomFieldsParamsWithContext creates a new GetInvoicePaymentCustomFieldsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetInvoicePaymentCustomFieldsParamsWithContext(ctx context.Context) *GetInvoicePaymentCustomFieldsParams {
	var (
		auditDefault = string("NONE")
	)
	return &GetInvoicePaymentCustomFieldsParams{
		Audit: &auditDefault,

		Context: ctx,
	}
}

// NewGetInvoicePaymentCustomFieldsParamsWithHTTPClient creates a new GetInvoicePaymentCustomFieldsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetInvoicePaymentCustomFieldsParamsWithHTTPClient(client *http.Client) *GetInvoicePaymentCustomFieldsParams {
	var (
		auditDefault = string("NONE")
	)
	return &GetInvoicePaymentCustomFieldsParams{
		Audit:      &auditDefault,
		HTTPClient: client,
	}
}

/*GetInvoicePaymentCustomFieldsParams contains all the parameters to send to the API endpoint
for the get invoice payment custom fields operation typically these are written to a http.Request
*/
type GetInvoicePaymentCustomFieldsParams struct {

	/*XKillbillAPIKey*/
	XKillbillAPIKey string
	/*XKillbillAPISecret*/
	XKillbillAPISecret string
	/*Audit*/
	Audit *string
	/*PaymentID*/
	PaymentID strfmt.UUID

	WithStackTrace        *bool // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithTimeout adds the timeout to the get invoice payment custom fields params
func (o *GetInvoicePaymentCustomFieldsParams) WithTimeout(timeout time.Duration) *GetInvoicePaymentCustomFieldsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get invoice payment custom fields params
func (o *GetInvoicePaymentCustomFieldsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get invoice payment custom fields params
func (o *GetInvoicePaymentCustomFieldsParams) WithContext(ctx context.Context) *GetInvoicePaymentCustomFieldsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get invoice payment custom fields params
func (o *GetInvoicePaymentCustomFieldsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get invoice payment custom fields params
func (o *GetInvoicePaymentCustomFieldsParams) WithHTTPClient(client *http.Client) *GetInvoicePaymentCustomFieldsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get invoice payment custom fields params
func (o *GetInvoicePaymentCustomFieldsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillAPIKey adds the xKillbillAPIKey to the get invoice payment custom fields params
func (o *GetInvoicePaymentCustomFieldsParams) WithXKillbillAPIKey(xKillbillAPIKey string) *GetInvoicePaymentCustomFieldsParams {
	o.SetXKillbillAPIKey(xKillbillAPIKey)
	return o
}

// SetXKillbillAPIKey adds the xKillbillApiKey to the get invoice payment custom fields params
func (o *GetInvoicePaymentCustomFieldsParams) SetXKillbillAPIKey(xKillbillAPIKey string) {
	o.XKillbillAPIKey = xKillbillAPIKey
}

// WithXKillbillAPISecret adds the xKillbillAPISecret to the get invoice payment custom fields params
func (o *GetInvoicePaymentCustomFieldsParams) WithXKillbillAPISecret(xKillbillAPISecret string) *GetInvoicePaymentCustomFieldsParams {
	o.SetXKillbillAPISecret(xKillbillAPISecret)
	return o
}

// SetXKillbillAPISecret adds the xKillbillApiSecret to the get invoice payment custom fields params
func (o *GetInvoicePaymentCustomFieldsParams) SetXKillbillAPISecret(xKillbillAPISecret string) {
	o.XKillbillAPISecret = xKillbillAPISecret
}

// WithAudit adds the audit to the get invoice payment custom fields params
func (o *GetInvoicePaymentCustomFieldsParams) WithAudit(audit *string) *GetInvoicePaymentCustomFieldsParams {
	o.SetAudit(audit)
	return o
}

// SetAudit adds the audit to the get invoice payment custom fields params
func (o *GetInvoicePaymentCustomFieldsParams) SetAudit(audit *string) {
	o.Audit = audit
}

// WithPaymentID adds the paymentID to the get invoice payment custom fields params
func (o *GetInvoicePaymentCustomFieldsParams) WithPaymentID(paymentID strfmt.UUID) *GetInvoicePaymentCustomFieldsParams {
	o.SetPaymentID(paymentID)
	return o
}

// SetPaymentID adds the paymentId to the get invoice payment custom fields params
func (o *GetInvoicePaymentCustomFieldsParams) SetPaymentID(paymentID strfmt.UUID) {
	o.PaymentID = paymentID
}

// WriteToRequest writes these params to a swagger request
func (o *GetInvoicePaymentCustomFieldsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param paymentId
	if err := r.SetPathParam("paymentId", o.PaymentID.String()); err != nil {
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
