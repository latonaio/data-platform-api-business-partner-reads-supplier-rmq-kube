package dpfm_api_input_reader

import (
	"data-platform-api-business-partner-reads-supplier-rmq-kube/DPFM_API_Caller/requests"
)

func (sdc *SDC) ConvertToSupplier() *requests.Supplier {
	data := sdc.Supplier
	return &requests.Supplier{
		BusinessPartner:               data.BusinessPartner,
		Supplier:                      data.Supplier,
		Currency:                      data.Currency,
		PaymentTerms:                  data.PaymentTerms,
		PaymentMethod:                 data.PaymentMethod,
		Incoterms:                     data.Incoterms,
		BPAccountAssignmentGroup:      data.BPAccountAssignmentGroup,
		InvoiceIsGoodsReceiptBased:    data.InvoiceIsGoodsReceiptBased,
		PurOrdAutoGenerationIsAllowed: data.PurOrdAutoGenerationIsAllowed,
		CreationDate:                  data.CreationDate,
		QuotationIsBlockedForSupplier: data.QuotationIsBlockedForSupplier,
		OrderIsBlockedForSupplier:     data.OrderIsBlockedForSupplier,
		DeliveryIsBlockedForSupplier:  data.DeliveryIsBlockedForSupplier,
		BillingIsBlockedForSupplier:   data.BillingIsBlockedForSupplier,
		PostingIsBlockedForSupplier:   data.PostingIsBlockedForSupplier,
		PaymentIsBlockedForSupplier:   data.PaymentIsBlockedForSupplier,
		IsMarkedForDeletion:           data.IsMarkedForDeletion,
	}
}

func (sdc *SDC) ConvertToContact() *requests.Contact {
	dataSupplier := sdc.Supplier
	data := sdc.Supplier.Contact
	return &requests.Contact{
		BusinessPartner:     dataSupplier.BusinessPartner,
		Supplier:            dataSupplier.Supplier,
		ContactID:           data.ContactID,
		ContactPersonName:   data.ContactPersonName,
		EmailAddress:        data.EmailAddress,
		PhoneNumber:         data.PhoneNumber,
		MobilePhoneNumber:   data.MobilePhoneNumber,
		FaxNumber:           data.FaxNumber,
		ContactTag1:         data.ContactTag1,
		ContactTag2:         data.ContactTag2,
		ContactTag3:         data.ContactTag3,
		ContactTag4:         data.ContactTag4,
		DefaultContact:      data.DefaultContact,
		IsMarkedForDeletion: data.IsMarkedForDeletion,
	}
}

func (sdc *SDC) ConvertToPartnerFunction() *requests.PartnerFunction {
	dataSupplier := sdc.Supplier
	data := sdc.Supplier.PartnerFunction
	return &requests.PartnerFunction{
		BusinessPartner:                dataSupplier.BusinessPartner,
		Supplier:                       dataSupplier.Supplier,
		PartnerCounter:                 data.PartnerCounter,
		PartnerFunction:                data.PartnerFunction,
		PartnerFunctionBusinessPartner: data.PartnerFunctionBusinessPartner,
		DefaultPartner:                 data.DefaultPartner,
		CreationDate:                   data.CreationDate,
		IsMarkedForDeletion:            data.IsMarkedForDeletion,
	}
}

func (sdc *SDC) ConvertToPartnerFunctionContact() *requests.PartnerFunctionContact {
	dataSupplier := sdc.Supplier
	dataPartnerFunction := sdc.Supplier.PartnerFunction
	data := sdc.Supplier.PartnerFunction.PartnerFunctionContact
	return &requests.PartnerFunctionContact{
		BusinessPartner:                dataSupplier.BusinessPartner,
		Supplier:                       dataSupplier.Supplier,
		PartnerCounter:                 dataPartnerFunction.PartnerCounter,
		ContactID:                      data.ContactID,
		PartnerFunction:                data.PartnerFunction,
		PartnerFunctionBusinessPartner: data.PartnerFunctionBusinessPartner,
		DefaultPartner:                 data.DefaultPartner,
		ContactPersonName:              data.ContactPersonName,
		EmailAddress:                   data.EmailAddress,
		PhoneNumber:                    data.PhoneNumber,
		MobilePhoneNumber:              data.MobilePhoneNumber,
		FaxNumber:                      data.FaxNumber,
		ContactTag1:                    data.ContactTag1,
		ContactTag2:                    data.ContactTag2,
		ContactTag3:                    data.ContactTag3,
		ContactTag4:                    data.ContactTag4,
		DefaultContact:                 data.DefaultContact,
		IsMarkedForDeletion:            data.IsMarkedForDeletion,
	}
}

func (sdc *SDC) ConvertToPartnerPlant() *requests.PartnerPlant {
	dataSupplier := sdc.Supplier
	dataPartnerFunction := sdc.Supplier.PartnerFunction
	data := sdc.Supplier.PartnerFunction.PartnerPlant
	return &requests.PartnerPlant{
		BusinessPartner:                dataSupplier.BusinessPartner,
		Supplier:                       dataSupplier.Supplier,
		PartnerCounter:                 dataPartnerFunction.PartnerCounter,
		PartnerFunction:                dataPartnerFunction.PartnerFunction,
		PartnerFunctionBusinessPartner: dataPartnerFunction.PartnerFunctionBusinessPartner,
		PlantCounter:                   data.PlantCounter,
		Plant:                          data.Plant,
		DefaultPlant:                   data.DefaultPlant,
		DefaultStockConfirmationPlant:  data.DefaultStockConfirmationPlant,
		IsMarkedForDeletion:            data.IsMarkedForDeletion,
	}
}

func (sdc *SDC) ConvertToFinInst() *requests.FinInst {
	dataSupplier := sdc.Supplier
	data := sdc.Supplier.FinInst
	return &requests.FinInst{
		BusinessPartner:           dataSupplier.BusinessPartner,
		Supplier:                  dataSupplier.Supplier,
		FinInstIdentification:     data.FinInstIdentification,
		ValidityEndDate:           data.ValidityEndDate,
		ValidityStartDate:         data.ValidityStartDate,
		FinInstCountry:            data.FinInstCountry,
		FinInstNumber:             data.FinInstNumber,
		FinInstName:               data.FinInstName,
		FinInstBranchName:         data.FinInstBranchName,
		SWIFTCode:                 data.SWIFTCode,
		InternalFinInstCustomerID: data.InternalFinInstCustomerID,
		InternalFinInstAccountID:  data.InternalFinInstAccountID,
		FinInstControlKey:         data.FinInstControlKey,
		FinInstAccountName:        data.FinInstAccountName,
		FinInstAccount:            data.FinInstAccount,
		IsMarkedForDeletion:       data.IsMarkedForDeletion,
	}
}

func (sdc *SDC) ConvertToTax() *requests.Tax {
	dataSupplier := sdc.Supplier
	data := sdc.Supplier.Tax
	return &requests.Tax{
		BusinessPartner:     dataSupplier.BusinessPartner,
		Supplier:            dataSupplier.Supplier,
		DepartureCountry:    data.DepartureCountry,
		TaxCategory:         data.TaxCategory,
		BPTaxClassification: data.BPTaxClassification,
	}
}

func (sdc *SDC) ConvertToAccounting() *requests.Accounting {
	dataSupplier := sdc.Supplier
	data := sdc.Supplier.Accounting
	return &requests.Accounting{
		BusinessPartner:       dataSupplier.BusinessPartner,
		Supplier:              dataSupplier.Supplier,
		HouseBank:             data.HouseBank,
		ClearCustomerSupplier: data.ClearCustomerSupplier,
		ReconciliationAccount: data.ReconciliationAccount,
		IsMarkedForDeletion:   data.IsMarkedForDeletion,
	}
}
