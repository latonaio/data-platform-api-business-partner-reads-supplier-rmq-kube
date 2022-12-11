package dpfm_api_caller

import (
	"context"
	dpfm_api_input_reader "data-platform-api-business-partner-reads-supplier-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-business-partner-reads-supplier-rmq-kube/DPFM_API_Output_Formatter"
	"sync"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func (c *DPFMAPICaller) readSqlProcess(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	accepter []string,
	errs *[]error,
	log *logger.Logger,
) interface{} {
	var supplier *dpfm_api_output_formatter.Supplier
	var partnerFunction *dpfm_api_output_formatter.PartnerFunction
	for _, fn := range accepter {
		switch fn {
		case "Supplier":
			func() {
				supplier = c.Supplier(mtx, input, output, errs, log)
			}()
		case "PartnerFunction":
			func() {
				partnerFunction = c.PartnerFunction(mtx, input, output, errs, log)
			}()
		default:
		}
	}

	data := &dpfm_api_output_formatter.Message{
		Supplier:        supplier,
		PartnerFunction: partnerFunction,
	}

	return data
}

func (c *DPFMAPICaller) Supplier(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *dpfm_api_output_formatter.Supplier {
	businessPartner := input.Supplier.BusinessPartner
	supplier := input.Supplier.Supplier

	rows, err := c.db.Query(
		`SELECT BusinessPartner, Supplier, Currency, PaymentTerms, PaymentMethod, Incoterms, 
		BPAccountAssignmentGroup, InvoiceIsGoodsReceiptBased, PurOrdAutoGenerationIsAllowed, CreationDate, 
		QuotationIsBlockedForSupplier, OrderIsBlockedForSupplier, DeliveryIsBlockedForSupplier, 
		BillingIsBlockedForSupplier, PostingIsBlockedForSupplier, PaymentIsBlockedForSupplier, IsMarkedForDeletion
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_business_partner_supplier_data
		WHERE (BusinessPartner, Supplier) = (?, ?);`, businessPartner, supplier,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	data, err := dpfm_api_output_formatter.ConvertToSupplier(input, rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) Contact(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *dpfm_api_output_formatter.Contact {
	businessPartner := input.Supplier.BusinessPartner
	supplier := input.Supplier.Supplier
	contactID := input.Supplier.Contact.ContactID

	rows, err := c.db.Query(
		`SELECT BusinessPartner, Supplier, ContactID, ContactPersonName, EmailAddress, PhoneNumber, MobilePhoneNumber, FaxNumber, ContactTag1, ContactTag2, ContactTag3, ContactTag4, DefaultContact, IsMarkedForDeletion
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_business_partner_supplier_contact_data
		WHERE (BusinessPartner, Supplier, ContactID) = (?, ?, ?);`, businessPartner, supplier, contactID,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	data, err := dpfm_api_output_formatter.ConvertToContact(input, rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) PartnerFunction(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *dpfm_api_output_formatter.PartnerFunction {
	businessPartner := input.Supplier.BusinessPartner
	supplier := input.Supplier.Supplier
	partnerCounter := input.Supplier.PartnerFunction.PartnerCounter

	rows, err := c.db.Query(
		`SELECT BusinessPartner, Supplier, PartnerCounter, PartnerFunction, PartnerFunctionBusinessPartner, 
		DefaultPartner, CreationDate, IsMarkedForDeletion
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_business_partner_supplier_partner_function_data
		WHERE (BusinessPartner, Supplier, PartnerCounter) = (?, ?, ?);`, businessPartner, supplier, partnerCounter,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	data, err := dpfm_api_output_formatter.ConvertToPartnerFunction(input, rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) PartnerFunctionContact(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *dpfm_api_output_formatter.PartnerFunctionContact {
	businessPartner := input.Supplier.BusinessPartner
	supplier := input.Supplier.Supplier
	partnerCounter := input.Supplier.PartnerFunction.PartnerCounter
	contactID := input.Supplier.PartnerFunction.PartnerFunctionContact.ContactID

	rows, err := c.db.Query(
		`SELECT BusinessPartner, Supplier, PartnerCounter, ContactID, PartnerFunction, PartnerFunctionBusinessPartner, 
		DefaultPartner, ContactPersonName, EmailAddress, PhoneNumber, MobilePhoneNumber, FaxNumber, ContactTag1, 
		ContactTag2, ContactTag3, ContactTag4, DefaultContact, IsMarkedForDeletion
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_business_partner_supplier_partner_function_contact_data
		WHERE (BusinessPartner, Supplier, PartnerCounter, ContactID) = (?, ?, ?, ?);`, businessPartner, supplier, partnerCounter, contactID,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	data, err := dpfm_api_output_formatter.ConvertToPartnerFunctionContact(input, rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) PartnerPlant(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *dpfm_api_output_formatter.PartnerPlant {
	businessPartner := input.Supplier.BusinessPartner
	supplier := input.Supplier.Supplier
	partnerCounter := input.Supplier.PartnerFunction.PartnerCounter
	partnerFunction := input.Supplier.PartnerFunction.PartnerFunction
	partnerFunctionBusinessPartner := input.Supplier.PartnerFunction.PartnerFunctionBusinessPartner
	plantCounter := input.Supplier.PartnerFunction.PartnerPlant.PlantCounter

	rows, err := c.db.Query(
		`SELECT BusinessPartner, Supplier, PartnerCounter, PartnerFunction, PartnerFunctionBusinessPartner, PlantCounter, 
		Plant, DefaultPlant, DefaultStockConfirmationPlant, IsMarkedForDeletion
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_business_partner_supplier_partner_plant_data
		WHERE (BusinessPartner, Supplier, PartnerCounter, PartnerFunction, PartnerFunctionBusinessPartner, PlantCounter) = (?, ?, ?, ?, ?, ?);`, businessPartner, supplier, partnerCounter, partnerFunction, partnerFunctionBusinessPartner, plantCounter,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	data, err := dpfm_api_output_formatter.ConvertToPartnerPlant(input, rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) FinInst(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *dpfm_api_output_formatter.FinInst {
	businessPartner := input.Supplier.BusinessPartner
	supplier := input.Supplier.Supplier
	finInstIdentification := input.Supplier.FinInst.FinInstIdentification
	validityEndDate := input.Supplier.FinInst.ValidityEndDate
	validityStartDate := input.Supplier.FinInst.ValidityStartDate

	rows, err := c.db.Query(
		`SELECT BusinessPartner, Supplier, FinInstIdentification, ValidityEndDate, ValidityStartDate, 
		FinInstCountry, FinInstNumber, FinInstName, FinInstBranchName, SWIFTCode, InternalFinInstCustomerID, 
		InternalFinInstAccountID, FinInstControlKey, FinInstAccountName, FinInstAccount, IsMarkedForDeletion
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_business_partner_supplier_fin_inst_data
		WHERE (BusinessPartner, Supplier, FinInstIdentification, ValidityEndDate, ValidityStartDate) = (?, ?, ?, ?);`, businessPartner, supplier, finInstIdentification, validityEndDate, validityStartDate,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	data, err := dpfm_api_output_formatter.ConvertToFinInst(input, rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) Tax(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *dpfm_api_output_formatter.Tax {
	businessPartner := input.Supplier.BusinessPartner
	supplier := input.Supplier.Supplier
	departureCountry := input.Supplier.Tax.DepartureCountry

	rows, err := c.db.Query(
		`SELECT BusinessPartner, Supplier, DepartureCountry, TaxCategory, BPTaxClassification
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_business_partner_supplier_tax_data
		WHERE (BusinessPartner, Supplier, DepartureCountry) = (?, ?, ?);`, businessPartner, supplier, departureCountry,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	data, err := dpfm_api_output_formatter.ConvertToTax(input, rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) Accounting(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *dpfm_api_output_formatter.Accounting {
	businessPartner := input.Supplier.BusinessPartner
	supplier := input.Supplier.Supplier

	rows, err := c.db.Query(
		`SELECT BusinessPartner, Supplier, HouseBank, ClearCustomerSupplier, ReconciliationAccount, IsMarkedForDeletion
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_business_partner_supplier_data
		WHERE (BusinessPartner, Supplier) = ?;`, businessPartner, supplier,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	data, err := dpfm_api_output_formatter.ConvertToAccounting(input, rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}
