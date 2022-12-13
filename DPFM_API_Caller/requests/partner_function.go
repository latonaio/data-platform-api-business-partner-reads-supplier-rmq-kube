package requests

type PartnerFunction struct {
	BusinessPartner                int     `json:"BusinessPartner"`
	Supplier                       int     `json:"Supplier"`
	PartnerCounter                 int     `json:"PartnerCounter"`
	PartnerFunction                *string `json:"PartnerFunction"`
	PartnerFunctionBusinessPartner *int    `json:"PartnerFunctionBusinessPartner"`
	DefaultPartner                 *bool   `json:"DefaultPartner"`
	CreationDate                   *string `json:"CreationDate"`
	IsMarkedForDeletion            *bool   `json:"IsMarkedForDeletion"`
}
