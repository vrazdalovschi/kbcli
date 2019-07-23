// Code generated by go-swagger; DO NOT EDIT.

package kbmodel

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Credit credit
// swagger:model Credit
type Credit struct {

	// account Id
	// Required: true
	AccountID *strfmt.UUID `json:"accountId"`

	// audit logs
	AuditLogs []*AuditLog `json:"auditLogs"`

	// credit amount
	// Required: true
	CreditAmount *float64 `json:"creditAmount"`

	// credit Id
	CreditID strfmt.UUID `json:"creditId,omitempty"`

	// currency
	Currency CreditCurrencyEnum `json:"currency,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// effective date
	EffectiveDate strfmt.Date `json:"effectiveDate,omitempty"`

	// invoice Id
	InvoiceID strfmt.UUID `json:"invoiceId,omitempty"`

	// invoice number
	InvoiceNumber string `json:"invoiceNumber,omitempty"`

	// item details
	ItemDetails string `json:"itemDetails,omitempty"`
}

// Validate validates this credit
func (m *Credit) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccountID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateAuditLogs(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCreditAmount(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCreditID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCurrency(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateEffectiveDate(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateInvoiceID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Credit) validateAccountID(formats strfmt.Registry) error {

	if err := validate.Required("accountId", "body", m.AccountID); err != nil {
		return err
	}

	if err := validate.FormatOf("accountId", "body", "uuid", m.AccountID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Credit) validateAuditLogs(formats strfmt.Registry) error {

	if swag.IsZero(m.AuditLogs) { // not required
		return nil
	}

	for i := 0; i < len(m.AuditLogs); i++ {

		if swag.IsZero(m.AuditLogs[i]) { // not required
			continue
		}

		if m.AuditLogs[i] != nil {

			if err := m.AuditLogs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("auditLogs" + "." + strconv.Itoa(i))
				}
				return err
			}

		}

	}

	return nil
}

func (m *Credit) validateCreditAmount(formats strfmt.Registry) error {

	if err := validate.Required("creditAmount", "body", m.CreditAmount); err != nil {
		return err
	}

	return nil
}

func (m *Credit) validateCreditID(formats strfmt.Registry) error {

	if swag.IsZero(m.CreditID) { // not required
		return nil
	}

	if err := validate.FormatOf("creditId", "body", "uuid", m.CreditID.String(), formats); err != nil {
		return err
	}

	return nil
}

var creditTypeCurrencyPropEnum []interface{}

func init() {
	var res []CreditCurrencyEnum
	if err := json.Unmarshal([]byte(`["AED","AFN","ALL","AMD","ANG","AOA","ARS","AUD","AWG","AZN","BAM","BBD","BDT","BGN","BHD","BIF","BMD","BND","BOB","BRL","BSD","BTN","BWP","BYR","BZD","CAD","CDF","CHF","CLP","CNY","COP","CRC","CUC","CUP","CVE","CZK","DJF","DKK","DOP","DZD","EGP","ERN","ETB","EUR","FJD","FKP","GBP","GEL","GGP","GHS","GIP","GMD","GNF","GTQ","GYD","HKD","HNL","HRK","HTG","HUF","IDR","ILS","IMP","INR","IQD","IRR","ISK","JEP","JMD","JOD","JPY","KES","KGS","KHR","KMF","KPW","KRW","KWD","KYD","KZT","LAK","LBP","LKR","LRD","LSL","LTL","LVL","LYD","MAD","MDL","MGA","MKD","MMK","MNT","MOP","MRO","MUR","MVR","MWK","MXN","MYR","MZN","NAD","NGN","NIO","NOK","NPR","NZD","OMR","PAB","PEN","PGK","PHP","PKR","PLN","PYG","QAR","RON","RSD","RUB","RWF","SAR","SBD","SCR","SDG","SEK","SGD","SHP","SLL","SOS","SPL","SRD","STD","SVC","SYP","SZL","THB","TJS","TMT","TND","TOP","TRY","TTD","TVD","TWD","TZS","UAH","UGX","USD","UYU","UZS","VEF","VND","VUV","WST","XAF","XCD","XDR","XOF","XPF","YER","ZAR","ZMW","ZWD","BTC"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		creditTypeCurrencyPropEnum = append(creditTypeCurrencyPropEnum, v)
	}
}

type CreditCurrencyEnum string

const (

	// CreditCurrencyAED captures enum value "AED"
	CreditCurrencyAED CreditCurrencyEnum = "AED"

	// CreditCurrencyAFN captures enum value "AFN"
	CreditCurrencyAFN CreditCurrencyEnum = "AFN"

	// CreditCurrencyALL captures enum value "ALL"
	CreditCurrencyALL CreditCurrencyEnum = "ALL"

	// CreditCurrencyAMD captures enum value "AMD"
	CreditCurrencyAMD CreditCurrencyEnum = "AMD"

	// CreditCurrencyANG captures enum value "ANG"
	CreditCurrencyANG CreditCurrencyEnum = "ANG"

	// CreditCurrencyAOA captures enum value "AOA"
	CreditCurrencyAOA CreditCurrencyEnum = "AOA"

	// CreditCurrencyARS captures enum value "ARS"
	CreditCurrencyARS CreditCurrencyEnum = "ARS"

	// CreditCurrencyAUD captures enum value "AUD"
	CreditCurrencyAUD CreditCurrencyEnum = "AUD"

	// CreditCurrencyAWG captures enum value "AWG"
	CreditCurrencyAWG CreditCurrencyEnum = "AWG"

	// CreditCurrencyAZN captures enum value "AZN"
	CreditCurrencyAZN CreditCurrencyEnum = "AZN"

	// CreditCurrencyBAM captures enum value "BAM"
	CreditCurrencyBAM CreditCurrencyEnum = "BAM"

	// CreditCurrencyBBD captures enum value "BBD"
	CreditCurrencyBBD CreditCurrencyEnum = "BBD"

	// CreditCurrencyBDT captures enum value "BDT"
	CreditCurrencyBDT CreditCurrencyEnum = "BDT"

	// CreditCurrencyBGN captures enum value "BGN"
	CreditCurrencyBGN CreditCurrencyEnum = "BGN"

	// CreditCurrencyBHD captures enum value "BHD"
	CreditCurrencyBHD CreditCurrencyEnum = "BHD"

	// CreditCurrencyBIF captures enum value "BIF"
	CreditCurrencyBIF CreditCurrencyEnum = "BIF"

	// CreditCurrencyBMD captures enum value "BMD"
	CreditCurrencyBMD CreditCurrencyEnum = "BMD"

	// CreditCurrencyBND captures enum value "BND"
	CreditCurrencyBND CreditCurrencyEnum = "BND"

	// CreditCurrencyBOB captures enum value "BOB"
	CreditCurrencyBOB CreditCurrencyEnum = "BOB"

	// CreditCurrencyBRL captures enum value "BRL"
	CreditCurrencyBRL CreditCurrencyEnum = "BRL"

	// CreditCurrencyBSD captures enum value "BSD"
	CreditCurrencyBSD CreditCurrencyEnum = "BSD"

	// CreditCurrencyBTN captures enum value "BTN"
	CreditCurrencyBTN CreditCurrencyEnum = "BTN"

	// CreditCurrencyBWP captures enum value "BWP"
	CreditCurrencyBWP CreditCurrencyEnum = "BWP"

	// CreditCurrencyBYR captures enum value "BYR"
	CreditCurrencyBYR CreditCurrencyEnum = "BYR"

	// CreditCurrencyBZD captures enum value "BZD"
	CreditCurrencyBZD CreditCurrencyEnum = "BZD"

	// CreditCurrencyCAD captures enum value "CAD"
	CreditCurrencyCAD CreditCurrencyEnum = "CAD"

	// CreditCurrencyCDF captures enum value "CDF"
	CreditCurrencyCDF CreditCurrencyEnum = "CDF"

	// CreditCurrencyCHF captures enum value "CHF"
	CreditCurrencyCHF CreditCurrencyEnum = "CHF"

	// CreditCurrencyCLP captures enum value "CLP"
	CreditCurrencyCLP CreditCurrencyEnum = "CLP"

	// CreditCurrencyCNY captures enum value "CNY"
	CreditCurrencyCNY CreditCurrencyEnum = "CNY"

	// CreditCurrencyCOP captures enum value "COP"
	CreditCurrencyCOP CreditCurrencyEnum = "COP"

	// CreditCurrencyCRC captures enum value "CRC"
	CreditCurrencyCRC CreditCurrencyEnum = "CRC"

	// CreditCurrencyCUC captures enum value "CUC"
	CreditCurrencyCUC CreditCurrencyEnum = "CUC"

	// CreditCurrencyCUP captures enum value "CUP"
	CreditCurrencyCUP CreditCurrencyEnum = "CUP"

	// CreditCurrencyCVE captures enum value "CVE"
	CreditCurrencyCVE CreditCurrencyEnum = "CVE"

	// CreditCurrencyCZK captures enum value "CZK"
	CreditCurrencyCZK CreditCurrencyEnum = "CZK"

	// CreditCurrencyDJF captures enum value "DJF"
	CreditCurrencyDJF CreditCurrencyEnum = "DJF"

	// CreditCurrencyDKK captures enum value "DKK"
	CreditCurrencyDKK CreditCurrencyEnum = "DKK"

	// CreditCurrencyDOP captures enum value "DOP"
	CreditCurrencyDOP CreditCurrencyEnum = "DOP"

	// CreditCurrencyDZD captures enum value "DZD"
	CreditCurrencyDZD CreditCurrencyEnum = "DZD"

	// CreditCurrencyEGP captures enum value "EGP"
	CreditCurrencyEGP CreditCurrencyEnum = "EGP"

	// CreditCurrencyERN captures enum value "ERN"
	CreditCurrencyERN CreditCurrencyEnum = "ERN"

	// CreditCurrencyETB captures enum value "ETB"
	CreditCurrencyETB CreditCurrencyEnum = "ETB"

	// CreditCurrencyEUR captures enum value "EUR"
	CreditCurrencyEUR CreditCurrencyEnum = "EUR"

	// CreditCurrencyFJD captures enum value "FJD"
	CreditCurrencyFJD CreditCurrencyEnum = "FJD"

	// CreditCurrencyFKP captures enum value "FKP"
	CreditCurrencyFKP CreditCurrencyEnum = "FKP"

	// CreditCurrencyGBP captures enum value "GBP"
	CreditCurrencyGBP CreditCurrencyEnum = "GBP"

	// CreditCurrencyGEL captures enum value "GEL"
	CreditCurrencyGEL CreditCurrencyEnum = "GEL"

	// CreditCurrencyGGP captures enum value "GGP"
	CreditCurrencyGGP CreditCurrencyEnum = "GGP"

	// CreditCurrencyGHS captures enum value "GHS"
	CreditCurrencyGHS CreditCurrencyEnum = "GHS"

	// CreditCurrencyGIP captures enum value "GIP"
	CreditCurrencyGIP CreditCurrencyEnum = "GIP"

	// CreditCurrencyGMD captures enum value "GMD"
	CreditCurrencyGMD CreditCurrencyEnum = "GMD"

	// CreditCurrencyGNF captures enum value "GNF"
	CreditCurrencyGNF CreditCurrencyEnum = "GNF"

	// CreditCurrencyGTQ captures enum value "GTQ"
	CreditCurrencyGTQ CreditCurrencyEnum = "GTQ"

	// CreditCurrencyGYD captures enum value "GYD"
	CreditCurrencyGYD CreditCurrencyEnum = "GYD"

	// CreditCurrencyHKD captures enum value "HKD"
	CreditCurrencyHKD CreditCurrencyEnum = "HKD"

	// CreditCurrencyHNL captures enum value "HNL"
	CreditCurrencyHNL CreditCurrencyEnum = "HNL"

	// CreditCurrencyHRK captures enum value "HRK"
	CreditCurrencyHRK CreditCurrencyEnum = "HRK"

	// CreditCurrencyHTG captures enum value "HTG"
	CreditCurrencyHTG CreditCurrencyEnum = "HTG"

	// CreditCurrencyHUF captures enum value "HUF"
	CreditCurrencyHUF CreditCurrencyEnum = "HUF"

	// CreditCurrencyIDR captures enum value "IDR"
	CreditCurrencyIDR CreditCurrencyEnum = "IDR"

	// CreditCurrencyILS captures enum value "ILS"
	CreditCurrencyILS CreditCurrencyEnum = "ILS"

	// CreditCurrencyIMP captures enum value "IMP"
	CreditCurrencyIMP CreditCurrencyEnum = "IMP"

	// CreditCurrencyINR captures enum value "INR"
	CreditCurrencyINR CreditCurrencyEnum = "INR"

	// CreditCurrencyIQD captures enum value "IQD"
	CreditCurrencyIQD CreditCurrencyEnum = "IQD"

	// CreditCurrencyIRR captures enum value "IRR"
	CreditCurrencyIRR CreditCurrencyEnum = "IRR"

	// CreditCurrencyISK captures enum value "ISK"
	CreditCurrencyISK CreditCurrencyEnum = "ISK"

	// CreditCurrencyJEP captures enum value "JEP"
	CreditCurrencyJEP CreditCurrencyEnum = "JEP"

	// CreditCurrencyJMD captures enum value "JMD"
	CreditCurrencyJMD CreditCurrencyEnum = "JMD"

	// CreditCurrencyJOD captures enum value "JOD"
	CreditCurrencyJOD CreditCurrencyEnum = "JOD"

	// CreditCurrencyJPY captures enum value "JPY"
	CreditCurrencyJPY CreditCurrencyEnum = "JPY"

	// CreditCurrencyKES captures enum value "KES"
	CreditCurrencyKES CreditCurrencyEnum = "KES"

	// CreditCurrencyKGS captures enum value "KGS"
	CreditCurrencyKGS CreditCurrencyEnum = "KGS"

	// CreditCurrencyKHR captures enum value "KHR"
	CreditCurrencyKHR CreditCurrencyEnum = "KHR"

	// CreditCurrencyKMF captures enum value "KMF"
	CreditCurrencyKMF CreditCurrencyEnum = "KMF"

	// CreditCurrencyKPW captures enum value "KPW"
	CreditCurrencyKPW CreditCurrencyEnum = "KPW"

	// CreditCurrencyKRW captures enum value "KRW"
	CreditCurrencyKRW CreditCurrencyEnum = "KRW"

	// CreditCurrencyKWD captures enum value "KWD"
	CreditCurrencyKWD CreditCurrencyEnum = "KWD"

	// CreditCurrencyKYD captures enum value "KYD"
	CreditCurrencyKYD CreditCurrencyEnum = "KYD"

	// CreditCurrencyKZT captures enum value "KZT"
	CreditCurrencyKZT CreditCurrencyEnum = "KZT"

	// CreditCurrencyLAK captures enum value "LAK"
	CreditCurrencyLAK CreditCurrencyEnum = "LAK"

	// CreditCurrencyLBP captures enum value "LBP"
	CreditCurrencyLBP CreditCurrencyEnum = "LBP"

	// CreditCurrencyLKR captures enum value "LKR"
	CreditCurrencyLKR CreditCurrencyEnum = "LKR"

	// CreditCurrencyLRD captures enum value "LRD"
	CreditCurrencyLRD CreditCurrencyEnum = "LRD"

	// CreditCurrencyLSL captures enum value "LSL"
	CreditCurrencyLSL CreditCurrencyEnum = "LSL"

	// CreditCurrencyLTL captures enum value "LTL"
	CreditCurrencyLTL CreditCurrencyEnum = "LTL"

	// CreditCurrencyLVL captures enum value "LVL"
	CreditCurrencyLVL CreditCurrencyEnum = "LVL"

	// CreditCurrencyLYD captures enum value "LYD"
	CreditCurrencyLYD CreditCurrencyEnum = "LYD"

	// CreditCurrencyMAD captures enum value "MAD"
	CreditCurrencyMAD CreditCurrencyEnum = "MAD"

	// CreditCurrencyMDL captures enum value "MDL"
	CreditCurrencyMDL CreditCurrencyEnum = "MDL"

	// CreditCurrencyMGA captures enum value "MGA"
	CreditCurrencyMGA CreditCurrencyEnum = "MGA"

	// CreditCurrencyMKD captures enum value "MKD"
	CreditCurrencyMKD CreditCurrencyEnum = "MKD"

	// CreditCurrencyMMK captures enum value "MMK"
	CreditCurrencyMMK CreditCurrencyEnum = "MMK"

	// CreditCurrencyMNT captures enum value "MNT"
	CreditCurrencyMNT CreditCurrencyEnum = "MNT"

	// CreditCurrencyMOP captures enum value "MOP"
	CreditCurrencyMOP CreditCurrencyEnum = "MOP"

	// CreditCurrencyMRO captures enum value "MRO"
	CreditCurrencyMRO CreditCurrencyEnum = "MRO"

	// CreditCurrencyMUR captures enum value "MUR"
	CreditCurrencyMUR CreditCurrencyEnum = "MUR"

	// CreditCurrencyMVR captures enum value "MVR"
	CreditCurrencyMVR CreditCurrencyEnum = "MVR"

	// CreditCurrencyMWK captures enum value "MWK"
	CreditCurrencyMWK CreditCurrencyEnum = "MWK"

	// CreditCurrencyMXN captures enum value "MXN"
	CreditCurrencyMXN CreditCurrencyEnum = "MXN"

	// CreditCurrencyMYR captures enum value "MYR"
	CreditCurrencyMYR CreditCurrencyEnum = "MYR"

	// CreditCurrencyMZN captures enum value "MZN"
	CreditCurrencyMZN CreditCurrencyEnum = "MZN"

	// CreditCurrencyNAD captures enum value "NAD"
	CreditCurrencyNAD CreditCurrencyEnum = "NAD"

	// CreditCurrencyNGN captures enum value "NGN"
	CreditCurrencyNGN CreditCurrencyEnum = "NGN"

	// CreditCurrencyNIO captures enum value "NIO"
	CreditCurrencyNIO CreditCurrencyEnum = "NIO"

	// CreditCurrencyNOK captures enum value "NOK"
	CreditCurrencyNOK CreditCurrencyEnum = "NOK"

	// CreditCurrencyNPR captures enum value "NPR"
	CreditCurrencyNPR CreditCurrencyEnum = "NPR"

	// CreditCurrencyNZD captures enum value "NZD"
	CreditCurrencyNZD CreditCurrencyEnum = "NZD"

	// CreditCurrencyOMR captures enum value "OMR"
	CreditCurrencyOMR CreditCurrencyEnum = "OMR"

	// CreditCurrencyPAB captures enum value "PAB"
	CreditCurrencyPAB CreditCurrencyEnum = "PAB"

	// CreditCurrencyPEN captures enum value "PEN"
	CreditCurrencyPEN CreditCurrencyEnum = "PEN"

	// CreditCurrencyPGK captures enum value "PGK"
	CreditCurrencyPGK CreditCurrencyEnum = "PGK"

	// CreditCurrencyPHP captures enum value "PHP"
	CreditCurrencyPHP CreditCurrencyEnum = "PHP"

	// CreditCurrencyPKR captures enum value "PKR"
	CreditCurrencyPKR CreditCurrencyEnum = "PKR"

	// CreditCurrencyPLN captures enum value "PLN"
	CreditCurrencyPLN CreditCurrencyEnum = "PLN"

	// CreditCurrencyPYG captures enum value "PYG"
	CreditCurrencyPYG CreditCurrencyEnum = "PYG"

	// CreditCurrencyQAR captures enum value "QAR"
	CreditCurrencyQAR CreditCurrencyEnum = "QAR"

	// CreditCurrencyRON captures enum value "RON"
	CreditCurrencyRON CreditCurrencyEnum = "RON"

	// CreditCurrencyRSD captures enum value "RSD"
	CreditCurrencyRSD CreditCurrencyEnum = "RSD"

	// CreditCurrencyRUB captures enum value "RUB"
	CreditCurrencyRUB CreditCurrencyEnum = "RUB"

	// CreditCurrencyRWF captures enum value "RWF"
	CreditCurrencyRWF CreditCurrencyEnum = "RWF"

	// CreditCurrencySAR captures enum value "SAR"
	CreditCurrencySAR CreditCurrencyEnum = "SAR"

	// CreditCurrencySBD captures enum value "SBD"
	CreditCurrencySBD CreditCurrencyEnum = "SBD"

	// CreditCurrencySCR captures enum value "SCR"
	CreditCurrencySCR CreditCurrencyEnum = "SCR"

	// CreditCurrencySDG captures enum value "SDG"
	CreditCurrencySDG CreditCurrencyEnum = "SDG"

	// CreditCurrencySEK captures enum value "SEK"
	CreditCurrencySEK CreditCurrencyEnum = "SEK"

	// CreditCurrencySGD captures enum value "SGD"
	CreditCurrencySGD CreditCurrencyEnum = "SGD"

	// CreditCurrencySHP captures enum value "SHP"
	CreditCurrencySHP CreditCurrencyEnum = "SHP"

	// CreditCurrencySLL captures enum value "SLL"
	CreditCurrencySLL CreditCurrencyEnum = "SLL"

	// CreditCurrencySOS captures enum value "SOS"
	CreditCurrencySOS CreditCurrencyEnum = "SOS"

	// CreditCurrencySPL captures enum value "SPL"
	CreditCurrencySPL CreditCurrencyEnum = "SPL"

	// CreditCurrencySRD captures enum value "SRD"
	CreditCurrencySRD CreditCurrencyEnum = "SRD"

	// CreditCurrencySTD captures enum value "STD"
	CreditCurrencySTD CreditCurrencyEnum = "STD"

	// CreditCurrencySVC captures enum value "SVC"
	CreditCurrencySVC CreditCurrencyEnum = "SVC"

	// CreditCurrencySYP captures enum value "SYP"
	CreditCurrencySYP CreditCurrencyEnum = "SYP"

	// CreditCurrencySZL captures enum value "SZL"
	CreditCurrencySZL CreditCurrencyEnum = "SZL"

	// CreditCurrencyTHB captures enum value "THB"
	CreditCurrencyTHB CreditCurrencyEnum = "THB"

	// CreditCurrencyTJS captures enum value "TJS"
	CreditCurrencyTJS CreditCurrencyEnum = "TJS"

	// CreditCurrencyTMT captures enum value "TMT"
	CreditCurrencyTMT CreditCurrencyEnum = "TMT"

	// CreditCurrencyTND captures enum value "TND"
	CreditCurrencyTND CreditCurrencyEnum = "TND"

	// CreditCurrencyTOP captures enum value "TOP"
	CreditCurrencyTOP CreditCurrencyEnum = "TOP"

	// CreditCurrencyTRY captures enum value "TRY"
	CreditCurrencyTRY CreditCurrencyEnum = "TRY"

	// CreditCurrencyTTD captures enum value "TTD"
	CreditCurrencyTTD CreditCurrencyEnum = "TTD"

	// CreditCurrencyTVD captures enum value "TVD"
	CreditCurrencyTVD CreditCurrencyEnum = "TVD"

	// CreditCurrencyTWD captures enum value "TWD"
	CreditCurrencyTWD CreditCurrencyEnum = "TWD"

	// CreditCurrencyTZS captures enum value "TZS"
	CreditCurrencyTZS CreditCurrencyEnum = "TZS"

	// CreditCurrencyUAH captures enum value "UAH"
	CreditCurrencyUAH CreditCurrencyEnum = "UAH"

	// CreditCurrencyUGX captures enum value "UGX"
	CreditCurrencyUGX CreditCurrencyEnum = "UGX"

	// CreditCurrencyUSD captures enum value "USD"
	CreditCurrencyUSD CreditCurrencyEnum = "USD"

	// CreditCurrencyUYU captures enum value "UYU"
	CreditCurrencyUYU CreditCurrencyEnum = "UYU"

	// CreditCurrencyUZS captures enum value "UZS"
	CreditCurrencyUZS CreditCurrencyEnum = "UZS"

	// CreditCurrencyVEF captures enum value "VEF"
	CreditCurrencyVEF CreditCurrencyEnum = "VEF"

	// CreditCurrencyVND captures enum value "VND"
	CreditCurrencyVND CreditCurrencyEnum = "VND"

	// CreditCurrencyVUV captures enum value "VUV"
	CreditCurrencyVUV CreditCurrencyEnum = "VUV"

	// CreditCurrencyWST captures enum value "WST"
	CreditCurrencyWST CreditCurrencyEnum = "WST"

	// CreditCurrencyXAF captures enum value "XAF"
	CreditCurrencyXAF CreditCurrencyEnum = "XAF"

	// CreditCurrencyXCD captures enum value "XCD"
	CreditCurrencyXCD CreditCurrencyEnum = "XCD"

	// CreditCurrencyXDR captures enum value "XDR"
	CreditCurrencyXDR CreditCurrencyEnum = "XDR"

	// CreditCurrencyXOF captures enum value "XOF"
	CreditCurrencyXOF CreditCurrencyEnum = "XOF"

	// CreditCurrencyXPF captures enum value "XPF"
	CreditCurrencyXPF CreditCurrencyEnum = "XPF"

	// CreditCurrencyYER captures enum value "YER"
	CreditCurrencyYER CreditCurrencyEnum = "YER"

	// CreditCurrencyZAR captures enum value "ZAR"
	CreditCurrencyZAR CreditCurrencyEnum = "ZAR"

	// CreditCurrencyZMW captures enum value "ZMW"
	CreditCurrencyZMW CreditCurrencyEnum = "ZMW"

	// CreditCurrencyZWD captures enum value "ZWD"
	CreditCurrencyZWD CreditCurrencyEnum = "ZWD"

	// CreditCurrencyBTC captures enum value "BTC"
	CreditCurrencyBTC CreditCurrencyEnum = "BTC"
)

var CreditCurrencyEnumValues = []string{
	"AED",
	"AFN",
	"ALL",
	"AMD",
	"ANG",
	"AOA",
	"ARS",
	"AUD",
	"AWG",
	"AZN",
	"BAM",
	"BBD",
	"BDT",
	"BGN",
	"BHD",
	"BIF",
	"BMD",
	"BND",
	"BOB",
	"BRL",
	"BSD",
	"BTN",
	"BWP",
	"BYR",
	"BZD",
	"CAD",
	"CDF",
	"CHF",
	"CLP",
	"CNY",
	"COP",
	"CRC",
	"CUC",
	"CUP",
	"CVE",
	"CZK",
	"DJF",
	"DKK",
	"DOP",
	"DZD",
	"EGP",
	"ERN",
	"ETB",
	"EUR",
	"FJD",
	"FKP",
	"GBP",
	"GEL",
	"GGP",
	"GHS",
	"GIP",
	"GMD",
	"GNF",
	"GTQ",
	"GYD",
	"HKD",
	"HNL",
	"HRK",
	"HTG",
	"HUF",
	"IDR",
	"ILS",
	"IMP",
	"INR",
	"IQD",
	"IRR",
	"ISK",
	"JEP",
	"JMD",
	"JOD",
	"JPY",
	"KES",
	"KGS",
	"KHR",
	"KMF",
	"KPW",
	"KRW",
	"KWD",
	"KYD",
	"KZT",
	"LAK",
	"LBP",
	"LKR",
	"LRD",
	"LSL",
	"LTL",
	"LVL",
	"LYD",
	"MAD",
	"MDL",
	"MGA",
	"MKD",
	"MMK",
	"MNT",
	"MOP",
	"MRO",
	"MUR",
	"MVR",
	"MWK",
	"MXN",
	"MYR",
	"MZN",
	"NAD",
	"NGN",
	"NIO",
	"NOK",
	"NPR",
	"NZD",
	"OMR",
	"PAB",
	"PEN",
	"PGK",
	"PHP",
	"PKR",
	"PLN",
	"PYG",
	"QAR",
	"RON",
	"RSD",
	"RUB",
	"RWF",
	"SAR",
	"SBD",
	"SCR",
	"SDG",
	"SEK",
	"SGD",
	"SHP",
	"SLL",
	"SOS",
	"SPL",
	"SRD",
	"STD",
	"SVC",
	"SYP",
	"SZL",
	"THB",
	"TJS",
	"TMT",
	"TND",
	"TOP",
	"TRY",
	"TTD",
	"TVD",
	"TWD",
	"TZS",
	"UAH",
	"UGX",
	"USD",
	"UYU",
	"UZS",
	"VEF",
	"VND",
	"VUV",
	"WST",
	"XAF",
	"XCD",
	"XDR",
	"XOF",
	"XPF",
	"YER",
	"ZAR",
	"ZMW",
	"ZWD",
	"BTC",
}

func (e CreditCurrencyEnum) IsValid() bool {
	for _, v := range CreditCurrencyEnumValues {
		if v == string(e) {
			return true
		}
	}
	return false
}

// prop value enum
func (m *Credit) validateCurrencyEnum(path, location string, value CreditCurrencyEnum) error {
	if err := validate.Enum(path, location, value, creditTypeCurrencyPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Credit) validateCurrency(formats strfmt.Registry) error {

	if swag.IsZero(m.Currency) { // not required
		return nil
	}

	// value enum
	if err := m.validateCurrencyEnum("currency", "body", m.Currency); err != nil {
		return err
	}

	return nil
}

func (m *Credit) validateEffectiveDate(formats strfmt.Registry) error {

	if swag.IsZero(m.EffectiveDate) { // not required
		return nil
	}

	if err := validate.FormatOf("effectiveDate", "body", "date", m.EffectiveDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Credit) validateInvoiceID(formats strfmt.Registry) error {

	if swag.IsZero(m.InvoiceID) { // not required
		return nil
	}

	if err := validate.FormatOf("invoiceId", "body", "uuid", m.InvoiceID.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Credit) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Credit) UnmarshalBinary(b []byte) error {
	var res Credit
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}