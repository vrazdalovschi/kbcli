// Code generated by go-swagger; DO NOT EDIT.

package invoice

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

// NewGetInvoiceByNumberParams creates a new GetInvoiceByNumberParams object
// with the default values initialized.
func NewGetInvoiceByNumberParams() *GetInvoiceByNumberParams {
	var (
		auditDefault             = string("NONE")
		withChildrenItemsDefault = bool(false)
		withItemsDefault         = bool(false)
	)
	return &GetInvoiceByNumberParams{
		Audit:             &auditDefault,
		WithChildrenItems: &withChildrenItemsDefault,
		WithItems:         &withItemsDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetInvoiceByNumberParamsWithTimeout creates a new GetInvoiceByNumberParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetInvoiceByNumberParamsWithTimeout(timeout time.Duration) *GetInvoiceByNumberParams {
	var (
		auditDefault             = string("NONE")
		withChildrenItemsDefault = bool(false)
		withItemsDefault         = bool(false)
	)
	return &GetInvoiceByNumberParams{
		Audit:             &auditDefault,
		WithChildrenItems: &withChildrenItemsDefault,
		WithItems:         &withItemsDefault,

		timeout: timeout,
	}
}

// NewGetInvoiceByNumberParamsWithContext creates a new GetInvoiceByNumberParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetInvoiceByNumberParamsWithContext(ctx context.Context) *GetInvoiceByNumberParams {
	var (
		auditDefault             = string("NONE")
		withChildrenItemsDefault = bool(false)
		withItemsDefault         = bool(false)
	)
	return &GetInvoiceByNumberParams{
		Audit:             &auditDefault,
		WithChildrenItems: &withChildrenItemsDefault,
		WithItems:         &withItemsDefault,

		Context: ctx,
	}
}

// NewGetInvoiceByNumberParamsWithHTTPClient creates a new GetInvoiceByNumberParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetInvoiceByNumberParamsWithHTTPClient(client *http.Client) *GetInvoiceByNumberParams {
	var (
		auditDefault             = string("NONE")
		withChildrenItemsDefault = bool(false)
		withItemsDefault         = bool(false)
	)
	return &GetInvoiceByNumberParams{
		Audit:             &auditDefault,
		WithChildrenItems: &withChildrenItemsDefault,
		WithItems:         &withItemsDefault,
		HTTPClient:        client,
	}
}

/*GetInvoiceByNumberParams contains all the parameters to send to the API endpoint
for the get invoice by number operation typically these are written to a http.Request
*/
type GetInvoiceByNumberParams struct {

	/*XKillbillAPIKey*/
	XKillbillAPIKey string
	/*XKillbillAPISecret*/
	XKillbillAPISecret string
	/*Audit*/
	Audit *string
	/*InvoiceNumber*/
	InvoiceNumber int32
	/*WithChildrenItems*/
	WithChildrenItems *bool
	/*WithItems*/
	WithItems *bool

	WithStackTrace        *bool // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithTimeout adds the timeout to the get invoice by number params
func (o *GetInvoiceByNumberParams) WithTimeout(timeout time.Duration) *GetInvoiceByNumberParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get invoice by number params
func (o *GetInvoiceByNumberParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get invoice by number params
func (o *GetInvoiceByNumberParams) WithContext(ctx context.Context) *GetInvoiceByNumberParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get invoice by number params
func (o *GetInvoiceByNumberParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get invoice by number params
func (o *GetInvoiceByNumberParams) WithHTTPClient(client *http.Client) *GetInvoiceByNumberParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get invoice by number params
func (o *GetInvoiceByNumberParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillAPIKey adds the xKillbillAPIKey to the get invoice by number params
func (o *GetInvoiceByNumberParams) WithXKillbillAPIKey(xKillbillAPIKey string) *GetInvoiceByNumberParams {
	o.SetXKillbillAPIKey(xKillbillAPIKey)
	return o
}

// SetXKillbillAPIKey adds the xKillbillApiKey to the get invoice by number params
func (o *GetInvoiceByNumberParams) SetXKillbillAPIKey(xKillbillAPIKey string) {
	o.XKillbillAPIKey = xKillbillAPIKey
}

// WithXKillbillAPISecret adds the xKillbillAPISecret to the get invoice by number params
func (o *GetInvoiceByNumberParams) WithXKillbillAPISecret(xKillbillAPISecret string) *GetInvoiceByNumberParams {
	o.SetXKillbillAPISecret(xKillbillAPISecret)
	return o
}

// SetXKillbillAPISecret adds the xKillbillApiSecret to the get invoice by number params
func (o *GetInvoiceByNumberParams) SetXKillbillAPISecret(xKillbillAPISecret string) {
	o.XKillbillAPISecret = xKillbillAPISecret
}

// WithAudit adds the audit to the get invoice by number params
func (o *GetInvoiceByNumberParams) WithAudit(audit *string) *GetInvoiceByNumberParams {
	o.SetAudit(audit)
	return o
}

// SetAudit adds the audit to the get invoice by number params
func (o *GetInvoiceByNumberParams) SetAudit(audit *string) {
	o.Audit = audit
}

// WithInvoiceNumber adds the invoiceNumber to the get invoice by number params
func (o *GetInvoiceByNumberParams) WithInvoiceNumber(invoiceNumber int32) *GetInvoiceByNumberParams {
	o.SetInvoiceNumber(invoiceNumber)
	return o
}

// SetInvoiceNumber adds the invoiceNumber to the get invoice by number params
func (o *GetInvoiceByNumberParams) SetInvoiceNumber(invoiceNumber int32) {
	o.InvoiceNumber = invoiceNumber
}

// WithWithChildrenItems adds the withChildrenItems to the get invoice by number params
func (o *GetInvoiceByNumberParams) WithWithChildrenItems(withChildrenItems *bool) *GetInvoiceByNumberParams {
	o.SetWithChildrenItems(withChildrenItems)
	return o
}

// SetWithChildrenItems adds the withChildrenItems to the get invoice by number params
func (o *GetInvoiceByNumberParams) SetWithChildrenItems(withChildrenItems *bool) {
	o.WithChildrenItems = withChildrenItems
}

// WithWithItems adds the withItems to the get invoice by number params
func (o *GetInvoiceByNumberParams) WithWithItems(withItems *bool) *GetInvoiceByNumberParams {
	o.SetWithItems(withItems)
	return o
}

// SetWithItems adds the withItems to the get invoice by number params
func (o *GetInvoiceByNumberParams) SetWithItems(withItems *bool) {
	o.WithItems = withItems
}

// WriteToRequest writes these params to a swagger request
func (o *GetInvoiceByNumberParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param invoiceNumber
	if err := r.SetPathParam("invoiceNumber", swag.FormatInt32(o.InvoiceNumber)); err != nil {
		return err
	}

	if o.WithChildrenItems != nil {

		// query param withChildrenItems
		var qrWithChildrenItems bool
		if o.WithChildrenItems != nil {
			qrWithChildrenItems = *o.WithChildrenItems
		}
		qWithChildrenItems := swag.FormatBool(qrWithChildrenItems)
		if qWithChildrenItems != "" {
			if err := r.SetQueryParam("withChildrenItems", qWithChildrenItems); err != nil {
				return err
			}
		}

	}

	if o.WithItems != nil {

		// query param withItems
		var qrWithItems bool
		if o.WithItems != nil {
			qrWithItems = *o.WithItems
		}
		qWithItems := swag.FormatBool(qrWithItems)
		if qWithItems != "" {
			if err := r.SetQueryParam("withItems", qWithItems); err != nil {
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
