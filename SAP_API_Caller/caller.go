package sap_api_caller

import (
	"fmt"
	"io/ioutil"
	sap_api_output_formatter "sap-api-integrations-bill-of-material-reads/SAP_API_Output_Formatter"
	"strings"
	"sync"

	sap_api_request_client_header_setup "github.com/latonaio/sap-api-request-client-header-setup"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
)

type SAPAPICaller struct {
	baseURL         string
	sapClientNumber string
	requestClient   *sap_api_request_client_header_setup.SAPRequestClient
	log             *logger.Logger
}

func NewSAPAPICaller(baseUrl, sapClientNumber string, requestClient *sap_api_request_client_header_setup.SAPRequestClient, l *logger.Logger) *SAPAPICaller {
	return &SAPAPICaller{
		baseURL:         baseUrl,
		requestClient:   requestClient,
		sapClientNumber: sapClientNumber,
		log:             l,
	}
}

func (c *SAPAPICaller) AsyncGetBillOfMaterial(material, plant, productDescription, billOfMaterialComponent, componentDescription string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "Header":
			func() {
				c.Header(material, plant)
				wg.Done()
			}()
		case "Item":
			func() {
				c.Item(material, plant)
				wg.Done()
			}()
		case "ProductDescription":
			func() {
				c.ProductDescription(plant, productDescription)
				wg.Done()
			}()
		case "Component":
			func() {
				c.Component(plant, billOfMaterialComponent)
				wg.Done()
			}()
		case "ComponentDescription":
			func() {
				c.ComponentDescription(plant, componentDescription)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}

func (c *SAPAPICaller) Header(material, plant string) {
	headerData, err := c.callBillOfMaterialSrvAPIRequirementHeader("MaterialBOM", material, plant)
	if err != nil {
		c.log.Error(err)
	} else {
		c.log.Info(headerData)
	}

	itemData, err := c.callToItem(headerData[0].ToItem)
	if err != nil {
		c.log.Error(err)
	} else {
	     c.log.Info(itemData)
	}
	return	 
}

func (c *SAPAPICaller) callBillOfMaterialSrvAPIRequirementHeader(api, material, plant string) ([]sap_api_output_formatter.Header, error) {
	url := strings.Join([]string{c.baseURL, "API_BILL_OF_MATERIAL_SRV;v=0002", api}, "/")
	param := c.getQueryWithHeader(map[string]string{}, material, plant)

	resp, err := c.requestClient.Request("GET", url, param, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToHeader(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToItem(url string) ([]sap_api_output_formatter.ToItem, error) {
	resp, err := c.requestClient.Request("GET", url, map[string]string{}, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToItem(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) Item(material, plant string) {
	data, err := c.callBillOfMaterialSrvAPIRequirementItem("MaterialBOMItem", material, plant)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)
}

func (c *SAPAPICaller) callBillOfMaterialSrvAPIRequirementItem(api, material, plant string) ([]sap_api_output_formatter.Item, error) {
	url := strings.Join([]string{c.baseURL, "API_BILL_OF_MATERIAL_SRV;v=0002", api}, "/")

	param := c.getQueryWithItem(map[string]string{}, material, plant)

	resp, err := c.requestClient.Request("GET", url, param, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToItem(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) ProductDescription(plant, productDescription string) {
	data, err := c.callBillOfMaterialSrvAPIRequirementProductDescription("MaterialBOM", plant, productDescription)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)
}

func (c *SAPAPICaller) callBillOfMaterialSrvAPIRequirementProductDescription(api, plant, productDescription string) ([]sap_api_output_formatter.Header, error) {
	url := strings.Join([]string{c.baseURL, "API_BILL_OF_MATERIAL_SRV;v=0002", api}, "/")

	param := c.getQueryWithProductDescription(map[string]string{}, plant, productDescription)

	resp, err := c.requestClient.Request("GET", url, param, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToHeader(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) Component(plant, billOfMaterialComponent string) {
	data, err := c.callBillOfMaterialSrvAPIRequirementComponent("MaterialBOMItem", plant, billOfMaterialComponent)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)
}

func (c *SAPAPICaller) callBillOfMaterialSrvAPIRequirementComponent(api, plant, billOfMaterialComponent string) ([]sap_api_output_formatter.Item, error) {
	url := strings.Join([]string{c.baseURL, "API_BILL_OF_MATERIAL_SRV;v=0002", api}, "/")

	param := c.getQueryWithComponent(map[string]string{}, plant, billOfMaterialComponent)

	resp, err := c.requestClient.Request("GET", url, param, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToItem(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) ComponentDescription(plant, componentDescription string) {
	data, err := c.callBillOfMaterialSrvAPIRequirementComponentDescription("MaterialBOMItem", plant, componentDescription)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)
}

func (c *SAPAPICaller) callBillOfMaterialSrvAPIRequirementComponentDescription(api, plant, componentDescription string) ([]sap_api_output_formatter.Item, error) {
	url := strings.Join([]string{c.baseURL, "API_BILL_OF_MATERIAL_SRV;v=0002", api}, "/")

	param := c.getQueryWithComponentDescription(map[string]string{}, plant, componentDescription)

	resp, err := c.requestClient.Request("GET", url, param, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToItem(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) getQueryWithHeader(params map[string]string, material, plant string) map[string]string {
	if len(params) == 0 {
		params = make(map[string]string, 1)
	}
	params["$filter"] = fmt.Sprintf("Material eq '%s' and Plant eq '%s'", material, plant)
	return params
}

func (c *SAPAPICaller) getQueryWithItem(params map[string]string, material, plant string) map[string]string {
	if len(params) == 0 {
		params = make(map[string]string, 1)
	}
	params["$filter"] = fmt.Sprintf("Material eq '%s' and Plant eq '%s'", material, plant)
	return params
}

func (c *SAPAPICaller) getQueryWithProductDescription(params map[string]string, plant, productDescription string) map[string]string {
	if len(params) == 0 {
		params = make(map[string]string, 1)
	}
	params["$filter"] = fmt.Sprintf("Plant eq '%s' and substringof('%s', ProductDescription)", plant, productDescription)
	return params
}

func (c *SAPAPICaller) getQueryWithComponent(params map[string]string, plant, billOfMaterialComponent string) map[string]string {
	if len(params) == 0 {
		params = make(map[string]string, 1)
	}
	params["$filter"] = fmt.Sprintf("Plant eq '%s' and BillOfMaterialComponent eq '%s'", plant, billOfMaterialComponent)
	return params
}

func (c *SAPAPICaller) getQueryWithComponentDescription(params map[string]string, plant, componentDescription string) map[string]string {
	if len(params) == 0 {
		params = make(map[string]string, 1)
	}
	params["$filter"] = fmt.Sprintf("Plant eq '%s' and substringof('%s', ComponentDescription)", plant, componentDescription)
	return params
}
