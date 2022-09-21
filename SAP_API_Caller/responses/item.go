package responses

type Item struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
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
		} `json:"results"`
	} `json:"d"`
}
