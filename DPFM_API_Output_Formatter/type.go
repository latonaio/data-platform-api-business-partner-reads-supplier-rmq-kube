package dpfm_api_output_formatter

type SDC struct {
	ConnectionKey       string      `json:"connection_key"`
	Result              bool        `json:"result"`
	RedisKey            string      `json:"redis_key"`
	Filepath            string      `json:"filepath"`
	APIStatusCode       int         `json:"api_status_code"`
	RuntimeSessionID    string      `json:"runtime_session_id"`
	BusinessPartnerID   *int        `json:"business_partner"`
	ServiceLabel        string      `json:"service_label"`
	APIType             string      `json:"api_type"`
	Message             interface{} `json:"message"`
	APISchema           string      `json:"api_schema"`
	Accepter            []string    `json:"accepter"`
	Deleted             bool        `json:"deleted"`
	SQLUpdateResult     *bool       `json:"sql_update_result"`
	SQLUpdateError      string      `json:"sql_update_error"`
	SubfuncResult       *bool       `json:"subfunc_result"`
	SubfuncError        string      `json:"subfunc_error"`
	ExconfResult        *bool       `json:"exconf_result"`
	ExconfError         string      `json:"exconf_error"`
	APIProcessingResult *bool       `json:"api_processing_result"`
	APIProcessingError  string      `json:"api_processing_error"`
}

type Message struct {
	Supplier               *Supplier               `json:"Supplier"`
	Contact                *Contact                `json:"Contact"`
	PartnerFunction        *PartnerFunction        `json:"PartnerFunction"`
	PartnerFunctionContact *PartnerFunctionContact `json:"PartnerFunctionContact"`
	PartnerPlant           *PartnerPlant           `json:"PartnerPlant"`
	FinInst                *FinInst                `json:"FinInst"`
	Tax                    *Tax                    `json:"Tax"`
	Accounting             *Accounting             `json:"Accounting"`
}

type BusinessPartnerSupplier struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	Product       string `json:"Product"`
	APISchema     string `json:"api_schema"`
	MaterialCode  string `json:"material_code"`
	Deleted       string `json:"deleted"`
}

type Supplier struct {
	BusinessPartner               *int    `json:"BusinessPartner"`
	Supplier                      *int    `json:"Supplier"`
	Currency                      string  `json:"Currency"`
	PaymentTerms                  string  `json:"PaymentTerms"`
	PaymentMethod                 string  `json:"PaymentMethod"`
	Incoterms                     string  `json:"Incoterms"`
	BPAccountAssignmentGroup      string  `json:"BPAccountAssignmentGroup"`
	InvoiceIsGoodsReceiptBased    *bool   `json:"InvoiceIsGoodsReceiptBased"`
	PurOrdAutoGenerationIsAllowed *bool   `json:"PurOrdAutoGenerationIsAllowed"`
	CreationDate                  *string `json:"CreationDate"`
	QuotationIsBlockedForSupplier *bool   `json:"QuotationIsBlockedForSupplier"`
	OrderIsBlockedForSupplier     *bool   `json:"OrderIsBlockedForSupplier"`
	DeliveryIsBlockedForSupplier  *bool   `json:"DeliveryIsBlockedForSupplier"`
	BillingIsBlockedForSupplier   *bool   `json:"BillingIsBlockedForSupplier"`
	PostingIsBlockedForSupplier   *bool   `json:"PostingIsBlockedForSupplier"`
	PaymentIsBlockedForSupplier   *bool   `json:"PaymentIsBlockedForSupplier"`
	IsMarkedForDeletion           *bool   `json:"IsMarkedForDeletion"`
}

type Contact struct {
	BusinessPartner     *int   `json:"BusinessPartner"`
	Supplier            *int   `json:"Supplier"`
	ContactID           *int   `json:"ContactID"`
	ContactPersonName   string `json:"ContactPersonName"`
	EmailAddress        string `json:"EmailAddress"`
	PhoneNumber         string `json:"PhoneNumber"`
	MobilePhoneNumber   string `json:"MobilePhoneNumber"`
	FaxNumber           string `json:"FaxNumber"`
	ContactTag1         string `json:"ContactTag1"`
	ContactTag2         string `json:"ContactTag2"`
	ContactTag3         string `json:"ContactTag3"`
	ContactTag4         string `json:"ContactTag4"`
	DefaultContact      *bool  `json:"DefaultContact"`
	IsMarkedForDeletion *bool  `json:"IsMarkedForDeletion"`
}

type PartnerFunction struct {
	BusinessPartner                *int    `json:"BusinessPartner"`
	Supplier                       *int    `json:"Supplier"`
	PartnerCounter                 *int    `json:"PartnerCounter"`
	PartnerFunction                string  `json:"PartnerFunction"`
	PartnerFunctionBusinessPartner *int    `json:"PartnerFunctionBusinessPartner"`
	DefaultPartner                 *bool   `json:"DefaultPartner"`
	CreationDate                   *string `json:"CreationDate"`
	IsMarkedForDeletion            *bool   `json:"IsMarkedForDeletion"`
}

type PartnerFunctionContact struct {
	BusinessPartner                *int   `json:"BusinessPartner"`
	Supplier                       *int   `json:"Supplier"`
	PartnerCounter                 *int   `json:"PartnerCounter"`
	ContactID                      *int   `json:"ContactID"`
	PartnerFunction                string `json:"PartnerFunction"`
	PartnerFunctionBusinessPartner *int   `json:"PartnerFunctionBusinessPartner"`
	DefaultPartner                 *bool  `json:"DefaultPartner"`
	ContactPersonName              string `json:"ContactPersonName"`
	EmailAddress                   string `json:"EmailAddress"`
	PhoneNumber                    string `json:"PhoneNumber"`
	MobilePhoneNumber              string `json:"MobilePhoneNumber"`
	FaxNumber                      string `json:"FaxNumber"`
	ContactTag1                    string `json:"ContactTag1"`
	ContactTag2                    string `json:"ContactTag2"`
	ContactTag3                    string `json:"ContactTag3"`
	ContactTag4                    string `json:"ContactTag4"`
	DefaultContact                 *bool  `json:"DefaultContact"`
	IsMarkedForDeletion            *bool  `json:"IsMarkedForDeletion"`
}

type PartnerPlant struct {
	BusinessPartner                *int   `json:"BusinessPartner"`
	Supplier                       *int   `json:"Supplier"`
	PartnerCounter                 *int   `json:"PartnerCounter"`
	PartnerFunction                string `json:"PartnerFunction"`
	PartnerFunctionBusinessPartner *int   `json:"PartnerFunctionBusinessPartner"`
	PlantCounter                   *int   `json:"PlantCounter"`
	Plant                          string `json:"Plant"`
	DefaultPlant                   *bool  `json:"DefaultPlant"`
	DefaultStockConfirmationPlant  *bool  `json:"DefaultStockConfirmationPlant"`
	IsMarkedForDeletion            *bool  `json:"IsMarkedForDeletion"`
}

type FinInst struct {
	BusinessPartner           *int    `json:"BusinessPartner"`
	Supplier                  *int    `json:"Supplier"`
	FinInstIdentification     *int    `json:"FinInstIdentification"`
	ValidityEndDate           *string `json:"ValidityEndDate"`
	ValidityStartDate         *string `json:"ValidityStartDate"`
	FinInstCountry            string  `json:"FinInstCountry"`
	FinInstNumber             string  `json:"FinInstNumber"`
	FinInstName               string  `json:"FinInstName"`
	FinInstBranchName         string  `json:"FinInstBranchName"`
	SWIFTCode                 string  `json:"SWIFTCode"`
	InternalFinInstCustomerID *int    `json:"InternalFinInstCustomerID"`
	InternalFinInstAccountID  *int    `json:"InternalFinInstAccountID"`
	FinInstControlKey         string  `json:"FinInstControlKey"`
	FinInstAccountName        string  `json:"FinInstAccountName"`
	FinInstAccount            string  `json:"FinInstAccount"`
	IsMarkedForDeletion       *bool   `json:"IsMarkedForDeletion"`
}

type Tax struct {
	BusinessPartner     *int   `json:"BusinessPartner"`
	Supplier            *int   `json:"Supplier"`
	DepartureCountry    string `json:"DepartureCountry"`
	TaxCategory         string `json:"TaxCategory"`
	BPTaxClassification string `json:"BPTaxClassification"`
}

type Accounting struct {
	BusinessPartner       *int   `json:"BusinessPartner"`
	Supplier              *int   `json:"Supplier"`
	HouseBank             string `json:"HouseBank"`
	ClearCustomerSupplier *bool  `json:"ClearCustomerSupplier"`
	ReconciliationAccount string `json:"ReconciliationAccount"`
	IsMarkedForDeletion   *bool  `json:"IsMarkedForDeletion"`
}
