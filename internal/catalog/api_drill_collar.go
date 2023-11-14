package catalog

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"os"
)

type Item struct {
	ODNominal             string `xml:"od_nominal,attr"`
	CompTypeCode          string `xml:"comp_type_code,attr"`
	CreateUserID          string `xml:"create_user_id,attr"`
	ApproximateWeight     string `xml:"approximate_weight,attr"`
	ClosedEndDisplacement string `xml:"closed_end_displacement,attr"`
	AverageJointLength    string `xml:"average_joint_length,attr"`
	IDBody                string `xml:"id_body,attr"`
	Connection            string `xml:"connection,attr"`
	IDNominal             string `xml:"id_nominal,attr"`
	MakeupTorque          string `xml:"makeup_torque,attr"`
	ODBody                string `xml:"od_body,attr"`
	CatalogItemID         string `xml:"catalog_item_id,attr"`
	LinearCapacity        string `xml:"linear_capacity,attr"`
	APIIndicator          string `xml:"api_indicator,attr"`
	GradeID               string `xml:"grade_id,attr"`
	SectTypeCode          string `xml:"sect_type_code,attr"`
	CatalogID             string `xml:"catalog_id,attr"`
}

type ApiDrillCollarCatalog struct {
	CatalogID        string `xml:"CATALOG_ID,attr"`
	CatalogTypeCode  string `xml:"CATALOG_TYPE_CODE,attr"`
	Name             string `xml:"NAME,attr"`
	IsAPI            string `xml:"IS_API,attr"`
	Description      string `xml:"DESCRIPTION,attr"`
	ReadOnly         string `xml:"READ_ONLY,attr"`
	Comments         string `xml:"COMMENTS,attr"`
	Author           string `xml:"AUTHOR,attr"`
	CreateDate       string `xml:"CREATE_DATE,attr"`
	CreateUserID     string `xml:"CREATE_USER_ID,attr"`
	CreateAppID      string `xml:"CREATE_APP_ID,attr"`
	UpdateDate       string `xml:"UPDATE_DATE,attr"`
	UpdateUserID     string `xml:"UPDATE_USER_ID,attr"`
	UpdateAppID      string `xml:"UPDATE_APP_ID,attr"`
	Cache            map[string]*Item
}

type ApiDrillCollarCatalogWrapper struct {
	XMLName xml.Name               `xml:"export"`
	Catalog ApiDrillCollarCatalog  `xml:"CD_CATALOG"`
	Items   []Item                 `xml:"CD_DC_CATALOG"`
}

func NewApiDrillCollarCatalogWrapper(filePath string) *ApiDrillCollarCatalogWrapper {
	catalog := &ApiDrillCollarCatalogWrapper{}
	if err := catalog.parseCatalog(filePath); err != nil {
		log.Fatal(err)
		return nil
	}
	fmt.Println(catalog.Catalog.Name)
	return catalog
}

// parseCatalog reads and parses the XML file into the ApiDrillCollarCatalogWrapper.
func (a *ApiDrillCollarCatalogWrapper) parseCatalog(filePath string) error {
	// Open the XML file.
	xmlFile, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer xmlFile.Close()

	// Read the XML content.
	xmlData, err := io.ReadAll(xmlFile)
	if err != nil {
		return err
	}

	// Unmarshal the XML data into the ApiDrillCollarCatalogWrapper struct.
	if err := xml.Unmarshal(xmlData, a); err != nil {
		return err
	}

	// Initialize the cache map.
	a.Catalog.Cache = make(map[string]*Item)

	// Iterate over the items and add them to the cache.
	for i := range a.Items {
		a.Catalog.Cache[a.Items[i].CatalogItemID] = &a.Items[i]
	}

	return nil
}

func (c *ApiDrillCollarCatalogWrapper) PrintCatalog() {
	for k, v := range c.Catalog.Cache {
		fmt.Println(k, "value is", v)
	}
}