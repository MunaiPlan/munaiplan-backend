package parser

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/xuri/excelize/v2"
)

type ApiDrillPipeCatalog interface {
	GetCatalogName() string
	GetCatalogItemByID(id string) *ApiDrillPipeCatalogItem
}

type ApiDrillPipeCatalogWrapper struct {
    XMLName xml.Name      `xml:"export"`
    Header  *CatalogHeader `xml:"CD_CATALOG"`
    Cache   map[string]*ApiDrillPipeCatalogItem
}

func NewApiDrillPipeCatalogWrapper(filePath, catalogCode, catalogItemCode string, file *excelize.File,) *ApiDrillPipeCatalogWrapper {
    wrapper := &ApiDrillPipeCatalogWrapper{
        Cache:  make(map[string]*ApiDrillPipeCatalogItem),
    }
    if err := wrapper.parseCatalog(filePath, catalogCode, catalogItemCode); err != nil {
        log.Fatal(err)
        return nil
    }
	writeCatalogToExcel(file, wrapper.Header, wrapper.Cache, wrapper.Header.Name, ApiDrillPipeCatalogItem{})
	println(wrapper.Header.Name)
    return wrapper
}

func (a *ApiDrillPipeCatalogWrapper) parseCatalog(filePath, catalogCode, catalogItemCode string) error {
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
                item := ApiDrillPipeCatalogItem{}
                if err := decoder.DecodeElement(&item, &se); err != nil {
                    return fmt.Errorf("error decoding item: %w", err)
                }
                a.Cache[item.CatalogItemId] = &item
            }
        }
    }

    return nil
}

func (a *ApiDrillPipeCatalogWrapper) GetCatalogName() string {
    return a.Header.Name
}

func (a *ApiDrillPipeCatalogWrapper) GetCatalogItemByID(id string) *ApiDrillPipeCatalogItem {
    return a.Cache[id]
}
