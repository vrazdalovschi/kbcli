// Code generated by go-swagger; DO NOT EDIT.

package subscription

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

	kbmodel "github.com/killbill/kbcli/kbmodel"
)

// NewCreateSubscriptionWithAddOnsParams creates a new CreateSubscriptionWithAddOnsParams object
// with the default values initialized.
func NewCreateSubscriptionWithAddOnsParams() *CreateSubscriptionWithAddOnsParams {
	var (
		callCompletionDefault             = bool(false)
		callTimeoutSecDefault             = int64(3)
		migratedDefault                   = bool(false)
		renameKeyIfExistsAndUnusedDefault = bool(true)
	)
	return &CreateSubscriptionWithAddOnsParams{
		CallCompletion:             &callCompletionDefault,
		CallTimeoutSec:             &callTimeoutSecDefault,
		Migrated:                   &migratedDefault,
		RenameKeyIfExistsAndUnused: &renameKeyIfExistsAndUnusedDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateSubscriptionWithAddOnsParamsWithTimeout creates a new CreateSubscriptionWithAddOnsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateSubscriptionWithAddOnsParamsWithTimeout(timeout time.Duration) *CreateSubscriptionWithAddOnsParams {
	var (
		callCompletionDefault             = bool(false)
		callTimeoutSecDefault             = int64(3)
		migratedDefault                   = bool(false)
		renameKeyIfExistsAndUnusedDefault = bool(true)
	)
	return &CreateSubscriptionWithAddOnsParams{
		CallCompletion:             &callCompletionDefault,
		CallTimeoutSec:             &callTimeoutSecDefault,
		Migrated:                   &migratedDefault,
		RenameKeyIfExistsAndUnused: &renameKeyIfExistsAndUnusedDefault,

		timeout: timeout,
	}
}

// NewCreateSubscriptionWithAddOnsParamsWithContext creates a new CreateSubscriptionWithAddOnsParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateSubscriptionWithAddOnsParamsWithContext(ctx context.Context) *CreateSubscriptionWithAddOnsParams {
	var (
		callCompletionDefault             = bool(false)
		callTimeoutSecDefault             = int64(3)
		migratedDefault                   = bool(false)
		renameKeyIfExistsAndUnusedDefault = bool(true)
	)
	return &CreateSubscriptionWithAddOnsParams{
		CallCompletion:             &callCompletionDefault,
		CallTimeoutSec:             &callTimeoutSecDefault,
		Migrated:                   &migratedDefault,
		RenameKeyIfExistsAndUnused: &renameKeyIfExistsAndUnusedDefault,

		Context: ctx,
	}
}

// NewCreateSubscriptionWithAddOnsParamsWithHTTPClient creates a new CreateSubscriptionWithAddOnsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateSubscriptionWithAddOnsParamsWithHTTPClient(client *http.Client) *CreateSubscriptionWithAddOnsParams {
	var (
		callCompletionDefault             = bool(false)
		callTimeoutSecDefault             = int64(3)
		migratedDefault                   = bool(false)
		renameKeyIfExistsAndUnusedDefault = bool(true)
	)
	return &CreateSubscriptionWithAddOnsParams{
		CallCompletion:             &callCompletionDefault,
		CallTimeoutSec:             &callTimeoutSecDefault,
		Migrated:                   &migratedDefault,
		RenameKeyIfExistsAndUnused: &renameKeyIfExistsAndUnusedDefault,
		HTTPClient:                 client,
	}
}

/*CreateSubscriptionWithAddOnsParams contains all the parameters to send to the API endpoint
for the create subscription with add ons operation typically these are written to a http.Request
*/
type CreateSubscriptionWithAddOnsParams struct {

	/*XKillbillComment*/
	XKillbillComment *string
	/*XKillbillCreatedBy*/
	XKillbillCreatedBy string
	/*XKillbillReason*/
	XKillbillReason *string
	/*BillingDate*/
	BillingDate *strfmt.Date
	/*Body*/
	Body []*kbmodel.Subscription
	/*CallCompletion*/
	CallCompletion *bool
	/*CallTimeoutSec*/
	CallTimeoutSec *int64
	/*EntitlementDate*/
	EntitlementDate *strfmt.Date
	/*Migrated*/
	Migrated *bool
	/*PluginProperty*/
	PluginProperty []string
	/*RenameKeyIfExistsAndUnused*/
	RenameKeyIfExistsAndUnused *bool

	WithStackTrace        *bool // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithTimeout adds the timeout to the create subscription with add ons params
func (o *CreateSubscriptionWithAddOnsParams) WithTimeout(timeout time.Duration) *CreateSubscriptionWithAddOnsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create subscription with add ons params
func (o *CreateSubscriptionWithAddOnsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create subscription with add ons params
func (o *CreateSubscriptionWithAddOnsParams) WithContext(ctx context.Context) *CreateSubscriptionWithAddOnsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create subscription with add ons params
func (o *CreateSubscriptionWithAddOnsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create subscription with add ons params
func (o *CreateSubscriptionWithAddOnsParams) WithHTTPClient(client *http.Client) *CreateSubscriptionWithAddOnsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create subscription with add ons params
func (o *CreateSubscriptionWithAddOnsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillComment adds the xKillbillComment to the create subscription with add ons params
func (o *CreateSubscriptionWithAddOnsParams) WithXKillbillComment(xKillbillComment *string) *CreateSubscriptionWithAddOnsParams {
	o.SetXKillbillComment(xKillbillComment)
	return o
}

// SetXKillbillComment adds the xKillbillComment to the create subscription with add ons params
func (o *CreateSubscriptionWithAddOnsParams) SetXKillbillComment(xKillbillComment *string) {
	o.XKillbillComment = xKillbillComment
}

// WithXKillbillCreatedBy adds the xKillbillCreatedBy to the create subscription with add ons params
func (o *CreateSubscriptionWithAddOnsParams) WithXKillbillCreatedBy(xKillbillCreatedBy string) *CreateSubscriptionWithAddOnsParams {
	o.SetXKillbillCreatedBy(xKillbillCreatedBy)
	return o
}

// SetXKillbillCreatedBy adds the xKillbillCreatedBy to the create subscription with add ons params
func (o *CreateSubscriptionWithAddOnsParams) SetXKillbillCreatedBy(xKillbillCreatedBy string) {
	o.XKillbillCreatedBy = xKillbillCreatedBy
}

// WithXKillbillReason adds the xKillbillReason to the create subscription with add ons params
func (o *CreateSubscriptionWithAddOnsParams) WithXKillbillReason(xKillbillReason *string) *CreateSubscriptionWithAddOnsParams {
	o.SetXKillbillReason(xKillbillReason)
	return o
}

// SetXKillbillReason adds the xKillbillReason to the create subscription with add ons params
func (o *CreateSubscriptionWithAddOnsParams) SetXKillbillReason(xKillbillReason *string) {
	o.XKillbillReason = xKillbillReason
}

// WithBillingDate adds the billingDate to the create subscription with add ons params
func (o *CreateSubscriptionWithAddOnsParams) WithBillingDate(billingDate *strfmt.Date) *CreateSubscriptionWithAddOnsParams {
	o.SetBillingDate(billingDate)
	return o
}

// SetBillingDate adds the billingDate to the create subscription with add ons params
func (o *CreateSubscriptionWithAddOnsParams) SetBillingDate(billingDate *strfmt.Date) {
	o.BillingDate = billingDate
}

// WithBody adds the body to the create subscription with add ons params
func (o *CreateSubscriptionWithAddOnsParams) WithBody(body []*kbmodel.Subscription) *CreateSubscriptionWithAddOnsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create subscription with add ons params
func (o *CreateSubscriptionWithAddOnsParams) SetBody(body []*kbmodel.Subscription) {
	o.Body = body
}

// WithCallCompletion adds the callCompletion to the create subscription with add ons params
func (o *CreateSubscriptionWithAddOnsParams) WithCallCompletion(callCompletion *bool) *CreateSubscriptionWithAddOnsParams {
	o.SetCallCompletion(callCompletion)
	return o
}

// SetCallCompletion adds the callCompletion to the create subscription with add ons params
func (o *CreateSubscriptionWithAddOnsParams) SetCallCompletion(callCompletion *bool) {
	o.CallCompletion = callCompletion
}

// WithCallTimeoutSec adds the callTimeoutSec to the create subscription with add ons params
func (o *CreateSubscriptionWithAddOnsParams) WithCallTimeoutSec(callTimeoutSec *int64) *CreateSubscriptionWithAddOnsParams {
	o.SetCallTimeoutSec(callTimeoutSec)
	return o
}

// SetCallTimeoutSec adds the callTimeoutSec to the create subscription with add ons params
func (o *CreateSubscriptionWithAddOnsParams) SetCallTimeoutSec(callTimeoutSec *int64) {
	o.CallTimeoutSec = callTimeoutSec
}

// WithEntitlementDate adds the entitlementDate to the create subscription with add ons params
func (o *CreateSubscriptionWithAddOnsParams) WithEntitlementDate(entitlementDate *strfmt.Date) *CreateSubscriptionWithAddOnsParams {
	o.SetEntitlementDate(entitlementDate)
	return o
}

// SetEntitlementDate adds the entitlementDate to the create subscription with add ons params
func (o *CreateSubscriptionWithAddOnsParams) SetEntitlementDate(entitlementDate *strfmt.Date) {
	o.EntitlementDate = entitlementDate
}

// WithMigrated adds the migrated to the create subscription with add ons params
func (o *CreateSubscriptionWithAddOnsParams) WithMigrated(migrated *bool) *CreateSubscriptionWithAddOnsParams {
	o.SetMigrated(migrated)
	return o
}

// SetMigrated adds the migrated to the create subscription with add ons params
func (o *CreateSubscriptionWithAddOnsParams) SetMigrated(migrated *bool) {
	o.Migrated = migrated
}

// WithPluginProperty adds the pluginProperty to the create subscription with add ons params
func (o *CreateSubscriptionWithAddOnsParams) WithPluginProperty(pluginProperty []string) *CreateSubscriptionWithAddOnsParams {
	o.SetPluginProperty(pluginProperty)
	return o
}

// SetPluginProperty adds the pluginProperty to the create subscription with add ons params
func (o *CreateSubscriptionWithAddOnsParams) SetPluginProperty(pluginProperty []string) {
	o.PluginProperty = pluginProperty
}

// WithRenameKeyIfExistsAndUnused adds the renameKeyIfExistsAndUnused to the create subscription with add ons params
func (o *CreateSubscriptionWithAddOnsParams) WithRenameKeyIfExistsAndUnused(renameKeyIfExistsAndUnused *bool) *CreateSubscriptionWithAddOnsParams {
	o.SetRenameKeyIfExistsAndUnused(renameKeyIfExistsAndUnused)
	return o
}

// SetRenameKeyIfExistsAndUnused adds the renameKeyIfExistsAndUnused to the create subscription with add ons params
func (o *CreateSubscriptionWithAddOnsParams) SetRenameKeyIfExistsAndUnused(renameKeyIfExistsAndUnused *bool) {
	o.RenameKeyIfExistsAndUnused = renameKeyIfExistsAndUnused
}

// WriteToRequest writes these params to a swagger request
func (o *CreateSubscriptionWithAddOnsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.XKillbillComment != nil {

		// header param X-Killbill-Comment
		if err := r.SetHeaderParam("X-Killbill-Comment", *o.XKillbillComment); err != nil {
			return err
		}

	}

	// header param X-Killbill-CreatedBy
	if err := r.SetHeaderParam("X-Killbill-CreatedBy", o.XKillbillCreatedBy); err != nil {
		return err
	}

	if o.XKillbillReason != nil {

		// header param X-Killbill-Reason
		if err := r.SetHeaderParam("X-Killbill-Reason", *o.XKillbillReason); err != nil {
			return err
		}

	}

	if o.BillingDate != nil {

		// query param billingDate
		var qrBillingDate strfmt.Date
		if o.BillingDate != nil {
			qrBillingDate = *o.BillingDate
		}
		qBillingDate := qrBillingDate.String()
		if qBillingDate != "" {
			if err := r.SetQueryParam("billingDate", qBillingDate); err != nil {
				return err
			}
		}

	}

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if o.CallCompletion != nil {

		// query param callCompletion
		var qrCallCompletion bool
		if o.CallCompletion != nil {
			qrCallCompletion = *o.CallCompletion
		}
		qCallCompletion := swag.FormatBool(qrCallCompletion)
		if qCallCompletion != "" {
			if err := r.SetQueryParam("callCompletion", qCallCompletion); err != nil {
				return err
			}
		}

	}

	if o.CallTimeoutSec != nil {

		// query param callTimeoutSec
		var qrCallTimeoutSec int64
		if o.CallTimeoutSec != nil {
			qrCallTimeoutSec = *o.CallTimeoutSec
		}
		qCallTimeoutSec := swag.FormatInt64(qrCallTimeoutSec)
		if qCallTimeoutSec != "" {
			if err := r.SetQueryParam("callTimeoutSec", qCallTimeoutSec); err != nil {
				return err
			}
		}

	}

	if o.EntitlementDate != nil {

		// query param entitlementDate
		var qrEntitlementDate strfmt.Date
		if o.EntitlementDate != nil {
			qrEntitlementDate = *o.EntitlementDate
		}
		qEntitlementDate := qrEntitlementDate.String()
		if qEntitlementDate != "" {
			if err := r.SetQueryParam("entitlementDate", qEntitlementDate); err != nil {
				return err
			}
		}

	}

	if o.Migrated != nil {

		// query param migrated
		var qrMigrated bool
		if o.Migrated != nil {
			qrMigrated = *o.Migrated
		}
		qMigrated := swag.FormatBool(qrMigrated)
		if qMigrated != "" {
			if err := r.SetQueryParam("migrated", qMigrated); err != nil {
				return err
			}
		}

	}

	valuesPluginProperty := o.PluginProperty

	joinedPluginProperty := swag.JoinByFormat(valuesPluginProperty, "multi")
	// query array param pluginProperty
	if err := r.SetQueryParam("pluginProperty", joinedPluginProperty...); err != nil {
		return err
	}

	if o.RenameKeyIfExistsAndUnused != nil {

		// query param renameKeyIfExistsAndUnused
		var qrRenameKeyIfExistsAndUnused bool
		if o.RenameKeyIfExistsAndUnused != nil {
			qrRenameKeyIfExistsAndUnused = *o.RenameKeyIfExistsAndUnused
		}
		qRenameKeyIfExistsAndUnused := swag.FormatBool(qrRenameKeyIfExistsAndUnused)
		if qRenameKeyIfExistsAndUnused != "" {
			if err := r.SetQueryParam("renameKeyIfExistsAndUnused", qRenameKeyIfExistsAndUnused); err != nil {
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
