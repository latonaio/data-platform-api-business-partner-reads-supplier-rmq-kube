package requests

type Tax struct {
	BusinessPartner     int     `json:"BusinessPartner"`
	Supplier            int     `json:"Supplier"`
	DepartureCountry    string  `json:"DepartureCountry"`
	TaxCategory         *string `json:"TaxCategory"`
	BPTaxClassification *string `json:"BPTaxClassification"`
}
