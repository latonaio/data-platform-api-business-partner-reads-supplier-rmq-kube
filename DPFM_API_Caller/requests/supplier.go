package requests

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
