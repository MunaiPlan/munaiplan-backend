package parser

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/xuri/excelize/v2"
)

type AdjustableGaugeStabilizersCatalog interface {
	GetCatalogName() string
	GetCatalogItemByID(id string) *ApiStabCatalogItem
}

type AdjustableGaugeStabilizersCatalogWrapper struct {
    XMLName xml.Name      `xml:"export"`
    Header  *CatalogHeader `xml:"CD_CATALOG"`
    Cache   map[string]*ApiStabCatalogItem
}

func NewAdjustableGaugeStabilizersCatalogWrapper(filePath, catalogCode, catalogItemCode string, file *excelize.File,) *AdjustableGaugeStabilizersCatalogWrapper {
    wrapper := &AdjustableGaugeStabilizersCatalogWrapper{
        Cache:  make(map[string]*ApiStabCatalogItem),
    }
    if err := wrapper.parseCatalog(filePath, catalogCode, catalogItemCode); err != nil {
        log.Fatal(err)
        return nil
    }
	writeCatalogToExcel(file, wrapper.Header, wrapper.Cache, wrapper.Header.Name, ApiStabCatalogItem{})
	println(wrapper.Header.Name)
    return wrapper
}

func (a *AdjustableGaugeStabilizersCatalogWrapper) parseCatalog(filePath, catalogCode, catalogItemCode string) error {
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
                item := ApiStabCatalogItem{}
                if err := decoder.DecodeElement(&item, &se); err != nil {
                    return fmt.Errorf("error decoding item: %w", err)
                }
                a.Cache[item.CatalogItemId] = &item
            }
        }
    }

    return nil
}

func (a *AdjustableGaugeStabilizersCatalogWrapper) GetCatalogName() string {
    return a.Header.Name
}

func (a *AdjustableGaugeStabilizersCatalogWrapper) GetCatalogItemByID(id string) *ApiStabCatalogItem {
    return a.Cache[id]
}