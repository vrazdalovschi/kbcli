// Code generated by go-swagger; DO NOT EDIT.

package catalog

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

// NewGetAvailableAddonsParams creates a new GetAvailableAddonsParams object
// with the default values initialized.
func NewGetAvailableAddonsParams() *GetAvailableAddonsParams {
	var ()
	return &GetAvailableAddonsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAvailableAddonsParamsWithTimeout creates a new GetAvailableAddonsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAvailableAddonsParamsWithTimeout(timeout time.Duration) *GetAvailableAddonsParams {
	var ()
	return &GetAvailableAddonsParams{

		timeout: timeout,
	}
}

// NewGetAvailableAddonsParamsWithContext creates a new GetAvailableAddonsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAvailableAddonsParamsWithContext(ctx context.Context) *GetAvailableAddonsParams {
	var ()
	return &GetAvailableAddonsParams{

		Context: ctx,
	}
}

// NewGetAvailableAddonsParamsWithHTTPClient creates a new GetAvailableAddonsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAvailableAddonsParamsWithHTTPClient(client *http.Client) *GetAvailableAddonsParams {
	var ()
	return &GetAvailableAddonsParams{
		HTTPClient: client,
	}
}

/*GetAvailableAddonsParams contains all the parameters to send to the API endpoint
for the get available addons operation typically these are written to a http.Request
*/
type GetAvailableAddonsParams struct {

	/*XKillbillAPIKey*/
	XKillbillAPIKey string
	/*XKillbillAPISecret*/
	XKillbillAPISecret string
	/*AccountID*/
	AccountID *strfmt.UUID
	/*BaseProductName*/
	BaseProductName *string
	/*PriceListName*/
	PriceListName *string

	WithStackTrace        *bool // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithTimeout adds the timeout to the get available addons params
func (o *GetAvailableAddonsParams) WithTimeout(timeout time.Duration) *GetAvailableAddonsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get available addons params
func (o *GetAvailableAddonsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get available addons params
func (o *GetAvailableAddonsParams) WithContext(ctx context.Context) *GetAvailableAddonsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get available addons params
func (o *GetAvailableAddonsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get available addons params
func (o *GetAvailableAddonsParams) WithHTTPClient(client *http.Client) *GetAvailableAddonsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get available addons params
func (o *GetAvailableAddonsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillAPIKey adds the xKillbillAPIKey to the get available addons params
func (o *GetAvailableAddonsParams) WithXKillbillAPIKey(xKillbillAPIKey string) *GetAvailableAddonsParams {
	o.SetXKillbillAPIKey(xKillbillAPIKey)
	return o
}

// SetXKillbillAPIKey adds the xKillbillApiKey to the get available addons params
func (o *GetAvailableAddonsParams) SetXKillbillAPIKey(xKillbillAPIKey string) {
	o.XKillbillAPIKey = xKillbillAPIKey
}

// WithXKillbillAPISecret adds the xKillbillAPISecret to the get available addons params
func (o *GetAvailableAddonsParams) WithXKillbillAPISecret(xKillbillAPISecret string) *GetAvailableAddonsParams {
	o.SetXKillbillAPISecret(xKillbillAPISecret)
	return o
}

// SetXKillbillAPISecret adds the xKillbillApiSecret to the get available addons params
func (o *GetAvailableAddonsParams) SetXKillbillAPISecret(xKillbillAPISecret string) {
	o.XKillbillAPISecret = xKillbillAPISecret
}

// WithAccountID adds the accountID to the get available addons params
func (o *GetAvailableAddonsParams) WithAccountID(accountID *strfmt.UUID) *GetAvailableAddonsParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the get available addons params
func (o *GetAvailableAddonsParams) SetAccountID(accountID *strfmt.UUID) {
	o.AccountID = accountID
}

// WithBaseProductName adds the baseProductName to the get available addons params
func (o *GetAvailableAddonsParams) WithBaseProductName(baseProductName *string) *GetAvailableAddonsParams {
	o.SetBaseProductName(baseProductName)
	return o
}

// SetBaseProductName adds the baseProductName to the get available addons params
func (o *GetAvailableAddonsParams) SetBaseProductName(baseProductName *string) {
	o.BaseProductName = baseProductName
}

// WithPriceListName adds the priceListName to the get available addons params
func (o *GetAvailableAddonsParams) WithPriceListName(priceListName *string) *GetAvailableAddonsParams {
	o.SetPriceListName(priceListName)
	return o
}

// SetPriceListName adds the priceListName to the get available addons params
func (o *GetAvailableAddonsParams) SetPriceListName(priceListName *string) {
	o.PriceListName = priceListName
}

// WriteToRequest writes these params to a swagger request
func (o *GetAvailableAddonsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.AccountID != nil {

		// query param accountId
		var qrAccountID strfmt.UUID
		if o.AccountID != nil {
			qrAccountID = *o.AccountID
		}
		qAccountID := qrAccountID.String()
		if qAccountID != "" {
			if err := r.SetQueryParam("accountId", qAccountID); err != nil {
				return err
			}
		}

	}

	if o.BaseProductName != nil {

		// query param baseProductName
		var qrBaseProductName string
		if o.BaseProductName != nil {
			qrBaseProductName = *o.BaseProductName
		}
		qBaseProductName := qrBaseProductName
		if qBaseProductName != "" {
			if err := r.SetQueryParam("baseProductName", qBaseProductName); err != nil {
				return err
			}
		}

	}

	if o.PriceListName != nil {

		// query param priceListName
		var qrPriceListName string
		if o.PriceListName != nil {
			qrPriceListName = *o.PriceListName
		}
		qPriceListName := qrPriceListName
		if qPriceListName != "" {
			if err := r.SetQueryParam("priceListName", qPriceListName); err != nil {
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
