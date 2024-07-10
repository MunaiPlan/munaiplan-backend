package parser

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/xuri/excelize/v2"
)

type ApiDrillCollarCatalog interface {
	GetCatalogName() string
	GetCatalogItemByID(id string) *ApiDrillCollarCatalogItem
}

type ApiDrillCollarCatalogWrapper struct {
	XMLName xml.Name       `xml:"export"`
	Header  *CatalogHeader `xml:"CD_CATALOG"`
	Cache   map[string]*ApiDrillCollarCatalogItem
}

func NewApiDrillCollarCatalogWrapper(catalogFilePath, catalogCode, catalogItemCode string, file *excelize.File) *ApiDrillCollarCatalogWrapper {
	wrapper := &ApiDrillCollarCatalogWrapper{
		Cache: make(map[string]*ApiDrillCollarCatalogItem),
	}
	if err := wrapper.parseCatalog(catalogFilePath, catalogCode, catalogItemCode); err != nil {
		log.Fatal(err)
		return nil
	}
	writeCatalogToExcel(file, wrapper.Header, wrapper.Cache, "Api Drill Collar", ApiDrillCollarCatalogItem{})
	println(wrapper.Header.Name)
	return wrapper
}

func (a *ApiDrillCollarCatalogWrapper) parseCatalog(filePath, catalogCode, catalogItemCode string) error {
	xmlFile, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer xmlFile.Close()

	decoder := xml.NewDecoder(xmlFile)

	for {
		token, err := decoder.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			return fmt.Errorf("error reading token: %w", err)
		}

		switch se := token.(type) {
		case xml.StartElement:
			if se.Name.Local == catalogCode {
				header := CatalogHeader{}
				if err := decoder.DecodeElement(&header, &se); err != nil {
					return fmt.Errorf("error decoding header: %w", err)
				}
				a.Header = &header
			}

			if se.Name.Local == catalogItemCode {
				item := ApiDrillCollarCatalogItem{}
				if err := decoder.DecodeElement(&item, &se); err != nil {
					return fmt.Errorf("error decoding item: %w", err)
				}
				a.Cache[item.CatalogItemId] = &item
			}
		}
	}

	return nil
}

func (a *ApiDrillCollarCatalogWrapper) GetCatalogName() string {
	return a.Header.Name
}

func (a *ApiDrillCollarCatalogWrapper) GetCatalogItemByID(id string) *ApiDrillCollarCatalogItem {
	return a.Cache[id]
}
