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

// Catalog catalog
// swagger:model Catalog
type Catalog struct {

	// currencies
	Currencies []CatalogCurrenciesEnum `json:"currencies"`

	// effective date
	// Format: date-time
	EffectiveDate strfmt.DateTime `json:"effectiveDate,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// price lists
	PriceLists []*PriceList `json:"priceLists"`

	// products
	Products []*Product `json:"products"`

	// units
	Units []*Unit `json:"units"`
}

// Validate validates this catalog
func (m *Catalog) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCurrencies(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEffectiveDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePriceLists(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProducts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUnits(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

type CatalogCurrenciesEnum string

const (

	// CatalogCurrenciesAED captures enum value "AED"
	CatalogCurrenciesAED CatalogCurrenciesEnum = "AED"

	// CatalogCurrenciesAFN captures enum value "AFN"
	CatalogCurrenciesAFN CatalogCurrenciesEnum = "AFN"

	// CatalogCurrenciesALL captures enum value "ALL"
	CatalogCurrenciesALL CatalogCurrenciesEnum = "ALL"

	// CatalogCurrenciesAMD captures enum value "AMD"
	CatalogCurrenciesAMD CatalogCurrenciesEnum = "AMD"

	// CatalogCurrenciesANG captures enum value "ANG"
	CatalogCurrenciesANG CatalogCurrenciesEnum = "ANG"

	// CatalogCurrenciesAOA captures enum value "AOA"
	CatalogCurrenciesAOA CatalogCurrenciesEnum = "AOA"

	// CatalogCurrenciesARS captures enum value "ARS"
	CatalogCurrenciesARS CatalogCurrenciesEnum = "ARS"

	// CatalogCurrenciesAUD captures enum value "AUD"
	CatalogCurrenciesAUD CatalogCurrenciesEnum = "AUD"

	// CatalogCurrenciesAWG captures enum value "AWG"
	CatalogCurrenciesAWG CatalogCurrenciesEnum = "AWG"

	// CatalogCurrenciesAZN captures enum value "AZN"
	CatalogCurrenciesAZN CatalogCurrenciesEnum = "AZN"

	// CatalogCurrenciesBAM captures enum value "BAM"
	CatalogCurrenciesBAM CatalogCurrenciesEnum = "BAM"

	// CatalogCurrenciesBBD captures enum value "BBD"
	CatalogCurrenciesBBD CatalogCurrenciesEnum = "BBD"

	// CatalogCurrenciesBDT captures enum value "BDT"
	CatalogCurrenciesBDT CatalogCurrenciesEnum = "BDT"

	// CatalogCurrenciesBGN captures enum value "BGN"
	CatalogCurrenciesBGN CatalogCurrenciesEnum = "BGN"

	// CatalogCurrenciesBHD captures enum value "BHD"
	CatalogCurrenciesBHD CatalogCurrenciesEnum = "BHD"

	// CatalogCurrenciesBIF captures enum value "BIF"
	CatalogCurrenciesBIF CatalogCurrenciesEnum = "BIF"

	// CatalogCurrenciesBMD captures enum value "BMD"
	CatalogCurrenciesBMD CatalogCurrenciesEnum = "BMD"

	// CatalogCurrenciesBND captures enum value "BND"
	CatalogCurrenciesBND CatalogCurrenciesEnum = "BND"

	// CatalogCurrenciesBOB captures enum value "BOB"
	CatalogCurrenciesBOB CatalogCurrenciesEnum = "BOB"

	// CatalogCurrenciesBRL captures enum value "BRL"
	CatalogCurrenciesBRL CatalogCurrenciesEnum = "BRL"

	// CatalogCurrenciesBSD captures enum value "BSD"
	CatalogCurrenciesBSD CatalogCurrenciesEnum = "BSD"

	// CatalogCurrenciesBTN captures enum value "BTN"
	CatalogCurrenciesBTN CatalogCurrenciesEnum = "BTN"

	// CatalogCurrenciesBWP captures enum value "BWP"
	CatalogCurrenciesBWP CatalogCurrenciesEnum = "BWP"

	// CatalogCurrenciesBYR captures enum value "BYR"
	CatalogCurrenciesBYR CatalogCurrenciesEnum = "BYR"

	// CatalogCurrenciesBZD captures enum value "BZD"
	CatalogCurrenciesBZD CatalogCurrenciesEnum = "BZD"

	// CatalogCurrenciesCAD captures enum value "CAD"
	CatalogCurrenciesCAD CatalogCurrenciesEnum = "CAD"

	// CatalogCurrenciesCDF captures enum value "CDF"
	CatalogCurrenciesCDF CatalogCurrenciesEnum = "CDF"

	// CatalogCurrenciesCHF captures enum value "CHF"
	CatalogCurrenciesCHF CatalogCurrenciesEnum = "CHF"

	// CatalogCurrenciesCLP captures enum value "CLP"
	CatalogCurrenciesCLP CatalogCurrenciesEnum = "CLP"

	// CatalogCurrenciesCNY captures enum value "CNY"
	CatalogCurrenciesCNY CatalogCurrenciesEnum = "CNY"

	// CatalogCurrenciesCOP captures enum value "COP"
	CatalogCurrenciesCOP CatalogCurrenciesEnum = "COP"

	// CatalogCurrenciesCRC captures enum value "CRC"
	CatalogCurrenciesCRC CatalogCurrenciesEnum = "CRC"

	// CatalogCurrenciesCUC captures enum value "CUC"
	CatalogCurrenciesCUC CatalogCurrenciesEnum = "CUC"

	// CatalogCurrenciesCUP captures enum value "CUP"
	CatalogCurrenciesCUP CatalogCurrenciesEnum = "CUP"

	// CatalogCurrenciesCVE captures enum value "CVE"
	CatalogCurrenciesCVE CatalogCurrenciesEnum = "CVE"

	// CatalogCurrenciesCZK captures enum value "CZK"
	CatalogCurrenciesCZK CatalogCurrenciesEnum = "CZK"

	// CatalogCurrenciesDJF captures enum value "DJF"
	CatalogCurrenciesDJF CatalogCurrenciesEnum = "DJF"

	// CatalogCurrenciesDKK captures enum value "DKK"
	CatalogCurrenciesDKK CatalogCurrenciesEnum = "DKK"

	// CatalogCurrenciesDOP captures enum value "DOP"
	CatalogCurrenciesDOP CatalogCurrenciesEnum = "DOP"

	// CatalogCurrenciesDZD captures enum value "DZD"
	CatalogCurrenciesDZD CatalogCurrenciesEnum = "DZD"

	// CatalogCurrenciesEGP captures enum value "EGP"
	CatalogCurrenciesEGP CatalogCurrenciesEnum = "EGP"

	// CatalogCurrenciesERN captures enum value "ERN"
	CatalogCurrenciesERN CatalogCurrenciesEnum = "ERN"

	// CatalogCurrenciesETB captures enum value "ETB"
	CatalogCurrenciesETB CatalogCurrenciesEnum = "ETB"

	// CatalogCurrenciesEUR captures enum value "EUR"
	CatalogCurrenciesEUR CatalogCurrenciesEnum = "EUR"

	// CatalogCurrenciesFJD captures enum value "FJD"
	CatalogCurrenciesFJD CatalogCurrenciesEnum = "FJD"

	// CatalogCurrenciesFKP captures enum value "FKP"
	CatalogCurrenciesFKP CatalogCurrenciesEnum = "FKP"

	// CatalogCurrenciesGBP captures enum value "GBP"
	CatalogCurrenciesGBP CatalogCurrenciesEnum = "GBP"

	// CatalogCurrenciesGEL captures enum value "GEL"
	CatalogCurrenciesGEL CatalogCurrenciesEnum = "GEL"

	// CatalogCurrenciesGGP captures enum value "GGP"
	CatalogCurrenciesGGP CatalogCurrenciesEnum = "GGP"

	// CatalogCurrenciesGHS captures enum value "GHS"
	CatalogCurrenciesGHS CatalogCurrenciesEnum = "GHS"

	// CatalogCurrenciesGIP captures enum value "GIP"
	CatalogCurrenciesGIP CatalogCurrenciesEnum = "GIP"

	// CatalogCurrenciesGMD captures enum value "GMD"
	CatalogCurrenciesGMD CatalogCurrenciesEnum = "GMD"

	// CatalogCurrenciesGNF captures enum value "GNF"
	CatalogCurrenciesGNF CatalogCurrenciesEnum = "GNF"

	// CatalogCurrenciesGTQ captures enum value "GTQ"
	CatalogCurrenciesGTQ CatalogCurrenciesEnum = "GTQ"

	// CatalogCurrenciesGYD captures enum value "GYD"
	CatalogCurrenciesGYD CatalogCurrenciesEnum = "GYD"

	// CatalogCurrenciesHKD captures enum value "HKD"
	CatalogCurrenciesHKD CatalogCurrenciesEnum = "HKD"

	// CatalogCurrenciesHNL captures enum value "HNL"
	CatalogCurrenciesHNL CatalogCurrenciesEnum = "HNL"

	// CatalogCurrenciesHRK captures enum value "HRK"
	CatalogCurrenciesHRK CatalogCurrenciesEnum = "HRK"

	// CatalogCurrenciesHTG captures enum value "HTG"
	CatalogCurrenciesHTG CatalogCurrenciesEnum = "HTG"

	// CatalogCurrenciesHUF captures enum value "HUF"
	CatalogCurrenciesHUF CatalogCurrenciesEnum = "HUF"

	// CatalogCurrenciesIDR captures enum value "IDR"
	CatalogCurrenciesIDR CatalogCurrenciesEnum = "IDR"

	// CatalogCurrenciesILS captures enum value "ILS"
	CatalogCurrenciesILS CatalogCurrenciesEnum = "ILS"

	// CatalogCurrenciesIMP captures enum value "IMP"
	CatalogCurrenciesIMP CatalogCurrenciesEnum = "IMP"

	// CatalogCurrenciesINR captures enum value "INR"
	CatalogCurrenciesINR CatalogCurrenciesEnum = "INR"

	// CatalogCurrenciesIQD captures enum value "IQD"
	CatalogCurrenciesIQD CatalogCurrenciesEnum = "IQD"

	// CatalogCurrenciesIRR captures enum value "IRR"
	CatalogCurrenciesIRR CatalogCurrenciesEnum = "IRR"

	// CatalogCurrenciesISK captures enum value "ISK"
	CatalogCurrenciesISK CatalogCurrenciesEnum = "ISK"

	// CatalogCurrenciesJEP captures enum value "JEP"
	CatalogCurrenciesJEP CatalogCurrenciesEnum = "JEP"

	// CatalogCurrenciesJMD captures enum value "JMD"
	CatalogCurrenciesJMD CatalogCurrenciesEnum = "JMD"

	// CatalogCurrenciesJOD captures enum value "JOD"
	CatalogCurrenciesJOD CatalogCurrenciesEnum = "JOD"

	// CatalogCurrenciesJPY captures enum value "JPY"
	CatalogCurrenciesJPY CatalogCurrenciesEnum = "JPY"

	// CatalogCurrenciesKES captures enum value "KES"
	CatalogCurrenciesKES CatalogCurrenciesEnum = "KES"

	// CatalogCurrenciesKGS captures enum value "KGS"
	CatalogCurrenciesKGS CatalogCurrenciesEnum = "KGS"

	// CatalogCurrenciesKHR captures enum value "KHR"
	CatalogCurrenciesKHR CatalogCurrenciesEnum = "KHR"

	// CatalogCurrenciesKMF captures enum value "KMF"
	CatalogCurrenciesKMF CatalogCurrenciesEnum = "KMF"

	// CatalogCurrenciesKPW captures enum value "KPW"
	CatalogCurrenciesKPW CatalogCurrenciesEnum = "KPW"

	// CatalogCurrenciesKRW captures enum value "KRW"
	CatalogCurrenciesKRW CatalogCurrenciesEnum = "KRW"

	// CatalogCurrenciesKWD captures enum value "KWD"
	CatalogCurrenciesKWD CatalogCurrenciesEnum = "KWD"

	// CatalogCurrenciesKYD captures enum value "KYD"
	CatalogCurrenciesKYD CatalogCurrenciesEnum = "KYD"

	// CatalogCurrenciesKZT captures enum value "KZT"
	CatalogCurrenciesKZT CatalogCurrenciesEnum = "KZT"

	// CatalogCurrenciesLAK captures enum value "LAK"
	CatalogCurrenciesLAK CatalogCurrenciesEnum = "LAK"

	// CatalogCurrenciesLBP captures enum value "LBP"
	CatalogCurrenciesLBP CatalogCurrenciesEnum = "LBP"

	// CatalogCurrenciesLKR captures enum value "LKR"
	CatalogCurrenciesLKR CatalogCurrenciesEnum = "LKR"

	// CatalogCurrenciesLRD captures enum value "LRD"
	CatalogCurrenciesLRD CatalogCurrenciesEnum = "LRD"

	// CatalogCurrenciesLSL captures enum value "LSL"
	CatalogCurrenciesLSL CatalogCurrenciesEnum = "LSL"

	// CatalogCurrenciesLTL captures enum value "LTL"
	CatalogCurrenciesLTL CatalogCurrenciesEnum = "LTL"

	// CatalogCurrenciesLVL captures enum value "LVL"
	CatalogCurrenciesLVL CatalogCurrenciesEnum = "LVL"

	// CatalogCurrenciesLYD captures enum value "LYD"
	CatalogCurrenciesLYD CatalogCurrenciesEnum = "LYD"

	// CatalogCurrenciesMAD captures enum value "MAD"
	CatalogCurrenciesMAD CatalogCurrenciesEnum = "MAD"

	// CatalogCurrenciesMDL captures enum value "MDL"
	CatalogCurrenciesMDL CatalogCurrenciesEnum = "MDL"

	// CatalogCurrenciesMGA captures enum value "MGA"
	CatalogCurrenciesMGA CatalogCurrenciesEnum = "MGA"

	// CatalogCurrenciesMKD captures enum value "MKD"
	CatalogCurrenciesMKD CatalogCurrenciesEnum = "MKD"

	// CatalogCurrenciesMMK captures enum value "MMK"
	CatalogCurrenciesMMK CatalogCurrenciesEnum = "MMK"

	// CatalogCurrenciesMNT captures enum value "MNT"
	CatalogCurrenciesMNT CatalogCurrenciesEnum = "MNT"

	// CatalogCurrenciesMOP captures enum value "MOP"
	CatalogCurrenciesMOP CatalogCurrenciesEnum = "MOP"

	// CatalogCurrenciesMRO captures enum value "MRO"
	CatalogCurrenciesMRO CatalogCurrenciesEnum = "MRO"

	// CatalogCurrenciesMUR captures enum value "MUR"
	CatalogCurrenciesMUR CatalogCurrenciesEnum = "MUR"

	// CatalogCurrenciesMVR captures enum value "MVR"
	CatalogCurrenciesMVR CatalogCurrenciesEnum = "MVR"

	// CatalogCurrenciesMWK captures enum value "MWK"
	CatalogCurrenciesMWK CatalogCurrenciesEnum = "MWK"

	// CatalogCurrenciesMXN captures enum value "MXN"
	CatalogCurrenciesMXN CatalogCurrenciesEnum = "MXN"

	// CatalogCurrenciesMYR captures enum value "MYR"
	CatalogCurrenciesMYR CatalogCurrenciesEnum = "MYR"

	// CatalogCurrenciesMZN captures enum value "MZN"
	CatalogCurrenciesMZN CatalogCurrenciesEnum = "MZN"

	// CatalogCurrenciesNAD captures enum value "NAD"
	CatalogCurrenciesNAD CatalogCurrenciesEnum = "NAD"

	// CatalogCurrenciesNGN captures enum value "NGN"
	CatalogCurrenciesNGN CatalogCurrenciesEnum = "NGN"

	// CatalogCurrenciesNIO captures enum value "NIO"
	CatalogCurrenciesNIO CatalogCurrenciesEnum = "NIO"

	// CatalogCurrenciesNOK captures enum value "NOK"
	CatalogCurrenciesNOK CatalogCurrenciesEnum = "NOK"

	// CatalogCurrenciesNPR captures enum value "NPR"
	CatalogCurrenciesNPR CatalogCurrenciesEnum = "NPR"

	// CatalogCurrenciesNZD captures enum value "NZD"
	CatalogCurrenciesNZD CatalogCurrenciesEnum = "NZD"

	// CatalogCurrenciesOMR captures enum value "OMR"
	CatalogCurrenciesOMR CatalogCurrenciesEnum = "OMR"

	// CatalogCurrenciesPAB captures enum value "PAB"
	CatalogCurrenciesPAB CatalogCurrenciesEnum = "PAB"

	// CatalogCurrenciesPEN captures enum value "PEN"
	CatalogCurrenciesPEN CatalogCurrenciesEnum = "PEN"

	// CatalogCurrenciesPGK captures enum value "PGK"
	CatalogCurrenciesPGK CatalogCurrenciesEnum = "PGK"

	// CatalogCurrenciesPHP captures enum value "PHP"
	CatalogCurrenciesPHP CatalogCurrenciesEnum = "PHP"

	// CatalogCurrenciesPKR captures enum value "PKR"
	CatalogCurrenciesPKR CatalogCurrenciesEnum = "PKR"

	// CatalogCurrenciesPLN captures enum value "PLN"
	CatalogCurrenciesPLN CatalogCurrenciesEnum = "PLN"

	// CatalogCurrenciesPYG captures enum value "PYG"
	CatalogCurrenciesPYG CatalogCurrenciesEnum = "PYG"

	// CatalogCurrenciesQAR captures enum value "QAR"
	CatalogCurrenciesQAR CatalogCurrenciesEnum = "QAR"

	// CatalogCurrenciesRON captures enum value "RON"
	CatalogCurrenciesRON CatalogCurrenciesEnum = "RON"

	// CatalogCurrenciesRSD captures enum value "RSD"
	CatalogCurrenciesRSD CatalogCurrenciesEnum = "RSD"

	// CatalogCurrenciesRUB captures enum value "RUB"
	CatalogCurrenciesRUB CatalogCurrenciesEnum = "RUB"

	// CatalogCurrenciesRWF captures enum value "RWF"
	CatalogCurrenciesRWF CatalogCurrenciesEnum = "RWF"

	// CatalogCurrenciesSAR captures enum value "SAR"
	CatalogCurrenciesSAR CatalogCurrenciesEnum = "SAR"

	// CatalogCurrenciesSBD captures enum value "SBD"
	CatalogCurrenciesSBD CatalogCurrenciesEnum = "SBD"

	// CatalogCurrenciesSCR captures enum value "SCR"
	CatalogCurrenciesSCR CatalogCurrenciesEnum = "SCR"

	// CatalogCurrenciesSDG captures enum value "SDG"
	CatalogCurrenciesSDG CatalogCurrenciesEnum = "SDG"

	// CatalogCurrenciesSEK captures enum value "SEK"
	CatalogCurrenciesSEK CatalogCurrenciesEnum = "SEK"

	// CatalogCurrenciesSGD captures enum value "SGD"
	CatalogCurrenciesSGD CatalogCurrenciesEnum = "SGD"

	// CatalogCurrenciesSHP captures enum value "SHP"
	CatalogCurrenciesSHP CatalogCurrenciesEnum = "SHP"

	// CatalogCurrenciesSLL captures enum value "SLL"
	CatalogCurrenciesSLL CatalogCurrenciesEnum = "SLL"

	// CatalogCurrenciesSOS captures enum value "SOS"
	CatalogCurrenciesSOS CatalogCurrenciesEnum = "SOS"

	// CatalogCurrenciesSPL captures enum value "SPL"
	CatalogCurrenciesSPL CatalogCurrenciesEnum = "SPL"

	// CatalogCurrenciesSRD captures enum value "SRD"
	CatalogCurrenciesSRD CatalogCurrenciesEnum = "SRD"

	// CatalogCurrenciesSTD captures enum value "STD"
	CatalogCurrenciesSTD CatalogCurrenciesEnum = "STD"

	// CatalogCurrenciesSVC captures enum value "SVC"
	CatalogCurrenciesSVC CatalogCurrenciesEnum = "SVC"

	// CatalogCurrenciesSYP captures enum value "SYP"
	CatalogCurrenciesSYP CatalogCurrenciesEnum = "SYP"

	// CatalogCurrenciesSZL captures enum value "SZL"
	CatalogCurrenciesSZL CatalogCurrenciesEnum = "SZL"

	// CatalogCurrenciesTHB captures enum value "THB"
	CatalogCurrenciesTHB CatalogCurrenciesEnum = "THB"

	// CatalogCurrenciesTJS captures enum value "TJS"
	CatalogCurrenciesTJS CatalogCurrenciesEnum = "TJS"

	// CatalogCurrenciesTMT captures enum value "TMT"
	CatalogCurrenciesTMT CatalogCurrenciesEnum = "TMT"

	// CatalogCurrenciesTND captures enum value "TND"
	CatalogCurrenciesTND CatalogCurrenciesEnum = "TND"

	// CatalogCurrenciesTOP captures enum value "TOP"
	CatalogCurrenciesTOP CatalogCurrenciesEnum = "TOP"

	// CatalogCurrenciesTRY captures enum value "TRY"
	CatalogCurrenciesTRY CatalogCurrenciesEnum = "TRY"

	// CatalogCurrenciesTTD captures enum value "TTD"
	CatalogCurrenciesTTD CatalogCurrenciesEnum = "TTD"

	// CatalogCurrenciesTVD captures enum value "TVD"
	CatalogCurrenciesTVD CatalogCurrenciesEnum = "TVD"

	// CatalogCurrenciesTWD captures enum value "TWD"
	CatalogCurrenciesTWD CatalogCurrenciesEnum = "TWD"

	// CatalogCurrenciesTZS captures enum value "TZS"
	CatalogCurrenciesTZS CatalogCurrenciesEnum = "TZS"

	// CatalogCurrenciesUAH captures enum value "UAH"
	CatalogCurrenciesUAH CatalogCurrenciesEnum = "UAH"

	// CatalogCurrenciesUGX captures enum value "UGX"
	CatalogCurrenciesUGX CatalogCurrenciesEnum = "UGX"

	// CatalogCurrenciesUSD captures enum value "USD"
	CatalogCurrenciesUSD CatalogCurrenciesEnum = "USD"

	// CatalogCurrenciesUYU captures enum value "UYU"
	CatalogCurrenciesUYU CatalogCurrenciesEnum = "UYU"

	// CatalogCurrenciesUZS captures enum value "UZS"
	CatalogCurrenciesUZS CatalogCurrenciesEnum = "UZS"

	// CatalogCurrenciesVEF captures enum value "VEF"
	CatalogCurrenciesVEF CatalogCurrenciesEnum = "VEF"

	// CatalogCurrenciesVND captures enum value "VND"
	CatalogCurrenciesVND CatalogCurrenciesEnum = "VND"

	// CatalogCurrenciesVUV captures enum value "VUV"
	CatalogCurrenciesVUV CatalogCurrenciesEnum = "VUV"

	// CatalogCurrenciesWST captures enum value "WST"
	CatalogCurrenciesWST CatalogCurrenciesEnum = "WST"

	// CatalogCurrenciesXAF captures enum value "XAF"
	CatalogCurrenciesXAF CatalogCurrenciesEnum = "XAF"

	// CatalogCurrenciesXCD captures enum value "XCD"
	CatalogCurrenciesXCD CatalogCurrenciesEnum = "XCD"

	// CatalogCurrenciesXDR captures enum value "XDR"
	CatalogCurrenciesXDR CatalogCurrenciesEnum = "XDR"

	// CatalogCurrenciesXOF captures enum value "XOF"
	CatalogCurrenciesXOF CatalogCurrenciesEnum = "XOF"

	// CatalogCurrenciesXPF captures enum value "XPF"
	CatalogCurrenciesXPF CatalogCurrenciesEnum = "XPF"

	// CatalogCurrenciesYER captures enum value "YER"
	CatalogCurrenciesYER CatalogCurrenciesEnum = "YER"

	// CatalogCurrenciesZAR captures enum value "ZAR"
	CatalogCurrenciesZAR CatalogCurrenciesEnum = "ZAR"

	// CatalogCurrenciesZMW captures enum value "ZMW"
	CatalogCurrenciesZMW CatalogCurrenciesEnum = "ZMW"

	// CatalogCurrenciesZWD captures enum value "ZWD"
	CatalogCurrenciesZWD CatalogCurrenciesEnum = "ZWD"

	// CatalogCurrenciesBTC captures enum value "BTC"
	CatalogCurrenciesBTC CatalogCurrenciesEnum = "BTC"
)

var CatalogCurrenciesEnumValues = []string{
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

func (e CatalogCurrenciesEnum) IsValid() bool {
	for _, v := range CatalogCurrenciesEnumValues {
		if v == string(e) {
			return true
		}
	}
	return false
}

var catalogCurrenciesItemsEnum []interface{}

func init() {
	var res []CatalogCurrenciesEnum
	if err := json.Unmarshal([]byte(`["AED","AFN","ALL","AMD","ANG","AOA","ARS","AUD","AWG","AZN","BAM","BBD","BDT","BGN","BHD","BIF","BMD","BND","BOB","BRL","BSD","BTN","BWP","BYR","BZD","CAD","CDF","CHF","CLP","CNY","COP","CRC","CUC","CUP","CVE","CZK","DJF","DKK","DOP","DZD","EGP","ERN","ETB","EUR","FJD","FKP","GBP","GEL","GGP","GHS","GIP","GMD","GNF","GTQ","GYD","HKD","HNL","HRK","HTG","HUF","IDR","ILS","IMP","INR","IQD","IRR","ISK","JEP","JMD","JOD","JPY","KES","KGS","KHR","KMF","KPW","KRW","KWD","KYD","KZT","LAK","LBP","LKR","LRD","LSL","LTL","LVL","LYD","MAD","MDL","MGA","MKD","MMK","MNT","MOP","MRO","MUR","MVR","MWK","MXN","MYR","MZN","NAD","NGN","NIO","NOK","NPR","NZD","OMR","PAB","PEN","PGK","PHP","PKR","PLN","PYG","QAR","RON","RSD","RUB","RWF","SAR","SBD","SCR","SDG","SEK","SGD","SHP","SLL","SOS","SPL","SRD","STD","SVC","SYP","SZL","THB","TJS","TMT","TND","TOP","TRY","TTD","TVD","TWD","TZS","UAH","UGX","USD","UYU","UZS","VEF","VND","VUV","WST","XAF","XCD","XDR","XOF","XPF","YER","ZAR","ZMW","ZWD","BTC"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		catalogCurrenciesItemsEnum = append(catalogCurrenciesItemsEnum, v)
	}
}

func (m *Catalog) validateCurrenciesItemsEnum(path, location string, value CatalogCurrenciesEnum) error {
	if err := validate.Enum(path, location, value, catalogCurrenciesItemsEnum); err != nil {
		return err
	}
	return nil
}

func (m *Catalog) validateCurrencies(formats strfmt.Registry) error {

	if swag.IsZero(m.Currencies) { // not required
		return nil
	}

	for i := 0; i < len(m.Currencies); i++ {

		// value enum
		if err := m.validateCurrenciesItemsEnum("currencies"+"."+strconv.Itoa(i), "body", m.Currencies[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *Catalog) validateEffectiveDate(formats strfmt.Registry) error {

	if swag.IsZero(m.EffectiveDate) { // not required
		return nil
	}

	if err := validate.FormatOf("effectiveDate", "body", "date-time", m.EffectiveDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Catalog) validatePriceLists(formats strfmt.Registry) error {

	if swag.IsZero(m.PriceLists) { // not required
		return nil
	}

	for i := 0; i < len(m.PriceLists); i++ {
		if swag.IsZero(m.PriceLists[i]) { // not required
			continue
		}

		if m.PriceLists[i] != nil {
			if err := m.PriceLists[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("priceLists" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Catalog) validateProducts(formats strfmt.Registry) error {

	if swag.IsZero(m.Products) { // not required
		return nil
	}

	for i := 0; i < len(m.Products); i++ {
		if swag.IsZero(m.Products[i]) { // not required
			continue
		}

		if m.Products[i] != nil {
			if err := m.Products[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("products" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Catalog) validateUnits(formats strfmt.Registry) error {

	if swag.IsZero(m.Units) { // not required
		return nil
	}

	for i := 0; i < len(m.Units); i++ {
		if swag.IsZero(m.Units[i]) { // not required
			continue
		}

		if m.Units[i] != nil {
			if err := m.Units[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("units" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Catalog) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Catalog) UnmarshalBinary(b []byte) error {
	var res Catalog
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
