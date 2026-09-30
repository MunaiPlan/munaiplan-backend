package wellplan

import (
	"archive/zip"
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"strings"
)

// Upload and decompression limits guard against oversized files and zip bombs.
const (
	MaxUploadBytes      = 10 << 20
	maxDocumentXMLBytes = 64 << 20
)

// ErrInvalidFile is returned for input that is not a readable WellPlan export.
var ErrInvalidFile = errors.New("not a readable WellPlan export")

// block is a body-level paragraph (text) or table (rows of cell text), in document order.
type block struct {
	text string
	rows [][]string
}

func (b block) isTable() bool { return b.rows != nil }

const wordNS = "http://schemas.openxmlformats.org/wordprocessingml/2006/main"

// readDocx extracts paragraphs and tables from word/document.xml. Nested tables are
// flattened into the text of the enclosing cell.
func readDocx(data []byte) ([]block, error) {
	if len(data) > MaxUploadBytes {
		return nil, fmt.Errorf("%w: file exceeds %d MB", ErrInvalidFile, MaxUploadBytes>>20)
	}
	archive, err := zip.NewReader(bytes.NewReader(data), int64(len(data)))
	if err != nil {
		return nil, fmt.Errorf("%w: not a .docx archive", ErrInvalidFile)
	}
	var doc *zip.File
	for _, f := range archive.File {
		if f.Name == "word/document.xml" {
			doc = f
			break
		}
	}
	if doc == nil {
		return nil, fmt.Errorf("%w: word/document.xml is missing", ErrInvalidFile)
	}
	if doc.UncompressedSize64 > maxDocumentXMLBytes {
		return nil, fmt.Errorf("%w: document is too large", ErrInvalidFile)
	}
	rc, err := doc.Open()
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrInvalidFile, err)
	}
	defer rc.Close()
	return parseDocumentXML(io.LimitReader(rc, maxDocumentXMLBytes))
}

func parseDocumentXML(r io.Reader) ([]block, error) {
	dec := xml.NewDecoder(r)
	var (
		blocks     []block
		tableDepth int
		rows       [][]string
		row        []string
		cell       strings.Builder
		para       strings.Builder
		inText     bool
	)
	sink := func() *strings.Builder {
		if tableDepth > 0 {
			return &cell
		}
		return &para
	}
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("%w: malformed document XML", ErrInvalidFile)
		}
		switch t := tok.(type) {
		case xml.StartElement:
			if t.Name.Space != wordNS {
				continue
			}
			switch t.Name.Local {
			case "tbl":
				tableDepth++
				if tableDepth == 1 {
					rows = [][]string{}
				}
			case "tr":
				if tableDepth == 1 {
					row = nil
				}
			case "tc":
				if tableDepth == 1 {
					cell.Reset()
				}
			case "t":
				inText = true
			case "tab", "br":
				sink().WriteByte(' ')
			}
		case xml.EndElement:
			if t.Name.Space != wordNS {
				continue
			}
			switch t.Name.Local {
			case "t":
				inText = false
			case "p":
				if tableDepth > 0 {
					cell.WriteByte(' ')
				} else if text := clean(para.String()); text != "" {
					blocks = append(blocks, block{text: text})
					para.Reset()
				} else {
					para.Reset()
				}
			case "tc":
				if tableDepth == 1 {
					row = append(row, clean(cell.String()))
				}
			case "tr":
				if tableDepth == 1 {
					rows = append(rows, row)
				}
			case "tbl":
				if tableDepth == 1 {
					blocks = append(blocks, block{rows: rows})
				}
				tableDepth--
			}
		case xml.CharData:
			if inText {
				sink().Write(t)
			}
		}
	}
	return blocks, nil
}

// clean collapses whitespace (including non-breaking spaces) and trims.
func clean(s string) string {
	return strings.Join(strings.Fields(strings.ReplaceAll(s, " ", " ")), " ")
}
