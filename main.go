package main

import (
	sap_api_caller "sap-api-integrations-bill-of-material-reads/SAP_API_Caller"
	"sap-api-integrations-bill-of-material-reads/SAP_API_Input_Reader"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
)

func main() {
	l := logger.NewLogger()
	fr := sap_api_input_reader.NewFileReader()
	inoutSDC := fr.ReadSDC("./Inputs//SDC_Bill_Of_Material_Product_Description_sample.json")
	caller := sap_api_caller.NewSAPAPICaller(
		"https://sandbox.api.sap.com/s4hanacloud/sap/opu/odata/sap/", l,
	)

	accepter := inoutSDC.Accepter
	if len(accepter) == 0 || accepter[0] == "All" {
		accepter = []string{
			"Header", "Item", "ProductDescription", "Component", "ComponentDescription",
		}
	}

	caller.AsyncGetBillOfMaterial(
		inoutSDC.BillOfMaterial.Material,
		inoutSDC.BillOfMaterial.Plant,
		inoutSDC.BillOfMaterial.ProductDescription,
		inoutSDC.BillOfMaterial.BillOfMaterialItem.BillOfMaterialComponent,
		inoutSDC.BillOfMaterial.BillOfMaterialItem.ComponentDescription,
		accepter,
	)
}
