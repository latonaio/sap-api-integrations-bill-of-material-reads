package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-bill-of-material-reads/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

func ConvertToHeader(raw []byte, l *logger.Logger) ([]Header, error) {
	pm := &responses.Header{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Header. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	header := make([]Header, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		header = append(header, Header{
			BillOfMaterial:                data.BillOfMaterial,
			BillOfMaterialCategory:        data.BillOfMaterialCategory,
			BillOfMaterialVariant:         data.BillOfMaterialVariant,
			BillOfMaterialVersion:         data.BillOfMaterialVersion,
			EngineeringChangeDocument:     data.EngineeringChangeDocument,
			Material:                      data.Material,
			Plant:                         data.Plant,
			BillOfMaterialHeaderUUID:      data.BillOfMaterialHeaderUUID,
			BillOfMaterialVariantUsage:    data.BillOfMaterialVariantUsage,
			EngineeringChangeDocForEdit:   data.EngineeringChangeDocForEdit,
			IsMultipleBOMAlt:              data.IsMultipleBOMAlt,
			BOMHeaderInternalChangeCount:  data.BOMHeaderInternalChangeCount,
			BOMUsagePriority:              data.BOMUsagePriority,
			BillOfMaterialAuthsnGrp:       data.BillOfMaterialAuthsnGrp,
			BOMVersionStatus:              data.BOMVersionStatus,
			IsVersionBillOfMaterial:       data.IsVersionBillOfMaterial,
			IsLatestBOMVersion:            data.IsLatestBOMVersion,
			IsConfiguredMaterial:          data.IsConfiguredMaterial,
			BOMTechnicalType:              data.BOMTechnicalType,
			BOMGroup:                      data.BOMGroup,
			BOMHeaderText:                 data.BOMHeaderText,
			BOMAlternativeText:            data.BOMAlternativeText,
			BillOfMaterialStatus:          data.BillOfMaterialStatus,
			HeaderValidityStartDate:       data.HeaderValidityStartDate,
			HeaderValidityEndDate:         data.HeaderValidityEndDate,
			ChgToEngineeringChgDocument:   data.ChgToEngineeringChgDocument,
			IsMarkedForDeletion:           data.IsMarkedForDeletion,
			IsALE:                         data.IsALE,
			MatFromLotSizeQuantity:        data.MatFromLotSizeQuantity,
			MaterialToLotSizeQuantity:     data.MaterialToLotSizeQuantity,
			BOMHeaderBaseUnit:             data.BOMHeaderBaseUnit,
			BOMHeaderQuantityInBaseUnit:   data.BOMHeaderQuantityInBaseUnit,
			RecordCreationDate:            data.RecordCreationDate,
			LastChangeDate:                data.LastChangeDate,
			BOMIsToBeDeleted:              data.BOMIsToBeDeleted,
			DocumentIsCreatedByCAD:        data.DocumentIsCreatedByCAD,
			LaboratoryOrDesignOffice:      data.LaboratoryOrDesignOffice,
			LastChangeDateTime:            data.LastChangeDateTime,
			ProductDescription:            data.ProductDescription,
			PlantName:                     data.PlantName,
			BillOfMaterialHdrDetailsText:  data.BillOfMaterialHdrDetailsText,
			SelectedBillOfMaterialVersion: data.SelectedBillOfMaterialVersion,
			ToItem:                        data.ToItem.Deferred.URI,
		})
	}

	return header, nil
}

func ConvertToItem(raw []byte, l *logger.Logger) ([]Item, error) {
	pm := &responses.Item{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Item. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	item := make([]Item, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		item = append(item, Item{
			BillOfMaterial:               data.BillOfMaterial,
			BillOfMaterialVariant:        data.BillOfMaterialVariant,
			BillOfMaterialCategory:       data.BillOfMaterialCategory,
			BillOfMaterialVersion:        data.BillOfMaterialVersion,
			BillOfMaterialItemNodeNumber: data.BillOfMaterialItemNodeNumber,
			HeaderChangeDocument:         data.HeaderChangeDocument,
			Material:                     data.Material,
			Plant:                        data.Plant,
			ValidityStartDate:            data.ValidityStartDate,
			ValidityEndDate:              data.ValidityEndDate,
			BillOfMaterialComponent:      data.BillOfMaterialComponent,
			ComponentDescription:         data.ComponentDescription,
			BillOfMaterialItemQuantity:   data.BillOfMaterialItemQuantity,
			ComponentScrapInPercent:      data.ComponentScrapInPercent,
		})
	}

	return item, nil
}

func ConvertToToItem(raw []byte, l *logger.Logger) ([]ToItem, error) {
	pm := &responses.ToItem{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ToItem. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	toItem := make([]ToItem, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		toItem = append(toItem, ToItem{
			BillOfMaterial:               data.BillOfMaterial,
			BillOfMaterialVariant:        data.BillOfMaterialVariant,
			BillOfMaterialCategory:       data.BillOfMaterialCategory,
			BillOfMaterialVersion:        data.BillOfMaterialVersion,
			BillOfMaterialItemNodeNumber: data.BillOfMaterialItemNodeNumber,
			HeaderChangeDocument:         data.HeaderChangeDocument,
			Material:                     data.Material,
			Plant:                        data.Plant,
			ValidityStartDate:            data.ValidityStartDate,
			ValidityEndDate:              data.ValidityEndDate,
			BillOfMaterialComponent:      data.BillOfMaterialComponent,
			ComponentDescription:         data.ComponentDescription,
			BillOfMaterialItemQuantity:   data.BillOfMaterialItemQuantity,
			ComponentScrapInPercent:      data.ComponentScrapInPercent,
		})
	}

	return toItem, nil
}
