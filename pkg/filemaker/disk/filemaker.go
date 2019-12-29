package disk

import "encoding/xml"

// File is the outer master file format
type File struct {
	XMLName   xml.Name   `xml:"FMPXMLRESULT"`
	ErrorCode int64      `xml:"ERRORCODE"`
	Product   *Product   `xml:"PRODUCT"`
	Database  *Database  `xml:"DATABASE"`
	Metadata  *Metadata  `xml:"METADATA"`
	ResultSet *ResultSet `xml:"RESULTSET"`
}

// Product entity
type Product struct {
	XMLName xml.Name `xml:"PRODUCT"`
	Build   string   `xml:"BUILD,attr"`
	Name    string   `xml:"NAME,attr"`
	Version string   `xml:"VERSION,attr"`
}

// Database entity
type Database struct {
	XMLName    xml.Name `xml:"DATABASE"`
	DateFormat string   `xml:"DATEFORMAT,attr"`
	Name       string   `xml:"NAME,attr"`
	TimeFormat string   `xml:"TIMEFORMAT,attr"`
	Records    string   `xml:"RECORDS,attr"`
	Layout     string   `xml:"LAYOUT,attr"`
}

// Metadata entity
type Metadata struct {
	XMLName xml.Name `xml:"METADATA"`
	Fileds  []*Field `xml:"FIELD"`
}

// Field entity
type Field struct {
	XMLName   xml.Name `xml:"FIELD"`
	EmptyOK   string   `xml:"EMPTYOK,attr"`
	MaxRepeat string   `xml:"MAXREPEAT,attr"`
	Name      string   `xml:"NAME,attr"`
	Type      string   `xml:"TYPE,attr"`
}

// ResultSet entity
type ResultSet struct {
	XMLName xml.Name `xml:"RESULTSET"`
	Rows    []*Row   `xml:"ROW"`
	Found   string   `xml:"FOUND,attr"`
}

// Row entity
type Row struct {
	XMLName  xml.Name `xml:"ROW"`
	ModID    string   `xml:"MODID,attr"`
	RecordID string   `xml:"RECORDID,attr"`
	Cols     []*Col   `xml:"COL"`
}

// Col entity
type Col struct {
	XMLName xml.Name `xml:"COL"`
	Data    []string `xml:"DATA"`
}
