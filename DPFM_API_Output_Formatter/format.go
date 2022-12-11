package dpfm_api_output_formatter

import (
	"data-platform-api-business-partner-reads-supplier-rmq-kube/DPFM_API_Caller/requests"
	api_input_reader "data-platform-api-business-partner-reads-supplier-rmq-kube/DPFM_API_Input_Reader"
	"database/sql"
	"fmt"
)

func ConvertToSupplier(sdc *api_input_reader.SDC, rows *sql.Rows) (*Supplier, error) {
	pm := &requests.Supplier{}

	for i := 0; true; i++ {
		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("DBに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.BusinessPartner,
			&pm.Supplier,
			&pm.Currency,
			&pm.PaymentTerms,
			&pm.PaymentMethod,
			&pm.Incoterms,
			&pm.BPAccountAssignmentGroup,
			&pm.InvoiceIsGoodsReceiptBased,
			&pm.PurOrdAutoGenerationIsAllowed,
			&pm.CreationDate,
			&pm.QuotationIsBlockedForSupplier,
			&pm.OrderIsBlockedForSupplier,
			&pm.DeliveryIsBlockedForSupplier,
			&pm.BillingIsBlockedForSupplier,
			&pm.PostingIsBlockedForSupplier,
			&pm.PaymentIsBlockedForSupplier,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}
	}
	data := pm

	supplier := &Supplier{
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
	return supplier, nil
}

func ConvertToContact(sdc *api_input_reader.SDC, rows *sql.Rows) (*Contact, error) {
	pm := &requests.Contact{}

	for i := 0; true; i++ {
		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("DBに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.BusinessPartner,
			&pm.Supplier,
			&pm.ContactID,
			&pm.ContactPersonName,
			&pm.EmailAddress,
			&pm.PhoneNumber,
			&pm.MobilePhoneNumber,
			&pm.FaxNumber,
			&pm.ContactTag1,
			&pm.ContactTag2,
			&pm.ContactTag3,
			&pm.ContactTag4,
			&pm.DefaultContact,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}
	}
	data := pm

	contact := &Contact{
		BusinessPartner:     data.BusinessPartner,
		Supplier:            data.Supplier,
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
	return contact, nil
}

func ConvertToPartnerFunction(sdc *api_input_reader.SDC, rows *sql.Rows) (*PartnerFunction, error) {
	pm := &requests.PartnerFunction{}

	for i := 0; true; i++ {
		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("DBに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.BusinessPartner,
			&pm.Supplier,
			&pm.PartnerCounter,
			&pm.PartnerFunction,
			&pm.PartnerFunctionBusinessPartner,
			&pm.DefaultPartner,
			&pm.CreationDate,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}
	}
	data := pm

	partnerFunction := &PartnerFunction{
		BusinessPartner:                data.BusinessPartner,
		Supplier:                       data.Supplier,
		PartnerCounter:                 data.PartnerCounter,
		PartnerFunction:                data.PartnerFunction,
		PartnerFunctionBusinessPartner: data.PartnerFunctionBusinessPartner,
		DefaultPartner:                 data.DefaultPartner,
		CreationDate:                   data.CreationDate,
		IsMarkedForDeletion:            data.IsMarkedForDeletion,
	}
	return partnerFunction, nil
}

func ConvertToPartnerFunctionContact(sdc *api_input_reader.SDC, rows *sql.Rows) (*PartnerFunctionContact, error) {
	pm := &requests.PartnerFunctionContact{}

	for i := 0; true; i++ {
		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("DBに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.BusinessPartner,
			&pm.Supplier,
			&pm.PartnerCounter,
			&pm.ContactID,
			&pm.PartnerFunction,
			&pm.PartnerFunctionBusinessPartner,
			&pm.DefaultPartner,
			&pm.ContactPersonName,
			&pm.EmailAddress,
			&pm.PhoneNumber,
			&pm.MobilePhoneNumber,
			&pm.FaxNumber,
			&pm.ContactTag1,
			&pm.ContactTag2,
			&pm.ContactTag3,
			&pm.ContactTag4,
			&pm.DefaultContact,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}
	}
	data := pm

	partnerFunctionContact := &PartnerFunctionContact{
		BusinessPartner:                data.BusinessPartner,
		Supplier:                       data.Supplier,
		PartnerCounter:                 data.PartnerCounter,
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
	return partnerFunctionContact, nil
}

func ConvertToPartnerPlant(sdc *api_input_reader.SDC, rows *sql.Rows) (*PartnerPlant, error) {
	pm := &requests.PartnerPlant{}

	for i := 0; true; i++ {
		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("DBに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.BusinessPartner,
			&pm.Supplier,
			&pm.PartnerCounter,
			&pm.PartnerFunction,
			&pm.PartnerFunctionBusinessPartner,
			&pm.PlantCounter,
			&pm.Plant,
			&pm.DefaultPlant,
			&pm.DefaultStockConfirmationPlant,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}
	}
	data := pm

	partnerPlant := &PartnerPlant{
		BusinessPartner:                data.BusinessPartner,
		Supplier:                       data.Supplier,
		PartnerCounter:                 data.PartnerCounter,
		PartnerFunction:                data.PartnerFunction,
		PartnerFunctionBusinessPartner: data.PartnerFunctionBusinessPartner,
		PlantCounter:                   data.PlantCounter,
		Plant:                          data.Plant,
		DefaultPlant:                   data.DefaultPlant,
		DefaultStockConfirmationPlant:  data.DefaultStockConfirmationPlant,
		IsMarkedForDeletion:            data.IsMarkedForDeletion,
	}
	return partnerPlant, nil
}

func ConvertToFinInst(sdc *api_input_reader.SDC, rows *sql.Rows) (*FinInst, error) {
	pm := &requests.FinInst{}

	for i := 0; true; i++ {
		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("DBに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.BusinessPartner,
			&pm.Supplier,
			&pm.FinInstIdentification,
			&pm.ValidityEndDate,
			&pm.ValidityStartDate,
			&pm.FinInstCountry,
			&pm.FinInstNumber,
			&pm.FinInstName,
			&pm.FinInstBranchName,
			&pm.SWIFTCode,
			&pm.InternalFinInstCustomerID,
			&pm.InternalFinInstAccountID,
			&pm.FinInstControlKey,
			&pm.FinInstAccountName,
			&pm.FinInstAccount,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}
	}
	data := pm

	finInst := &FinInst{
		BusinessPartner:           data.BusinessPartner,
		Supplier:                  data.Supplier,
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
	return finInst, nil
}

func ConvertToTax(sdc *api_input_reader.SDC, rows *sql.Rows) (*Tax, error) {
	pm := &requests.Tax{}

	for i := 0; true; i++ {
		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("DBに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.BusinessPartner,
			&pm.Supplier,
			&pm.DepartureCountry,
			&pm.TaxCategory,
			&pm.BPTaxClassification,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}
	}
	data := pm

	tax := &Tax{
		BusinessPartner:     data.BusinessPartner,
		Supplier:            data.Supplier,
		DepartureCountry:    data.DepartureCountry,
		TaxCategory:         data.TaxCategory,
		BPTaxClassification: data.BPTaxClassification,
	}
	return tax, nil
}

func ConvertToAccounting(sdc *api_input_reader.SDC, rows *sql.Rows) (*Accounting, error) {
	pm := &requests.Accounting{}

	for i := 0; true; i++ {
		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("DBに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.BusinessPartner,
			&pm.Supplier,
			&pm.HouseBank,
			&pm.ClearCustomerSupplier,
			&pm.ReconciliationAccount,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}
	}
	data := pm

	accounting := &Accounting{
		BusinessPartner:       data.BusinessPartner,
		Supplier:              data.Supplier,
		HouseBank:             data.HouseBank,
		ClearCustomerSupplier: data.ClearCustomerSupplier,
		ReconciliationAccount: data.ReconciliationAccount,
		IsMarkedForDeletion:   data.IsMarkedForDeletion,
	}
	return accounting, nil
}
