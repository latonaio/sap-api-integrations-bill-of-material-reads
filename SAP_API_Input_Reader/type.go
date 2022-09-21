package sap_api_input_reader

type EC_MC struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	Document      struct {
		DocumentNo     string `json:"document_no"`
		DeliverTo      string `json:"deliver_to"`
		Quantity       string `json:"quantity"`
		PickedQuantity string `json:"picked_quantity"`
		Price          string `json:"price"`
		Batch          string `json:"batch"`
	} `json:"document"`
	ProductionOrder struct {
		DocumentNo           string `json:"document_no"`
		Status               string `json:"status"`
		DeliverTo            string `json:"deliver_to"`
		Quantity             string `json:"quantity"`
		CompletedQuantity    string `json:"completed_quantity"`
		PlannedStartDate     string `json:"planned_start_date"`
		PlannedValidatedDate string `json:"planned_validated_date"`
		ActualStartDate      string `json:"actual_start_date"`
		ActualValidatedDate  string `json:"actual_validated_date"`
		Batch                string `json:"batch"`
		Work                 struct {
			WorkNo                   string `json:"work_no"`
			Quantity                 string `json:"quantity"`
			CompletedQuantity        string `json:"completed_quantity"`
			ErroredQuantity          string `json:"errored_quantity"`
			Component                string `json:"component"`
			PlannedComponentQuantity string `json:"planned_component_quantity"`
			PlannedStartDate         string `json:"planned_start_date"`
			PlannedStartTime         string `json:"planned_start_time"`
			PlannedValidatedDate     string `json:"planned_validated_date"`
			PlannedValidatedTime     string `json:"planned_validated_time"`
			ActualStartDate          string `json:"actual_start_date"`
			ActualStartTime          string `json:"actual_start_time"`
			ActualValidatedDate      string `json:"actual_validated_date"`
			ActualValidatedTime      string `json:"actual_validated_time"`
		} `json:"work"`
	} `json:"production_order"`
	APISchema      string `json:"api_schema"`
	MaterialCode   string `json:"material_code"`
	Plant_Supplier string `json:"plant/supplier"`
	Stock          string `json:"stock"`
	DocumentType   string `json:"document_type"`
	DocumentNo     string `json:"document_no"`
	PlannedDate    string `json:"planned_date"`
	ValidatedDate  string `json:"validated_date"`
	Deleted        bool   `json:"deleted"`
}

type SDC struct {
	ConnectionKey  string `json:"connection_key"`
	Result         bool   `json:"result"`
	RedisKey       string `json:"redis_key"`
	Filepath       string `json:"filepath"`
	BillOfMaterial struct {
		BillOfMaterial                string `json:"BillOfMaterial"`
		BillOfMaterialCategory        string `json:"BillOfMaterialCategory"`
		BillOfMaterialVariant         string `json:"BillOfMaterialVariant"`
		BillOfMaterialVersion         string `json:"BillOfMaterialVersion"`
		EngineeringChangeDocument     string `json:"EngineeringChangeDocument"`
		Material                      string `json:"Material"`
		Plant                         string `json:"Plant"`
		BillOfMaterialHeaderUUID      string `json:"BillOfMaterialHeaderUUID"`
		BillOfMaterialVariantUsage    string `json:"BillOfMaterialVariantUsage"`
		EngineeringChangeDocForEdit   string `json:"EngineeringChangeDocForEdit"`
		IsMultipleBOMAlt              string `json:"IsMultipleBOMAlt"`
		BOMHeaderInternalChangeCount  string `json:"BOMHeaderInternalChangeCount"`
		BOMUsagePriority              string `json:"BOMUsagePriority"`
		BillOfMaterialAuthsnGrp       string `json:"BillOfMaterialAuthsnGrp"`
		BOMVersionStatus              string `json:"BOMVersionStatus"`
		IsVersionBillOfMaterial       string `json:"IsVersionBillOfMaterial"`
		IsLatestBOMVersion            string `json:"IsLatestBOMVersion"`
		IsConfiguredMaterial          string `json:"IsConfiguredMaterial"`
		BOMTechnicalType              string `json:"BOMTechnicalType"`
		BOMGroup                      string `json:"BOMGroup"`
		BOMHeaderText                 string `json:"BOMHeaderText"`
		BOMAlternativeText            string `json:"BOMAlternativeText"`
		BillOfMaterialStatus          string `json:"BillOfMaterialStatus"`
		HeaderValidityStartDate       string `json:"HeaderValidityStartDate"`
		HeaderValidityEndDate         string `json:"HeaderValidityEndDate"`
		ChgToEngineeringChgDocument   string `json:"ChgToEngineeringChgDocument"`
		IsMarkedForDeletion           string `json:"IsMarkedForDeletion"`
		IsALE                         string `json:"IsALE"`
		MatFromLotSizeQuantity        string `json:"MatFromLotSizeQuantity"`
		MaterialToLotSizeQuantity     string `json:"MaterialToLotSizeQuantity"`
		BOMHeaderBaseUnit             string `json:"BOMHeaderBaseUnit"`
		BOMHeaderQuantityInBaseUnit   string `json:"BOMHeaderQuantityInBaseUnit"`
		RecordCreationDate            string `json:"RecordCreationDate"`
		LastChangeDate                string `json:"LastChangeDate"`
		BOMIsToBeDeleted              string `json:"BOMIsToBeDeleted"`
		DocumentIsCreatedByCAD        string `json:"DocumentIsCreatedByCAD"`
		LaboratoryOrDesignOffice      string `json:"LaboratoryOrDesignOffice"`
		LastChangeDateTime            string `json:"LastChangeDateTime"`
		ProductDescription            string `json:"ProductDescription"`
		PlantName                     string `json:"PlantName"`
		BillOfMaterialHdrDetailsText  string `json:"BillOfMaterialHdrDetailsText"`
		SelectedBillOfMaterialVersion string `json:"SelectedBillOfMaterialVersion"`
		BillOfMaterialItem            struct {
			BillOfMaterialItemNodeNumber string `json:"BillOfMaterialItemNodeNumber"`
			HeaderChangeDocument         string `json:"HeaderChangeDocument"`
			Material                     string `json:"Material"`
			Plant                        string `json:"Plant"`
			ValidityStartDate            string `json:"ValidityStartDate"`
			ValidityEndDate              string `json:"ValidityEndDate"`
			BillOfMaterialComponent      string `json:"BillOfMaterialComponent"`
			ComponentDescription         string `json:"ComponentDescription"`
			BillOfMaterialItemQuantity   string `json:"BillOfMaterialItemQuantity"`
			ComponentScrapInPercent      string `json:"ComponentScrapInPercent"`
			IsDeleted                    bool   `json:"IsDeleted"`
		} `json:"BillOfMaterialItem"`
	} `json:"BillOfMaterial"`
	APISchema    string   `json:"api_schema"`
	Accepter     []string `json:"accepter"`
	MaterialCode string   `json:"material_code"`
	Plant        string   `json:"plant"`
	Deleted      bool     `json:"deleted"`
}
