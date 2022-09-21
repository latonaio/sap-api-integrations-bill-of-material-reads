package sap_api_output_formatter

type BillOfMaterial struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	Product       string `json:"Product"`
	APISchema     string `json:"api_schema"`
	MaterialCode  string `json:"material_code"`
	Plant         string `json:"plant"`
	Deleted       bool   `json:"deleted"`
}

type Header struct {
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
	IsMultipleBOMAlt              bool   `json:"IsMultipleBOMAlt"`
	BOMHeaderInternalChangeCount  string `json:"BOMHeaderInternalChangeCount"`
	BOMUsagePriority              string `json:"BOMUsagePriority"`
	BillOfMaterialAuthsnGrp       string `json:"BillOfMaterialAuthsnGrp"`
	BOMVersionStatus              string `json:"BOMVersionStatus"`
	IsVersionBillOfMaterial       bool   `json:"IsVersionBillOfMaterial"`
	IsLatestBOMVersion            bool   `json:"IsLatestBOMVersion"`
	IsConfiguredMaterial          bool   `json:"IsConfiguredMaterial"`
	BOMTechnicalType              string `json:"BOMTechnicalType"`
	BOMGroup                      string `json:"BOMGroup"`
	BOMHeaderText                 string `json:"BOMHeaderText"`
	BOMAlternativeText            string `json:"BOMAlternativeText"`
	BillOfMaterialStatus          string `json:"BillOfMaterialStatus"`
	HeaderValidityStartDate       string `json:"HeaderValidityStartDate"`
	HeaderValidityEndDate         string `json:"HeaderValidityEndDate"`
	ChgToEngineeringChgDocument   string `json:"ChgToEngineeringChgDocument"`
	IsMarkedForDeletion           bool   `json:"IsMarkedForDeletion"`
	IsALE                         bool   `json:"IsALE"`
	MatFromLotSizeQuantity        string `json:"MatFromLotSizeQuantity"`
	MaterialToLotSizeQuantity     string `json:"MaterialToLotSizeQuantity"`
	BOMHeaderBaseUnit             string `json:"BOMHeaderBaseUnit"`
	BOMHeaderQuantityInBaseUnit   string `json:"BOMHeaderQuantityInBaseUnit"`
	RecordCreationDate            string `json:"RecordCreationDate"`
	LastChangeDate                string `json:"LastChangeDate"`
	BOMIsToBeDeleted              string `json:"BOMIsToBeDeleted"`
	DocumentIsCreatedByCAD        bool   `json:"DocumentIsCreatedByCAD"`
	LaboratoryOrDesignOffice      string `json:"LaboratoryOrDesignOffice"`
	LastChangeDateTime            string `json:"LastChangeDateTime"`
	ProductDescription            string `json:"ProductDescription"`
	PlantName                     string `json:"PlantName"`
	BillOfMaterialHdrDetailsText  string `json:"BillOfMaterialHdrDetailsText"`
	SelectedBillOfMaterialVersion string `json:"SelectedBillOfMaterialVersion"`
	ToItem                        string `json:"to_BillOfMaterialItem"`
}

type Item struct {
	BillOfMaterial               string `json:"BillOfMaterial"`
	BillOfMaterialVariant        string `json:"BillOfMaterialVariant"`
	BillOfMaterialCategory       string `json:"BillOfMaterialCategory"`
	BillOfMaterialVersion        string `json:"BillOfMaterialVersion"`
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
}

type ToItem struct {
	BillOfMaterial               string `json:"BillOfMaterial"`
	BillOfMaterialVariant        string `json:"BillOfMaterialVariant"`
	BillOfMaterialCategory       string `json:"BillOfMaterialCategory"`
	BillOfMaterialVersion        string `json:"BillOfMaterialVersion"`
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
}
